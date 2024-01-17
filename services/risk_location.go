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

func (s *RiskLocationService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.RiskLocation)
	save, err := db.EntClient.RiskLocation.Create().
		SetName(e.Name).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating RiskLocation")
	}
	return save, nil
}

func (s *RiskLocationService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.RiskLocation)
	save, err := db.EntClient.RiskLocation.UpdateOneID(e.ID).
		SetName(e.Name).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating RiskLocation")
	}
	return save, nil
}

func (s *RiskLocationService) query(fit structs.IFilter) *dao.RiskLocationQuery {
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

func (s *RiskLocationService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *RiskLocationService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithCreator().
		WithUpdater().
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(risklocation.FieldID)).
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

func (s *RiskLocationService) Delete(ent structs.IEntity) error {
	e := ent.(*entities.ID)
	err := s.entClient.DeleteOneID(e.ID).
		Exec(s.Ctx)
	if err != nil {
		return utils.ErrorWrap(err, "failed deleting RiskLocation")
	}
	return nil
}
