// Code generated by ent, DO NOT EDIT.

package sweepresultdetails

import (
	"aisecurity/ent/dao/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldCreateTime, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldCreatorID, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldDeleteTime, v))
}

// UpdaterID applies equality check predicate on the "updater_id" field. It's identical to UpdaterIDEQ.
func UpdaterID(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldUpdateTime, v))
}

// SweepID applies equality check predicate on the "sweep_id" field. It's identical to SweepIDEQ.
func SweepID(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldSweepID, v))
}

// SweepScheduleID applies equality check predicate on the "sweep_schedule_id" field. It's identical to SweepScheduleIDEQ.
func SweepScheduleID(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldSweepScheduleID, v))
}

// SweepResultID applies equality check predicate on the "sweep_result_id" field. It's identical to SweepResultIDEQ.
func SweepResultID(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldSweepResultID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldTitle, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldResult, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLTE(FieldCreateTime, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldCreatorID, vs...))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotNull(FieldDeleteTime))
}

// UpdaterIDEQ applies the EQ predicate on the "updater_id" field.
func UpdaterIDEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdaterIDNEQ applies the NEQ predicate on the "updater_id" field.
func UpdaterIDNEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldUpdaterID, v))
}

// UpdaterIDIn applies the In predicate on the "updater_id" field.
func UpdaterIDIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldUpdaterID, vs...))
}

// UpdaterIDNotIn applies the NotIn predicate on the "updater_id" field.
func UpdaterIDNotIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldUpdaterID, vs...))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLTE(FieldUpdateTime, v))
}

// SweepIDEQ applies the EQ predicate on the "sweep_id" field.
func SweepIDEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldSweepID, v))
}

// SweepIDNEQ applies the NEQ predicate on the "sweep_id" field.
func SweepIDNEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldSweepID, v))
}

// SweepIDIn applies the In predicate on the "sweep_id" field.
func SweepIDIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldSweepID, vs...))
}

// SweepIDNotIn applies the NotIn predicate on the "sweep_id" field.
func SweepIDNotIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldSweepID, vs...))
}

// SweepScheduleIDEQ applies the EQ predicate on the "sweep_schedule_id" field.
func SweepScheduleIDEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldSweepScheduleID, v))
}

// SweepScheduleIDNEQ applies the NEQ predicate on the "sweep_schedule_id" field.
func SweepScheduleIDNEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldSweepScheduleID, v))
}

// SweepScheduleIDIn applies the In predicate on the "sweep_schedule_id" field.
func SweepScheduleIDIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldSweepScheduleID, vs...))
}

// SweepScheduleIDNotIn applies the NotIn predicate on the "sweep_schedule_id" field.
func SweepScheduleIDNotIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldSweepScheduleID, vs...))
}

// SweepResultIDEQ applies the EQ predicate on the "sweep_result_id" field.
func SweepResultIDEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldSweepResultID, v))
}

// SweepResultIDNEQ applies the NEQ predicate on the "sweep_result_id" field.
func SweepResultIDNEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldSweepResultID, v))
}

// SweepResultIDIn applies the In predicate on the "sweep_result_id" field.
func SweepResultIDIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldSweepResultID, vs...))
}

// SweepResultIDNotIn applies the NotIn predicate on the "sweep_result_id" field.
func SweepResultIDNotIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldSweepResultID, vs...))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldContainsFold(FieldTitle, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldEQ(FieldResult, v))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNEQ(FieldResult, v))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldIn(FieldResult, vs...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldNotIn(FieldResult, vs...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGT(FieldResult, v))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldGTE(FieldResult, v))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLT(FieldResult, v))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v int) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.FieldLTE(FieldResult, v))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.Admin) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdater applies the HasEdge predicate on the "updater" edge.
func HasUpdater() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdaterWith applies the HasEdge predicate on the "updater" edge with a given conditions (other predicates).
func HasUpdaterWith(preds ...predicate.Admin) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := newUpdaterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSweep applies the HasEdge predicate on the "sweep" edge.
func HasSweep() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SweepTable, SweepColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSweepWith applies the HasEdge predicate on the "sweep" edge with a given conditions (other predicates).
func HasSweepWith(preds ...predicate.Sweep) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := newSweepStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSweepSchedule applies the HasEdge predicate on the "sweep_schedule" edge.
func HasSweepSchedule() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SweepScheduleTable, SweepScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSweepScheduleWith applies the HasEdge predicate on the "sweep_schedule" edge with a given conditions (other predicates).
func HasSweepScheduleWith(preds ...predicate.SweepSchedule) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := newSweepScheduleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSweepResult applies the HasEdge predicate on the "sweep_result" edge.
func HasSweepResult() predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SweepResultTable, SweepResultColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSweepResultWith applies the HasEdge predicate on the "sweep_result" edge with a given conditions (other predicates).
func HasSweepResultWith(preds ...predicate.SweepResult) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(func(s *sql.Selector) {
		step := newSweepResultStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SweepResultDetails) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SweepResultDetails) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SweepResultDetails) predicate.SweepResultDetails {
	return predicate.SweepResultDetails(sql.NotPredicates(p))
}