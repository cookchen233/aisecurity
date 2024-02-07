package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/occupation"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
)

type OccupationService struct {
	Service
	entClient *dao.OccupationClient
}

func NewOccupationService(ctx context.Context) *OccupationService {
	s := &OccupationService{entClient: db.EntClient.Occupation}
	s.Ctx = ctx
	return s
}

var ()

func (s *OccupationService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Occupation)
	c := tx.Occupation.Create().
		SetName(e.Name)

	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Occupation"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.Occupation, entities.Occupation](saved), nil
}

func (s *OccupationService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Occupation)
	u := tx.Occupation.UpdateOneID(e.ID).
		SetName(e.Name)
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating Occupation"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.Occupation, entities.Occupation](saved), nil
}

func (s *OccupationService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
	fit.SetPage(1)
	fit.SetLimit(1)
	list, err := s.GetList(fit)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	if len(list) == 0 {
		return nil, utils.ErrorWithStack(expects.NewDataNotFound())
	}
	return list[0], nil
}

func (s *OccupationService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithCreator().
		Order(dao.Desc(occupation.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = structs.ConvertTo[*dao.Occupation, entities.Occupation](v)
	}
	return ents, nil
}

func (s *OccupationService) query(fit structs.IFilter) *dao.OccupationQuery {
	f := fit.(*filters.Occupation)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(occupation.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(occupation.NameEQ(f.Name))
	}
	return q.Clone()
}

func (s *OccupationService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}
