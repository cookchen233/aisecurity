package entities

import (
	"aisecurity/ent/dao"
)

type Admin struct {
	dao.Admin
	AdminRoleIDs []int `json:"admin_role_ids"`
}
