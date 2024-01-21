package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/area"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
)

type Area struct {
	Service
	entClient *dao.AreaClient
}

func NewAreaService(ctx context.Context) *Area {
	s := &Area{entClient: db.EntClient.Area}
	s.Ctx = ctx
	return s
}

func (s *Area) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Area)
	c := s.entClient.Create().
		SetName(e.Name).
		SetDescription(e.Description)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Area")
	}
	// ConvertTo now returns a non-pointer type that implements IEntity
	return structs.ConvertTo[*dao.Area, entities.Area](saved), nil
}

func (s *Area) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Area)
	u := s.entClient.UpdateOneID(e.ID).
		SetName(e.Name).
		SetDescription(e.Description)
	save, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating Area")
	}
	return structs.ConvertTo[*dao.Area, entities.Area](save), nil
}

func (s *Area) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *Area) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.Area)
		ents[i] = v
		err := gconv.Struct(v, v2)
		if err != nil {
			utils.Logger.Warn("convert error", zap.Error(err))
		} else {
			ents[i] = v2
		}
	}
	return ents, nil
}

func (s *Area) query(fit structs.IFilter) *dao.AreaQuery {
	f := fit.(*filters.Area)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(area.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(area.NameEQ(f.Name))
	}
	return q.Clone()
}

func (s *Area) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *Area) GetByName(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(area.NameEQ(name)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.Area, entities.Area](first), nil
}

func (s *Area) GetAnyOne() (structs.IEntity, error) {
	list, err := s.GetList(&filters.Area{})
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		saved, err := s.Create(&entities.Area{
			Area: dao.Area{
				Name: "默认区域",
			},
		})
		if err != nil {
			utils.Logger.Error("failed to create area", zap.Error(err))
		}
		return saved, nil
	}
	return list[0], nil
}
