// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/device"
	"aisecurity/ent/dao/ipcevent"
	"aisecurity/ent/dao/video"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// IPCEvent is the model entity for the IPCEvent schema.
type IPCEvent struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建者
	CreatedBy int `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 最后更新者
	UpdatedBy int `json:"updated_by"`
	// 最后更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// 设备ID
	DeviceID int `json:"device_id" validate:"required"`
	// 视频ID
	VideoID int `json:"video_id"`
	// 事件发生时间
	EventTime time.Time `json:"event_time" validate:"required"`
	// 事件类型
	EventType enums.EventType `json:"event_type" validate:"required"`
	// 事件状态
	EventStatus enums.EventStatus `json:"event_status" validate:"required"`
	// 图片
	Images []*types.UploadedImage `json:"images" validate:"required"`
	// 标记的图片
	LabeledImages []*types.UploadedImage `json:"labeled_images" validate:"required"`
	// 事件ID
	EventID string `json:"event_id" validate:"required"`
	// 描述
	Description string `json:"description"`
	// 设备商原始上报数据
	RawData string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IPCEventQuery when eager-loading is set.
	Edges        IPCEventEdges `json:"edges"`
	selectValues sql.SelectValues
}

// IPCEventEdges holds the relations/edges for other nodes in the graph.
type IPCEventEdges struct {
	// Creator holds the value of the creator edge.
	Creator *Admin `json:"creator,omitempty"`
	// Updater holds the value of the updater edge.
	Updater *Admin `json:"updater,omitempty"`
	// Video holds the value of the video edge.
	Video *Video `json:"video,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// Fixers holds the value of the fixers edge.
	Fixers []*Employee `json:"fixers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IPCEventEdges) CreatorOrErr() (*Admin, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// UpdaterOrErr returns the Updater value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IPCEventEdges) UpdaterOrErr() (*Admin, error) {
	if e.loadedTypes[1] {
		if e.Updater == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Updater, nil
	}
	return nil, &NotLoadedError{edge: "updater"}
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IPCEventEdges) VideoOrErr() (*Video, error) {
	if e.loadedTypes[2] {
		if e.Video == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: video.Label}
		}
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IPCEventEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[3] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// FixersOrErr returns the Fixers value or an error if the edge
// was not loaded in eager-loading.
func (e IPCEventEdges) FixersOrErr() ([]*Employee, error) {
	if e.loadedTypes[4] {
		return e.Fixers, nil
	}
	return nil, &NotLoadedError{edge: "fixers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IPCEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ipcevent.FieldImages, ipcevent.FieldLabeledImages:
			values[i] = new([]byte)
		case ipcevent.FieldID, ipcevent.FieldCreatedBy, ipcevent.FieldUpdatedBy, ipcevent.FieldDeviceID, ipcevent.FieldVideoID, ipcevent.FieldEventType, ipcevent.FieldEventStatus:
			values[i] = new(sql.NullInt64)
		case ipcevent.FieldEventID, ipcevent.FieldDescription, ipcevent.FieldRawData:
			values[i] = new(sql.NullString)
		case ipcevent.FieldCreatedAt, ipcevent.FieldDeletedAt, ipcevent.FieldUpdatedAt, ipcevent.FieldEventTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IPCEvent fields.
func (ie *IPCEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ipcevent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ie.ID = int(value.Int64)
		case ipcevent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ie.CreatedAt = value.Time
			}
		case ipcevent.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ie.CreatedBy = int(value.Int64)
			}
		case ipcevent.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ie.DeletedAt = new(time.Time)
				*ie.DeletedAt = value.Time
			}
		case ipcevent.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ie.UpdatedBy = int(value.Int64)
			}
		case ipcevent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ie.UpdatedAt = value.Time
			}
		case ipcevent.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				ie.DeviceID = int(value.Int64)
			}
		case ipcevent.FieldVideoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				ie.VideoID = int(value.Int64)
			}
		case ipcevent.FieldEventTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field event_time", values[i])
			} else if value.Valid {
				ie.EventTime = value.Time
			}
		case ipcevent.FieldEventType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field event_type", values[i])
			} else if value.Valid {
				ie.EventType = enums.EventType(value.Int64)
			}
		case ipcevent.FieldEventStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field event_status", values[i])
			} else if value.Valid {
				ie.EventStatus = enums.EventStatus(value.Int64)
			}
		case ipcevent.FieldImages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ie.Images); err != nil {
					return fmt.Errorf("unmarshal field images: %w", err)
				}
			}
		case ipcevent.FieldLabeledImages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labeled_images", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ie.LabeledImages); err != nil {
					return fmt.Errorf("unmarshal field labeled_images: %w", err)
				}
			}
		case ipcevent.FieldEventID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value.Valid {
				ie.EventID = value.String
			}
		case ipcevent.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ie.Description = value.String
			}
		case ipcevent.FieldRawData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw_data", values[i])
			} else if value.Valid {
				ie.RawData = value.String
			}
		default:
			ie.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IPCEvent.
// This includes values selected through modifiers, order, etc.
func (ie *IPCEvent) Value(name string) (ent.Value, error) {
	return ie.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the IPCEvent entity.
func (ie *IPCEvent) QueryCreator() *AdminQuery {
	return NewIPCEventClient(ie.config).QueryCreator(ie)
}

// QueryUpdater queries the "updater" edge of the IPCEvent entity.
func (ie *IPCEvent) QueryUpdater() *AdminQuery {
	return NewIPCEventClient(ie.config).QueryUpdater(ie)
}

// QueryVideo queries the "video" edge of the IPCEvent entity.
func (ie *IPCEvent) QueryVideo() *VideoQuery {
	return NewIPCEventClient(ie.config).QueryVideo(ie)
}

// QueryDevice queries the "device" edge of the IPCEvent entity.
func (ie *IPCEvent) QueryDevice() *DeviceQuery {
	return NewIPCEventClient(ie.config).QueryDevice(ie)
}

// QueryFixers queries the "fixers" edge of the IPCEvent entity.
func (ie *IPCEvent) QueryFixers() *EmployeeQuery {
	return NewIPCEventClient(ie.config).QueryFixers(ie)
}

// Update returns a builder for updating this IPCEvent.
// Note that you need to call IPCEvent.Unwrap() before calling this method if this IPCEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (ie *IPCEvent) Update() *IPCEventUpdateOne {
	return NewIPCEventClient(ie.config).UpdateOne(ie)
}

// Unwrap unwraps the IPCEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ie *IPCEvent) Unwrap() *IPCEvent {
	_tx, ok := ie.config.driver.(*txDriver)
	if !ok {
		panic("dao: IPCEvent is not a transactional entity")
	}
	ie.config.driver = _tx.drv
	return ie
}

// String implements the fmt.Stringer.
func (ie *IPCEvent) String() string {
	var builder strings.Builder
	builder.WriteString("IPCEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ie.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ie.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ie.CreatedBy))
	builder.WriteString(", ")
	if v := ie.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ie.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ie.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", ie.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("video_id=")
	builder.WriteString(fmt.Sprintf("%v", ie.VideoID))
	builder.WriteString(", ")
	builder.WriteString("event_time=")
	builder.WriteString(ie.EventTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("event_type=")
	builder.WriteString(fmt.Sprintf("%v", ie.EventType))
	builder.WriteString(", ")
	builder.WriteString("event_status=")
	builder.WriteString(fmt.Sprintf("%v", ie.EventStatus))
	builder.WriteString(", ")
	builder.WriteString("images=")
	builder.WriteString(fmt.Sprintf("%v", ie.Images))
	builder.WriteString(", ")
	builder.WriteString("labeled_images=")
	builder.WriteString(fmt.Sprintf("%v", ie.LabeledImages))
	builder.WriteString(", ")
	builder.WriteString("event_id=")
	builder.WriteString(ie.EventID)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ie.Description)
	builder.WriteString(", ")
	builder.WriteString("raw_data=")
	builder.WriteString(ie.RawData)
	builder.WriteByte(')')
	return builder.String()
}

// IPCEvents is a parsable slice of IPCEvent.
type IPCEvents []*IPCEvent