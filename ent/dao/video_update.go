// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/ipcreportevent"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/video"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoUpdate is the builder for updating Video entities.
type VideoUpdate struct {
	config
	hooks    []Hook
	mutation *VideoMutation
}

// Where appends a list predicates to the VideoUpdate builder.
func (vu *VideoUpdate) Where(ps ...predicate.Video) *VideoUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetDeletedAt sets the "deleted_at" field.
func (vu *VideoUpdate) SetDeletedAt(t time.Time) *VideoUpdate {
	vu.mutation.SetDeletedAt(t)
	return vu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableDeletedAt(t *time.Time) *VideoUpdate {
	if t != nil {
		vu.SetDeletedAt(*t)
	}
	return vu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (vu *VideoUpdate) ClearDeletedAt() *VideoUpdate {
	vu.mutation.ClearDeletedAt()
	return vu
}

// SetUpdatedBy sets the "updated_by" field.
func (vu *VideoUpdate) SetUpdatedBy(i int) *VideoUpdate {
	vu.mutation.SetUpdatedBy(i)
	return vu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableUpdatedBy(i *int) *VideoUpdate {
	if i != nil {
		vu.SetUpdatedBy(*i)
	}
	return vu
}

// SetUpdatedAt sets the "updated_at" field.
func (vu *VideoUpdate) SetUpdatedAt(t time.Time) *VideoUpdate {
	vu.mutation.SetUpdatedAt(t)
	return vu
}

// SetName sets the "name" field.
func (vu *VideoUpdate) SetName(s string) *VideoUpdate {
	vu.mutation.SetName(s)
	return vu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableName(s *string) *VideoUpdate {
	if s != nil {
		vu.SetName(*s)
	}
	return vu
}

// ClearName clears the value of the "name" field.
func (vu *VideoUpdate) ClearName() *VideoUpdate {
	vu.mutation.ClearName()
	return vu
}

// SetURL sets the "url" field.
func (vu *VideoUpdate) SetURL(s string) *VideoUpdate {
	vu.mutation.SetURL(s)
	return vu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableURL(s *string) *VideoUpdate {
	if s != nil {
		vu.SetURL(*s)
	}
	return vu
}

// ClearURL clears the value of the "url" field.
func (vu *VideoUpdate) ClearURL() *VideoUpdate {
	vu.mutation.ClearURL()
	return vu
}

// SetSize sets the "size" field.
func (vu *VideoUpdate) SetSize(i int64) *VideoUpdate {
	vu.mutation.ResetSize()
	vu.mutation.SetSize(i)
	return vu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableSize(i *int64) *VideoUpdate {
	if i != nil {
		vu.SetSize(*i)
	}
	return vu
}

// AddSize adds i to the "size" field.
func (vu *VideoUpdate) AddSize(i int64) *VideoUpdate {
	vu.mutation.AddSize(i)
	return vu
}

// SetDuration sets the "duration" field.
func (vu *VideoUpdate) SetDuration(s string) *VideoUpdate {
	vu.mutation.SetDuration(s)
	return vu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableDuration(s *string) *VideoUpdate {
	if s != nil {
		vu.SetDuration(*s)
	}
	return vu
}

// ClearDuration clears the value of the "duration" field.
func (vu *VideoUpdate) ClearDuration() *VideoUpdate {
	vu.mutation.ClearDuration()
	return vu
}

// SetUploadedAt sets the "uploaded_at" field.
func (vu *VideoUpdate) SetUploadedAt(t time.Time) *VideoUpdate {
	vu.mutation.SetUploadedAt(t)
	return vu
}

// SetNillableUploadedAt sets the "uploaded_at" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableUploadedAt(t *time.Time) *VideoUpdate {
	if t != nil {
		vu.SetUploadedAt(*t)
	}
	return vu
}

// ClearUploadedAt clears the value of the "uploaded_at" field.
func (vu *VideoUpdate) ClearUploadedAt() *VideoUpdate {
	vu.mutation.ClearUploadedAt()
	return vu
}

// SetUpdaterID sets the "updater" edge to the Admin entity by ID.
func (vu *VideoUpdate) SetUpdaterID(id int) *VideoUpdate {
	vu.mutation.SetUpdaterID(id)
	return vu
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (vu *VideoUpdate) SetUpdater(a *Admin) *VideoUpdate {
	return vu.SetUpdaterID(a.ID)
}

// AddIpcReportEventVideoIDs adds the "ipc_report_event_video" edge to the IPCReportEvent entity by IDs.
func (vu *VideoUpdate) AddIpcReportEventVideoIDs(ids ...int) *VideoUpdate {
	vu.mutation.AddIpcReportEventVideoIDs(ids...)
	return vu
}

// AddIpcReportEventVideo adds the "ipc_report_event_video" edges to the IPCReportEvent entity.
func (vu *VideoUpdate) AddIpcReportEventVideo(i ...*IPCReportEvent) *VideoUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return vu.AddIpcReportEventVideoIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vu *VideoUpdate) Mutation() *VideoMutation {
	return vu.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (vu *VideoUpdate) ClearUpdater() *VideoUpdate {
	vu.mutation.ClearUpdater()
	return vu
}

// ClearIpcReportEventVideo clears all "ipc_report_event_video" edges to the IPCReportEvent entity.
func (vu *VideoUpdate) ClearIpcReportEventVideo() *VideoUpdate {
	vu.mutation.ClearIpcReportEventVideo()
	return vu
}

// RemoveIpcReportEventVideoIDs removes the "ipc_report_event_video" edge to IPCReportEvent entities by IDs.
func (vu *VideoUpdate) RemoveIpcReportEventVideoIDs(ids ...int) *VideoUpdate {
	vu.mutation.RemoveIpcReportEventVideoIDs(ids...)
	return vu
}

// RemoveIpcReportEventVideo removes "ipc_report_event_video" edges to IPCReportEvent entities.
func (vu *VideoUpdate) RemoveIpcReportEventVideo(i ...*IPCReportEvent) *VideoUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return vu.RemoveIpcReportEventVideoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VideoUpdate) Save(ctx context.Context) (int, error) {
	if err := vu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VideoUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VideoUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VideoUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VideoUpdate) defaults() error {
	if _, ok := vu.mutation.UpdatedAt(); !ok {
		if video.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized video.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := video.UpdateDefaultUpdatedAt()
		vu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vu *VideoUpdate) check() error {
	if v, ok := vu.mutation.UpdatedBy(); ok {
		if err := video.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "Video.updated_by": %w`, err)}
		}
	}
	if _, ok := vu.mutation.CreatorID(); vu.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Video.creator"`)
	}
	if _, ok := vu.mutation.UpdaterID(); vu.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Video.updater"`)
	}
	return nil
}

func (vu *VideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.DeletedAt(); ok {
		_spec.SetField(video.FieldDeletedAt, field.TypeTime, value)
	}
	if vu.mutation.DeletedAtCleared() {
		_spec.ClearField(video.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := vu.mutation.UpdatedAt(); ok {
		_spec.SetField(video.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vu.mutation.Name(); ok {
		_spec.SetField(video.FieldName, field.TypeString, value)
	}
	if vu.mutation.NameCleared() {
		_spec.ClearField(video.FieldName, field.TypeString)
	}
	if value, ok := vu.mutation.URL(); ok {
		_spec.SetField(video.FieldURL, field.TypeString, value)
	}
	if vu.mutation.URLCleared() {
		_spec.ClearField(video.FieldURL, field.TypeString)
	}
	if value, ok := vu.mutation.Size(); ok {
		_spec.SetField(video.FieldSize, field.TypeInt64, value)
	}
	if value, ok := vu.mutation.AddedSize(); ok {
		_spec.AddField(video.FieldSize, field.TypeInt64, value)
	}
	if value, ok := vu.mutation.Duration(); ok {
		_spec.SetField(video.FieldDuration, field.TypeString, value)
	}
	if vu.mutation.DurationCleared() {
		_spec.ClearField(video.FieldDuration, field.TypeString)
	}
	if value, ok := vu.mutation.UploadedAt(); ok {
		_spec.SetField(video.FieldUploadedAt, field.TypeTime, value)
	}
	if vu.mutation.UploadedAtCleared() {
		_spec.ClearField(video.FieldUploadedAt, field.TypeTime)
	}
	if vu.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.UpdaterTable,
			Columns: []string{video.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.UpdaterTable,
			Columns: []string{video.UpdaterColumn},
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
	if vu.mutation.IpcReportEventVideoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.IpcReportEventVideoTable,
			Columns: []string{video.IpcReportEventVideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedIpcReportEventVideoIDs(); len(nodes) > 0 && !vu.mutation.IpcReportEventVideoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.IpcReportEventVideoTable,
			Columns: []string{video.IpcReportEventVideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.IpcReportEventVideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.IpcReportEventVideoTable,
			Columns: []string{video.IpcReportEventVideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VideoUpdateOne is the builder for updating a single Video entity.
type VideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VideoMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (vuo *VideoUpdateOne) SetDeletedAt(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetDeletedAt(t)
	return vuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableDeletedAt(t *time.Time) *VideoUpdateOne {
	if t != nil {
		vuo.SetDeletedAt(*t)
	}
	return vuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (vuo *VideoUpdateOne) ClearDeletedAt() *VideoUpdateOne {
	vuo.mutation.ClearDeletedAt()
	return vuo
}

// SetUpdatedBy sets the "updated_by" field.
func (vuo *VideoUpdateOne) SetUpdatedBy(i int) *VideoUpdateOne {
	vuo.mutation.SetUpdatedBy(i)
	return vuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableUpdatedBy(i *int) *VideoUpdateOne {
	if i != nil {
		vuo.SetUpdatedBy(*i)
	}
	return vuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vuo *VideoUpdateOne) SetUpdatedAt(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetUpdatedAt(t)
	return vuo
}

// SetName sets the "name" field.
func (vuo *VideoUpdateOne) SetName(s string) *VideoUpdateOne {
	vuo.mutation.SetName(s)
	return vuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableName(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetName(*s)
	}
	return vuo
}

// ClearName clears the value of the "name" field.
func (vuo *VideoUpdateOne) ClearName() *VideoUpdateOne {
	vuo.mutation.ClearName()
	return vuo
}

// SetURL sets the "url" field.
func (vuo *VideoUpdateOne) SetURL(s string) *VideoUpdateOne {
	vuo.mutation.SetURL(s)
	return vuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableURL(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetURL(*s)
	}
	return vuo
}

// ClearURL clears the value of the "url" field.
func (vuo *VideoUpdateOne) ClearURL() *VideoUpdateOne {
	vuo.mutation.ClearURL()
	return vuo
}

// SetSize sets the "size" field.
func (vuo *VideoUpdateOne) SetSize(i int64) *VideoUpdateOne {
	vuo.mutation.ResetSize()
	vuo.mutation.SetSize(i)
	return vuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableSize(i *int64) *VideoUpdateOne {
	if i != nil {
		vuo.SetSize(*i)
	}
	return vuo
}

// AddSize adds i to the "size" field.
func (vuo *VideoUpdateOne) AddSize(i int64) *VideoUpdateOne {
	vuo.mutation.AddSize(i)
	return vuo
}

// SetDuration sets the "duration" field.
func (vuo *VideoUpdateOne) SetDuration(s string) *VideoUpdateOne {
	vuo.mutation.SetDuration(s)
	return vuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableDuration(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetDuration(*s)
	}
	return vuo
}

// ClearDuration clears the value of the "duration" field.
func (vuo *VideoUpdateOne) ClearDuration() *VideoUpdateOne {
	vuo.mutation.ClearDuration()
	return vuo
}

// SetUploadedAt sets the "uploaded_at" field.
func (vuo *VideoUpdateOne) SetUploadedAt(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetUploadedAt(t)
	return vuo
}

// SetNillableUploadedAt sets the "uploaded_at" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableUploadedAt(t *time.Time) *VideoUpdateOne {
	if t != nil {
		vuo.SetUploadedAt(*t)
	}
	return vuo
}

// ClearUploadedAt clears the value of the "uploaded_at" field.
func (vuo *VideoUpdateOne) ClearUploadedAt() *VideoUpdateOne {
	vuo.mutation.ClearUploadedAt()
	return vuo
}

// SetUpdaterID sets the "updater" edge to the Admin entity by ID.
func (vuo *VideoUpdateOne) SetUpdaterID(id int) *VideoUpdateOne {
	vuo.mutation.SetUpdaterID(id)
	return vuo
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (vuo *VideoUpdateOne) SetUpdater(a *Admin) *VideoUpdateOne {
	return vuo.SetUpdaterID(a.ID)
}

// AddIpcReportEventVideoIDs adds the "ipc_report_event_video" edge to the IPCReportEvent entity by IDs.
func (vuo *VideoUpdateOne) AddIpcReportEventVideoIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.AddIpcReportEventVideoIDs(ids...)
	return vuo
}

// AddIpcReportEventVideo adds the "ipc_report_event_video" edges to the IPCReportEvent entity.
func (vuo *VideoUpdateOne) AddIpcReportEventVideo(i ...*IPCReportEvent) *VideoUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return vuo.AddIpcReportEventVideoIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vuo *VideoUpdateOne) Mutation() *VideoMutation {
	return vuo.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (vuo *VideoUpdateOne) ClearUpdater() *VideoUpdateOne {
	vuo.mutation.ClearUpdater()
	return vuo
}

// ClearIpcReportEventVideo clears all "ipc_report_event_video" edges to the IPCReportEvent entity.
func (vuo *VideoUpdateOne) ClearIpcReportEventVideo() *VideoUpdateOne {
	vuo.mutation.ClearIpcReportEventVideo()
	return vuo
}

// RemoveIpcReportEventVideoIDs removes the "ipc_report_event_video" edge to IPCReportEvent entities by IDs.
func (vuo *VideoUpdateOne) RemoveIpcReportEventVideoIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.RemoveIpcReportEventVideoIDs(ids...)
	return vuo
}

// RemoveIpcReportEventVideo removes "ipc_report_event_video" edges to IPCReportEvent entities.
func (vuo *VideoUpdateOne) RemoveIpcReportEventVideo(i ...*IPCReportEvent) *VideoUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return vuo.RemoveIpcReportEventVideoIDs(ids...)
}

// Where appends a list predicates to the VideoUpdate builder.
func (vuo *VideoUpdateOne) Where(ps ...predicate.Video) *VideoUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VideoUpdateOne) Select(field string, fields ...string) *VideoUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Video entity.
func (vuo *VideoUpdateOne) Save(ctx context.Context) (*Video, error) {
	if err := vuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VideoUpdateOne) SaveX(ctx context.Context) *Video {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VideoUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VideoUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VideoUpdateOne) defaults() error {
	if _, ok := vuo.mutation.UpdatedAt(); !ok {
		if video.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized video.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := video.UpdateDefaultUpdatedAt()
		vuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VideoUpdateOne) check() error {
	if v, ok := vuo.mutation.UpdatedBy(); ok {
		if err := video.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "Video.updated_by": %w`, err)}
		}
	}
	if _, ok := vuo.mutation.CreatorID(); vuo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Video.creator"`)
	}
	if _, ok := vuo.mutation.UpdaterID(); vuo.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "Video.updater"`)
	}
	return nil
}

func (vuo *VideoUpdateOne) sqlSave(ctx context.Context) (_node *Video, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "Video.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for _, f := range fields {
			if !video.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.DeletedAt(); ok {
		_spec.SetField(video.FieldDeletedAt, field.TypeTime, value)
	}
	if vuo.mutation.DeletedAtCleared() {
		_spec.ClearField(video.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := vuo.mutation.UpdatedAt(); ok {
		_spec.SetField(video.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.Name(); ok {
		_spec.SetField(video.FieldName, field.TypeString, value)
	}
	if vuo.mutation.NameCleared() {
		_spec.ClearField(video.FieldName, field.TypeString)
	}
	if value, ok := vuo.mutation.URL(); ok {
		_spec.SetField(video.FieldURL, field.TypeString, value)
	}
	if vuo.mutation.URLCleared() {
		_spec.ClearField(video.FieldURL, field.TypeString)
	}
	if value, ok := vuo.mutation.Size(); ok {
		_spec.SetField(video.FieldSize, field.TypeInt64, value)
	}
	if value, ok := vuo.mutation.AddedSize(); ok {
		_spec.AddField(video.FieldSize, field.TypeInt64, value)
	}
	if value, ok := vuo.mutation.Duration(); ok {
		_spec.SetField(video.FieldDuration, field.TypeString, value)
	}
	if vuo.mutation.DurationCleared() {
		_spec.ClearField(video.FieldDuration, field.TypeString)
	}
	if value, ok := vuo.mutation.UploadedAt(); ok {
		_spec.SetField(video.FieldUploadedAt, field.TypeTime, value)
	}
	if vuo.mutation.UploadedAtCleared() {
		_spec.ClearField(video.FieldUploadedAt, field.TypeTime)
	}
	if vuo.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.UpdaterTable,
			Columns: []string{video.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.UpdaterTable,
			Columns: []string{video.UpdaterColumn},
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
	if vuo.mutation.IpcReportEventVideoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.IpcReportEventVideoTable,
			Columns: []string{video.IpcReportEventVideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedIpcReportEventVideoIDs(); len(nodes) > 0 && !vuo.mutation.IpcReportEventVideoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.IpcReportEventVideoTable,
			Columns: []string{video.IpcReportEventVideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.IpcReportEventVideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.IpcReportEventVideoTable,
			Columns: []string{video.IpcReportEventVideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipcreportevent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Video{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}