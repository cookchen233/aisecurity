// Code generated by ent, DO NOT EDIT.

package event

import (
	"aisecurity/ent/dao/predicate"
	"aisecurity/enums"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreateTime, v))
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatorID, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDeleteTime, v))
}

// UpdaterID applies equality check predicate on the "updater_id" field. It's identical to UpdaterIDEQ.
func UpdaterID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdateTime, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDeviceID, v))
}

// VideoID applies equality check predicate on the "video_id" field. It's identical to VideoIDEQ.
func VideoID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldVideoID, v))
}

// EventTime applies equality check predicate on the "event_time" field. It's identical to EventTimeEQ.
func EventTime(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventTime, v))
}

// EventType applies equality check predicate on the "event_type" field. It's identical to EventTypeEQ.
func EventType(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldEQ(FieldEventType, vc))
}

// EventStatus applies equality check predicate on the "event_status" field. It's identical to EventStatusEQ.
func EventStatus(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldEQ(FieldEventStatus, vc))
}

// DataID applies equality check predicate on the "data_id" field. It's identical to DataIDEQ.
func DataID(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDataID, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// RawData applies equality check predicate on the "raw_data" field. It's identical to RawDataEQ.
func RawData(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldRawData, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreateTime, v))
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatorID, v))
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatorID, v))
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatorID, vs...))
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatorID, vs...))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDeleteTime))
}

// UpdaterIDEQ applies the EQ predicate on the "updater_id" field.
func UpdaterIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdaterID, v))
}

// UpdaterIDNEQ applies the NEQ predicate on the "updater_id" field.
func UpdaterIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUpdaterID, v))
}

// UpdaterIDIn applies the In predicate on the "updater_id" field.
func UpdaterIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUpdaterID, vs...))
}

// UpdaterIDNotIn applies the NotIn predicate on the "updater_id" field.
func UpdaterIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUpdaterID, vs...))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldUpdateTime, v))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDeviceID, vs...))
}

// VideoIDEQ applies the EQ predicate on the "video_id" field.
func VideoIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldVideoID, v))
}

// VideoIDNEQ applies the NEQ predicate on the "video_id" field.
func VideoIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldVideoID, v))
}

// VideoIDIn applies the In predicate on the "video_id" field.
func VideoIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldVideoID, vs...))
}

// VideoIDNotIn applies the NotIn predicate on the "video_id" field.
func VideoIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldVideoID, vs...))
}

// VideoIDIsNil applies the IsNil predicate on the "video_id" field.
func VideoIDIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldVideoID))
}

// VideoIDNotNil applies the NotNil predicate on the "video_id" field.
func VideoIDNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldVideoID))
}

// EventTimeEQ applies the EQ predicate on the "event_time" field.
func EventTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventTime, v))
}

// EventTimeNEQ applies the NEQ predicate on the "event_time" field.
func EventTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEventTime, v))
}

// EventTimeIn applies the In predicate on the "event_time" field.
func EventTimeIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEventTime, vs...))
}

// EventTimeNotIn applies the NotIn predicate on the "event_time" field.
func EventTimeNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEventTime, vs...))
}

// EventTimeGT applies the GT predicate on the "event_time" field.
func EventTimeGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEventTime, v))
}

// EventTimeGTE applies the GTE predicate on the "event_time" field.
func EventTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEventTime, v))
}

// EventTimeLT applies the LT predicate on the "event_time" field.
func EventTimeLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEventTime, v))
}

// EventTimeLTE applies the LTE predicate on the "event_time" field.
func EventTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEventTime, v))
}

// EventTypeEQ applies the EQ predicate on the "event_type" field.
func EventTypeEQ(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldEQ(FieldEventType, vc))
}

// EventTypeNEQ applies the NEQ predicate on the "event_type" field.
func EventTypeNEQ(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldNEQ(FieldEventType, vc))
}

// EventTypeIn applies the In predicate on the "event_type" field.
func EventTypeIn(vs ...enums.EventType) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Event(sql.FieldIn(FieldEventType, v...))
}

// EventTypeNotIn applies the NotIn predicate on the "event_type" field.
func EventTypeNotIn(vs ...enums.EventType) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Event(sql.FieldNotIn(FieldEventType, v...))
}

// EventTypeGT applies the GT predicate on the "event_type" field.
func EventTypeGT(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldGT(FieldEventType, vc))
}

// EventTypeGTE applies the GTE predicate on the "event_type" field.
func EventTypeGTE(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldGTE(FieldEventType, vc))
}

// EventTypeLT applies the LT predicate on the "event_type" field.
func EventTypeLT(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldLT(FieldEventType, vc))
}

// EventTypeLTE applies the LTE predicate on the "event_type" field.
func EventTypeLTE(v enums.EventType) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldLTE(FieldEventType, vc))
}

// EventStatusEQ applies the EQ predicate on the "event_status" field.
func EventStatusEQ(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldEQ(FieldEventStatus, vc))
}

// EventStatusNEQ applies the NEQ predicate on the "event_status" field.
func EventStatusNEQ(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldNEQ(FieldEventStatus, vc))
}

// EventStatusIn applies the In predicate on the "event_status" field.
func EventStatusIn(vs ...enums.EventStatus) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Event(sql.FieldIn(FieldEventStatus, v...))
}

// EventStatusNotIn applies the NotIn predicate on the "event_status" field.
func EventStatusNotIn(vs ...enums.EventStatus) predicate.Event {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Event(sql.FieldNotIn(FieldEventStatus, v...))
}

// EventStatusGT applies the GT predicate on the "event_status" field.
func EventStatusGT(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldGT(FieldEventStatus, vc))
}

// EventStatusGTE applies the GTE predicate on the "event_status" field.
func EventStatusGTE(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldGTE(FieldEventStatus, vc))
}

// EventStatusLT applies the LT predicate on the "event_status" field.
func EventStatusLT(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldLT(FieldEventStatus, vc))
}

// EventStatusLTE applies the LTE predicate on the "event_status" field.
func EventStatusLTE(v enums.EventStatus) predicate.Event {
	vc := int(v)
	return predicate.Event(sql.FieldLTE(FieldEventStatus, vc))
}

// ImagesIsNil applies the IsNil predicate on the "images" field.
func ImagesIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldImages))
}

// ImagesNotNil applies the NotNil predicate on the "images" field.
func ImagesNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldImages))
}

// LabeledImagesIsNil applies the IsNil predicate on the "labeled_images" field.
func LabeledImagesIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldLabeledImages))
}

// LabeledImagesNotNil applies the NotNil predicate on the "labeled_images" field.
func LabeledImagesNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldLabeledImages))
}

// DataIDEQ applies the EQ predicate on the "data_id" field.
func DataIDEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDataID, v))
}

// DataIDNEQ applies the NEQ predicate on the "data_id" field.
func DataIDNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDataID, v))
}

// DataIDIn applies the In predicate on the "data_id" field.
func DataIDIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDataID, vs...))
}

// DataIDNotIn applies the NotIn predicate on the "data_id" field.
func DataIDNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDataID, vs...))
}

// DataIDGT applies the GT predicate on the "data_id" field.
func DataIDGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDataID, v))
}

// DataIDGTE applies the GTE predicate on the "data_id" field.
func DataIDGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDataID, v))
}

// DataIDLT applies the LT predicate on the "data_id" field.
func DataIDLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDataID, v))
}

// DataIDLTE applies the LTE predicate on the "data_id" field.
func DataIDLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDataID, v))
}

// DataIDContains applies the Contains predicate on the "data_id" field.
func DataIDContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDataID, v))
}

// DataIDHasPrefix applies the HasPrefix predicate on the "data_id" field.
func DataIDHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDataID, v))
}

// DataIDHasSuffix applies the HasSuffix predicate on the "data_id" field.
func DataIDHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDataID, v))
}

// DataIDEqualFold applies the EqualFold predicate on the "data_id" field.
func DataIDEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDataID, v))
}

// DataIDContainsFold applies the ContainsFold predicate on the "data_id" field.
func DataIDContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDataID, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDescription, v))
}

// RawDataEQ applies the EQ predicate on the "raw_data" field.
func RawDataEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldRawData, v))
}

// RawDataNEQ applies the NEQ predicate on the "raw_data" field.
func RawDataNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldRawData, v))
}

// RawDataIn applies the In predicate on the "raw_data" field.
func RawDataIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldRawData, vs...))
}

// RawDataNotIn applies the NotIn predicate on the "raw_data" field.
func RawDataNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldRawData, vs...))
}

// RawDataGT applies the GT predicate on the "raw_data" field.
func RawDataGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldRawData, v))
}

// RawDataGTE applies the GTE predicate on the "raw_data" field.
func RawDataGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldRawData, v))
}

// RawDataLT applies the LT predicate on the "raw_data" field.
func RawDataLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldRawData, v))
}

// RawDataLTE applies the LTE predicate on the "raw_data" field.
func RawDataLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldRawData, v))
}

// RawDataContains applies the Contains predicate on the "raw_data" field.
func RawDataContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldRawData, v))
}

// RawDataHasPrefix applies the HasPrefix predicate on the "raw_data" field.
func RawDataHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldRawData, v))
}

// RawDataHasSuffix applies the HasSuffix predicate on the "raw_data" field.
func RawDataHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldRawData, v))
}

// RawDataIsNil applies the IsNil predicate on the "raw_data" field.
func RawDataIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldRawData))
}

// RawDataNotNil applies the NotNil predicate on the "raw_data" field.
func RawDataNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldRawData))
}

// RawDataEqualFold applies the EqualFold predicate on the "raw_data" field.
func RawDataEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldRawData, v))
}

// RawDataContainsFold applies the ContainsFold predicate on the "raw_data" field.
func RawDataContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldRawData, v))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.Admin) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdater applies the HasEdge predicate on the "updater" edge.
func HasUpdater() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdaterTable, UpdaterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdaterWith applies the HasEdge predicate on the "updater" edge with a given conditions (other predicates).
func HasUpdaterWith(preds ...predicate.Admin) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newUpdaterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newVideoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFixing applies the HasEdge predicate on the "fixing" edge.
func HasFixing() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FixingTable, FixingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixingWith applies the HasEdge predicate on the "fixing" edge with a given conditions (other predicates).
func HasFixingWith(preds ...predicate.Fixing) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newFixingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEventLog applies the HasEdge predicate on the "event_log" edge.
func HasEventLog() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventLogTable, EventLogColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventLogWith applies the HasEdge predicate on the "event_log" edge with a given conditions (other predicates).
func HasEventLogWith(preds ...predicate.EventLog) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newEventLogStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}