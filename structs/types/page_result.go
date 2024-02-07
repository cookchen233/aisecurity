package types

import (
	"aisecurity/structs"
)

type PageResult struct {
	Total int               `json:"total"`
	List  []structs.IEntity `json:"list"`
}
