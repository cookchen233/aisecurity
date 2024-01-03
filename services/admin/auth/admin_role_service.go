package auth

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/adminrole"
	"aisecurity/utils/db"
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"sync"
)

type AdminRoleService struct {
	entClient *dao.AdminRoleClient
}

func NewAdminRoleService() *AdminRoleService {
	return &AdminRoleService{
		entClient: db.EntClient.AdminRole,
	}
}

var (
	modules        []Module
	onceLoadModule sync.Once
)

type config struct {
	Modules []Module `toml:"Modules"`
}

type Module struct {
	Name    string `toml:"Name"`
	Title   string `toml:"Title"`
	Icon    string `toml:"Icon"`
	Enabled int    `toml:"IsEnabled"`
	Menus   []Menu `toml:"Menus"`
}

type Menu struct {
	ID          string    `toml:"id"`
	Title       string    `toml:"Title"`
	DisplayType int       `toml:"DisplayType"`
	Icon        string    `toml:"Icon"`
	Enabled     int       `toml:"IsEnabled"`
	Submenus    []SubMenu `toml:"Submenus"`
}

type SubMenu struct {
	ID          string `toml:"id"`
	Title       string `toml:"Title"`
	DisplayType int    `toml:"DisplayType"`
	Icon        string `toml:"Icon"`
	Enabled     int    `toml:"IsEnabled"`
}

func (service *AdminRoleService) GetModules() ([]Module, error) {
	onceLoadModule.Do(func() {
		var cfg config
		_, err := toml.DecodeFile("config/admin_permissions.toml", &cfg)
		if err != nil {
			log.Fatalf("error parsing TOML: %v", err)
		}
		modules = cfg.Modules
	})
	if modules == nil {
		return nil, fmt.Errorf("modules load failed")
	}
	return modules, nil
}

func (service *AdminRoleService) GetSuperRole() (*dao.AdminRole, error) {
	data, err := service.entClient.Query().Order(adminrole.ByID(sql.OrderAsc())).First(context.Background())
	if err != nil {
		if dao.IsNotFound(err) {
			fmt.Printf("there is no one role yet, create a new one")
			saved, err := service.entClient.Create().
				SetName("超级管理员").
				Save(context.Background())
			if err != nil {
				return nil, fmt.Errorf("failed creating super role: %w", err)
			}
			return saved, nil
		}
		return nil, fmt.Errorf("failed getting super role: %w", err)
	}
	// super role must be the first record.
	if data.ID != 1 {
		return nil, fmt.Errorf("the ID of super role is incorrect: %v", data.ID)
	}
	return data, nil
}

func (service *AdminRoleService) Create(data *dao.AdminRole) (*dao.AdminRole, error) {
	saved, err := service.entClient.Create().
		SetName(data.Name).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating AdminRole: %w", err)
	}
	return saved, nil
}

func (service *AdminRoleService) GetList() ([]*dao.AdminRole, error) {
	all, err := service.entClient.Query().WithCreator().All(context.Background())
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *AdminRoleService) GetByID(id int) (*dao.AdminRole, error) {
	one, err := service.entClient.Query().Where(adminrole.IDEQ(id)).WithCreator().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return one, nil
}
