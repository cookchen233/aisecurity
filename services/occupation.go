package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/occupation"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
)

type OccupationService struct {
	Service
	entClient *dao.OccupationClient
}

func NewOccupationService() *OccupationService {
	return &OccupationService{
		entClient: db.EntClient.Occupation,
	}
}

func (s *OccupationService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Occupation)
	save, err := db.EntClient.Occupation.Create().
		SetName(e.Name).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Occupation")
	}
	return save, nil
}

func (s *OccupationService) query(fit structs.IFilter) *dao.OccupationQuery {
	f := fit.(*filters.Occupation)
	q := db.EntClient.Occupation.Query()
	if f.ID != 0 {
		q = q.Where(occupation.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(occupation.NameContainsFold(f.Name))
	}
	return q.Clone()
}

func (s *OccupationService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *OccupationService) GetList(filter structs.IFilter) ([]structs.IEntity, error) {
	// list
	page := min(1000, max(1, filter.GetOffset()))
	limit := min(10000, max(1, filter.GetOffset()))
	offset := (page - 1) * limit
	list, err := s.query(filter).
		WithCreator().
		WithUpdater().
		Limit(limit). // Set the number of items to return
		Offset(offset).
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
