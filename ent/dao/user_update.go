// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/user"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetDeleteTime sets the "delete_time" field.
func (uu *UserUpdate) SetDeleteTime(t time.Time) *UserUpdate {
	uu.mutation.SetDeleteTime(t)
	return uu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeleteTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeleteTime(*t)
	}
	return uu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uu *UserUpdate) ClearDeleteTime() *UserUpdate {
	uu.mutation.ClearDeleteTime()
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetUpdaterID sets the "updater_id" field.
func (uu *UserUpdate) SetUpdaterID(i int) *UserUpdate {
	uu.mutation.SetUpdaterID(i)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetRealName sets the "real_name" field.
func (uu *UserUpdate) SetRealName(s string) *UserUpdate {
	uu.mutation.SetRealName(s)
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMobile(s *string) *UserUpdate {
	if s != nil {
		uu.SetMobile(*s)
	}
	return uu
}

// ClearMobile clears the value of the "mobile" field.
func (uu *UserUpdate) ClearMobile() *UserUpdate {
	uu.mutation.ClearMobile()
	return uu
}

// SetWechatOpenid sets the "wechat_openid" field.
func (uu *UserUpdate) SetWechatOpenid(s string) *UserUpdate {
	uu.mutation.SetWechatOpenid(s)
	return uu
}

// SetNillableWechatOpenid sets the "wechat_openid" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWechatOpenid(s *string) *UserUpdate {
	if s != nil {
		uu.SetWechatOpenid(*s)
	}
	return uu
}

// ClearWechatOpenid clears the value of the "wechat_openid" field.
func (uu *UserUpdate) ClearWechatOpenid() *UserUpdate {
	uu.mutation.ClearWechatOpenid()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(ti types.UploadedImage) *UserUpdate {
	uu.mutation.SetAvatar(ti)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(ti *types.UploadedImage) *UserUpdate {
	if ti != nil {
		uu.SetAvatar(*ti)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetUserStatus sets the "user_status" field.
func (uu *UserUpdate) SetUserStatus(es enums.EnabledStatus) *UserUpdate {
	uu.mutation.ResetUserStatus()
	uu.mutation.SetUserStatus(es)
	return uu
}

// SetNillableUserStatus sets the "user_status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserStatus(es *enums.EnabledStatus) *UserUpdate {
	if es != nil {
		uu.SetUserStatus(*es)
	}
	return uu
}

// AddUserStatus adds es to the "user_status" field.
func (uu *UserUpdate) AddUserStatus(es enums.EnabledStatus) *UserUpdate {
	uu.mutation.AddUserStatus(es)
	return uu
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (uu *UserUpdate) SetUpdater(a *Admin) *UserUpdate {
	return uu.SetUpdaterID(a.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (uu *UserUpdate) ClearUpdater() *UserUpdate {
	uu.mutation.ClearUpdater()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		if user.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("dao: uninitialized user.UpdateDefaultUpdateTime (forgotten import dao/runtime?)")
		}
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.UpdaterID(); ok {
		if err := user.UpdaterIDValidator(v); err != nil {
			return &ValidationError{Name: "updater_id", err: fmt.Errorf(`dao: validator failed for field "User.updater_id": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`dao: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`dao: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`dao: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.RealName(); ok {
		if err := user.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`dao: validator failed for field "User.real_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Mobile(); ok {
		if err := user.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`dao: validator failed for field "User.mobile": %w`, err)}
		}
	}
	if v, ok := uu.mutation.WechatOpenid(); ok {
		if err := user.WechatOpenidValidator(v); err != nil {
			return &ValidationError{Name: "wechat_openid", err: fmt.Errorf(`dao: validator failed for field "User.wechat_openid": %w`, err)}
		}
	}
	if v, ok := uu.mutation.UserStatus(); ok {
		if err := user.UserStatusValidator(int(v)); err != nil {
			return &ValidationError{Name: "user_status", err: fmt.Errorf(`dao: validator failed for field "User.user_status": %w`, err)}
		}
	}
	if _, ok := uu.mutation.UpdaterID(); uu.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "User.updater"`)
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeTime, value)
	}
	if uu.mutation.DeleteTimeCleared() {
		_spec.ClearField(user.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uu.mutation.RealName(); ok {
		_spec.SetField(user.FieldRealName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uu.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if value, ok := uu.mutation.WechatOpenid(); ok {
		_spec.SetField(user.FieldWechatOpenid, field.TypeString, value)
	}
	if uu.mutation.WechatOpenidCleared() {
		_spec.ClearField(user.FieldWechatOpenid, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeJSON, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeJSON)
	}
	if value, ok := uu.mutation.UserStatus(); ok {
		_spec.SetField(user.FieldUserStatus, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedUserStatus(); ok {
		_spec.AddField(user.FieldUserStatus, field.TypeInt, value)
	}
	if uu.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UpdaterTable,
			Columns: []string{user.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UpdaterTable,
			Columns: []string{user.UpdaterColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetDeleteTime sets the "delete_time" field.
func (uuo *UserUpdateOne) SetDeleteTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeleteTime(t)
	return uuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeleteTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeleteTime(*t)
	}
	return uuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uuo *UserUpdateOne) ClearDeleteTime() *UserUpdateOne {
	uuo.mutation.ClearDeleteTime()
	return uuo
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetUpdaterID sets the "updater_id" field.
func (uuo *UserUpdateOne) SetUpdaterID(i int) *UserUpdateOne {
	uuo.mutation.SetUpdaterID(i)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetRealName sets the "real_name" field.
func (uuo *UserUpdateOne) SetRealName(s string) *UserUpdateOne {
	uuo.mutation.SetRealName(s)
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMobile(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMobile(*s)
	}
	return uuo
}

// ClearMobile clears the value of the "mobile" field.
func (uuo *UserUpdateOne) ClearMobile() *UserUpdateOne {
	uuo.mutation.ClearMobile()
	return uuo
}

// SetWechatOpenid sets the "wechat_openid" field.
func (uuo *UserUpdateOne) SetWechatOpenid(s string) *UserUpdateOne {
	uuo.mutation.SetWechatOpenid(s)
	return uuo
}

// SetNillableWechatOpenid sets the "wechat_openid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWechatOpenid(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetWechatOpenid(*s)
	}
	return uuo
}

// ClearWechatOpenid clears the value of the "wechat_openid" field.
func (uuo *UserUpdateOne) ClearWechatOpenid() *UserUpdateOne {
	uuo.mutation.ClearWechatOpenid()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(ti types.UploadedImage) *UserUpdateOne {
	uuo.mutation.SetAvatar(ti)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(ti *types.UploadedImage) *UserUpdateOne {
	if ti != nil {
		uuo.SetAvatar(*ti)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetUserStatus sets the "user_status" field.
func (uuo *UserUpdateOne) SetUserStatus(es enums.EnabledStatus) *UserUpdateOne {
	uuo.mutation.ResetUserStatus()
	uuo.mutation.SetUserStatus(es)
	return uuo
}

// SetNillableUserStatus sets the "user_status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserStatus(es *enums.EnabledStatus) *UserUpdateOne {
	if es != nil {
		uuo.SetUserStatus(*es)
	}
	return uuo
}

// AddUserStatus adds es to the "user_status" field.
func (uuo *UserUpdateOne) AddUserStatus(es enums.EnabledStatus) *UserUpdateOne {
	uuo.mutation.AddUserStatus(es)
	return uuo
}

// SetUpdater sets the "updater" edge to the Admin entity.
func (uuo *UserUpdateOne) SetUpdater(a *Admin) *UserUpdateOne {
	return uuo.SetUpdaterID(a.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUpdater clears the "updater" edge to the Admin entity.
func (uuo *UserUpdateOne) ClearUpdater() *UserUpdateOne {
	uuo.mutation.ClearUpdater()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		if user.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("dao: uninitialized user.UpdateDefaultUpdateTime (forgotten import dao/runtime?)")
		}
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.UpdaterID(); ok {
		if err := user.UpdaterIDValidator(v); err != nil {
			return &ValidationError{Name: "updater_id", err: fmt.Errorf(`dao: validator failed for field "User.updater_id": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`dao: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`dao: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`dao: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.RealName(); ok {
		if err := user.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`dao: validator failed for field "User.real_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Mobile(); ok {
		if err := user.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`dao: validator failed for field "User.mobile": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.WechatOpenid(); ok {
		if err := user.WechatOpenidValidator(v); err != nil {
			return &ValidationError{Name: "wechat_openid", err: fmt.Errorf(`dao: validator failed for field "User.wechat_openid": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.UserStatus(); ok {
		if err := user.UserStatusValidator(int(v)); err != nil {
			return &ValidationError{Name: "user_status", err: fmt.Errorf(`dao: validator failed for field "User.user_status": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.UpdaterID(); uuo.mutation.UpdaterCleared() && !ok {
		return errors.New(`dao: clearing a required unique edge "User.updater"`)
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`dao: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeTime, value)
	}
	if uuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(user.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.RealName(); ok {
		_spec.SetField(user.FieldRealName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uuo.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if value, ok := uuo.mutation.WechatOpenid(); ok {
		_spec.SetField(user.FieldWechatOpenid, field.TypeString, value)
	}
	if uuo.mutation.WechatOpenidCleared() {
		_spec.ClearField(user.FieldWechatOpenid, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeJSON, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeJSON)
	}
	if value, ok := uuo.mutation.UserStatus(); ok {
		_spec.SetField(user.FieldUserStatus, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedUserStatus(); ok {
		_spec.AddField(user.FieldUserStatus, field.TypeInt, value)
	}
	if uuo.mutation.UpdaterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UpdaterTable,
			Columns: []string{user.UpdaterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UpdaterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UpdaterTable,
			Columns: []string{user.UpdaterColumn},
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}