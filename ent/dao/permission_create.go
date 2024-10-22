// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/department"
	"aisecurity/ent/dao/permission"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PermissionCreate) SetCreateTime(t time.Time) *PermissionCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreateTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetCreatorID sets the "creator_id" field.
func (pc *PermissionCreate) SetCreatorID(i int) *PermissionCreate {
	pc.mutation.SetCreatorID(i)
	return pc
}

// SetDeleteTime sets the "delete_time" field.
func (pc *PermissionCreate) SetDeleteTime(t time.Time) *PermissionCreate {
	pc.mutation.SetDeleteTime(t)
	return pc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDeleteTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetDeleteTime(*t)
	}
	return pc
}

// SetUpdaterID sets the "updater_id" field.
func (pc *PermissionCreate) SetUpdaterID(i int) *PermissionCreate {
	pc.mutation.SetUpdaterID(i)
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PermissionCreate) SetUpdateTime(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdateTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PermissionCreate) SetName(s string) *PermissionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetAccessIds sets the "access_ids" field.
func (pc *PermissionCreate) SetAccessIds(s []string) *PermissionCreate {
	pc.mutation.SetAccessIds(s)
	return pc
}

// SetCreator sets the "creator" edge to the Admin entity.
func (pc *PermissionCreate) SetCreator(a *Admin) *PermissionCreate {
	return pc.SetCreatorID(a.ID)
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (pc *PermissionCreate) SetUpdater(a *Admin) *PermissionCreate {
	return pc.SetUpdaterID(a.ID)
}

// AddAdminIDs adds the "admin" edge to the Admin entity by IDs.
func (pc *PermissionCreate) AddAdminIDs(ids ...int) *PermissionCreate {
	pc.mutation.AddAdminIDs(ids...)
	return pc
}

// AddAdmin adds the "admin" edges to the Admin entity.
func (pc *PermissionCreate) AddAdmin(a ...*Admin) *PermissionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddAdminIDs(ids...)
}

// AddDepartmentIDs adds the "department" edge to the Department entity by IDs.
func (pc *PermissionCreate) AddDepartmentIDs(ids ...int) *PermissionCreate {
	pc.mutation.AddDepartmentIDs(ids...)
	return pc
}

// AddDepartment adds the "department" edges to the Department entity.
func (pc *PermissionCreate) AddDepartment(d ...*Department) *PermissionCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDepartmentIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermissionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermissionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		if permission.DefaultCreateTime == nil {
			return fmt.Errorf("dao: uninitialized permission.DefaultCreateTime (forgotten import dao/runtime?)")
		}
		v := permission.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if permission.DefaultUpdateTime == nil {
			return fmt.Errorf("dao: uninitialized permission.DefaultUpdateTime (forgotten import dao/runtime?)")
		}
		v := permission.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`dao: missing required field "Permission.create_time"`)}
	}
	if _, ok := pc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator_id", err: errors.New(`dao: missing required field "Permission.creator_id"`)}
	}
	if v, ok := pc.mutation.CreatorID(); ok {
		if err := permission.CreatorIDValidator(v); err != nil {
			return &ValidationError{Name: "creator_id", err: fmt.Errorf(`dao: validator failed for field "Permission.creator_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.UpdaterID(); !ok {
		return &ValidationError{Name: "updater_id", err: errors.New(`dao: missing required field "Permission.updater_id"`)}
	}
	if v, ok := pc.mutation.UpdaterID(); ok {
		if err := permission.UpdaterIDValidator(v); err != nil {
			return &ValidationError{Name: "updater_id", err: fmt.Errorf(`dao: validator failed for field "Permission.updater_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`dao: missing required field "Permission.update_time"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`dao: missing required field "Permission.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`dao: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.AccessIds(); !ok {
		return &ValidationError{Name: "access_ids", err: errors.New(`dao: missing required field "Permission.access_ids"`)}
	}
	if _, ok := pc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`dao: missing required edge "Permission.creator"`)}
	}
	if _, ok := pc.mutation.UpdaterID(); !ok {
		return &ValidationError{Name: "updater", err: errors.New(`dao: missing required edge "Permission.updater"`)}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(permission.Table, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(permission.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.DeleteTime(); ok {
		_spec.SetField(permission.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(permission.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.AccessIds(); ok {
		_spec.SetField(permission.FieldAccessIds, field.TypeJSON, value)
		_node.AccessIds = value
	}
	if nodes := pc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.CreatorTable,
			Columns: []string{permission.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.UpdaterTable,
			Columns: []string{permission.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UpdaterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.AdminTable,
			Columns: permission.AdminPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.DepartmentTable,
			Columns: permission.DepartmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	err      error
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermissionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
