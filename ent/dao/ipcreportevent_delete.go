// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/ipcreportevent"
	"aisecurity/ent/dao/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IPCReportEventDelete is the builder for deleting a IPCReportEvent entity.
type IPCReportEventDelete struct {
	config
	hooks    []Hook
	mutation *IPCReportEventMutation
}

// Where appends a list predicates to the IPCReportEventDelete builder.
func (ired *IPCReportEventDelete) Where(ps ...predicate.IPCReportEvent) *IPCReportEventDelete {
	ired.mutation.Where(ps...)
	return ired
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ired *IPCReportEventDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ired.sqlExec, ired.mutation, ired.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ired *IPCReportEventDelete) ExecX(ctx context.Context) int {
	n, err := ired.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ired *IPCReportEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ipcreportevent.Table, sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt))
	if ps := ired.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ired.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ired.mutation.done = true
	return affected, err
}

// IPCReportEventDeleteOne is the builder for deleting a single IPCReportEvent entity.
type IPCReportEventDeleteOne struct {
	ired *IPCReportEventDelete
}

// Where appends a list predicates to the IPCReportEventDelete builder.
func (iredo *IPCReportEventDeleteOne) Where(ps ...predicate.IPCReportEvent) *IPCReportEventDeleteOne {
	iredo.ired.mutation.Where(ps...)
	return iredo
}

// Exec executes the deletion query.
func (iredo *IPCReportEventDeleteOne) Exec(ctx context.Context) error {
	n, err := iredo.ired.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ipcreportevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iredo *IPCReportEventDeleteOne) ExecX(ctx context.Context) {
	if err := iredo.Exec(ctx); err != nil {
		panic(err)
	}
}