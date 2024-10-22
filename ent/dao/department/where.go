// Code generated by ent, DO NOT EDIT.

package department

import (
	"aisecurity/ent/dao/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreateTime, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreatorID, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDeleteTime, v))
}

// UpdaterID applies equality check predicate on the "updater_id" field. It's identical to UpdaterIDEQ.
func UpdaterID(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// Notes applies equality check predicate on the "notes" field. It's identical to NotesEQ.
func Notes(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldNotes, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldCreateTime, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldCreatorID, vs...))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldDeleteTime))
}

// UpdaterIDEQ applies the EQ predicate on the "updater_id" field.
func UpdaterIDEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdaterIDNEQ applies the NEQ predicate on the "updater_id" field.
func UpdaterIDNEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldUpdaterID, v))
}

// UpdaterIDIn applies the In predicate on the "updater_id" field.
func UpdaterIDIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldUpdaterID, vs...))
}

// UpdaterIDNotIn applies the NotIn predicate on the "updater_id" field.
func UpdaterIDNotIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldUpdaterID, vs...))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldName, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldParentID))
}

// NotesEQ applies the EQ predicate on the "notes" field.
func NotesEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldNotes, v))
}

// NotesNEQ applies the NEQ predicate on the "notes" field.
func NotesNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldNotes, v))
}

// NotesIn applies the In predicate on the "notes" field.
func NotesIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldNotes, vs...))
}

// NotesNotIn applies the NotIn predicate on the "notes" field.
func NotesNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldNotes, vs...))
}

// NotesGT applies the GT predicate on the "notes" field.
func NotesGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldNotes, v))
}

// NotesGTE applies the GTE predicate on the "notes" field.
func NotesGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldNotes, v))
}

// NotesLT applies the LT predicate on the "notes" field.
func NotesLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldNotes, v))
}

// NotesLTE applies the LTE predicate on the "notes" field.
func NotesLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldNotes, v))
}

// NotesContains applies the Contains predicate on the "notes" field.
func NotesContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldNotes, v))
}

// NotesHasPrefix applies the HasPrefix predicate on the "notes" field.
func NotesHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldNotes, v))
}

// NotesHasSuffix applies the HasSuffix predicate on the "notes" field.
func NotesHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldNotes, v))
}

// NotesIsNil applies the IsNil predicate on the "notes" field.
func NotesIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldNotes))
}

// NotesNotNil applies the NotNil predicate on the "notes" field.
func NotesNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldNotes))
}

// NotesEqualFold applies the EqualFold predicate on the "notes" field.
func NotesEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldNotes, v))
}

// NotesContainsFold applies the ContainsFold predicate on the "notes" field.
func NotesContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldNotes, v))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.Admin) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdater applies the HasEdge predicate on the "updater" edge.
func HasUpdater() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdaterWith applies the HasEdge predicate on the "updater" edge with a given conditions (other predicates).
func HasUpdaterWith(preds ...predicate.Admin) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newUpdaterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployees applies the HasEdge predicate on the "employees" edge.
func HasEmployees() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmployeesTable, EmployeesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeesWith applies the HasEdge predicate on the "employees" edge with a given conditions (other predicates).
func HasEmployeesWith(preds ...predicate.Employee) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newEmployeesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(sql.NotPredicates(p))
}
