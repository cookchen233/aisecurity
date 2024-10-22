// Code generated by ent, DO NOT EDIT.

package fixing

import (
	"aisecurity/ent/dao/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCreateTime, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCreatorID, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldDeleteTime, v))
}

// UpdaterID applies equality check predicate on the "updater_id" field. It's identical to UpdaterIDEQ.
func UpdaterID(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldUpdateTime, v))
}

// FixerID applies equality check predicate on the "fixer_id" field. It's identical to FixerIDEQ.
func FixerID(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldFixerID, v))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldEventID, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldDeviceID, v))
}

// AssignNotes applies equality check predicate on the "assign_notes" field. It's identical to AssignNotesEQ.
func AssignNotes(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldAssignNotes, v))
}

// FixTime applies equality check predicate on the "fix_time" field. It's identical to FixTimeEQ.
func FixTime(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldFixTime, v))
}

// CompleteTime applies equality check predicate on the "complete_time" field. It's identical to CompleteTimeEQ.
func CompleteTime(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCompleteTime, v))
}

// CompleteNotes applies equality check predicate on the "complete_notes" field. It's identical to CompleteNotesEQ.
func CompleteNotes(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCompleteNotes, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldCreateTime, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldCreatorID, vs...))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldNotNull(FieldDeleteTime))
}

// UpdaterIDEQ applies the EQ predicate on the "updater_id" field.
func UpdaterIDEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdaterIDNEQ applies the NEQ predicate on the "updater_id" field.
func UpdaterIDNEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldUpdaterID, v))
}

// UpdaterIDIn applies the In predicate on the "updater_id" field.
func UpdaterIDIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldUpdaterID, vs...))
}

// UpdaterIDNotIn applies the NotIn predicate on the "updater_id" field.
func UpdaterIDNotIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldUpdaterID, vs...))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldUpdateTime, v))
}

// FixerIDEQ applies the EQ predicate on the "fixer_id" field.
func FixerIDEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldFixerID, v))
}

// FixerIDNEQ applies the NEQ predicate on the "fixer_id" field.
func FixerIDNEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldFixerID, v))
}

// FixerIDIn applies the In predicate on the "fixer_id" field.
func FixerIDIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldFixerID, vs...))
}

// FixerIDNotIn applies the NotIn predicate on the "fixer_id" field.
func FixerIDNotIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldFixerID, vs...))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldEventID, vs...))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...int) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldDeviceID, vs...))
}

// AssignNotesEQ applies the EQ predicate on the "assign_notes" field.
func AssignNotesEQ(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldAssignNotes, v))
}

// AssignNotesNEQ applies the NEQ predicate on the "assign_notes" field.
func AssignNotesNEQ(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldAssignNotes, v))
}

// AssignNotesIn applies the In predicate on the "assign_notes" field.
func AssignNotesIn(vs ...string) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldAssignNotes, vs...))
}

// AssignNotesNotIn applies the NotIn predicate on the "assign_notes" field.
func AssignNotesNotIn(vs ...string) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldAssignNotes, vs...))
}

// AssignNotesGT applies the GT predicate on the "assign_notes" field.
func AssignNotesGT(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldAssignNotes, v))
}

// AssignNotesGTE applies the GTE predicate on the "assign_notes" field.
func AssignNotesGTE(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldAssignNotes, v))
}

// AssignNotesLT applies the LT predicate on the "assign_notes" field.
func AssignNotesLT(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldAssignNotes, v))
}

// AssignNotesLTE applies the LTE predicate on the "assign_notes" field.
func AssignNotesLTE(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldAssignNotes, v))
}

// AssignNotesContains applies the Contains predicate on the "assign_notes" field.
func AssignNotesContains(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldContains(FieldAssignNotes, v))
}

// AssignNotesHasPrefix applies the HasPrefix predicate on the "assign_notes" field.
func AssignNotesHasPrefix(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldHasPrefix(FieldAssignNotes, v))
}

// AssignNotesHasSuffix applies the HasSuffix predicate on the "assign_notes" field.
func AssignNotesHasSuffix(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldHasSuffix(FieldAssignNotes, v))
}

// AssignNotesIsNil applies the IsNil predicate on the "assign_notes" field.
func AssignNotesIsNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldIsNull(FieldAssignNotes))
}

// AssignNotesNotNil applies the NotNil predicate on the "assign_notes" field.
func AssignNotesNotNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldNotNull(FieldAssignNotes))
}

// AssignNotesEqualFold applies the EqualFold predicate on the "assign_notes" field.
func AssignNotesEqualFold(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldEqualFold(FieldAssignNotes, v))
}

// AssignNotesContainsFold applies the ContainsFold predicate on the "assign_notes" field.
func AssignNotesContainsFold(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldContainsFold(FieldAssignNotes, v))
}

// FixTimeEQ applies the EQ predicate on the "fix_time" field.
func FixTimeEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldFixTime, v))
}

// FixTimeNEQ applies the NEQ predicate on the "fix_time" field.
func FixTimeNEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldFixTime, v))
}

// FixTimeIn applies the In predicate on the "fix_time" field.
func FixTimeIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldFixTime, vs...))
}

// FixTimeNotIn applies the NotIn predicate on the "fix_time" field.
func FixTimeNotIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldFixTime, vs...))
}

// FixTimeGT applies the GT predicate on the "fix_time" field.
func FixTimeGT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldFixTime, v))
}

// FixTimeGTE applies the GTE predicate on the "fix_time" field.
func FixTimeGTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldFixTime, v))
}

// FixTimeLT applies the LT predicate on the "fix_time" field.
func FixTimeLT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldFixTime, v))
}

// FixTimeLTE applies the LTE predicate on the "fix_time" field.
func FixTimeLTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldFixTime, v))
}

// FixTimeIsNil applies the IsNil predicate on the "fix_time" field.
func FixTimeIsNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldIsNull(FieldFixTime))
}

// FixTimeNotNil applies the NotNil predicate on the "fix_time" field.
func FixTimeNotNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldNotNull(FieldFixTime))
}

// CompleteTimeEQ applies the EQ predicate on the "complete_time" field.
func CompleteTimeEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCompleteTime, v))
}

// CompleteTimeNEQ applies the NEQ predicate on the "complete_time" field.
func CompleteTimeNEQ(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldCompleteTime, v))
}

// CompleteTimeIn applies the In predicate on the "complete_time" field.
func CompleteTimeIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldCompleteTime, vs...))
}

// CompleteTimeNotIn applies the NotIn predicate on the "complete_time" field.
func CompleteTimeNotIn(vs ...time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldCompleteTime, vs...))
}

// CompleteTimeGT applies the GT predicate on the "complete_time" field.
func CompleteTimeGT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldCompleteTime, v))
}

// CompleteTimeGTE applies the GTE predicate on the "complete_time" field.
func CompleteTimeGTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldCompleteTime, v))
}

// CompleteTimeLT applies the LT predicate on the "complete_time" field.
func CompleteTimeLT(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldCompleteTime, v))
}

// CompleteTimeLTE applies the LTE predicate on the "complete_time" field.
func CompleteTimeLTE(v time.Time) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldCompleteTime, v))
}

// CompleteTimeIsNil applies the IsNil predicate on the "complete_time" field.
func CompleteTimeIsNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldIsNull(FieldCompleteTime))
}

// CompleteTimeNotNil applies the NotNil predicate on the "complete_time" field.
func CompleteTimeNotNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldNotNull(FieldCompleteTime))
}

// CompleteNotesEQ applies the EQ predicate on the "complete_notes" field.
func CompleteNotesEQ(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldEQ(FieldCompleteNotes, v))
}

// CompleteNotesNEQ applies the NEQ predicate on the "complete_notes" field.
func CompleteNotesNEQ(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldNEQ(FieldCompleteNotes, v))
}

// CompleteNotesIn applies the In predicate on the "complete_notes" field.
func CompleteNotesIn(vs ...string) predicate.Fixing {
	return predicate.Fixing(sql.FieldIn(FieldCompleteNotes, vs...))
}

// CompleteNotesNotIn applies the NotIn predicate on the "complete_notes" field.
func CompleteNotesNotIn(vs ...string) predicate.Fixing {
	return predicate.Fixing(sql.FieldNotIn(FieldCompleteNotes, vs...))
}

// CompleteNotesGT applies the GT predicate on the "complete_notes" field.
func CompleteNotesGT(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldGT(FieldCompleteNotes, v))
}

// CompleteNotesGTE applies the GTE predicate on the "complete_notes" field.
func CompleteNotesGTE(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldGTE(FieldCompleteNotes, v))
}

// CompleteNotesLT applies the LT predicate on the "complete_notes" field.
func CompleteNotesLT(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldLT(FieldCompleteNotes, v))
}

// CompleteNotesLTE applies the LTE predicate on the "complete_notes" field.
func CompleteNotesLTE(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldLTE(FieldCompleteNotes, v))
}

// CompleteNotesContains applies the Contains predicate on the "complete_notes" field.
func CompleteNotesContains(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldContains(FieldCompleteNotes, v))
}

// CompleteNotesHasPrefix applies the HasPrefix predicate on the "complete_notes" field.
func CompleteNotesHasPrefix(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldHasPrefix(FieldCompleteNotes, v))
}

// CompleteNotesHasSuffix applies the HasSuffix predicate on the "complete_notes" field.
func CompleteNotesHasSuffix(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldHasSuffix(FieldCompleteNotes, v))
}

// CompleteNotesIsNil applies the IsNil predicate on the "complete_notes" field.
func CompleteNotesIsNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldIsNull(FieldCompleteNotes))
}

// CompleteNotesNotNil applies the NotNil predicate on the "complete_notes" field.
func CompleteNotesNotNil() predicate.Fixing {
	return predicate.Fixing(sql.FieldNotNull(FieldCompleteNotes))
}

// CompleteNotesEqualFold applies the EqualFold predicate on the "complete_notes" field.
func CompleteNotesEqualFold(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldEqualFold(FieldCompleteNotes, v))
}

// CompleteNotesContainsFold applies the ContainsFold predicate on the "complete_notes" field.
func CompleteNotesContainsFold(v string) predicate.Fixing {
	return predicate.Fixing(sql.FieldContainsFold(FieldCompleteNotes, v))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.Admin) predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdater applies the HasEdge predicate on the "updater" edge.
func HasUpdater() predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdaterWith applies the HasEdge predicate on the "updater" edge with a given conditions (other predicates).
func HasUpdaterWith(preds ...predicate.Admin) predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := newUpdaterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFixer applies the HasEdge predicate on the "fixer" edge.
func HasFixer() predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixerTable, FixerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixerWith applies the HasEdge predicate on the "fixer" edge with a given conditions (other predicates).
func HasFixerWith(preds ...predicate.Admin) predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := newFixerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.Fixing {
	return predicate.Fixing(func(s *sql.Selector) {
		step := newDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Fixing) predicate.Fixing {
	return predicate.Fixing(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Fixing) predicate.Fixing {
	return predicate.Fixing(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Fixing) predicate.Fixing {
	return predicate.Fixing(sql.NotPredicates(p))
}
