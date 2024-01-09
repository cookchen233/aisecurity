package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/riskcategory"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
)

type RiskCategoryService struct {
	Service
	entClient *dao.RiskCategoryClient
}

func NewRiskCategoryService() *RiskCategoryService {
	return &RiskCategoryService{
		entClient: db.EntClient.RiskCategory,
	}
}

func (service *RiskCategoryService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.RiskCategory)
	save, err := db.EntClient.RiskCategory.Create().
		SetName(e.Name).
		Save(service.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating RiskCategory")
	}
	return save, nil
}

func (service *RiskCategoryService) query(fit structs.IFilter) *dao.RiskCategoryQuery {
	f := fit.(*filters.RiskCategory)
	q := db.EntClient.RiskCategory.Query()
	if f.ID != 0 {
		q = q.Where(riskcategory.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(riskcategory.NameContainsFold(f.Name))
	}
	return q.Clone()
}

func (service *RiskCategoryService) GetToal(fit structs.IFilter) (int, error) {
	// total
	total, err := service.query(fit).Count(service.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (service *RiskCategoryService) GetList(filter structs.IFilter) ([]structs.IEntity, error) {
	// list
	page := min(1000, max(1, filter.GetOffset()))
	limit := min(10000, max(1, filter.GetOffset()))
	offset := (page - 1) * limit
	list, err := service.query(filter).
		WithCreator().
		WithUpdater().
		Limit(limit). // Set the number of items to return
		Offset(offset).
		All(service.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = v
	}
	return ents, nil
}
