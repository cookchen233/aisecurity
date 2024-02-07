package entities

import (
	"aisecurity/ent/dao"
)

type Fixing struct {
	dao.Fixing
	Fixer    *dao.Admin    `json:"fixer"`
	Event    *dao.Event    `json:"event"`
	Device   *dao.Device   `json:"device"`
	EventLog *dao.EventLog `json:"event_log"`
}
