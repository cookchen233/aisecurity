// Code generated by ent, DO NOT EDIT.

package sweepschedule

import (
	"aisecurity/enums"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sweepschedule type in the database.
	Label = "sweep_schedule"
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
	// FieldSweepID holds the string denoting the sweep_id field in the database.
	FieldSweepID = "sweep_id"
	// FieldEnabledStatus holds the string denoting the enabled_status field in the database.
	FieldEnabledStatus = "enabled_status"
	// FieldActionTime holds the string denoting the action_time field in the database.
	FieldActionTime = "action_time"
	// FieldRemind holds the string denoting the remind field in the database.
	FieldRemind = "remind"
	// FieldRepeat holds the string denoting the repeat field in the database.
	FieldRepeat = "repeat"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeUpdater holds the string denoting the updater edge name in mutations.
	EdgeUpdater = "updater"
	// EdgeSweep holds the string denoting the sweep edge name in mutations.
	EdgeSweep = "sweep"
	// EdgeWorkers holds the string denoting the workers edge name in mutations.
	EdgeWorkers = "workers"
	// EdgeSweepResult holds the string denoting the sweep_result edge name in mutations.
	EdgeSweepResult = "sweep_result"
	// EdgeSweepResultDetails holds the string denoting the sweep_result_details edge name in mutations.
	EdgeSweepResultDetails = "sweep_result_details"
	// Table holds the table name of the sweepschedule in the database.
	Table = "sweep_schedules"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "sweep_schedules"
	// CreatorInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	CreatorInverseTable = "admins"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "creator_id"
	// UpdaterTable is the table that holds the updater relation/edge.
	UpdaterTable = "sweep_schedules"
	// UpdaterInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	UpdaterInverseTable = "admins"
	// UpdaterColumn is the table column denoting the updater relation/edge.
	UpdaterColumn = "updater_id"
	// SweepTable is the table that holds the sweep relation/edge.
	SweepTable = "sweep_schedules"
	// SweepInverseTable is the table name for the Sweep entity.
	// It exists in this package in order to avoid circular dependency with the "sweep" package.
	SweepInverseTable = "sweeps"
	// SweepColumn is the table column denoting the sweep relation/edge.
	SweepColumn = "sweep_id"
	// WorkersTable is the table that holds the workers relation/edge. The primary key declared below.
	WorkersTable = "admin_sweep_schedule"
	// WorkersInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	WorkersInverseTable = "admins"
	// SweepResultTable is the table that holds the sweep_result relation/edge.
	SweepResultTable = "sweep_results"
	// SweepResultInverseTable is the table name for the SweepResult entity.
	// It exists in this package in order to avoid circular dependency with the "sweepresult" package.
	SweepResultInverseTable = "sweep_results"
	// SweepResultColumn is the table column denoting the sweep_result relation/edge.
	SweepResultColumn = "sweep_schedule_id"
	// SweepResultDetailsTable is the table that holds the sweep_result_details relation/edge.
	SweepResultDetailsTable = "sweep_result_details"
	// SweepResultDetailsInverseTable is the table name for the SweepResultDetails entity.
	// It exists in this package in order to avoid circular dependency with the "sweepresultdetails" package.
	SweepResultDetailsInverseTable = "sweep_result_details"
	// SweepResultDetailsColumn is the table column denoting the sweep_result_details relation/edge.
	SweepResultDetailsColumn = "sweep_schedule_id"
)

// Columns holds all SQL columns for sweepschedule fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldCreatorID,
	FieldDeleteTime,
	FieldUpdaterID,
	FieldUpdateTime,
	FieldName,
	FieldSweepID,
	FieldEnabledStatus,
	FieldActionTime,
	FieldRemind,
	FieldRepeat,
}

var (
	// WorkersPrimaryKey and WorkersColumn2 are the table columns denoting the
	// primary key for the workers relation (M2M).
	WorkersPrimaryKey = []string{"admin_id", "sweep_schedule_id"}
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
	// SweepIDValidator is a validator for the "sweep_id" field. It is called by the builders before save.
	SweepIDValidator func(int) error
	// DefaultEnabledStatus holds the default value on creation for the "enabled_status" field.
	DefaultEnabledStatus enums.EnabledStatus
	// EnabledStatusValidator is a validator for the "enabled_status" field. It is called by the builders before save.
	EnabledStatusValidator func(int) error
)

// OrderOption defines the ordering options for the SweepSchedule queries.
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

// BySweepID orders the results by the sweep_id field.
func BySweepID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSweepID, opts...).ToFunc()
}

// ByEnabledStatus orders the results by the enabled_status field.
func ByEnabledStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnabledStatus, opts...).ToFunc()
}

// ByActionTime orders the results by the action_time field.
func ByActionTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionTime, opts...).ToFunc()
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

// BySweepField orders the results by sweep field.
func BySweepField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSweepStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkersCount orders the results by workers count.
func ByWorkersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkersStep(), opts...)
	}
}

// ByWorkers orders the results by workers terms.
func ByWorkers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySweepResultCount orders the results by sweep_result count.
func BySweepResultCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSweepResultStep(), opts...)
	}
}

// BySweepResult orders the results by sweep_result terms.
func BySweepResult(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSweepResultStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySweepResultDetailsCount orders the results by sweep_result_details count.
func BySweepResultDetailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSweepResultDetailsStep(), opts...)
	}
}

// BySweepResultDetails orders the results by sweep_result_details terms.
func BySweepResultDetails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSweepResultDetailsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newSweepStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SweepInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SweepTable, SweepColumn),
	)
}
func newWorkersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WorkersTable, WorkersPrimaryKey...),
	)
}
func newSweepResultStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SweepResultInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SweepResultTable, SweepResultColumn),
	)
}
func newSweepResultDetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SweepResultDetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SweepResultDetailsTable, SweepResultDetailsColumn),
	)
}
