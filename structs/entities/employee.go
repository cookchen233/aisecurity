package entities

import (
	"aisecurity/ent/dao"
)

type Employee struct {
	dao.Employee
	OccupationID int           `json:"occupation_id"`
	Edges        EmployeeEdges `json:"edges"`
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	dao.EmployeeEdges
	TopDepartment *dao.Department `json:"top_department,omitempty"`
}
