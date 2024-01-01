package auth

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"sync"
)

type RoleService struct {
	msgs     []byte
	isOnline bool
}

func NewRoleService() *RoleService {
	return &RoleService{}
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

func (service *RoleService) GetModules() ([]Module, error) {
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
