// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/risk"
	"aisecurity/ent/dao/risklocation"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RiskLocationUpdate is the builder for updating RiskLocation entities.
type RiskLocationUpdate struct {
	config
	hooks    []Hook
	mutation *RiskLocationMutation
}

// Where appends a list predicates to the RiskLocationUpdate builder.
func (rlu *RiskLocationUpdate) Where(ps ...predicate.RiskLocation) *RiskLocationUpdate {
	rlu.mutation.Where(ps...)
	return rlu
}

// SetDeletedAt sets the "deleted_at" field.
func (rlu *RiskLocationUpdate) SetDeletedAt(t time.Time) *RiskLocationUpdate {
	rlu.mutation.SetDeletedAt(t)
	return rlu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rlu *RiskLocationUpdate) SetNillableDeletedAt(t *time.Time) *RiskLocationUpdate {
	if t != nil {
		rlu.SetDeletedAt(*t)
	}
	return rlu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rlu *RiskLocationUpdate) ClearDeletedAt() *RiskLocationUpdate {
	rlu.mutation.ClearDeletedAt()
	return rlu
}

// SetUpdatedBy sets the "updated_by" field.
func (rlu *RiskLocationUpdate) SetUpdatedBy(i int) *RiskLocationUpdate {
	rlu.mutation.ResetUpdatedBy()
	rlu.mutation.SetUpdatedBy(i)
	return rlu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rlu *RiskLocationUpdate) SetNillableUpdatedBy(i *int) *RiskLocationUpdate {
	if i != nil {
		rlu.SetUpdatedBy(*i)
	}
	return rlu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (rlu *RiskLocationUpdate) AddUpdatedBy(i int) *RiskLocationUpdate {
	rlu.mutation.AddUpdatedBy(i)
	return rlu
}

// SetUpdatedAt sets the "updated_at" field.
func (rlu *RiskLocationUpdate) SetUpdatedAt(t time.Time) *RiskLocationUpdate {
	rlu.mutation.SetUpdatedAt(t)
	return rlu
}

// SetTitle sets the "title" field.
func (rlu *RiskLocationUpdate) SetTitle(s string) *RiskLocationUpdate {
	rlu.mutation.SetTitle(s)
	return rlu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (rlu *RiskLocationUpdate) SetNillableTitle(s *string) *RiskLocationUpdate {
	if s != nil {
		rlu.SetTitle(*s)
	}
	return rlu
}

// AddRiskLocationIDs adds the "risk_location" edge to the Risk entity by IDs.
func (rlu *RiskLocationUpdate) AddRiskLocationIDs(ids ...int) *RiskLocationUpdate {
	rlu.mutation.AddRiskLocationIDs(ids...)
	return rlu
}

// AddRiskLocation adds the "risk_location" edges to the Risk entity.
func (rlu *RiskLocationUpdate) AddRiskLocation(r ...*Risk) *RiskLocationUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rlu.AddRiskLocationIDs(ids...)
}

// Mutation returns the RiskLocationMutation object of the builder.
func (rlu *RiskLocationUpdate) Mutation() *RiskLocationMutation {
	return rlu.mutation
}

// ClearRiskLocation clears all "risk_location" edges to the Risk entity.
func (rlu *RiskLocationUpdate) ClearRiskLocation() *RiskLocationUpdate {
	rlu.mutation.ClearRiskLocation()
	return rlu
}

// RemoveRiskLocationIDs removes the "risk_location" edge to Risk entities by IDs.
func (rlu *RiskLocationUpdate) RemoveRiskLocationIDs(ids ...int) *RiskLocationUpdate {
	rlu.mutation.RemoveRiskLocationIDs(ids...)
	return rlu
}

// RemoveRiskLocation removes "risk_location" edges to Risk entities.
func (rlu *RiskLocationUpdate) RemoveRiskLocation(r ...*Risk) *RiskLocationUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rlu.RemoveRiskLocationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rlu *RiskLocationUpdate) Save(ctx context.Context) (int, error) {
	if err := rlu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, rlu.sqlSave, rlu.mutation, rlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rlu *RiskLocationUpdate) SaveX(ctx context.Context) int {
	affected, err := rlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rlu *RiskLocationUpdate) Exec(ctx context.Context) error {
	_, err := rlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlu *RiskLocationUpdate) ExecX(ctx context.Context) {
	if err := rlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rlu *RiskLocationUpdate) defaults() error {
	if _, ok := rlu.mutation.UpdatedAt(); !ok {
		if risklocation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized risklocation.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := risklocation.UpdateDefaultUpdatedAt()
		rlu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rlu *RiskLocationUpdate) check() error {
	if v, ok := rlu.mutation.UpdatedBy(); ok {
		if err := risklocation.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "RiskLocation.updated_by": %w`, err)}
		}
	}
	if v, ok := rlu.mutation.Title(); ok {
		if err := risklocation.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "RiskLocation.title": %w`, err)}
		}
	}
	if _, ok := rlu.mutation.CreatorID(); rlu.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "RiskLocation.creator"`)
	}
	return nil
}

func (rlu *RiskLocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rlu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(risklocation.Table, risklocation.Columns, sqlgraph.NewFieldSpec(risklocation.FieldID, field.TypeInt))
	if ps := rlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rlu.mutation.DeletedAt(); ok {
		_spec.SetField(risklocation.FieldDeletedAt, field.TypeTime, value)
	}
	if rlu.mutation.DeletedAtCleared() {
		_spec.ClearField(risklocation.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rlu.mutation.UpdatedBy(); ok {
		_spec.SetField(risklocation.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rlu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(risklocation.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rlu.mutation.UpdatedAt(); ok {
		_spec.SetField(risklocation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rlu.mutation.Title(); ok {
		_spec.SetField(risklocation.FieldTitle, field.TypeString, value)
	}
	if rlu.mutation.RiskLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   risklocation.RiskLocationTable,
			Columns: []string{risklocation.RiskLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rlu.mutation.RemovedRiskLocationIDs(); len(nodes) > 0 && !rlu.mutation.RiskLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   risklocation.RiskLocationTable,
			Columns: []string{risklocation.RiskLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rlu.mutation.RiskLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   risklocation.RiskLocationTable,
			Columns: []string{risklocation.RiskLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{risklocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rlu.mutation.done = true
	return n, nil
}

// RiskLocationUpdateOne is the builder for updating a single RiskLocation entity.
type RiskLocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RiskLocationMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (rluo *RiskLocationUpdateOne) SetDeletedAt(t time.Time) *RiskLocationUpdateOne {
	rluo.mutation.SetDeletedAt(t)
	return rluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rluo *RiskLocationUpdateOne) SetNillableDeletedAt(t *time.Time) *RiskLocationUpdateOne {
	if t != nil {
		rluo.SetDeletedAt(*t)
	}
	return rluo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rluo *RiskLocationUpdateOne) ClearDeletedAt() *RiskLocationUpdateOne {
	rluo.mutation.ClearDeletedAt()
	return rluo
}

// SetUpdatedBy sets the "updated_by" field.
func (rluo *RiskLocationUpdateOne) SetUpdatedBy(i int) *RiskLocationUpdateOne {
	rluo.mutation.ResetUpdatedBy()
	rluo.mutation.SetUpdatedBy(i)
	return rluo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rluo *RiskLocationUpdateOne) SetNillableUpdatedBy(i *int) *RiskLocationUpdateOne {
	if i != nil {
		rluo.SetUpdatedBy(*i)
	}
	return rluo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (rluo *RiskLocationUpdateOne) AddUpdatedBy(i int) *RiskLocationUpdateOne {
	rluo.mutation.AddUpdatedBy(i)
	return rluo
}

// SetUpdatedAt sets the "updated_at" field.
func (rluo *RiskLocationUpdateOne) SetUpdatedAt(t time.Time) *RiskLocationUpdateOne {
	rluo.mutation.SetUpdatedAt(t)
	return rluo
}

// SetTitle sets the "title" field.
func (rluo *RiskLocationUpdateOne) SetTitle(s string) *RiskLocationUpdateOne {
	rluo.mutation.SetTitle(s)
	return rluo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (rluo *RiskLocationUpdateOne) SetNillableTitle(s *string) *RiskLocationUpdateOne {
	if s != nil {
		rluo.SetTitle(*s)
	}
	return rluo
}

// AddRiskLocationIDs adds the "risk_location" edge to the Risk entity by IDs.
func (rluo *RiskLocationUpdateOne) AddRiskLocationIDs(ids ...int) *RiskLocationUpdateOne {
	rluo.mutation.AddRiskLocationIDs(ids...)
	return rluo
}

// AddRiskLocation adds the "risk_location" edges to the Risk entity.
func (rluo *RiskLocationUpdateOne) AddRiskLocation(r ...*Risk) *RiskLocationUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rluo.AddRiskLocationIDs(ids...)
}

// Mutation returns the RiskLocationMutation object of the builder.
func (rluo *RiskLocationUpdateOne) Mutation() *RiskLocationMutation {
	return rluo.mutation
}

// ClearRiskLocation clears all "risk_location" edges to the Risk entity.
func (rluo *RiskLocationUpdateOne) ClearRiskLocation() *RiskLocationUpdateOne {
	rluo.mutation.ClearRiskLocation()
	return rluo
}

// RemoveRiskLocationIDs removes the "risk_location" edge to Risk entities by IDs.
func (rluo *RiskLocationUpdateOne) RemoveRiskLocationIDs(ids ...int) *RiskLocationUpdateOne {
	rluo.mutation.RemoveRiskLocationIDs(ids...)
	return rluo
}

// RemoveRiskLocation removes "risk_location" edges to Risk entities.
func (rluo *RiskLocationUpdateOne) RemoveRiskLocation(r ...*Risk) *RiskLocationUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rluo.RemoveRiskLocationIDs(ids...)
}

// Where appends a list predicates to the RiskLocationUpdate builder.
func (rluo *RiskLocationUpdateOne) Where(ps ...predicate.RiskLocation) *RiskLocationUpdateOne {
	rluo.mutation.Where(ps...)
	return rluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rluo *RiskLocationUpdateOne) Select(field string, fields ...string) *RiskLocationUpdateOne {
	rluo.fields = append([]string{field}, fields...)
	return rluo
}

// Save executes the query and returns the updated RiskLocation entity.
func (rluo *RiskLocationUpdateOne) Save(ctx context.Context) (*RiskLocation, error) {
	if err := rluo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rluo.sqlSave, rluo.mutation, rluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rluo *RiskLocationUpdateOne) SaveX(ctx context.Context) *RiskLocation {
	node, err := rluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rluo *RiskLocationUpdateOne) Exec(ctx context.Context) error {
	_, err := rluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rluo *RiskLocationUpdateOne) ExecX(ctx context.Context) {
	if err := rluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rluo *RiskLocationUpdateOne) defaults() error {
	if _, ok := rluo.mutation.UpdatedAt(); !ok {
		if risklocation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized risklocation.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := risklocation.UpdateDefaultUpdatedAt()
		rluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rluo *RiskLocationUpdateOne) check() error {
	if v, ok := rluo.mutation.UpdatedBy(); ok {
		if err := risklocation.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "RiskLocation.updated_by": %w`, err)}
		}
	}
	if v, ok := rluo.mutation.Title(); ok {
		if err := risklocation.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "RiskLocation.title": %w`, err)}
		}
	}
	if _, ok := rluo.mutation.CreatorID(); rluo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "RiskLocation.creator"`)
	}
	return nil
}

func (rluo *RiskLocationUpdateOne) sqlSave(ctx context.Context) (_node *RiskLocation, err error) {
	if err := rluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(risklocation.Table, risklocation.Columns, sqlgraph.NewFieldSpec(risklocation.FieldID, field.TypeInt))
	id, ok := rluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "RiskLocation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, risklocation.FieldID)
		for _, f := range fields {
			if !risklocation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != risklocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rluo.mutation.DeletedAt(); ok {
		_spec.SetField(risklocation.FieldDeletedAt, field.TypeTime, value)
	}
	if rluo.mutation.DeletedAtCleared() {
		_spec.ClearField(risklocation.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rluo.mutation.UpdatedBy(); ok {
		_spec.SetField(risklocation.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rluo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(risklocation.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rluo.mutation.UpdatedAt(); ok {
		_spec.SetField(risklocation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rluo.mutation.Title(); ok {
		_spec.SetField(risklocation.FieldTitle, field.TypeString, value)
	}
	if rluo.mutation.RiskLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   risklocation.RiskLocationTable,
			Columns: []string{risklocation.RiskLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rluo.mutation.RemovedRiskLocationIDs(); len(nodes) > 0 && !rluo.mutation.RiskLocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   risklocation.RiskLocationTable,
			Columns: []string{risklocation.RiskLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rluo.mutation.RiskLocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   risklocation.RiskLocationTable,
			Columns: []string{risklocation.RiskLocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RiskLocation{config: rluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{risklocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rluo.mutation.done = true
	return _node, nil
}