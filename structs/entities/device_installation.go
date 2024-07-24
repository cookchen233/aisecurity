package entities

import (
	"aisecurity/ent/dao"
)

type DeviceInstallation struct {
	dao.DeviceInstallation
	AreaName              string `json:"area_name"`
	LocationWithAliasName string `json:"location_with_alias_name"`
}
