// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/department"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Department is the model entity for the Department schema.
type Department struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 创建者
	CreatorID int `json:"creator_id,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 最后更新者
	UpdaterID int `json:"updater_id,omitempty"`
	// 最后更新时间
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 名称
	Name string `json:"name,omitempty" validate:"required"`
	// 上级部门id
	ParentID int `json:"parent_id,omitempty"`
	// 备注
	Notes string `json:"notes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges        DepartmentEdges `json:"edges"`
	selectValues sql.SelectValues

	AllEmployees []*Employee `json:"all_employees"`
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Creator holds the value of the creator edge.
	Creator *Admin `json:"creator,omitempty"`
	// Updater holds the value of the updater edge.
	Updater *Admin `json:"updater,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Department `json:"parent,omitempty"`
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// Employees holds the value of the employees edge.
	Employees []*Employee `json:"employees,omitempty"`
	// Children holds the value of the children edge.
	Children []*Department `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) CreatorOrErr() (*Admin, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// UpdaterOrErr returns the Updater value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) UpdaterOrErr() (*Admin, error) {
	if e.loadedTypes[1] {
		if e.Updater == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Updater, nil
	}
	return nil, &NotLoadedError{edge: "updater"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) ParentOrErr() (*Department, error) {
	if e.loadedTypes[2] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[3] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// EmployeesOrErr returns the Employees value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) EmployeesOrErr() ([]*Employee, error) {
	if e.loadedTypes[4] {
		return e.Employees, nil
	}
	return nil, &NotLoadedError{edge: "employees"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) ChildrenOrErr() ([]*Department, error) {
	if e.loadedTypes[5] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case department.FieldID, department.FieldCreatorID, department.FieldUpdaterID, department.FieldParentID:
			values[i] = new(sql.NullInt64)
		case department.FieldName, department.FieldNotes:
			values[i] = new(sql.NullString)
		case department.FieldCreateTime, department.FieldDeleteTime, department.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case department.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Time
			}
		case department.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				d.CreatorID = int(value.Int64)
			}
		case department.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				d.DeleteTime = new(time.Time)
				*d.DeleteTime = value.Time
			}
		case department.FieldUpdaterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updater_id", values[i])
			} else if value.Valid {
				d.UpdaterID = int(value.Int64)
			}
		case department.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				d.UpdateTime = value.Time
			}
		case department.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case department.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				d.ParentID = int(value.Int64)
			}
		case department.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				d.Notes = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Department.
// This includes values selected through modifiers, order, etc.
func (d *Department) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the Department entity.
func (d *Department) QueryCreator() *AdminQuery {
	return NewDepartmentClient(d.config).QueryCreator(d)
}

// QueryUpdater queries the "updater" edge of the Department entity.
func (d *Department) QueryUpdater() *AdminQuery {
	return NewDepartmentClient(d.config).QueryUpdater(d)
}

// QueryParent queries the "parent" edge of the Department entity.
func (d *Department) QueryParent() *DepartmentQuery {
	return NewDepartmentClient(d.config).QueryParent(d)
}

// QueryPermissions queries the "permissions" edge of the Department entity.
func (d *Department) QueryPermissions() *PermissionQuery {
	return NewDepartmentClient(d.config).QueryPermissions(d)
}

// QueryEmployees queries the "employees" edge of the Department entity.
func (d *Department) QueryEmployees() *EmployeeQuery {
	return NewDepartmentClient(d.config).QueryEmployees(d)
}

// QueryChildren queries the "children" edge of the Department entity.
func (d *Department) QueryChildren() *DepartmentQuery {
	return NewDepartmentClient(d.config).QueryChildren(d)
}

// Update returns a builder for updating this Department.
// Note that you need to call Department.Unwrap() before calling this method if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return NewDepartmentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Department entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("dao: Department is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("create_time=")
	builder.WriteString(d.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatorID))
	builder.WriteString(", ")
	if v := d.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("updater_id=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdaterID))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(d.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", d.ParentID))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(d.Notes)
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department
