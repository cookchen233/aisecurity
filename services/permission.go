package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/permission"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"log"
	"sync"
)

type PermissionService struct {
	Service
	entClient *dao.PermissionClient
}

func NewPermissionService(ctx context.Context) *PermissionService {
	s := &PermissionService{
		entClient: db.EntClient.Permission,
	}
	s.Ctx = ctx
	return s
}

var (
	modules        []AccessGroup
	onceLoadModule sync.Once
)

type config struct {
	AccessGroups []AccessGroup `toml:"access_group"`
}

type AccessGroup struct {
	ID       string         `toml:"id" json:"id"`
	Title    string         `toml:"title" json:"title"`
	Enabled  bool           `toml:"enabled" json:"enabled"`
	Accesses []types.Access `toml:"access" json:"accesses"`
}

func (s *PermissionService) GetAccessGroups() ([]AccessGroup, error) {
	onceLoadModule.Do(func() {
		var cfg config
		_, err := toml.DecodeFile("config/dashboard_accesses.toml", &cfg)
		if err != nil {
			log.Fatalf("error parsing TOML: %v", err)
		}
		modules = cfg.AccessGroups
	})
	if modules == nil {
		return nil, fmt.Errorf("modules load failed")
	}
	return modules, nil
}

func (s *PermissionService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Permission)
	saved, err := s.entClient.Create().
		SetName(e.Name).
		SetAccessIds(e.AccessIds).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Permission")
	}
	return saved, nil
}

func (s *PermissionService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Permission)
	saved, err := s.entClient.UpdateOneID(e.ID).
		SetName(e.Name).
		SetAccessIds(e.AccessIds).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating Permission")
	}
	return saved, nil
}

func (s *PermissionService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *PermissionService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(permission.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.Permission)
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

func (s *PermissionService) query(fit structs.IFilter) *dao.PermissionQuery {
	f := fit.(*filters.Permission)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(permission.IDEQ(f.ID))
	}
	if f.Keywords != "" {
		q = q.Where(permission.NameContainsFold(f.Keywords))
	}
	if f.Name != "" {
		q = q.Where(permission.NameEQ(f.Name))
	}
	return q.Clone()
}

func (s *PermissionService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}
