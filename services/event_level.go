package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/eventlevel"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	stdsql "database/sql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"reflect"
)

type EventLevelService struct {
	Service
	entClient *dao.EventLevelClient
}

func NewEventLevelService(ctx context.Context) *EventLevelService {
	s := &EventLevelService{entClient: db.EntClient.EventLevel}
	s.Ctx = ctx
	return s
}

func (s *EventLevelService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.EventLevel)
	c := s.entClient.Create().
		SetName(e.Name).
		SetDescription(e.Description).
		SetEventTypes(e.EventTypes).
		SetIsReport(e.IsReport)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating EventLevelService")
	}
	// ConvertTo now returns a non-pointer type that implements IEntity
	return structs.ConvertTo[*dao.EventLevel, entities.EventLevel](saved), nil
}

func (s *EventLevelService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.EventLevel)
	u := s.entClient.UpdateOneID(e.ID).
		SetDescription(e.Description).
		SetEventTypes(e.EventTypes).
		SetIsReport(e.IsReport)
	save, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating EventLevelService")
	}
	return structs.ConvertTo[*dao.EventLevel, entities.EventLevel](save), nil
}

func (s *EventLevelService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *EventLevelService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
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
		v2 := new(entities.EventLevel)
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

func (s *EventLevelService) query(fit structs.IFilter) *dao.EventLevelQuery {
	f := fit.(*filters.EventLevel)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(eventlevel.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(eventlevel.NameEQ(f.Name))
	}
	if f.EventType != 0 {
		q = q.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(s.C("event_types"), f.EventType))
		})
		//q = q.Where(func(s *sql.Selector) {
		//    s.Where(sqljson.ValueEQ(eventlevel.FieldEventTypes, key, value))
		//})
		//q = q.Where(func(s *sql.Selector) {
		//	// Using the jsonb array contains operator @>
		//	s.Where(sql.P(func(b *sql.Builder) {
		//		b.WriteString("event_types @> ?")
		//		b.Arg(18) // Argument is a JSON array as string
		//	}))
		//})
	}
	return q.Clone()
}

func (s *EventLevelService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *EventLevelService) GetByName(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(eventlevel.NameEQ(name)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.EventLevel, entities.EventLevel](first), nil
}
func (s *EventLevelService) GetByEventType2(eventType enums.EventType) (structs.IEntity, error) {
	result, err := s.entClient.QueryContext(s.Ctx, "SELECT * FROM event_level WHERE event_types = ANY($1) LIMIT 1", eventType)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	defer func(result *stdsql.Rows) {
		err := result.Close()
		if err != nil {
			utils.Logger.Error("failed closing rows", zap.Error(err))
		}
	}(result)

	for result.Next() {
		level := entities.EventLevel{}
		s := reflect.ValueOf(&level).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := result.Scan(columns...)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}
		return &level, nil
	}
	return nil, &dao.NotFoundError{}
}

func (s *EventLevelService) GetByEventType(eventType enums.EventType) (structs.IEntity, error) {

	result, err := s.entClient.QueryContext(s.Ctx, "SELECT * FROM event_level WHERE event_types = ANY(event_types) LIMIT 1", eventType)

	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	defer func(result *stdsql.Rows) {
		err := result.Close()
		if err != nil {
			utils.Logger.Error("failed closing rows", zap.Error(err))
		}
	}(result)

	//var rows []structs.IEntity
	for result.Next() {
		level := entities.EventLevel{}
		s := reflect.ValueOf(&level).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := result.Scan(columns...)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}
		return &level, nil
	}
	return nil, &dao.NotFoundError{}
}
