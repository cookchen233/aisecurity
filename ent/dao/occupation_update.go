// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/employee"
	"aisecurity/ent/dao/occupation"
	"aisecurity/ent/dao/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OccupationUpdate is the builder for updating Occupation entities.
type OccupationUpdate struct {
	config
	hooks    []Hook
	mutation *OccupationMutation
}

// Where appends a list predicates to the OccupationUpdate builder.
func (ou *OccupationUpdate) Where(ps ...predicate.Occupation) *OccupationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetDeletedAt sets the "deleted_at" field.
func (ou *OccupationUpdate) SetDeletedAt(t time.Time) *OccupationUpdate {
	ou.mutation.SetDeletedAt(t)
	return ou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ou *OccupationUpdate) SetNillableDeletedAt(t *time.Time) *OccupationUpdate {
	if t != nil {
		ou.SetDeletedAt(*t)
	}
	return ou
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ou *OccupationUpdate) ClearDeletedAt() *OccupationUpdate {
	ou.mutation.ClearDeletedAt()
	return ou
}

// SetUpdatedBy sets the "updated_by" field.
func (ou *OccupationUpdate) SetUpdatedBy(i int) *OccupationUpdate {
	ou.mutation.SetUpdatedBy(i)
	return ou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ou *OccupationUpdate) SetNillableUpdatedBy(i *int) *OccupationUpdate {
	if i != nil {
		ou.SetUpdatedBy(*i)
	}
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OccupationUpdate) SetUpdatedAt(t time.Time) *OccupationUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetName sets the "name" field.
func (ou *OccupationUpdate) SetName(s string) *OccupationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OccupationUpdate) SetNillableName(s *string) *OccupationUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// SetDescription sets the "description" field.
func (ou *OccupationUpdate) SetDescription(s string) *OccupationUpdate {
	ou.mutation.SetDescription(s)
	return ou
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ou *OccupationUpdate) SetNillableDescription(s *string) *OccupationUpdate {
	if s != nil {
		ou.SetDescription(*s)
	}
	return ou
}

// ClearDescription clears the value of the "description" field.
func (ou *OccupationUpdate) ClearDescription() *OccupationUpdate {
	ou.mutation.ClearDescription()
	return ou
}

// SetUpdaterID sets the "updater" edge to the Admin entity by ID.
func (ou *OccupationUpdate) SetUpdaterID(id int) *OccupationUpdate {
	ou.mutation.SetUpdaterID(id)
	return ou
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (ou *OccupationUpdate) SetUpdater(a *Admin) *OccupationUpdate {
	return ou.SetUpdaterID(a.ID)
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (ou *OccupationUpdate) AddEmployeeIDs(ids ...int) *OccupationUpdate {
	ou.mutation.AddEmployeeIDs(ids...)
	return ou
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (ou *OccupationUpdate) AddEmployees(e ...*Employee) *OccupationUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ou.AddEmployeeIDs(ids...)
}

// Mutation returns the OccupationMutation object of the builder.
func (ou *OccupationUpdate) Mutation() *OccupationMutation {
	return ou.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (ou *OccupationUpdate) ClearUpdater() *OccupationUpdate {
	ou.mutation.ClearUpdater()
	return ou
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (ou *OccupationUpdate) ClearEmployees() *OccupationUpdate {
	ou.mutation.ClearEmployees()
	return ou
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (ou *OccupationUpdate) RemoveEmployeeIDs(ids ...int) *OccupationUpdate {
	ou.mutation.RemoveEmployeeIDs(ids...)
	return ou
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (ou *OccupationUpdate) RemoveEmployees(e ...*Employee) *OccupationUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ou.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OccupationUpdate) Save(ctx context.Context) (int, error) {
	if err := ou.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OccupationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OccupationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OccupationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OccupationUpdate) defaults() error {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		if occupation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized occupation.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := occupation.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ou *OccupationUpdate) check() error {
	if v, ok := ou.mutation.UpdatedBy(); ok {
		if err := occupation.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "Occupation.updated_by": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Name(); ok {
		if err := occupation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`dao: validator failed for field "Occupation.name": %w`, err)}
		}
	}
	if _, ok := ou.mutation.CreatorID(); ou.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Occupation.creator"`)
	}
	if _, ok := ou.mutation.UpdaterID(); ou.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Occupation.updater"`)
	}
	return nil
}

func (ou *OccupationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(occupation.Table, occupation.Columns, sqlgraph.NewFieldSpec(occupation.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.DeletedAt(); ok {
		_spec.SetField(occupation.FieldDeletedAt, field.TypeTime, value)
	}
	if ou.mutation.DeletedAtCleared() {
		_spec.ClearField(occupation.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(occupation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(occupation.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Description(); ok {
		_spec.SetField(occupation.FieldDescription, field.TypeString, value)
	}
	if ou.mutation.DescriptionCleared() {
		_spec.ClearField(occupation.FieldDescription, field.TypeString)
	}
	if ou.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   occupation.UpdaterTable,
			Columns: []string{occupation.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   occupation.UpdaterTable,
			Columns: []string{occupation.UpdaterColumn},
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
	if ou.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   occupation.EmployeesTable,
			Columns: occupation.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !ou.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   occupation.EmployeesTable,
			Columns: occupation.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   occupation.EmployeesTable,
			Columns: occupation.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{occupation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OccupationUpdateOne is the builder for updating a single Occupation entity.
type OccupationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OccupationMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (ouo *OccupationUpdateOne) SetDeletedAt(t time.Time) *OccupationUpdateOne {
	ouo.mutation.SetDeletedAt(t)
	return ouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouo *OccupationUpdateOne) SetNillableDeletedAt(t *time.Time) *OccupationUpdateOne {
	if t != nil {
		ouo.SetDeletedAt(*t)
	}
	return ouo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ouo *OccupationUpdateOne) ClearDeletedAt() *OccupationUpdateOne {
	ouo.mutation.ClearDeletedAt()
	return ouo
}

// SetUpdatedBy sets the "updated_by" field.
func (ouo *OccupationUpdateOne) SetUpdatedBy(i int) *OccupationUpdateOne {
	ouo.mutation.SetUpdatedBy(i)
	return ouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouo *OccupationUpdateOne) SetNillableUpdatedBy(i *int) *OccupationUpdateOne {
	if i != nil {
		ouo.SetUpdatedBy(*i)
	}
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OccupationUpdateOne) SetUpdatedAt(t time.Time) *OccupationUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetName sets the "name" field.
func (ouo *OccupationUpdateOne) SetName(s string) *OccupationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OccupationUpdateOne) SetNillableName(s *string) *OccupationUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// SetDescription sets the "description" field.
func (ouo *OccupationUpdateOne) SetDescription(s string) *OccupationUpdateOne {
	ouo.mutation.SetDescription(s)
	return ouo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ouo *OccupationUpdateOne) SetNillableDescription(s *string) *OccupationUpdateOne {
	if s != nil {
		ouo.SetDescription(*s)
	}
	return ouo
}

// ClearDescription clears the value of the "description" field.
func (ouo *OccupationUpdateOne) ClearDescription() *OccupationUpdateOne {
	ouo.mutation.ClearDescription()
	return ouo
}

// SetUpdaterID sets the "updater" edge to the Admin entity by ID.
func (ouo *OccupationUpdateOne) SetUpdaterID(id int) *OccupationUpdateOne {
	ouo.mutation.SetUpdaterID(id)
	return ouo
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (ouo *OccupationUpdateOne) SetUpdater(a *Admin) *OccupationUpdateOne {
	return ouo.SetUpdaterID(a.ID)
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (ouo *OccupationUpdateOne) AddEmployeeIDs(ids ...int) *OccupationUpdateOne {
	ouo.mutation.AddEmployeeIDs(ids...)
	return ouo
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (ouo *OccupationUpdateOne) AddEmployees(e ...*Employee) *OccupationUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ouo.AddEmployeeIDs(ids...)
}

// Mutation returns the OccupationMutation object of the builder.
func (ouo *OccupationUpdateOne) Mutation() *OccupationMutation {
	return ouo.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (ouo *OccupationUpdateOne) ClearUpdater() *OccupationUpdateOne {
	ouo.mutation.ClearUpdater()
	return ouo
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (ouo *OccupationUpdateOne) ClearEmployees() *OccupationUpdateOne {
	ouo.mutation.ClearEmployees()
	return ouo
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (ouo *OccupationUpdateOne) RemoveEmployeeIDs(ids ...int) *OccupationUpdateOne {
	ouo.mutation.RemoveEmployeeIDs(ids...)
	return ouo
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (ouo *OccupationUpdateOne) RemoveEmployees(e ...*Employee) *OccupationUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ouo.RemoveEmployeeIDs(ids...)
}

// Where appends a list predicates to the OccupationUpdate builder.
func (ouo *OccupationUpdateOne) Where(ps ...predicate.Occupation) *OccupationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OccupationUpdateOne) Select(field string, fields ...string) *OccupationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Occupation entity.
func (ouo *OccupationUpdateOne) Save(ctx context.Context) (*Occupation, error) {
	if err := ouo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OccupationUpdateOne) SaveX(ctx context.Context) *Occupation {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OccupationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OccupationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OccupationUpdateOne) defaults() error {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		if occupation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized occupation.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := occupation.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OccupationUpdateOne) check() error {
	if v, ok := ouo.mutation.UpdatedBy(); ok {
		if err := occupation.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "Occupation.updated_by": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Name(); ok {
		if err := occupation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`dao: validator failed for field "Occupation.name": %w`, err)}
		}
	}
	if _, ok := ouo.mutation.CreatorID(); ouo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Occupation.creator"`)
	}
	if _, ok := ouo.mutation.UpdaterID(); ouo.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Occupation.updater"`)
	}
	return nil
}

func (ouo *OccupationUpdateOne) sqlSave(ctx context.Context) (_node *Occupation, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(occupation.Table, occupation.Columns, sqlgraph.NewFieldSpec(occupation.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "Occupation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, occupation.FieldID)
		for _, f := range fields {
			if !occupation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != occupation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.DeletedAt(); ok {
		_spec.SetField(occupation.FieldDeletedAt, field.TypeTime, value)
	}
	if ouo.mutation.DeletedAtCleared() {
		_spec.ClearField(occupation.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(occupation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(occupation.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Description(); ok {
		_spec.SetField(occupation.FieldDescription, field.TypeString, value)
	}
	if ouo.mutation.DescriptionCleared() {
		_spec.ClearField(occupation.FieldDescription, field.TypeString)
	}
	if ouo.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   occupation.UpdaterTable,
			Columns: []string{occupation.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   occupation.UpdaterTable,
			Columns: []string{occupation.UpdaterColumn},
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
	if ouo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   occupation.EmployeesTable,
			Columns: occupation.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !ouo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   occupation.EmployeesTable,
			Columns: occupation.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   occupation.EmployeesTable,
			Columns: occupation.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Occupation{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{occupation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}