// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/risk"
	"aisecurity/ent/dao/riskcategory"
	"aisecurity/ent/dao/risklocation"
	"aisecurity/properties"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// RiskUpdate is the builder for updating Risk entities.
type RiskUpdate struct {
	config
	hooks    []Hook
	mutation *RiskMutation
}

// Where appends a list predicates to the RiskUpdate builder.
func (ru *RiskUpdate) Where(ps ...predicate.Risk) *RiskUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RiskUpdate) SetDeletedAt(t time.Time) *RiskUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableDeletedAt(t *time.Time) *RiskUpdate {
	if t != nil {
		ru.SetDeletedAt(*t)
	}
	return ru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ru *RiskUpdate) ClearDeletedAt() *RiskUpdate {
	ru.mutation.ClearDeletedAt()
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RiskUpdate) SetUpdatedBy(i int) *RiskUpdate {
	ru.mutation.ResetUpdatedBy()
	ru.mutation.SetUpdatedBy(i)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableUpdatedBy(i *int) *RiskUpdate {
	if i != nil {
		ru.SetUpdatedBy(*i)
	}
	return ru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ru *RiskUpdate) AddUpdatedBy(i int) *RiskUpdate {
	ru.mutation.AddUpdatedBy(i)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RiskUpdate) SetUpdatedAt(t time.Time) *RiskUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetTitle sets the "title" field.
func (ru *RiskUpdate) SetTitle(s string) *RiskUpdate {
	ru.mutation.SetTitle(s)
	return ru
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableTitle(s *string) *RiskUpdate {
	if s != nil {
		ru.SetTitle(*s)
	}
	return ru
}

// SetContent sets the "content" field.
func (ru *RiskUpdate) SetContent(s string) *RiskUpdate {
	ru.mutation.SetContent(s)
	return ru
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableContent(s *string) *RiskUpdate {
	if s != nil {
		ru.SetContent(*s)
	}
	return ru
}

// SetImages sets the "images" field.
func (ru *RiskUpdate) SetImages(s []struct {
	Title string "json:\"title\""
	URL   string "json:\"url\""
}) *RiskUpdate {
	ru.mutation.SetImages(s)
	return ru
}

// AppendImages appends s to the "images" field.
func (ru *RiskUpdate) AppendImages(s []struct {
	Title string "json:\"title\""
	URL   string "json:\"url\""
}) *RiskUpdate {
	ru.mutation.AppendImages(s)
	return ru
}

// ClearImages clears the value of the "images" field.
func (ru *RiskUpdate) ClearImages() *RiskUpdate {
	ru.mutation.ClearImages()
	return ru
}

// SetRiskCategoryID sets the "risk_category_id" field.
func (ru *RiskUpdate) SetRiskCategoryID(i int) *RiskUpdate {
	ru.mutation.SetRiskCategoryID(i)
	return ru
}

// SetNillableRiskCategoryID sets the "risk_category_id" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableRiskCategoryID(i *int) *RiskUpdate {
	if i != nil {
		ru.SetRiskCategoryID(*i)
	}
	return ru
}

// SetRiskLocationID sets the "risk_location_id" field.
func (ru *RiskUpdate) SetRiskLocationID(i int) *RiskUpdate {
	ru.mutation.SetRiskLocationID(i)
	return ru
}

// SetNillableRiskLocationID sets the "risk_location_id" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableRiskLocationID(i *int) *RiskUpdate {
	if i != nil {
		ru.SetRiskLocationID(*i)
	}
	return ru
}

// SetMaintainer sets the "maintainer" field.
func (ru *RiskUpdate) SetMaintainer(i int) *RiskUpdate {
	ru.mutation.SetMaintainer(i)
	return ru
}

// SetNillableMaintainer sets the "maintainer" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableMaintainer(i *int) *RiskUpdate {
	if i != nil {
		ru.SetMaintainer(*i)
	}
	return ru
}

// SetMeasures sets the "measures" field.
func (ru *RiskUpdate) SetMeasures(s string) *RiskUpdate {
	ru.mutation.SetMeasures(s)
	return ru
}

// SetNillableMeasures sets the "measures" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableMeasures(s *string) *RiskUpdate {
	if s != nil {
		ru.SetMeasures(*s)
	}
	return ru
}

// ClearMeasures clears the value of the "measures" field.
func (ru *RiskUpdate) ClearMeasures() *RiskUpdate {
	ru.mutation.ClearMeasures()
	return ru
}

// SetMaintainStatus sets the "maintain_status" field.
func (ru *RiskUpdate) SetMaintainStatus(ps properties.MaintainStatus) *RiskUpdate {
	ru.mutation.ResetMaintainStatus()
	ru.mutation.SetMaintainStatus(ps)
	return ru
}

// SetNillableMaintainStatus sets the "maintain_status" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableMaintainStatus(ps *properties.MaintainStatus) *RiskUpdate {
	if ps != nil {
		ru.SetMaintainStatus(*ps)
	}
	return ru
}

// AddMaintainStatus adds ps to the "maintain_status" field.
func (ru *RiskUpdate) AddMaintainStatus(ps properties.MaintainStatus) *RiskUpdate {
	ru.mutation.AddMaintainStatus(ps)
	return ru
}

// SetDueTime sets the "due_time" field.
func (ru *RiskUpdate) SetDueTime(t time.Time) *RiskUpdate {
	ru.mutation.SetDueTime(t)
	return ru
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (ru *RiskUpdate) SetNillableDueTime(t *time.Time) *RiskUpdate {
	if t != nil {
		ru.SetDueTime(*t)
	}
	return ru
}

// SetMaintainerAdminID sets the "maintainer_admin" edge to the Admin entity by ID.
func (ru *RiskUpdate) SetMaintainerAdminID(id int) *RiskUpdate {
	ru.mutation.SetMaintainerAdminID(id)
	return ru
}

// SetMaintainerAdmin sets the "maintainer_admin" edge to the Admin entity.
func (ru *RiskUpdate) SetMaintainerAdmin(a *Admin) *RiskUpdate {
	return ru.SetMaintainerAdminID(a.ID)
}

// SetCategoryID sets the "category" edge to the RiskCategory entity by ID.
func (ru *RiskUpdate) SetCategoryID(id int) *RiskUpdate {
	ru.mutation.SetCategoryID(id)
	return ru
}

// SetCategory sets the "category" edge to the RiskCategory entity.
func (ru *RiskUpdate) SetCategory(r *RiskCategory) *RiskUpdate {
	return ru.SetCategoryID(r.ID)
}

// SetLocationID sets the "location" edge to the RiskLocation entity by ID.
func (ru *RiskUpdate) SetLocationID(id int) *RiskUpdate {
	ru.mutation.SetLocationID(id)
	return ru
}

// SetLocation sets the "location" edge to the RiskLocation entity.
func (ru *RiskUpdate) SetLocation(r *RiskLocation) *RiskUpdate {
	return ru.SetLocationID(r.ID)
}

// Mutation returns the RiskMutation object of the builder.
func (ru *RiskUpdate) Mutation() *RiskMutation {
	return ru.mutation
}

// ClearMaintainerAdmin clears the "maintainer_admin" edge to the Admin entity.
func (ru *RiskUpdate) ClearMaintainerAdmin() *RiskUpdate {
	ru.mutation.ClearMaintainerAdmin()
	return ru
}

// ClearCategory clears the "category" edge to the RiskCategory entity.
func (ru *RiskUpdate) ClearCategory() *RiskUpdate {
	ru.mutation.ClearCategory()
	return ru
}

// ClearLocation clears the "location" edge to the RiskLocation entity.
func (ru *RiskUpdate) ClearLocation() *RiskUpdate {
	ru.mutation.ClearLocation()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RiskUpdate) Save(ctx context.Context) (int, error) {
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RiskUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RiskUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RiskUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RiskUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if risk.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized risk.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := risk.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ru *RiskUpdate) check() error {
	if v, ok := ru.mutation.UpdatedBy(); ok {
		if err := risk.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "Risk.updated_by": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Title(); ok {
		if err := risk.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "Risk.title": %w`, err)}
		}
	}
	if v, ok := ru.mutation.RiskCategoryID(); ok {
		if err := risk.RiskCategoryIDValidator(v); err != nil {
			return &ValidationError{Name: "risk_category_id", err: fmt.Errorf(`dao: validator failed for field "Risk.risk_category_id": %w`, err)}
		}
	}
	if v, ok := ru.mutation.RiskLocationID(); ok {
		if err := risk.RiskLocationIDValidator(v); err != nil {
			return &ValidationError{Name: "risk_location_id", err: fmt.Errorf(`dao: validator failed for field "Risk.risk_location_id": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Maintainer(); ok {
		if err := risk.MaintainerValidator(v); err != nil {
			return &ValidationError{Name: "maintainer", err: fmt.Errorf(`dao: validator failed for field "Risk.maintainer": %w`, err)}
		}
	}
	if v, ok := ru.mutation.MaintainStatus(); ok {
		if err := risk.MaintainStatusValidator(int(v)); err != nil {
			return &ValidationError{Name: "maintain_status", err: fmt.Errorf(`dao: validator failed for field "Risk.maintain_status": %w`, err)}
		}
	}
	if _, ok := ru.mutation.CreatorID(); ru.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.creator"`)
	}
	if _, ok := ru.mutation.MaintainerAdminID(); ru.mutation.MaintainerAdminCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.maintainer_admin"`)
	}
	if _, ok := ru.mutation.CategoryID(); ru.mutation.CategoryCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.category"`)
	}
	if _, ok := ru.mutation.LocationID(); ru.mutation.LocationCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.location"`)
	}
	return nil
}

func (ru *RiskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(risk.Table, risk.Columns, sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(risk.FieldDeletedAt, field.TypeTime, value)
	}
	if ru.mutation.DeletedAtCleared() {
		_spec.ClearField(risk.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.SetField(risk.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(risk.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(risk.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Title(); ok {
		_spec.SetField(risk.FieldTitle, field.TypeString, value)
	}
	if value, ok := ru.mutation.Content(); ok {
		_spec.SetField(risk.FieldContent, field.TypeString, value)
	}
	if value, ok := ru.mutation.Images(); ok {
		_spec.SetField(risk.FieldImages, field.TypeJSON, value)
	}
	if value, ok := ru.mutation.AppendedImages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, risk.FieldImages, value)
		})
	}
	if ru.mutation.ImagesCleared() {
		_spec.ClearField(risk.FieldImages, field.TypeJSON)
	}
	if value, ok := ru.mutation.Measures(); ok {
		_spec.SetField(risk.FieldMeasures, field.TypeString, value)
	}
	if ru.mutation.MeasuresCleared() {
		_spec.ClearField(risk.FieldMeasures, field.TypeString)
	}
	if value, ok := ru.mutation.MaintainStatus(); ok {
		_spec.SetField(risk.FieldMaintainStatus, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedMaintainStatus(); ok {
		_spec.AddField(risk.FieldMaintainStatus, field.TypeInt, value)
	}
	if value, ok := ru.mutation.DueTime(); ok {
		_spec.SetField(risk.FieldDueTime, field.TypeTime, value)
	}
	if ru.mutation.MaintainerAdminCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.MaintainerAdminTable,
			Columns: []string{risk.MaintainerAdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MaintainerAdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.MaintainerAdminTable,
			Columns: []string{risk.MaintainerAdminColumn},
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
	if ru.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.CategoryTable,
			Columns: []string{risk.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(riskcategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.CategoryTable,
			Columns: []string{risk.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(riskcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.LocationTable,
			Columns: []string{risk.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risklocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.LocationTable,
			Columns: []string{risk.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risklocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{risk.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RiskUpdateOne is the builder for updating a single Risk entity.
type RiskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RiskMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RiskUpdateOne) SetDeletedAt(t time.Time) *RiskUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableDeletedAt(t *time.Time) *RiskUpdateOne {
	if t != nil {
		ruo.SetDeletedAt(*t)
	}
	return ruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruo *RiskUpdateOne) ClearDeletedAt() *RiskUpdateOne {
	ruo.mutation.ClearDeletedAt()
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RiskUpdateOne) SetUpdatedBy(i int) *RiskUpdateOne {
	ruo.mutation.ResetUpdatedBy()
	ruo.mutation.SetUpdatedBy(i)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableUpdatedBy(i *int) *RiskUpdateOne {
	if i != nil {
		ruo.SetUpdatedBy(*i)
	}
	return ruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ruo *RiskUpdateOne) AddUpdatedBy(i int) *RiskUpdateOne {
	ruo.mutation.AddUpdatedBy(i)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RiskUpdateOne) SetUpdatedAt(t time.Time) *RiskUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetTitle sets the "title" field.
func (ruo *RiskUpdateOne) SetTitle(s string) *RiskUpdateOne {
	ruo.mutation.SetTitle(s)
	return ruo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableTitle(s *string) *RiskUpdateOne {
	if s != nil {
		ruo.SetTitle(*s)
	}
	return ruo
}

// SetContent sets the "content" field.
func (ruo *RiskUpdateOne) SetContent(s string) *RiskUpdateOne {
	ruo.mutation.SetContent(s)
	return ruo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableContent(s *string) *RiskUpdateOne {
	if s != nil {
		ruo.SetContent(*s)
	}
	return ruo
}

// SetImages sets the "images" field.
func (ruo *RiskUpdateOne) SetImages(s []struct {
	Title string "json:\"title\""
	URL   string "json:\"url\""
}) *RiskUpdateOne {
	ruo.mutation.SetImages(s)
	return ruo
}

// AppendImages appends s to the "images" field.
func (ruo *RiskUpdateOne) AppendImages(s []struct {
	Title string "json:\"title\""
	URL   string "json:\"url\""
}) *RiskUpdateOne {
	ruo.mutation.AppendImages(s)
	return ruo
}

// ClearImages clears the value of the "images" field.
func (ruo *RiskUpdateOne) ClearImages() *RiskUpdateOne {
	ruo.mutation.ClearImages()
	return ruo
}

// SetRiskCategoryID sets the "risk_category_id" field.
func (ruo *RiskUpdateOne) SetRiskCategoryID(i int) *RiskUpdateOne {
	ruo.mutation.SetRiskCategoryID(i)
	return ruo
}

// SetNillableRiskCategoryID sets the "risk_category_id" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableRiskCategoryID(i *int) *RiskUpdateOne {
	if i != nil {
		ruo.SetRiskCategoryID(*i)
	}
	return ruo
}

// SetRiskLocationID sets the "risk_location_id" field.
func (ruo *RiskUpdateOne) SetRiskLocationID(i int) *RiskUpdateOne {
	ruo.mutation.SetRiskLocationID(i)
	return ruo
}

// SetNillableRiskLocationID sets the "risk_location_id" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableRiskLocationID(i *int) *RiskUpdateOne {
	if i != nil {
		ruo.SetRiskLocationID(*i)
	}
	return ruo
}

// SetMaintainer sets the "maintainer" field.
func (ruo *RiskUpdateOne) SetMaintainer(i int) *RiskUpdateOne {
	ruo.mutation.SetMaintainer(i)
	return ruo
}

// SetNillableMaintainer sets the "maintainer" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableMaintainer(i *int) *RiskUpdateOne {
	if i != nil {
		ruo.SetMaintainer(*i)
	}
	return ruo
}

// SetMeasures sets the "measures" field.
func (ruo *RiskUpdateOne) SetMeasures(s string) *RiskUpdateOne {
	ruo.mutation.SetMeasures(s)
	return ruo
}

// SetNillableMeasures sets the "measures" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableMeasures(s *string) *RiskUpdateOne {
	if s != nil {
		ruo.SetMeasures(*s)
	}
	return ruo
}

// ClearMeasures clears the value of the "measures" field.
func (ruo *RiskUpdateOne) ClearMeasures() *RiskUpdateOne {
	ruo.mutation.ClearMeasures()
	return ruo
}

// SetMaintainStatus sets the "maintain_status" field.
func (ruo *RiskUpdateOne) SetMaintainStatus(ps properties.MaintainStatus) *RiskUpdateOne {
	ruo.mutation.ResetMaintainStatus()
	ruo.mutation.SetMaintainStatus(ps)
	return ruo
}

// SetNillableMaintainStatus sets the "maintain_status" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableMaintainStatus(ps *properties.MaintainStatus) *RiskUpdateOne {
	if ps != nil {
		ruo.SetMaintainStatus(*ps)
	}
	return ruo
}

// AddMaintainStatus adds ps to the "maintain_status" field.
func (ruo *RiskUpdateOne) AddMaintainStatus(ps properties.MaintainStatus) *RiskUpdateOne {
	ruo.mutation.AddMaintainStatus(ps)
	return ruo
}

// SetDueTime sets the "due_time" field.
func (ruo *RiskUpdateOne) SetDueTime(t time.Time) *RiskUpdateOne {
	ruo.mutation.SetDueTime(t)
	return ruo
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (ruo *RiskUpdateOne) SetNillableDueTime(t *time.Time) *RiskUpdateOne {
	if t != nil {
		ruo.SetDueTime(*t)
	}
	return ruo
}

// SetMaintainerAdminID sets the "maintainer_admin" edge to the Admin entity by ID.
func (ruo *RiskUpdateOne) SetMaintainerAdminID(id int) *RiskUpdateOne {
	ruo.mutation.SetMaintainerAdminID(id)
	return ruo
}

// SetMaintainerAdmin sets the "maintainer_admin" edge to the Admin entity.
func (ruo *RiskUpdateOne) SetMaintainerAdmin(a *Admin) *RiskUpdateOne {
	return ruo.SetMaintainerAdminID(a.ID)
}

// SetCategoryID sets the "category" edge to the RiskCategory entity by ID.
func (ruo *RiskUpdateOne) SetCategoryID(id int) *RiskUpdateOne {
	ruo.mutation.SetCategoryID(id)
	return ruo
}

// SetCategory sets the "category" edge to the RiskCategory entity.
func (ruo *RiskUpdateOne) SetCategory(r *RiskCategory) *RiskUpdateOne {
	return ruo.SetCategoryID(r.ID)
}

// SetLocationID sets the "location" edge to the RiskLocation entity by ID.
func (ruo *RiskUpdateOne) SetLocationID(id int) *RiskUpdateOne {
	ruo.mutation.SetLocationID(id)
	return ruo
}

// SetLocation sets the "location" edge to the RiskLocation entity.
func (ruo *RiskUpdateOne) SetLocation(r *RiskLocation) *RiskUpdateOne {
	return ruo.SetLocationID(r.ID)
}

// Mutation returns the RiskMutation object of the builder.
func (ruo *RiskUpdateOne) Mutation() *RiskMutation {
	return ruo.mutation
}

// ClearMaintainerAdmin clears the "maintainer_admin" edge to the Admin entity.
func (ruo *RiskUpdateOne) ClearMaintainerAdmin() *RiskUpdateOne {
	ruo.mutation.ClearMaintainerAdmin()
	return ruo
}

// ClearCategory clears the "category" edge to the RiskCategory entity.
func (ruo *RiskUpdateOne) ClearCategory() *RiskUpdateOne {
	ruo.mutation.ClearCategory()
	return ruo
}

// ClearLocation clears the "location" edge to the RiskLocation entity.
func (ruo *RiskUpdateOne) ClearLocation() *RiskUpdateOne {
	ruo.mutation.ClearLocation()
	return ruo
}

// Where appends a list predicates to the RiskUpdate builder.
func (ruo *RiskUpdateOne) Where(ps ...predicate.Risk) *RiskUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RiskUpdateOne) Select(field string, fields ...string) *RiskUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Risk entity.
func (ruo *RiskUpdateOne) Save(ctx context.Context) (*Risk, error) {
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RiskUpdateOne) SaveX(ctx context.Context) *Risk {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RiskUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RiskUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RiskUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if risk.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized risk.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := risk.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RiskUpdateOne) check() error {
	if v, ok := ruo.mutation.UpdatedBy(); ok {
		if err := risk.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "Risk.updated_by": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Title(); ok {
		if err := risk.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`dao: validator failed for field "Risk.title": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.RiskCategoryID(); ok {
		if err := risk.RiskCategoryIDValidator(v); err != nil {
			return &ValidationError{Name: "risk_category_id", err: fmt.Errorf(`dao: validator failed for field "Risk.risk_category_id": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.RiskLocationID(); ok {
		if err := risk.RiskLocationIDValidator(v); err != nil {
			return &ValidationError{Name: "risk_location_id", err: fmt.Errorf(`dao: validator failed for field "Risk.risk_location_id": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Maintainer(); ok {
		if err := risk.MaintainerValidator(v); err != nil {
			return &ValidationError{Name: "maintainer", err: fmt.Errorf(`dao: validator failed for field "Risk.maintainer": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.MaintainStatus(); ok {
		if err := risk.MaintainStatusValidator(int(v)); err != nil {
			return &ValidationError{Name: "maintain_status", err: fmt.Errorf(`dao: validator failed for field "Risk.maintain_status": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.CreatorID(); ruo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.creator"`)
	}
	if _, ok := ruo.mutation.MaintainerAdminID(); ruo.mutation.MaintainerAdminCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.maintainer_admin"`)
	}
	if _, ok := ruo.mutation.CategoryID(); ruo.mutation.CategoryCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.category"`)
	}
	if _, ok := ruo.mutation.LocationID(); ruo.mutation.LocationCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Risk.location"`)
	}
	return nil
}

func (ruo *RiskUpdateOne) sqlSave(ctx context.Context) (_node *Risk, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(risk.Table, risk.Columns, sqlgraph.NewFieldSpec(risk.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "Risk.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, risk.FieldID)
		for _, f := range fields {
			if !risk.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != risk.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(risk.FieldDeletedAt, field.TypeTime, value)
	}
	if ruo.mutation.DeletedAtCleared() {
		_spec.ClearField(risk.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.SetField(risk.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(risk.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(risk.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Title(); ok {
		_spec.SetField(risk.FieldTitle, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Content(); ok {
		_spec.SetField(risk.FieldContent, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Images(); ok {
		_spec.SetField(risk.FieldImages, field.TypeJSON, value)
	}
	if value, ok := ruo.mutation.AppendedImages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, risk.FieldImages, value)
		})
	}
	if ruo.mutation.ImagesCleared() {
		_spec.ClearField(risk.FieldImages, field.TypeJSON)
	}
	if value, ok := ruo.mutation.Measures(); ok {
		_spec.SetField(risk.FieldMeasures, field.TypeString, value)
	}
	if ruo.mutation.MeasuresCleared() {
		_spec.ClearField(risk.FieldMeasures, field.TypeString)
	}
	if value, ok := ruo.mutation.MaintainStatus(); ok {
		_spec.SetField(risk.FieldMaintainStatus, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedMaintainStatus(); ok {
		_spec.AddField(risk.FieldMaintainStatus, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.DueTime(); ok {
		_spec.SetField(risk.FieldDueTime, field.TypeTime, value)
	}
	if ruo.mutation.MaintainerAdminCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.MaintainerAdminTable,
			Columns: []string{risk.MaintainerAdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MaintainerAdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.MaintainerAdminTable,
			Columns: []string{risk.MaintainerAdminColumn},
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
	if ruo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.CategoryTable,
			Columns: []string{risk.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(riskcategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.CategoryTable,
			Columns: []string{risk.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(riskcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.LocationTable,
			Columns: []string{risk.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risklocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   risk.LocationTable,
			Columns: []string{risk.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risklocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Risk{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{risk.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}