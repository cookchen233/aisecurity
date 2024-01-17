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

func (s *RiskCategoryService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.RiskCategory)
	save, err := db.EntClient.RiskCategory.Create().
		SetName(e.Name).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating RiskCategory")
	}
	return save, nil
}

func (s *RiskCategoryService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.RiskCategory)
	save, err := db.EntClient.RiskCategory.UpdateOneID(e.ID).
		SetName(e.Name).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating RiskCategory")
	}
	return save, nil
}

func (s *RiskCategoryService) query(fit structs.IFilter) *dao.RiskCategoryQuery {
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

func (s *RiskCategoryService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *RiskCategoryService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithCreator().
		WithUpdater().
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(riskcategory.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = v
	}
	return ents, nil
}

func (s *RiskCategoryService) Delete(ent structs.IEntity) error {
	e := ent.(*entities.ID)
	err := s.entClient.DeleteOneID(e.ID).
		Exec(s.Ctx)
	if err != nil {
		return utils.ErrorWrap(err, "failed deleting RiskCategory")
	}
	return nil
}
