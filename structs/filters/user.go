package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
)

type User struct {
	structs.StandardFilter
	UserStatus enums.EnabledStatus `form:"user_status"`
}
