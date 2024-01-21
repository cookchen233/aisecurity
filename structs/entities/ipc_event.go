package entities

import (
	"aisecurity/ent/dao"
)

type IPCEvent struct {
	dao.IPCEvent
	EventTypeLabel     string             `json:"event_type_label"`
	EventStatusLabel   string             `json:"event_status_label"`
	DeviceInstallation DeviceInstallation `json:"device_installation"`
	EventLevel         EventLevel         `json:"event_level"`
}
