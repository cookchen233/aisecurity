// Code generated by ent, DO NOT EDIT.

package eventlog

import (
	"aisecurity/ent/dao/predicate"
	"aisecurity/enums"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EventLog {
	return predicate.EventLog(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldCreateTime, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldCreatorID, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldDeleteTime, v))
}

// UpdaterID applies equality check predicate on the "updater_id" field. It's identical to UpdaterIDEQ.
func UpdaterID(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldUpdateTime, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldDeviceID, v))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldEventID, v))
}

// ActorID applies equality check predicate on the "actor_id" field. It's identical to ActorIDEQ.
func ActorID(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldActorID, v))
}

// Actor2ID applies equality check predicate on the "actor2_id" field. It's identical to Actor2IDEQ.
func Actor2ID(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldActor2ID, v))
}

// LogType applies equality check predicate on the "log_type" field. It's identical to LogTypeEQ.
func LogType(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldEQ(FieldLogType, vc))
}

// Notes applies equality check predicate on the "notes" field. It's identical to NotesEQ.
func Notes(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldNotes, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldLTE(FieldCreateTime, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldCreatorID, vs...))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldNotNull(FieldDeleteTime))
}

// UpdaterIDEQ applies the EQ predicate on the "updater_id" field.
func UpdaterIDEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdaterIDNEQ applies the NEQ predicate on the "updater_id" field.
func UpdaterIDNEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldUpdaterID, v))
}

// UpdaterIDIn applies the In predicate on the "updater_id" field.
func UpdaterIDIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldUpdaterID, vs...))
}

// UpdaterIDNotIn applies the NotIn predicate on the "updater_id" field.
func UpdaterIDNotIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldUpdaterID, vs...))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(sql.FieldLTE(FieldUpdateTime, v))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldDeviceID, vs...))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldEventID, vs...))
}

// ActorIDEQ applies the EQ predicate on the "actor_id" field.
func ActorIDEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldActorID, v))
}

// ActorIDNEQ applies the NEQ predicate on the "actor_id" field.
func ActorIDNEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldActorID, v))
}

// ActorIDIn applies the In predicate on the "actor_id" field.
func ActorIDIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldActorID, vs...))
}

// ActorIDNotIn applies the NotIn predicate on the "actor_id" field.
func ActorIDNotIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldActorID, vs...))
}

// ActorIDIsNil applies the IsNil predicate on the "actor_id" field.
func ActorIDIsNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldIsNull(FieldActorID))
}

// ActorIDNotNil applies the NotNil predicate on the "actor_id" field.
func ActorIDNotNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldNotNull(FieldActorID))
}

// Actor2IDEQ applies the EQ predicate on the "actor2_id" field.
func Actor2IDEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldActor2ID, v))
}

// Actor2IDNEQ applies the NEQ predicate on the "actor2_id" field.
func Actor2IDNEQ(v int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldActor2ID, v))
}

// Actor2IDIn applies the In predicate on the "actor2_id" field.
func Actor2IDIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldActor2ID, vs...))
}

// Actor2IDNotIn applies the NotIn predicate on the "actor2_id" field.
func Actor2IDNotIn(vs ...int) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldActor2ID, vs...))
}

// Actor2IDIsNil applies the IsNil predicate on the "actor2_id" field.
func Actor2IDIsNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldIsNull(FieldActor2ID))
}

// Actor2IDNotNil applies the NotNil predicate on the "actor2_id" field.
func Actor2IDNotNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldNotNull(FieldActor2ID))
}

// LogTypeEQ applies the EQ predicate on the "log_type" field.
func LogTypeEQ(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldEQ(FieldLogType, vc))
}

// LogTypeNEQ applies the NEQ predicate on the "log_type" field.
func LogTypeNEQ(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldNEQ(FieldLogType, vc))
}

// LogTypeIn applies the In predicate on the "log_type" field.
func LogTypeIn(vs ...enums.EventLogType) predicate.EventLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.EventLog(sql.FieldIn(FieldLogType, v...))
}

// LogTypeNotIn applies the NotIn predicate on the "log_type" field.
func LogTypeNotIn(vs ...enums.EventLogType) predicate.EventLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.EventLog(sql.FieldNotIn(FieldLogType, v...))
}

// LogTypeGT applies the GT predicate on the "log_type" field.
func LogTypeGT(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldGT(FieldLogType, vc))
}

// LogTypeGTE applies the GTE predicate on the "log_type" field.
func LogTypeGTE(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldGTE(FieldLogType, vc))
}

// LogTypeLT applies the LT predicate on the "log_type" field.
func LogTypeLT(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldLT(FieldLogType, vc))
}

// LogTypeLTE applies the LTE predicate on the "log_type" field.
func LogTypeLTE(v enums.EventLogType) predicate.EventLog {
	vc := int(v)
	return predicate.EventLog(sql.FieldLTE(FieldLogType, vc))
}

// NotesEQ applies the EQ predicate on the "notes" field.
func NotesEQ(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldEQ(FieldNotes, v))
}

// NotesNEQ applies the NEQ predicate on the "notes" field.
func NotesNEQ(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldNEQ(FieldNotes, v))
}

// NotesIn applies the In predicate on the "notes" field.
func NotesIn(vs ...string) predicate.EventLog {
	return predicate.EventLog(sql.FieldIn(FieldNotes, vs...))
}

// NotesNotIn applies the NotIn predicate on the "notes" field.
func NotesNotIn(vs ...string) predicate.EventLog {
	return predicate.EventLog(sql.FieldNotIn(FieldNotes, vs...))
}

// NotesGT applies the GT predicate on the "notes" field.
func NotesGT(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldGT(FieldNotes, v))
}

// NotesGTE applies the GTE predicate on the "notes" field.
func NotesGTE(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldGTE(FieldNotes, v))
}

// NotesLT applies the LT predicate on the "notes" field.
func NotesLT(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldLT(FieldNotes, v))
}

// NotesLTE applies the LTE predicate on the "notes" field.
func NotesLTE(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldLTE(FieldNotes, v))
}

// NotesContains applies the Contains predicate on the "notes" field.
func NotesContains(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldContains(FieldNotes, v))
}

// NotesHasPrefix applies the HasPrefix predicate on the "notes" field.
func NotesHasPrefix(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldHasPrefix(FieldNotes, v))
}

// NotesHasSuffix applies the HasSuffix predicate on the "notes" field.
func NotesHasSuffix(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldHasSuffix(FieldNotes, v))
}

// NotesIsNil applies the IsNil predicate on the "notes" field.
func NotesIsNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldIsNull(FieldNotes))
}

// NotesNotNil applies the NotNil predicate on the "notes" field.
func NotesNotNil() predicate.EventLog {
	return predicate.EventLog(sql.FieldNotNull(FieldNotes))
}

// NotesEqualFold applies the EqualFold predicate on the "notes" field.
func NotesEqualFold(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldEqualFold(FieldNotes, v))
}

// NotesContainsFold applies the ContainsFold predicate on the "notes" field.
func NotesContainsFold(v string) predicate.EventLog {
	return predicate.EventLog(sql.FieldContainsFold(FieldNotes, v))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.Admin) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdater applies the HasEdge predicate on the "updater" edge.
func HasUpdater() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdaterWith applies the HasEdge predicate on the "updater" edge with a given conditions (other predicates).
func HasUpdaterWith(preds ...predicate.Admin) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := newUpdaterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := newDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActor applies the HasEdge predicate on the "actor" edge.
func HasActor() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActorTable, ActorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActorWith applies the HasEdge predicate on the "actor" edge with a given conditions (other predicates).
func HasActorWith(preds ...predicate.Admin) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := newActorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActor2 applies the HasEdge predicate on the "actor2" edge.
func HasActor2() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, Actor2Table, Actor2Column),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActor2With applies the HasEdge predicate on the "actor2" edge with a given conditions (other predicates).
func HasActor2With(preds ...predicate.Admin) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := newActor2Step()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EventLog) predicate.EventLog {
	return predicate.EventLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EventLog) predicate.EventLog {
	return predicate.EventLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EventLog) predicate.EventLog {
	return predicate.EventLog(sql.NotPredicates(p))
}
