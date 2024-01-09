package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/risklocation"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
)

type RiskLocationService struct {
	Service
	entClient *dao.RiskLocationClient
}

func NewRiskLocationService() *RiskLocationService {
	return &RiskLocationService{
		entClient: db.EntClient.RiskLocation,
	}
}

func (service *RiskLocationService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.RiskLocation)
	save, err := db.EntClient.RiskLocation.Create().
		SetName(e.Name).
		Save(service.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating RiskLocation")
	}
	return save, nil
}

func (service *RiskLocationService) query(fit structs.IFilter) *dao.RiskLocationQuery {
	f := fit.(*filters.RiskLocation)
	q := db.EntClient.RiskLocation.Query()
	if f.ID != 0 {
		q = q.Where(risklocation.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(risklocation.NameContainsFold(f.Name))
	}
	return q.Clone()
}

func (service *RiskLocationService) GetToal(fit structs.IFilter) (int, error) {
	// total
	total, err := service.query(fit).Count(service.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (service *RiskLocationService) GetList(filter structs.IFilter) ([]structs.IEntity, error) {
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
