// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/area"
	"aisecurity/ent/dao/device"
	"aisecurity/ent/dao/deviceinstallation"
	"aisecurity/ent/dao/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceInstallationUpdate is the builder for updating DeviceInstallation entities.
type DeviceInstallationUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceInstallationMutation
}

// Where appends a list predicates to the DeviceInstallationUpdate builder.
func (diu *DeviceInstallationUpdate) Where(ps ...predicate.DeviceInstallation) *DeviceInstallationUpdate {
	diu.mutation.Where(ps...)
	return diu
}

// SetDeletedAt sets the "deleted_at" field.
func (diu *DeviceInstallationUpdate) SetDeletedAt(t time.Time) *DeviceInstallationUpdate {
	diu.mutation.SetDeletedAt(t)
	return diu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableDeletedAt(t *time.Time) *DeviceInstallationUpdate {
	if t != nil {
		diu.SetDeletedAt(*t)
	}
	return diu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (diu *DeviceInstallationUpdate) ClearDeletedAt() *DeviceInstallationUpdate {
	diu.mutation.ClearDeletedAt()
	return diu
}

// SetUpdatedBy sets the "updated_by" field.
func (diu *DeviceInstallationUpdate) SetUpdatedBy(i int) *DeviceInstallationUpdate {
	diu.mutation.SetUpdatedBy(i)
	return diu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableUpdatedBy(i *int) *DeviceInstallationUpdate {
	if i != nil {
		diu.SetUpdatedBy(*i)
	}
	return diu
}

// SetUpdatedAt sets the "updated_at" field.
func (diu *DeviceInstallationUpdate) SetUpdatedAt(t time.Time) *DeviceInstallationUpdate {
	diu.mutation.SetUpdatedAt(t)
	return diu
}

// SetDeviceID sets the "device_id" field.
func (diu *DeviceInstallationUpdate) SetDeviceID(i int) *DeviceInstallationUpdate {
	diu.mutation.SetDeviceID(i)
	return diu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableDeviceID(i *int) *DeviceInstallationUpdate {
	if i != nil {
		diu.SetDeviceID(*i)
	}
	return diu
}

// SetAreaID sets the "area_id" field.
func (diu *DeviceInstallationUpdate) SetAreaID(i int) *DeviceInstallationUpdate {
	diu.mutation.SetAreaID(i)
	return diu
}

// SetNillableAreaID sets the "area_id" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableAreaID(i *int) *DeviceInstallationUpdate {
	if i != nil {
		diu.SetAreaID(*i)
	}
	return diu
}

// SetAliasName sets the "alias_name" field.
func (diu *DeviceInstallationUpdate) SetAliasName(s string) *DeviceInstallationUpdate {
	diu.mutation.SetAliasName(s)
	return diu
}

// SetNillableAliasName sets the "alias_name" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableAliasName(s *string) *DeviceInstallationUpdate {
	if s != nil {
		diu.SetAliasName(*s)
	}
	return diu
}

// ClearAliasName clears the value of the "alias_name" field.
func (diu *DeviceInstallationUpdate) ClearAliasName() *DeviceInstallationUpdate {
	diu.mutation.ClearAliasName()
	return diu
}

// SetLongitude sets the "longitude" field.
func (diu *DeviceInstallationUpdate) SetLongitude(f float64) *DeviceInstallationUpdate {
	diu.mutation.ResetLongitude()
	diu.mutation.SetLongitude(f)
	return diu
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableLongitude(f *float64) *DeviceInstallationUpdate {
	if f != nil {
		diu.SetLongitude(*f)
	}
	return diu
}

// AddLongitude adds f to the "longitude" field.
func (diu *DeviceInstallationUpdate) AddLongitude(f float64) *DeviceInstallationUpdate {
	diu.mutation.AddLongitude(f)
	return diu
}

// SetLatitude sets the "latitude" field.
func (diu *DeviceInstallationUpdate) SetLatitude(f float64) *DeviceInstallationUpdate {
	diu.mutation.ResetLatitude()
	diu.mutation.SetLatitude(f)
	return diu
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableLatitude(f *float64) *DeviceInstallationUpdate {
	if f != nil {
		diu.SetLatitude(*f)
	}
	return diu
}

// AddLatitude adds f to the "latitude" field.
func (diu *DeviceInstallationUpdate) AddLatitude(f float64) *DeviceInstallationUpdate {
	diu.mutation.AddLatitude(f)
	return diu
}

// SetLocationData sets the "location_data" field.
func (diu *DeviceInstallationUpdate) SetLocationData(s string) *DeviceInstallationUpdate {
	diu.mutation.SetLocationData(s)
	return diu
}

// SetNillableLocationData sets the "location_data" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableLocationData(s *string) *DeviceInstallationUpdate {
	if s != nil {
		diu.SetLocationData(*s)
	}
	return diu
}

// ClearLocationData clears the value of the "location_data" field.
func (diu *DeviceInstallationUpdate) ClearLocationData() *DeviceInstallationUpdate {
	diu.mutation.ClearLocationData()
	return diu
}

// SetLocation sets the "location" field.
func (diu *DeviceInstallationUpdate) SetLocation(s string) *DeviceInstallationUpdate {
	diu.mutation.SetLocation(s)
	return diu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableLocation(s *string) *DeviceInstallationUpdate {
	if s != nil {
		diu.SetLocation(*s)
	}
	return diu
}

// ClearLocation clears the value of the "location" field.
func (diu *DeviceInstallationUpdate) ClearLocation() *DeviceInstallationUpdate {
	diu.mutation.ClearLocation()
	return diu
}

// SetInstaller sets the "installer" field.
func (diu *DeviceInstallationUpdate) SetInstaller(s string) *DeviceInstallationUpdate {
	diu.mutation.SetInstaller(s)
	return diu
}

// SetNillableInstaller sets the "installer" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableInstaller(s *string) *DeviceInstallationUpdate {
	if s != nil {
		diu.SetInstaller(*s)
	}
	return diu
}

// ClearInstaller clears the value of the "installer" field.
func (diu *DeviceInstallationUpdate) ClearInstaller() *DeviceInstallationUpdate {
	diu.mutation.ClearInstaller()
	return diu
}

// SetInstallTime sets the "install_time" field.
func (diu *DeviceInstallationUpdate) SetInstallTime(t time.Time) *DeviceInstallationUpdate {
	diu.mutation.SetInstallTime(t)
	return diu
}

// SetNillableInstallTime sets the "install_time" field if the given value is not nil.
func (diu *DeviceInstallationUpdate) SetNillableInstallTime(t *time.Time) *DeviceInstallationUpdate {
	if t != nil {
		diu.SetInstallTime(*t)
	}
	return diu
}

// ClearInstallTime clears the value of the "install_time" field.
func (diu *DeviceInstallationUpdate) ClearInstallTime() *DeviceInstallationUpdate {
	diu.mutation.ClearInstallTime()
	return diu
}

// SetUpdaterID sets the "updater" edge to the Admin entity by ID.
func (diu *DeviceInstallationUpdate) SetUpdaterID(id int) *DeviceInstallationUpdate {
	diu.mutation.SetUpdaterID(id)
	return diu
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (diu *DeviceInstallationUpdate) SetUpdater(a *Admin) *DeviceInstallationUpdate {
	return diu.SetUpdaterID(a.ID)
}

// SetArea sets the "area" edge to the Area entity.
func (diu *DeviceInstallationUpdate) SetArea(a *Area) *DeviceInstallationUpdate {
	return diu.SetAreaID(a.ID)
}

// SetDevice sets the "device" edge to the Device entity.
func (diu *DeviceInstallationUpdate) SetDevice(d *Device) *DeviceInstallationUpdate {
	return diu.SetDeviceID(d.ID)
}

// Mutation returns the DeviceInstallationMutation object of the builder.
func (diu *DeviceInstallationUpdate) Mutation() *DeviceInstallationMutation {
	return diu.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (diu *DeviceInstallationUpdate) ClearUpdater() *DeviceInstallationUpdate {
	diu.mutation.ClearUpdater()
	return diu
}

// ClearArea clears the "area" edge to the Area entity.
func (diu *DeviceInstallationUpdate) ClearArea() *DeviceInstallationUpdate {
	diu.mutation.ClearArea()
	return diu
}

// ClearDevice clears the "device" edge to the Device entity.
func (diu *DeviceInstallationUpdate) ClearDevice() *DeviceInstallationUpdate {
	diu.mutation.ClearDevice()
	return diu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (diu *DeviceInstallationUpdate) Save(ctx context.Context) (int, error) {
	if err := diu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, diu.sqlSave, diu.mutation, diu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (diu *DeviceInstallationUpdate) SaveX(ctx context.Context) int {
	affected, err := diu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (diu *DeviceInstallationUpdate) Exec(ctx context.Context) error {
	_, err := diu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diu *DeviceInstallationUpdate) ExecX(ctx context.Context) {
	if err := diu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diu *DeviceInstallationUpdate) defaults() error {
	if _, ok := diu.mutation.UpdatedAt(); !ok {
		if deviceinstallation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized deviceinstallation.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := deviceinstallation.UpdateDefaultUpdatedAt()
		diu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (diu *DeviceInstallationUpdate) check() error {
	if v, ok := diu.mutation.UpdatedBy(); ok {
		if err := deviceinstallation.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.updated_by": %w`, err)}
		}
	}
	if v, ok := diu.mutation.DeviceID(); ok {
		if err := deviceinstallation.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.device_id": %w`, err)}
		}
	}
	if v, ok := diu.mutation.AreaID(); ok {
		if err := deviceinstallation.AreaIDValidator(v); err != nil {
			return &ValidationError{Name: "area_id", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.area_id": %w`, err)}
		}
	}
	if v, ok := diu.mutation.AliasName(); ok {
		if err := deviceinstallation.AliasNameValidator(v); err != nil {
			return &ValidationError{Name: "alias_name", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.alias_name": %w`, err)}
		}
	}
	if _, ok := diu.mutation.CreatorID(); diu.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.creator"`)
	}
	if _, ok := diu.mutation.UpdaterID(); diu.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.updater"`)
	}
	if _, ok := diu.mutation.AreaID(); diu.mutation.AreaCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.area"`)
	}
	if _, ok := diu.mutation.DeviceID(); diu.mutation.DeviceCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.device"`)
	}
	return nil
}

func (diu *DeviceInstallationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := diu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deviceinstallation.Table, deviceinstallation.Columns, sqlgraph.NewFieldSpec(deviceinstallation.FieldID, field.TypeInt))
	if ps := diu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diu.mutation.DeletedAt(); ok {
		_spec.SetField(deviceinstallation.FieldDeletedAt, field.TypeTime, value)
	}
	if diu.mutation.DeletedAtCleared() {
		_spec.ClearField(deviceinstallation.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := diu.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceinstallation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := diu.mutation.AliasName(); ok {
		_spec.SetField(deviceinstallation.FieldAliasName, field.TypeString, value)
	}
	if diu.mutation.AliasNameCleared() {
		_spec.ClearField(deviceinstallation.FieldAliasName, field.TypeString)
	}
	if value, ok := diu.mutation.Longitude(); ok {
		_spec.SetField(deviceinstallation.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := diu.mutation.AddedLongitude(); ok {
		_spec.AddField(deviceinstallation.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := diu.mutation.Latitude(); ok {
		_spec.SetField(deviceinstallation.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := diu.mutation.AddedLatitude(); ok {
		_spec.AddField(deviceinstallation.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := diu.mutation.LocationData(); ok {
		_spec.SetField(deviceinstallation.FieldLocationData, field.TypeString, value)
	}
	if diu.mutation.LocationDataCleared() {
		_spec.ClearField(deviceinstallation.FieldLocationData, field.TypeString)
	}
	if value, ok := diu.mutation.Location(); ok {
		_spec.SetField(deviceinstallation.FieldLocation, field.TypeString, value)
	}
	if diu.mutation.LocationCleared() {
		_spec.ClearField(deviceinstallation.FieldLocation, field.TypeString)
	}
	if value, ok := diu.mutation.Installer(); ok {
		_spec.SetField(deviceinstallation.FieldInstaller, field.TypeString, value)
	}
	if diu.mutation.InstallerCleared() {
		_spec.ClearField(deviceinstallation.FieldInstaller, field.TypeString)
	}
	if value, ok := diu.mutation.InstallTime(); ok {
		_spec.SetField(deviceinstallation.FieldInstallTime, field.TypeTime, value)
	}
	if diu.mutation.InstallTimeCleared() {
		_spec.ClearField(deviceinstallation.FieldInstallTime, field.TypeTime)
	}
	if diu.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.UpdaterTable,
			Columns: []string{deviceinstallation.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diu.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.UpdaterTable,
			Columns: []string{deviceinstallation.UpdaterColumn},
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
	if diu.mutation.AreaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.AreaTable,
			Columns: []string{deviceinstallation.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(area.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diu.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.AreaTable,
			Columns: []string{deviceinstallation.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(area.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if diu.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.DeviceTable,
			Columns: []string{deviceinstallation.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diu.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.DeviceTable,
			Columns: []string{deviceinstallation.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, diu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deviceinstallation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	diu.mutation.done = true
	return n, nil
}

// DeviceInstallationUpdateOne is the builder for updating a single DeviceInstallation entity.
type DeviceInstallationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceInstallationMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (diuo *DeviceInstallationUpdateOne) SetDeletedAt(t time.Time) *DeviceInstallationUpdateOne {
	diuo.mutation.SetDeletedAt(t)
	return diuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableDeletedAt(t *time.Time) *DeviceInstallationUpdateOne {
	if t != nil {
		diuo.SetDeletedAt(*t)
	}
	return diuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (diuo *DeviceInstallationUpdateOne) ClearDeletedAt() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearDeletedAt()
	return diuo
}

// SetUpdatedBy sets the "updated_by" field.
func (diuo *DeviceInstallationUpdateOne) SetUpdatedBy(i int) *DeviceInstallationUpdateOne {
	diuo.mutation.SetUpdatedBy(i)
	return diuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableUpdatedBy(i *int) *DeviceInstallationUpdateOne {
	if i != nil {
		diuo.SetUpdatedBy(*i)
	}
	return diuo
}

// SetUpdatedAt sets the "updated_at" field.
func (diuo *DeviceInstallationUpdateOne) SetUpdatedAt(t time.Time) *DeviceInstallationUpdateOne {
	diuo.mutation.SetUpdatedAt(t)
	return diuo
}

// SetDeviceID sets the "device_id" field.
func (diuo *DeviceInstallationUpdateOne) SetDeviceID(i int) *DeviceInstallationUpdateOne {
	diuo.mutation.SetDeviceID(i)
	return diuo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableDeviceID(i *int) *DeviceInstallationUpdateOne {
	if i != nil {
		diuo.SetDeviceID(*i)
	}
	return diuo
}

// SetAreaID sets the "area_id" field.
func (diuo *DeviceInstallationUpdateOne) SetAreaID(i int) *DeviceInstallationUpdateOne {
	diuo.mutation.SetAreaID(i)
	return diuo
}

// SetNillableAreaID sets the "area_id" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableAreaID(i *int) *DeviceInstallationUpdateOne {
	if i != nil {
		diuo.SetAreaID(*i)
	}
	return diuo
}

// SetAliasName sets the "alias_name" field.
func (diuo *DeviceInstallationUpdateOne) SetAliasName(s string) *DeviceInstallationUpdateOne {
	diuo.mutation.SetAliasName(s)
	return diuo
}

// SetNillableAliasName sets the "alias_name" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableAliasName(s *string) *DeviceInstallationUpdateOne {
	if s != nil {
		diuo.SetAliasName(*s)
	}
	return diuo
}

// ClearAliasName clears the value of the "alias_name" field.
func (diuo *DeviceInstallationUpdateOne) ClearAliasName() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearAliasName()
	return diuo
}

// SetLongitude sets the "longitude" field.
func (diuo *DeviceInstallationUpdateOne) SetLongitude(f float64) *DeviceInstallationUpdateOne {
	diuo.mutation.ResetLongitude()
	diuo.mutation.SetLongitude(f)
	return diuo
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableLongitude(f *float64) *DeviceInstallationUpdateOne {
	if f != nil {
		diuo.SetLongitude(*f)
	}
	return diuo
}

// AddLongitude adds f to the "longitude" field.
func (diuo *DeviceInstallationUpdateOne) AddLongitude(f float64) *DeviceInstallationUpdateOne {
	diuo.mutation.AddLongitude(f)
	return diuo
}

// SetLatitude sets the "latitude" field.
func (diuo *DeviceInstallationUpdateOne) SetLatitude(f float64) *DeviceInstallationUpdateOne {
	diuo.mutation.ResetLatitude()
	diuo.mutation.SetLatitude(f)
	return diuo
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableLatitude(f *float64) *DeviceInstallationUpdateOne {
	if f != nil {
		diuo.SetLatitude(*f)
	}
	return diuo
}

// AddLatitude adds f to the "latitude" field.
func (diuo *DeviceInstallationUpdateOne) AddLatitude(f float64) *DeviceInstallationUpdateOne {
	diuo.mutation.AddLatitude(f)
	return diuo
}

// SetLocationData sets the "location_data" field.
func (diuo *DeviceInstallationUpdateOne) SetLocationData(s string) *DeviceInstallationUpdateOne {
	diuo.mutation.SetLocationData(s)
	return diuo
}

// SetNillableLocationData sets the "location_data" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableLocationData(s *string) *DeviceInstallationUpdateOne {
	if s != nil {
		diuo.SetLocationData(*s)
	}
	return diuo
}

// ClearLocationData clears the value of the "location_data" field.
func (diuo *DeviceInstallationUpdateOne) ClearLocationData() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearLocationData()
	return diuo
}

// SetLocation sets the "location" field.
func (diuo *DeviceInstallationUpdateOne) SetLocation(s string) *DeviceInstallationUpdateOne {
	diuo.mutation.SetLocation(s)
	return diuo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableLocation(s *string) *DeviceInstallationUpdateOne {
	if s != nil {
		diuo.SetLocation(*s)
	}
	return diuo
}

// ClearLocation clears the value of the "location" field.
func (diuo *DeviceInstallationUpdateOne) ClearLocation() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearLocation()
	return diuo
}

// SetInstaller sets the "installer" field.
func (diuo *DeviceInstallationUpdateOne) SetInstaller(s string) *DeviceInstallationUpdateOne {
	diuo.mutation.SetInstaller(s)
	return diuo
}

// SetNillableInstaller sets the "installer" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableInstaller(s *string) *DeviceInstallationUpdateOne {
	if s != nil {
		diuo.SetInstaller(*s)
	}
	return diuo
}

// ClearInstaller clears the value of the "installer" field.
func (diuo *DeviceInstallationUpdateOne) ClearInstaller() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearInstaller()
	return diuo
}

// SetInstallTime sets the "install_time" field.
func (diuo *DeviceInstallationUpdateOne) SetInstallTime(t time.Time) *DeviceInstallationUpdateOne {
	diuo.mutation.SetInstallTime(t)
	return diuo
}

// SetNillableInstallTime sets the "install_time" field if the given value is not nil.
func (diuo *DeviceInstallationUpdateOne) SetNillableInstallTime(t *time.Time) *DeviceInstallationUpdateOne {
	if t != nil {
		diuo.SetInstallTime(*t)
	}
	return diuo
}

// ClearInstallTime clears the value of the "install_time" field.
func (diuo *DeviceInstallationUpdateOne) ClearInstallTime() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearInstallTime()
	return diuo
}

// SetUpdaterID sets the "updater" edge to the Admin entity by ID.
func (diuo *DeviceInstallationUpdateOne) SetUpdaterID(id int) *DeviceInstallationUpdateOne {
	diuo.mutation.SetUpdaterID(id)
	return diuo
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (diuo *DeviceInstallationUpdateOne) SetUpdater(a *Admin) *DeviceInstallationUpdateOne {
	return diuo.SetUpdaterID(a.ID)
}

// SetArea sets the "area" edge to the Area entity.
func (diuo *DeviceInstallationUpdateOne) SetArea(a *Area) *DeviceInstallationUpdateOne {
	return diuo.SetAreaID(a.ID)
}

// SetDevice sets the "device" edge to the Device entity.
func (diuo *DeviceInstallationUpdateOne) SetDevice(d *Device) *DeviceInstallationUpdateOne {
	return diuo.SetDeviceID(d.ID)
}

// Mutation returns the DeviceInstallationMutation object of the builder.
func (diuo *DeviceInstallationUpdateOne) Mutation() *DeviceInstallationMutation {
	return diuo.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (diuo *DeviceInstallationUpdateOne) ClearUpdater() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearUpdater()
	return diuo
}

// ClearArea clears the "area" edge to the Area entity.
func (diuo *DeviceInstallationUpdateOne) ClearArea() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearArea()
	return diuo
}

// ClearDevice clears the "device" edge to the Device entity.
func (diuo *DeviceInstallationUpdateOne) ClearDevice() *DeviceInstallationUpdateOne {
	diuo.mutation.ClearDevice()
	return diuo
}

// Where appends a list predicates to the DeviceInstallationUpdate builder.
func (diuo *DeviceInstallationUpdateOne) Where(ps ...predicate.DeviceInstallation) *DeviceInstallationUpdateOne {
	diuo.mutation.Where(ps...)
	return diuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (diuo *DeviceInstallationUpdateOne) Select(field string, fields ...string) *DeviceInstallationUpdateOne {
	diuo.fields = append([]string{field}, fields...)
	return diuo
}

// Save executes the query and returns the updated DeviceInstallation entity.
func (diuo *DeviceInstallationUpdateOne) Save(ctx context.Context) (*DeviceInstallation, error) {
	if err := diuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, diuo.sqlSave, diuo.mutation, diuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (diuo *DeviceInstallationUpdateOne) SaveX(ctx context.Context) *DeviceInstallation {
	node, err := diuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (diuo *DeviceInstallationUpdateOne) Exec(ctx context.Context) error {
	_, err := diuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diuo *DeviceInstallationUpdateOne) ExecX(ctx context.Context) {
	if err := diuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diuo *DeviceInstallationUpdateOne) defaults() error {
	if _, ok := diuo.mutation.UpdatedAt(); !ok {
		if deviceinstallation.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("dao: uninitialized deviceinstallation.UpdateDefaultUpdatedAt (forgotten import dao/runtime?)")
		}
		v := deviceinstallation.UpdateDefaultUpdatedAt()
		diuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (diuo *DeviceInstallationUpdateOne) check() error {
	if v, ok := diuo.mutation.UpdatedBy(); ok {
		if err := deviceinstallation.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.updated_by": %w`, err)}
		}
	}
	if v, ok := diuo.mutation.DeviceID(); ok {
		if err := deviceinstallation.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.device_id": %w`, err)}
		}
	}
	if v, ok := diuo.mutation.AreaID(); ok {
		if err := deviceinstallation.AreaIDValidator(v); err != nil {
			return &ValidationError{Name: "area_id", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.area_id": %w`, err)}
		}
	}
	if v, ok := diuo.mutation.AliasName(); ok {
		if err := deviceinstallation.AliasNameValidator(v); err != nil {
			return &ValidationError{Name: "alias_name", err: fmt.Errorf(`dao: validator failed for field "DeviceInstallation.alias_name": %w`, err)}
		}
	}
	if _, ok := diuo.mutation.CreatorID(); diuo.mutation.CreatorCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.creator"`)
	}
	if _, ok := diuo.mutation.UpdaterID(); diuo.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.updater"`)
	}
	if _, ok := diuo.mutation.AreaID(); diuo.mutation.AreaCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.area"`)
	}
	if _, ok := diuo.mutation.DeviceID(); diuo.mutation.DeviceCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "DeviceInstallation.device"`)
	}
	return nil
}

func (diuo *DeviceInstallationUpdateOne) sqlSave(ctx context.Context) (_node *DeviceInstallation, err error) {
	if err := diuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deviceinstallation.Table, deviceinstallation.Columns, sqlgraph.NewFieldSpec(deviceinstallation.FieldID, field.TypeInt))
	id, ok := diuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "DeviceInstallation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := diuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deviceinstallation.FieldID)
		for _, f := range fields {
			if !deviceinstallation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != deviceinstallation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := diuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diuo.mutation.DeletedAt(); ok {
		_spec.SetField(deviceinstallation.FieldDeletedAt, field.TypeTime, value)
	}
	if diuo.mutation.DeletedAtCleared() {
		_spec.ClearField(deviceinstallation.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := diuo.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceinstallation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := diuo.mutation.AliasName(); ok {
		_spec.SetField(deviceinstallation.FieldAliasName, field.TypeString, value)
	}
	if diuo.mutation.AliasNameCleared() {
		_spec.ClearField(deviceinstallation.FieldAliasName, field.TypeString)
	}
	if value, ok := diuo.mutation.Longitude(); ok {
		_spec.SetField(deviceinstallation.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := diuo.mutation.AddedLongitude(); ok {
		_spec.AddField(deviceinstallation.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := diuo.mutation.Latitude(); ok {
		_spec.SetField(deviceinstallation.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := diuo.mutation.AddedLatitude(); ok {
		_spec.AddField(deviceinstallation.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := diuo.mutation.LocationData(); ok {
		_spec.SetField(deviceinstallation.FieldLocationData, field.TypeString, value)
	}
	if diuo.mutation.LocationDataCleared() {
		_spec.ClearField(deviceinstallation.FieldLocationData, field.TypeString)
	}
	if value, ok := diuo.mutation.Location(); ok {
		_spec.SetField(deviceinstallation.FieldLocation, field.TypeString, value)
	}
	if diuo.mutation.LocationCleared() {
		_spec.ClearField(deviceinstallation.FieldLocation, field.TypeString)
	}
	if value, ok := diuo.mutation.Installer(); ok {
		_spec.SetField(deviceinstallation.FieldInstaller, field.TypeString, value)
	}
	if diuo.mutation.InstallerCleared() {
		_spec.ClearField(deviceinstallation.FieldInstaller, field.TypeString)
	}
	if value, ok := diuo.mutation.InstallTime(); ok {
		_spec.SetField(deviceinstallation.FieldInstallTime, field.TypeTime, value)
	}
	if diuo.mutation.InstallTimeCleared() {
		_spec.ClearField(deviceinstallation.FieldInstallTime, field.TypeTime)
	}
	if diuo.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.UpdaterTable,
			Columns: []string{deviceinstallation.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diuo.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.UpdaterTable,
			Columns: []string{deviceinstallation.UpdaterColumn},
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
	if diuo.mutation.AreaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.AreaTable,
			Columns: []string{deviceinstallation.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(area.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diuo.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.AreaTable,
			Columns: []string{deviceinstallation.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(area.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if diuo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.DeviceTable,
			Columns: []string{deviceinstallation.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diuo.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceinstallation.DeviceTable,
			Columns: []string{deviceinstallation.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DeviceInstallation{config: diuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, diuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deviceinstallation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	diuo.mutation.done = true
	return _node, nil
}