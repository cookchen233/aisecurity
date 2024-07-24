package entities

import (
	"aisecurity/ent/dao"
)

type Admin struct {
	dao.Admin
	SetPassword string            `json:"set_password"`
	Occupation  *dao.Occupation   `json:"occupation"`
	Permissions []*dao.Permission `json:"permissions"`
	Edges       AdminEdges        `json:"edges"`
	DisplayName string            `json:"displayName"`
}

type AdminEdges struct {
	dao.AdminEdges
	Employee Employee `json:"employee"`
}
