package entities

import (
	"aisecurity/ent/dao"
)

type User struct {
	dao.User
	SetPassword string `json:"set_password"`
}
