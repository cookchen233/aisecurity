// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/sweep"
	"aisecurity/ent/dao/sweepresult"
	"aisecurity/ent/dao/sweepresultdetails"
	"aisecurity/ent/dao/sweepschedule"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SweepResultDetailsUpdate is the builder for updating SweepResultDetails entities.
type SweepResultDetailsUpdate struct {
	config
	hooks    []Hook
	mutation *SweepResultDetailsMutation
}

// Where appends a list predicates to the SweepResultDetailsUpdate builder.
func (srdu *SweepResultDetailsUpdate) Where(ps ...predicate.SweepResultDetails) *SweepResultDetailsUpdate {
	srdu.mutation.Where(ps...)
	return srdu
}

// SetDeleteTime sets the "delete_time" field.
func (srdu *SweepResultDetailsUpdate) SetDeleteTime(t time.Time) *SweepResultDetailsUpdate {
	srdu.mutation.SetDeleteTime(t)
	return srdu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableDeleteTime(t *time.Time) *SweepResultDetailsUpdate {
	if t != nil {
		srdu.SetDeleteTime(*t)
	}
	return srdu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (srdu *SweepResultDetailsUpdate) ClearDeleteTime() *SweepResultDetailsUpdate {
	srdu.mutation.ClearDeleteTime()
	return srdu
}

// SetUpdaterID sets the "updater_id" field.
func (srdu *SweepResultDetailsUpdate) SetUpdaterID(i int) *SweepResultDetailsUpdate {
	srdu.mutation.SetUpdaterID(i)
	return srdu
}

// SetNillableUpdaterID sets the "updater_id" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableUpdaterID(i *int) *SweepResultDetailsUpdate {
	if i != nil {
		srdu.SetUpdaterID(*i)
	}
	return srdu
}

// SetUpdateTime sets the "update_time" field.
func (srdu *SweepResultDetailsUpdate) SetUpdateTime(t time.Time) *SweepResultDetailsUpdate {
	srdu.mutation.SetUpdateTime(t)
	return srdu
}

// SetSweepID sets the "sweep_id" field.
func (srdu *SweepResultDetailsUpdate) SetSweepID(i int) *SweepResultDetailsUpdate {
	srdu.mutation.SetSweepID(i)
	return srdu
}

// SetNillableSweepID sets the "sweep_id" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableSweepID(i *int) *SweepResultDetailsUpdate {
	if i != nil {
		srdu.SetSweepID(*i)
	}
	return srdu
}

// SetSweepScheduleID sets the "sweep_schedule_id" field.
func (srdu *SweepResultDetailsUpdate) SetSweepScheduleID(i int) *SweepResultDetailsUpdate {
	srdu.mutation.SetSweepScheduleID(i)
	return srdu
}

// SetNillableSweepScheduleID sets the "sweep_schedule_id" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableSweepScheduleID(i *int) *SweepResultDetailsUpdate {
	if i != nil {
		srdu.SetSweepScheduleID(*i)
	}
	return srdu
}

// SetSweepResultID sets the "sweep_result_id" field.
func (srdu *SweepResultDetailsUpdate) SetSweepResultID(i int) *SweepResultDetailsUpdate {
	srdu.mutation.SetSweepResultID(i)
	return srdu
}

// SetNillableSweepResultID sets the "sweep_result_id" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableSweepResultID(i *int) *SweepResultDetailsUpdate {
	if i != nil {
		srdu.SetSweepResultID(*i)
	}
	return srdu
}

// SetTitle sets the "title" field.
func (srdu *SweepResultDetailsUpdate) SetTitle(s string) *SweepResultDetailsUpdate {
	srdu.mutation.SetTitle(s)
	return srdu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableTitle(s *string) *SweepResultDetailsUpdate {
	if s != nil {
		srdu.SetTitle(*s)
	}
	return srdu
}

// SetResult sets the "result" field.
func (srdu *SweepResultDetailsUpdate) SetResult(i int) *SweepResultDetailsUpdate {
	srdu.mutation.ResetResult()
	srdu.mutation.SetResult(i)
	return srdu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (srdu *SweepResultDetailsUpdate) SetNillableResult(i *int) *SweepResultDetailsUpdate {
	if i != nil {
		srdu.SetResult(*i)
	}
	return srdu
}

// AddResult adds i to the "result" field.
func (srdu *SweepResultDetailsUpdate) AddResult(i int) *SweepResultDetailsUpdate {
	srdu.mutation.AddResult(i)
	return srdu
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (srdu *SweepResultDetailsUpdate) SetUpdater(a *Admin) *SweepResultDetailsUpdate {
	return srdu.SetUpdaterID(a.ID)
}

// SetSweep sets the "sweep" edge to the Sweep entity.
func (srdu *SweepResultDetailsUpdate) SetSweep(s *Sweep) *SweepResultDetailsUpdate {
	return srdu.SetSweepID(s.ID)
}

// SetSweepSchedule sets the "sweep_schedule" edge to the SweepSchedule entity.
func (srdu *SweepResultDetailsUpdate) SetSweepSchedule(s *SweepSchedule) *SweepResultDetailsUpdate {
	return srdu.SetSweepScheduleID(s.ID)
}

// SetSweepResult sets the "sweep_result" edge to the SweepResult entity.
func (srdu *SweepResultDetailsUpdate) SetSweepResult(s *SweepResult) *SweepResultDetailsUpdate {
	return srdu.SetSweepResultID(s.ID)
}

// Mutation returns the SweepResultDetailsMutation object of the builder.
func (srdu *SweepResultDetailsUpdate) Mutation() *SweepResultDetailsMutation {
	return srdu.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (srdu *SweepResultDetailsUpdate) ClearUpdater() *SweepResultDetailsUpdate {
	srdu.mutation.ClearUpdater()
	return srdu
}

// ClearSweep clears the "sweep" edge to the Sweep entity.
func (srdu *SweepResultDetailsUpdate) ClearSweep() *SweepResultDetailsUpdate {
	srdu.mutation.ClearSweep()
	return srdu
}

// ClearSweepSchedule clears the "sweep_schedule" edge to the SweepSchedule entity.
func (srdu *SweepResultDetailsUpdate) ClearSweepSchedule() *SweepResultDetailsUpdate {
	srdu.mutation.ClearSweepSchedule()
	return srdu
}

// ClearSweepResult clears the "sweep_result" edge to the SweepResult entity.
func (srdu *SweepResultDetailsUpdate) ClearSweepResult() *SweepResultDetailsUpdate {
	srdu.mutation.ClearSweepResult()
	return srdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (srdu *SweepResultDetailsUpdate) Save(ctx context.Context) (int, error) {
	if err := srdu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, srdu.sqlSave, srdu.mutation, srdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srdu *SweepResultDetailsUpdate) SaveX(ctx context.Context) int {
	affected, err := srdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (srdu *SweepResultDetailsUpdate) Exec(ctx context.Context) error {
	_, err := srdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srdu *SweepResultDetailsUpdate) ExecX(ctx context.Context) {
	if err := srdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srdu *SweepResultDetailsUpdate) defaults() error {
	if _, ok := srdu.mutation.UpdateTime(); !ok {
		if sweepresultdetails.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("dao: uninitialized sweepresultdetails.UpdateDefaultUpdateTime (forgotten import dao/runtime?)")
		}
		v := sweepresultdetails.UpdateDefaultUpdateTime()
		srdu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (srdu *SweepResultDetailsUpdate) check() error {
	if v, ok := srdu.mutation.UpdaterID(); ok {
		if err := sweepresultdetails.UpdaterIDValidator(v); err != nil {
			return &ValidationError{Name: "updater_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.updater_id": %w`, err)}
		}
	}
	if v, ok := srdu.mutation.SweepID(); ok {
		if err := sweepresultdetails.SweepIDValidator(v); err != nil {
			return &ValidationError{Name: "sweep_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.sweep_id": %w`, err)}
		}
	}
	if v, ok := srdu.mutation.SweepScheduleID(); ok {
		if err := sweepresultdetails.SweepScheduleIDValidator(v); err != nil {
			return &ValidationError{Name: "sweep_schedule_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.sweep_schedule_id": %w`, err)}
		}
	}
	if v, ok := srdu.mutation.SweepResultID(); ok {
		if err := sweepresultdetails.SweepResultIDValidator(v); err != nil {
			return &ValidationError{Name: "sweep_result_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.sweep_result_id": %w`, err)}
		}
	}
	if v, ok := srdu.mutation.Title(); ok {
		if err := sweepresultdetails.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.title": %w`, err)}
		}
	}
	if v, ok := srdu.mutation.Result(); ok {
		if err := sweepresultdetails.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.result": %w`, err)}
		}
	}
	if _, ok := srdu.mutation.CreatorID(); srdu.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.creator"`)
	}
	if _, ok := srdu.mutation.UpdaterID(); srdu.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.updater"`)
	}
	if _, ok := srdu.mutation.SweepID(); srdu.mutation.SweepCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.sweep"`)
	}
	if _, ok := srdu.mutation.SweepScheduleID(); srdu.mutation.SweepScheduleCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.sweep_schedule"`)
	}
	if _, ok := srdu.mutation.SweepResultID(); srdu.mutation.SweepResultCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.sweep_result"`)
	}
	return nil
}

func (srdu *SweepResultDetailsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := srdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sweepresultdetails.Table, sweepresultdetails.Columns, sqlgraph.NewFieldSpec(sweepresultdetails.FieldID, field.TypeInt))
	if ps := srdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srdu.mutation.DeleteTime(); ok {
		_spec.SetField(sweepresultdetails.FieldDeleteTime, field.TypeTime, value)
	}
	if srdu.mutation.DeleteTimeCleared() {
		_spec.ClearField(sweepresultdetails.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := srdu.mutation.UpdateTime(); ok {
		_spec.SetField(sweepresultdetails.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := srdu.mutation.Title(); ok {
		_spec.SetField(sweepresultdetails.FieldTitle, field.TypeString, value)
	}
	if value, ok := srdu.mutation.Result(); ok {
		_spec.SetField(sweepresultdetails.FieldResult, field.TypeInt, value)
	}
	if value, ok := srdu.mutation.AddedResult(); ok {
		_spec.AddField(sweepresultdetails.FieldResult, field.TypeInt, value)
	}
	if srdu.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.UpdaterTable,
			Columns: []string{sweepresultdetails.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srdu.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.UpdaterTable,
			Columns: []string{sweepresultdetails.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if srdu.mutation.SweepCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepTable,
			Columns: []string{sweepresultdetails.SweepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweep.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srdu.mutation.SweepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepTable,
			Columns: []string{sweepresultdetails.SweepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweep.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if srdu.mutation.SweepScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepScheduleTable,
			Columns: []string{sweepresultdetails.SweepScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepschedule.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srdu.mutation.SweepScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepScheduleTable,
			Columns: []string{sweepresultdetails.SweepScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepschedule.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if srdu.mutation.SweepResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepResultTable,
			Columns: []string{sweepresultdetails.SweepResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepresult.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srdu.mutation.SweepResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepResultTable,
			Columns: []string{sweepresultdetails.SweepResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, srdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sweepresultdetails.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	srdu.mutation.done = true
	return n, nil
}

// SweepResultDetailsUpdateOne is the builder for updating a single SweepResultDetails entity.
type SweepResultDetailsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SweepResultDetailsMutation
}

// SetDeleteTime sets the "delete_time" field.
func (srduo *SweepResultDetailsUpdateOne) SetDeleteTime(t time.Time) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetDeleteTime(t)
	return srduo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableDeleteTime(t *time.Time) *SweepResultDetailsUpdateOne {
	if t != nil {
		srduo.SetDeleteTime(*t)
	}
	return srduo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (srduo *SweepResultDetailsUpdateOne) ClearDeleteTime() *SweepResultDetailsUpdateOne {
	srduo.mutation.ClearDeleteTime()
	return srduo
}

// SetUpdaterID sets the "updater_id" field.
func (srduo *SweepResultDetailsUpdateOne) SetUpdaterID(i int) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetUpdaterID(i)
	return srduo
}

// SetNillableUpdaterID sets the "updater_id" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableUpdaterID(i *int) *SweepResultDetailsUpdateOne {
	if i != nil {
		srduo.SetUpdaterID(*i)
	}
	return srduo
}

// SetUpdateTime sets the "update_time" field.
func (srduo *SweepResultDetailsUpdateOne) SetUpdateTime(t time.Time) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetUpdateTime(t)
	return srduo
}

// SetSweepID sets the "sweep_id" field.
func (srduo *SweepResultDetailsUpdateOne) SetSweepID(i int) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetSweepID(i)
	return srduo
}

// SetNillableSweepID sets the "sweep_id" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableSweepID(i *int) *SweepResultDetailsUpdateOne {
	if i != nil {
		srduo.SetSweepID(*i)
	}
	return srduo
}

// SetSweepScheduleID sets the "sweep_schedule_id" field.
func (srduo *SweepResultDetailsUpdateOne) SetSweepScheduleID(i int) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetSweepScheduleID(i)
	return srduo
}

// SetNillableSweepScheduleID sets the "sweep_schedule_id" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableSweepScheduleID(i *int) *SweepResultDetailsUpdateOne {
	if i != nil {
		srduo.SetSweepScheduleID(*i)
	}
	return srduo
}

// SetSweepResultID sets the "sweep_result_id" field.
func (srduo *SweepResultDetailsUpdateOne) SetSweepResultID(i int) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetSweepResultID(i)
	return srduo
}

// SetNillableSweepResultID sets the "sweep_result_id" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableSweepResultID(i *int) *SweepResultDetailsUpdateOne {
	if i != nil {
		srduo.SetSweepResultID(*i)
	}
	return srduo
}

// SetTitle sets the "title" field.
func (srduo *SweepResultDetailsUpdateOne) SetTitle(s string) *SweepResultDetailsUpdateOne {
	srduo.mutation.SetTitle(s)
	return srduo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableTitle(s *string) *SweepResultDetailsUpdateOne {
	if s != nil {
		srduo.SetTitle(*s)
	}
	return srduo
}

// SetResult sets the "result" field.
func (srduo *SweepResultDetailsUpdateOne) SetResult(i int) *SweepResultDetailsUpdateOne {
	srduo.mutation.ResetResult()
	srduo.mutation.SetResult(i)
	return srduo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (srduo *SweepResultDetailsUpdateOne) SetNillableResult(i *int) *SweepResultDetailsUpdateOne {
	if i != nil {
		srduo.SetResult(*i)
	}
	return srduo
}

// AddResult adds i to the "result" field.
func (srduo *SweepResultDetailsUpdateOne) AddResult(i int) *SweepResultDetailsUpdateOne {
	srduo.mutation.AddResult(i)
	return srduo
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (srduo *SweepResultDetailsUpdateOne) SetUpdater(a *Admin) *SweepResultDetailsUpdateOne {
	return srduo.SetUpdaterID(a.ID)
}

// SetSweep sets the "sweep" edge to the Sweep entity.
func (srduo *SweepResultDetailsUpdateOne) SetSweep(s *Sweep) *SweepResultDetailsUpdateOne {
	return srduo.SetSweepID(s.ID)
}

// SetSweepSchedule sets the "sweep_schedule" edge to the SweepSchedule entity.
func (srduo *SweepResultDetailsUpdateOne) SetSweepSchedule(s *SweepSchedule) *SweepResultDetailsUpdateOne {
	return srduo.SetSweepScheduleID(s.ID)
}

// SetSweepResult sets the "sweep_result" edge to the SweepResult entity.
func (srduo *SweepResultDetailsUpdateOne) SetSweepResult(s *SweepResult) *SweepResultDetailsUpdateOne {
	return srduo.SetSweepResultID(s.ID)
}

// Mutation returns the SweepResultDetailsMutation object of the builder.
func (srduo *SweepResultDetailsUpdateOne) Mutation() *SweepResultDetailsMutation {
	return srduo.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (srduo *SweepResultDetailsUpdateOne) ClearUpdater() *SweepResultDetailsUpdateOne {
	srduo.mutation.ClearUpdater()
	return srduo
}

// ClearSweep clears the "sweep" edge to the Sweep entity.
func (srduo *SweepResultDetailsUpdateOne) ClearSweep() *SweepResultDetailsUpdateOne {
	srduo.mutation.ClearSweep()
	return srduo
}

// ClearSweepSchedule clears the "sweep_schedule" edge to the SweepSchedule entity.
func (srduo *SweepResultDetailsUpdateOne) ClearSweepSchedule() *SweepResultDetailsUpdateOne {
	srduo.mutation.ClearSweepSchedule()
	return srduo
}

// ClearSweepResult clears the "sweep_result" edge to the SweepResult entity.
func (srduo *SweepResultDetailsUpdateOne) ClearSweepResult() *SweepResultDetailsUpdateOne {
	srduo.mutation.ClearSweepResult()
	return srduo
}

// Where appends a list predicates to the SweepResultDetailsUpdate builder.
func (srduo *SweepResultDetailsUpdateOne) Where(ps ...predicate.SweepResultDetails) *SweepResultDetailsUpdateOne {
	srduo.mutation.Where(ps...)
	return srduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (srduo *SweepResultDetailsUpdateOne) Select(field string, fields ...string) *SweepResultDetailsUpdateOne {
	srduo.fields = append([]string{field}, fields...)
	return srduo
}

// Save executes the query and returns the updated SweepResultDetails entity.
func (srduo *SweepResultDetailsUpdateOne) Save(ctx context.Context) (*SweepResultDetails, error) {
	if err := srduo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, srduo.sqlSave, srduo.mutation, srduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srduo *SweepResultDetailsUpdateOne) SaveX(ctx context.Context) *SweepResultDetails {
	node, err := srduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (srduo *SweepResultDetailsUpdateOne) Exec(ctx context.Context) error {
	_, err := srduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srduo *SweepResultDetailsUpdateOne) ExecX(ctx context.Context) {
	if err := srduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srduo *SweepResultDetailsUpdateOne) defaults() error {
	if _, ok := srduo.mutation.UpdateTime(); !ok {
		if sweepresultdetails.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("dao: uninitialized sweepresultdetails.UpdateDefaultUpdateTime (forgotten import dao/runtime?)")
		}
		v := sweepresultdetails.UpdateDefaultUpdateTime()
		srduo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (srduo *SweepResultDetailsUpdateOne) check() error {
	if v, ok := srduo.mutation.UpdaterID(); ok {
		if err := sweepresultdetails.UpdaterIDValidator(v); err != nil {
			return &ValidationError{Name: "updater_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.updater_id": %w`, err)}
		}
	}
	if v, ok := srduo.mutation.SweepID(); ok {
		if err := sweepresultdetails.SweepIDValidator(v); err != nil {
			return &ValidationError{Name: "sweep_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.sweep_id": %w`, err)}
		}
	}
	if v, ok := srduo.mutation.SweepScheduleID(); ok {
		if err := sweepresultdetails.SweepScheduleIDValidator(v); err != nil {
			return &ValidationError{Name: "sweep_schedule_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.sweep_schedule_id": %w`, err)}
		}
	}
	if v, ok := srduo.mutation.SweepResultID(); ok {
		if err := sweepresultdetails.SweepResultIDValidator(v); err != nil {
			return &ValidationError{Name: "sweep_result_id", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.sweep_result_id": %w`, err)}
		}
	}
	if v, ok := srduo.mutation.Title(); ok {
		if err := sweepresultdetails.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.title": %w`, err)}
		}
	}
	if v, ok := srduo.mutation.Result(); ok {
		if err := sweepresultdetails.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`dao: validator failed for field "SweepResultDetails.result": %w`, err)}
		}
	}
	if _, ok := srduo.mutation.CreatorID(); srduo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.creator"`)
	}
	if _, ok := srduo.mutation.UpdaterID(); srduo.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.updater"`)
	}
	if _, ok := srduo.mutation.SweepID(); srduo.mutation.SweepCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.sweep"`)
	}
	if _, ok := srduo.mutation.SweepScheduleID(); srduo.mutation.SweepScheduleCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.sweep_schedule"`)
	}
	if _, ok := srduo.mutation.SweepResultID(); srduo.mutation.SweepResultCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "SweepResultDetails.sweep_result"`)
	}
	return nil
}

func (srduo *SweepResultDetailsUpdateOne) sqlSave(ctx context.Context) (_node *SweepResultDetails, err error) {
	if err := srduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sweepresultdetails.Table, sweepresultdetails.Columns, sqlgraph.NewFieldSpec(sweepresultdetails.FieldID, field.TypeInt))
	id, ok := srduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "SweepResultDetails.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := srduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sweepresultdetails.FieldID)
		for _, f := range fields {
			if !sweepresultdetails.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != sweepresultdetails.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := srduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srduo.mutation.DeleteTime(); ok {
		_spec.SetField(sweepresultdetails.FieldDeleteTime, field.TypeTime, value)
	}
	if srduo.mutation.DeleteTimeCleared() {
		_spec.ClearField(sweepresultdetails.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := srduo.mutation.UpdateTime(); ok {
		_spec.SetField(sweepresultdetails.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := srduo.mutation.Title(); ok {
		_spec.SetField(sweepresultdetails.FieldTitle, field.TypeString, value)
	}
	if value, ok := srduo.mutation.Result(); ok {
		_spec.SetField(sweepresultdetails.FieldResult, field.TypeInt, value)
	}
	if value, ok := srduo.mutation.AddedResult(); ok {
		_spec.AddField(sweepresultdetails.FieldResult, field.TypeInt, value)
	}
	if srduo.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.UpdaterTable,
			Columns: []string{sweepresultdetails.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srduo.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.UpdaterTable,
			Columns: []string{sweepresultdetails.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if srduo.mutation.SweepCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepTable,
			Columns: []string{sweepresultdetails.SweepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweep.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srduo.mutation.SweepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepTable,
			Columns: []string{sweepresultdetails.SweepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweep.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if srduo.mutation.SweepScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepScheduleTable,
			Columns: []string{sweepresultdetails.SweepScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepschedule.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srduo.mutation.SweepScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepScheduleTable,
			Columns: []string{sweepresultdetails.SweepScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepschedule.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if srduo.mutation.SweepResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepResultTable,
			Columns: []string{sweepresultdetails.SweepResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepresult.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := srduo.mutation.SweepResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sweepresultdetails.SweepResultTable,
			Columns: []string{sweepresultdetails.SweepResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sweepresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SweepResultDetails{config: srduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, srduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sweepresultdetails.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	srduo.mutation.done = true
	return _node, nil
}