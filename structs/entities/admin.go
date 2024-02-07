package entities

import (
	"aisecurity/ent/dao"
)

type Admin struct {
	dao.Admin
	Occupation  *dao.Occupation   `json:"occupation"`
	Permissions []*dao.Permission `json:"permissions"`
	Edges       AdminEdges        `json:"edges"`
}

type AdminEdges struct {
	dao.AdminEdges
	Employee Employee `json:"employee"`
}
