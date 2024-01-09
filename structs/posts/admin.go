package posts

import "aisecurity/structs"

type DashboardLogin struct {
	structs.StandardPost
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PersistDays int    `json:"persist_days"`
}
