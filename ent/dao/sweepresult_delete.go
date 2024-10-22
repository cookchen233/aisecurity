// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/sweepresult"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SweepResultDelete is the builder for deleting a SweepResult entity.
type SweepResultDelete struct {
	config
	hooks    []Hook
	mutation *SweepResultMutation
}

// Where appends a list predicates to the SweepResultDelete builder.
func (srd *SweepResultDelete) Where(ps ...predicate.SweepResult) *SweepResultDelete {
	srd.mutation.Where(ps...)
	return srd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srd *SweepResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srd.sqlExec, srd.mutation, srd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srd *SweepResultDelete) ExecX(ctx context.Context) int {
	n, err := srd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srd *SweepResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sweepresult.Table, sqlgraph.NewFieldSpec(sweepresult.FieldID, field.TypeInt))
	if ps := srd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srd.mutation.done = true
	return affected, err
}

// SweepResultDeleteOne is the builder for deleting a single SweepResult entity.
type SweepResultDeleteOne struct {
	srd *SweepResultDelete
}

// Where appends a list predicates to the SweepResultDelete builder.
func (srdo *SweepResultDeleteOne) Where(ps ...predicate.SweepResult) *SweepResultDeleteOne {
	srdo.srd.mutation.Where(ps...)
	return srdo
}

// Exec executes the deletion query.
func (srdo *SweepResultDeleteOne) Exec(ctx context.Context) error {
	n, err := srdo.srd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sweepresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srdo *SweepResultDeleteOne) ExecX(ctx context.Context) {
	if err := srdo.Exec(ctx); err != nil {
		panic(err)
	}
}
