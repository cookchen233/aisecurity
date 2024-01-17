package entities

import "aisecurity/structs"

type ID struct {
	structs.StandardIEntity
	ID int `json:"id"  validate:"required"`
}
