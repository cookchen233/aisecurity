// Code generated by ent, DO NOT EDIT.

package user

import (
	"aisecurity/ent/dao/predicate"
	"aisecurity/enums"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleteTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdaterID applies equality check predicate on the "updater_id" field. It's identical to UpdaterIDEQ.
func UpdaterID(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdaterID, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// RealName applies equality check predicate on the "real_name" field. It's identical to RealNameEQ.
func RealName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRealName, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// WechatOpenid applies equality check predicate on the "wechat_openid" field. It's identical to WechatOpenidEQ.
func WechatOpenid(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWechatOpenid, v))
}

// UserStatus applies equality check predicate on the "user_status" field. It's identical to UserStatusEQ.
func UserStatus(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldEQ(FieldUserStatus, vc))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDeleteTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdaterIDEQ applies the EQ predicate on the "updater_id" field.
func UpdaterIDEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdaterIDNEQ applies the NEQ predicate on the "updater_id" field.
func UpdaterIDNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdaterID, v))
}

// UpdaterIDIn applies the In predicate on the "updater_id" field.
func UpdaterIDIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdaterID, vs...))
}

// UpdaterIDNotIn applies the NotIn predicate on the "updater_id" field.
func UpdaterIDNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdaterID, vs...))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickname, v))
}

// RealNameEQ applies the EQ predicate on the "real_name" field.
func RealNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRealName, v))
}

// RealNameNEQ applies the NEQ predicate on the "real_name" field.
func RealNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRealName, v))
}

// RealNameIn applies the In predicate on the "real_name" field.
func RealNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRealName, vs...))
}

// RealNameNotIn applies the NotIn predicate on the "real_name" field.
func RealNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRealName, vs...))
}

// RealNameGT applies the GT predicate on the "real_name" field.
func RealNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRealName, v))
}

// RealNameGTE applies the GTE predicate on the "real_name" field.
func RealNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRealName, v))
}

// RealNameLT applies the LT predicate on the "real_name" field.
func RealNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRealName, v))
}

// RealNameLTE applies the LTE predicate on the "real_name" field.
func RealNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRealName, v))
}

// RealNameContains applies the Contains predicate on the "real_name" field.
func RealNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRealName, v))
}

// RealNameHasPrefix applies the HasPrefix predicate on the "real_name" field.
func RealNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRealName, v))
}

// RealNameHasSuffix applies the HasSuffix predicate on the "real_name" field.
func RealNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRealName, v))
}

// RealNameEqualFold applies the EqualFold predicate on the "real_name" field.
func RealNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRealName, v))
}

// RealNameContainsFold applies the ContainsFold predicate on the "real_name" field.
func RealNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRealName, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMobile, v))
}

// WechatOpenidEQ applies the EQ predicate on the "wechat_openid" field.
func WechatOpenidEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWechatOpenid, v))
}

// WechatOpenidNEQ applies the NEQ predicate on the "wechat_openid" field.
func WechatOpenidNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldWechatOpenid, v))
}

// WechatOpenidIn applies the In predicate on the "wechat_openid" field.
func WechatOpenidIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldWechatOpenid, vs...))
}

// WechatOpenidNotIn applies the NotIn predicate on the "wechat_openid" field.
func WechatOpenidNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldWechatOpenid, vs...))
}

// WechatOpenidGT applies the GT predicate on the "wechat_openid" field.
func WechatOpenidGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldWechatOpenid, v))
}

// WechatOpenidGTE applies the GTE predicate on the "wechat_openid" field.
func WechatOpenidGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldWechatOpenid, v))
}

// WechatOpenidLT applies the LT predicate on the "wechat_openid" field.
func WechatOpenidLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldWechatOpenid, v))
}

// WechatOpenidLTE applies the LTE predicate on the "wechat_openid" field.
func WechatOpenidLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldWechatOpenid, v))
}

// WechatOpenidContains applies the Contains predicate on the "wechat_openid" field.
func WechatOpenidContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldWechatOpenid, v))
}

// WechatOpenidHasPrefix applies the HasPrefix predicate on the "wechat_openid" field.
func WechatOpenidHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldWechatOpenid, v))
}

// WechatOpenidHasSuffix applies the HasSuffix predicate on the "wechat_openid" field.
func WechatOpenidHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldWechatOpenid, v))
}

// WechatOpenidIsNil applies the IsNil predicate on the "wechat_openid" field.
func WechatOpenidIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldWechatOpenid))
}

// WechatOpenidNotNil applies the NotNil predicate on the "wechat_openid" field.
func WechatOpenidNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldWechatOpenid))
}

// WechatOpenidEqualFold applies the EqualFold predicate on the "wechat_openid" field.
func WechatOpenidEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldWechatOpenid, v))
}

// WechatOpenidContainsFold applies the ContainsFold predicate on the "wechat_openid" field.
func WechatOpenidContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldWechatOpenid, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatar))
}

// UserStatusEQ applies the EQ predicate on the "user_status" field.
func UserStatusEQ(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldEQ(FieldUserStatus, vc))
}

// UserStatusNEQ applies the NEQ predicate on the "user_status" field.
func UserStatusNEQ(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldNEQ(FieldUserStatus, vc))
}

// UserStatusIn applies the In predicate on the "user_status" field.
func UserStatusIn(vs ...enums.EnabledStatus) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.User(sql.FieldIn(FieldUserStatus, v...))
}

// UserStatusNotIn applies the NotIn predicate on the "user_status" field.
func UserStatusNotIn(vs ...enums.EnabledStatus) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.User(sql.FieldNotIn(FieldUserStatus, v...))
}

// UserStatusGT applies the GT predicate on the "user_status" field.
func UserStatusGT(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldGT(FieldUserStatus, vc))
}

// UserStatusGTE applies the GTE predicate on the "user_status" field.
func UserStatusGTE(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldGTE(FieldUserStatus, vc))
}

// UserStatusLT applies the LT predicate on the "user_status" field.
func UserStatusLT(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldLT(FieldUserStatus, vc))
}

// UserStatusLTE applies the LTE predicate on the "user_status" field.
func UserStatusLTE(v enums.EnabledStatus) predicate.User {
	vc := int(v)
	return predicate.User(sql.FieldLTE(FieldUserStatus, vc))
}

// HasUpdater applies the HasEdge predicate on the "updater" edge.
func HasUpdater() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdaterWith applies the HasEdge predicate on the "updater" edge with a given conditions (other predicates).
func HasUpdaterWith(preds ...predicate.Admin) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUpdaterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}