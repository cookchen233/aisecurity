package filters

import (
	"aisecurity/structs"
)

type Video struct {
	structs.StandardFilter
	Name string `form:"name"`
	URL  string `form:"url"`
}
