// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/risk"
	"aisecurity/ent/dao/riskcategory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RiskCategoryUpdate is the builder for updating RiskCategory entities.
type RiskCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *RiskCategoryMutation
}

// Where appends a list predicates to the RiskCategoryUpdate builder.
func (rcu *RiskCategoryUpdate) Where(ps ...predicate.RiskCategory) *RiskCategoryUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetDeletedAt sets the "deleted_at" field.
func (rcu *RiskCategoryUpdate) SetDeletedAt(t time.Time) *RiskCategoryUpdate {
	rcu.mutation.SetDeletedAt(t)
	return rcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rcu *RiskCategoryUpdate) SetNillableDeletedAt(t *time.Time) *RiskCategoryUpdate {
	if t != nil {
		rcu.SetDeletedAt(*t)
	}
	return rcu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rcu *RiskCategoryUpdate) ClearDeletedAt() *RiskCategoryUpdate {
	rcu.mutation.ClearDeletedAt()
	return rcu
}

// SetUpdatedBy sets the "updated_by" field.
func (rcu *RiskCategoryUpdate) SetUpdatedBy(i int) *RiskCategoryUpdate {
	rcu.mutation.ResetUpdatedBy()
	rcu.mutation.SetUpdatedBy(i)
	return rcu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rcu *RiskCategoryUpdate) SetNillableUpdatedBy(i *int) *RiskCategoryUpdate {
	if i != nil {
		rcu.SetUpdatedBy(*i)
	}
	return rcu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (rcu *RiskCategoryUpdate) AddUpdatedBy(i int) *RiskCategoryUpdate {
	rcu.mutation.AddUpdatedBy(i)
	return rcu
}

// SetUpdatedAt sets the "updated_at" field.
func (rcu *RiskCategoryUpdate) SetUpdatedAt(t time.Time) *RiskCategoryUpdate {
	rcu.mutation.SetUpdatedAt(t)
	return rcu
}

// SetTitle sets the "title" field.
func (rcu *RiskCategoryUpdate) SetTitle(s string) *RiskCategoryUpdate {
	rcu.mutation.SetTitle(s)
	return rcu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (rcu *RiskCategoryUpdate) SetNillableTitle(s *string) *RiskCategoryUpdate {
	if s != nil {
		rcu.SetTitle(*s)
	}
	return rcu
}

// AddRiskCategoryIDs adds the "risk_category" edge to the Risk entity by IDs.
func (rcu *RiskCategoryUpdate) AddRiskCategoryIDs(ids ...int) *RiskCategoryUpdate {
	rcu.mutation.AddRiskCategoryIDs(ids...)
	return rcu
}

// AddRiskCategory adds the "risk_category" edges to the Risk entity.
func (rcu *RiskCategoryUpdate) AddRiskCategory(r ...*Risk) *RiskCategoryUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.AddRiskCategoryIDs(ids...)
}

// Mutation returns the RiskCategoryMutation object of the builder.
func (rcu *RiskCategoryUpdate) Mutation() *RiskCategoryMutation {
	return rcu.mutation
}

// ClearRiskCategory clears all "risk_category" edges to the Risk entity.
func (rcu *RiskCategoryUpdate) ClearRiskCategory() *RiskCategoryUpdate {
	rcu.mutation.ClearRiskCategory()
	return rcu
}

// RemoveRiskCategoryIDs removes the "risk_category" edge to Risk entities by IDs.
func (rcu *RiskCategoryUpdate) RemoveRiskCategoryIDs(ids ...int) *RiskCategoryUpdate {
	rcu.mutation.RemoveRiskCategoryIDs(ids...)
	return rcu
}

// RemoveRiskCategory removes "risk_category" edges to Risk entities.
func (rcu *RiskCategoryUpdate) RemoveRiskCategory(r ...*Risk) *RiskCategoryUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcu.RemoveRiskCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *RiskCategoryUpdate) Save(ctx context.Context) (int, error) {
	if err := rcu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, rcu.sqlSave, rcu.mutation, rcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *RiskCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *RiskCategoryUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *RiskCategoryUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcu *RiskCategoryUpdate) defaults() error {
	if _, ok := rcu.mutation.UpdatedAt(); !ok {
		if riskcategory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized riskcategory.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := riskcategory.UpdateDefaultUpdatedAt()
		rcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcu *RiskCategoryUpdate) check() error {
	if v, ok := rcu.mutation.UpdatedBy(); ok {
		if err := riskcategory.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "RiskCategory.updated_by": %w`, err)}
		}
	}
	if v, ok := rcu.mutation.Title(); ok {
		if err := riskcategory.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "RiskCategory.title": %w`, err)}
		}
	}
	if _, ok := rcu.mutation.CreatorID(); rcu.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "RiskCategory.creator"`)
	}
	return nil
}

func (rcu *RiskCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(riskcategory.Table, riskcategory.Columns, sqlgraph.NewFieldSpec(riskcategory.FieldID, field.TypeInt))
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.DeletedAt(); ok {
		_spec.SetField(riskcategory.FieldDeletedAt, field.TypeTime, value)
	}
	if rcu.mutation.DeletedAtCleared() {
		_spec.ClearField(riskcategory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rcu.mutation.UpdatedBy(); ok {
		_spec.SetField(riskcategory.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rcu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(riskcategory.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rcu.mutation.UpdatedAt(); ok {
		_spec.SetField(riskcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rcu.mutation.Title(); ok {
		_spec.SetField(riskcategory.FieldTitle, field.TypeString, value)
	}
	if rcu.mutation.RiskCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   riskcategory.RiskCategoryTable,
			Columns: []string{riskcategory.RiskCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.RemovedRiskCategoryIDs(); len(nodes) > 0 && !rcu.mutation.RiskCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   riskcategory.RiskCategoryTable,
			Columns: []string{riskcategory.RiskCategoryColumn},
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
	if nodes := rcu.mutation.RiskCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   riskcategory.RiskCategoryTable,
			Columns: []string{riskcategory.RiskCategoryColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{riskcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rcu.mutation.done = true
	return n, nil
}

// RiskCategoryUpdateOne is the builder for updating a single RiskCategory entity.
type RiskCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RiskCategoryMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (rcuo *RiskCategoryUpdateOne) SetDeletedAt(t time.Time) *RiskCategoryUpdateOne {
	rcuo.mutation.SetDeletedAt(t)
	return rcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rcuo *RiskCategoryUpdateOne) SetNillableDeletedAt(t *time.Time) *RiskCategoryUpdateOne {
	if t != nil {
		rcuo.SetDeletedAt(*t)
	}
	return rcuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rcuo *RiskCategoryUpdateOne) ClearDeletedAt() *RiskCategoryUpdateOne {
	rcuo.mutation.ClearDeletedAt()
	return rcuo
}

// SetUpdatedBy sets the "updated_by" field.
func (rcuo *RiskCategoryUpdateOne) SetUpdatedBy(i int) *RiskCategoryUpdateOne {
	rcuo.mutation.ResetUpdatedBy()
	rcuo.mutation.SetUpdatedBy(i)
	return rcuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rcuo *RiskCategoryUpdateOne) SetNillableUpdatedBy(i *int) *RiskCategoryUpdateOne {
	if i != nil {
		rcuo.SetUpdatedBy(*i)
	}
	return rcuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (rcuo *RiskCategoryUpdateOne) AddUpdatedBy(i int) *RiskCategoryUpdateOne {
	rcuo.mutation.AddUpdatedBy(i)
	return rcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (rcuo *RiskCategoryUpdateOne) SetUpdatedAt(t time.Time) *RiskCategoryUpdateOne {
	rcuo.mutation.SetUpdatedAt(t)
	return rcuo
}

// SetTitle sets the "title" field.
func (rcuo *RiskCategoryUpdateOne) SetTitle(s string) *RiskCategoryUpdateOne {
	rcuo.mutation.SetTitle(s)
	return rcuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (rcuo *RiskCategoryUpdateOne) SetNillableTitle(s *string) *RiskCategoryUpdateOne {
	if s != nil {
		rcuo.SetTitle(*s)
	}
	return rcuo
}

// AddRiskCategoryIDs adds the "risk_category" edge to the Risk entity by IDs.
func (rcuo *RiskCategoryUpdateOne) AddRiskCategoryIDs(ids ...int) *RiskCategoryUpdateOne {
	rcuo.mutation.AddRiskCategoryIDs(ids...)
	return rcuo
}

// AddRiskCategory adds the "risk_category" edges to the Risk entity.
func (rcuo *RiskCategoryUpdateOne) AddRiskCategory(r ...*Risk) *RiskCategoryUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.AddRiskCategoryIDs(ids...)
}

// Mutation returns the RiskCategoryMutation object of the builder.
func (rcuo *RiskCategoryUpdateOne) Mutation() *RiskCategoryMutation {
	return rcuo.mutation
}

// ClearRiskCategory clears all "risk_category" edges to the Risk entity.
func (rcuo *RiskCategoryUpdateOne) ClearRiskCategory() *RiskCategoryUpdateOne {
	rcuo.mutation.ClearRiskCategory()
	return rcuo
}

// RemoveRiskCategoryIDs removes the "risk_category" edge to Risk entities by IDs.
func (rcuo *RiskCategoryUpdateOne) RemoveRiskCategoryIDs(ids ...int) *RiskCategoryUpdateOne {
	rcuo.mutation.RemoveRiskCategoryIDs(ids...)
	return rcuo
}

// RemoveRiskCategory removes "risk_category" edges to Risk entities.
func (rcuo *RiskCategoryUpdateOne) RemoveRiskCategory(r ...*Risk) *RiskCategoryUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcuo.RemoveRiskCategoryIDs(ids...)
}

// Where appends a list predicates to the RiskCategoryUpdate builder.
func (rcuo *RiskCategoryUpdateOne) Where(ps ...predicate.RiskCategory) *RiskCategoryUpdateOne {
	rcuo.mutation.Where(ps...)
	return rcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *RiskCategoryUpdateOne) Select(field string, fields ...string) *RiskCategoryUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated RiskCategory entity.
func (rcuo *RiskCategoryUpdateOne) Save(ctx context.Context) (*RiskCategory, error) {
	if err := rcuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rcuo.sqlSave, rcuo.mutation, rcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *RiskCategoryUpdateOne) SaveX(ctx context.Context) *RiskCategory {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *RiskCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *RiskCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcuo *RiskCategoryUpdateOne) defaults() error {
	if _, ok := rcuo.mutation.UpdatedAt(); !ok {
		if riskcategory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized riskcategory.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := riskcategory.UpdateDefaultUpdatedAt()
		rcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcuo *RiskCategoryUpdateOne) check() error {
	if v, ok := rcuo.mutation.UpdatedBy(); ok {
		if err := riskcategory.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "RiskCategory.updated_by": %w`, err)}
		}
	}
	if v, ok := rcuo.mutation.Title(); ok {
		if err := riskcategory.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "RiskCategory.title": %w`, err)}
		}
	}
	if _, ok := rcuo.mutation.CreatorID(); rcuo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "RiskCategory.creator"`)
	}
	return nil
}

func (rcuo *RiskCategoryUpdateOne) sqlSave(ctx context.Context) (_node *RiskCategory, err error) {
	if err := rcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(riskcategory.Table, riskcategory.Columns, sqlgraph.NewFieldSpec(riskcategory.FieldID, field.TypeInt))
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "RiskCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, riskcategory.FieldID)
		for _, f := range fields {
			if !riskcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != riskcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.DeletedAt(); ok {
		_spec.SetField(riskcategory.FieldDeletedAt, field.TypeTime, value)
	}
	if rcuo.mutation.DeletedAtCleared() {
		_spec.ClearField(riskcategory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rcuo.mutation.UpdatedBy(); ok {
		_spec.SetField(riskcategory.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rcuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(riskcategory.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := rcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(riskcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rcuo.mutation.Title(); ok {
		_spec.SetField(riskcategory.FieldTitle, field.TypeString, value)
	}
	if rcuo.mutation.RiskCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   riskcategory.RiskCategoryTable,
			Columns: []string{riskcategory.RiskCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.RemovedRiskCategoryIDs(); len(nodes) > 0 && !rcuo.mutation.RiskCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   riskcategory.RiskCategoryTable,
			Columns: []string{riskcategory.RiskCategoryColumn},
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
	if nodes := rcuo.mutation.RiskCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   riskcategory.RiskCategoryTable,
			Columns: []string{riskcategory.RiskCategoryColumn},
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
	_node = &RiskCategory{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{riskcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rcuo.mutation.done = true
	return _node, nil
}