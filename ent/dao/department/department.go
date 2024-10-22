// Code generated by ent, DO NOT EDIT.

package department

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldCreatorID holds the string denoting the creator_id field in the database.
	FieldCreatorID = "creator_id"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldUpdaterID holds the string denoting the updater_id field in the database.
	FieldUpdaterID = "updater_id"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldNotes holds the string denoting the notes field in the database.
	FieldNotes = "notes"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeUpdater holds the string denoting the updater edge name in mutations.
	EdgeUpdater = "updater"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgeEmployees holds the string denoting the employees edge name in mutations.
	EdgeEmployees = "employees"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the department in the database.
	Table = "departments"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "departments"
	// CreatorInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	CreatorInverseTable = "admins"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "creator_id"
	// UpdaterTable is the table that holds the updater relation/edge.
	UpdaterTable = "departments"
	// UpdaterInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	UpdaterInverseTable = "admins"
	// UpdaterColumn is the table column denoting the updater relation/edge.
	UpdaterColumn = "updater_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "departments"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "permission_department"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permissions"
	// EmployeesTable is the table that holds the employees relation/edge.
	EmployeesTable = "employees"
	// EmployeesInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeesInverseTable = "employees"
	// EmployeesColumn is the table column denoting the employees relation/edge.
	EmployeesColumn = "department_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "departments"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldCreatorID,
	FieldDeleteTime,
	FieldUpdaterID,
	FieldUpdateTime,
	FieldName,
	FieldParentID,
	FieldNotes,
}

var (
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"permission_id", "department_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "aisecurity/ent/dao/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// CreatorIDValidator is a validator for the "creator_id" field. It is called by the builders before save.
	CreatorIDValidator func(int) error
	// UpdaterIDValidator is a validator for the "updater_id" field. It is called by the builders before save.
	UpdaterIDValidator func(int) error
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(int) error
)

// OrderOption defines the ordering options for the Department queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByCreatorID orders the results by the creator_id field.
func ByCreatorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatorID, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByUpdaterID orders the results by the updater_id field.
func ByUpdaterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdaterID, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByNotes orders the results by the notes field.
func ByNotes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotes, opts...).ToFunc()
}

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByUpdaterField orders the results by updater field.
func ByUpdaterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpdaterStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByPermissionsCount orders the results by permissions count.
func ByPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionsStep(), opts...)
	}
}

// ByPermissions orders the results by permissions terms.
func ByPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEmployeesCount orders the results by employees count.
func ByEmployeesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmployeesStep(), opts...)
	}
}

// ByEmployees orders the results by employees terms.
func ByEmployees(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmployeesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
	)
}
func newUpdaterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpdaterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PermissionsTable, PermissionsPrimaryKey...),
	)
}
func newEmployeesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmployeesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmployeesTable, EmployeesColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
