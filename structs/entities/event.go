package entities

import (
	"aisecurity/ent/dao"
)

type Event struct {
	dao.Event
	EventTypeLabel        string `json:"event_type_label"`
	EventStatusLabel      string `json:"event_status_label"`
	Location              string `json:"location"`
	LocationWithAliasName string `json:"location_with_alias_name"`
}
