// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/video"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Video is the model entity for the Video schema.
type Video struct {
	config `json:"-"`
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
	// 名称
	Name string `json:"name"`
	// 文件地址
	URL string `json:"url"`
	// 文件大小
	Size int64 `json:"size"`
	// 视频时长
	Duration string `json:"duration"`
	// 上传时间
	UploadedAt *time.Time `json:"uploaded_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoQuery when eager-loading is set.
	Edges        VideoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VideoEdges holds the relations/edges for other nodes in the graph.
type VideoEdges struct {
	// Creator holds the value of the creator edge.
	Creator *Admin `json:"creator,omitempty"`
	// Updater holds the value of the updater edge.
	Updater *Admin `json:"updater,omitempty"`
	// IpcReportEventVideo holds the value of the ipc_report_event_video edge.
	IpcReportEventVideo []*IPCReportEvent `json:"ipc_report_event_video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VideoEdges) CreatorOrErr() (*Admin, error) {
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
func (e VideoEdges) UpdaterOrErr() (*Admin, error) {
	if e.loadedTypes[1] {
		if e.Updater == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Updater, nil
	}
	return nil, &NotLoadedError{edge: "updater"}
}

// IpcReportEventVideoOrErr returns the IpcReportEventVideo value or an error if the edge
// was not loaded in eager-loading.
func (e VideoEdges) IpcReportEventVideoOrErr() ([]*IPCReportEvent, error) {
	if e.loadedTypes[2] {
		return e.IpcReportEventVideo, nil
	}
	return nil, &NotLoadedError{edge: "ipc_report_event_video"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Video) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case video.FieldID, video.FieldCreatedBy, video.FieldUpdatedBy, video.FieldSize:
			values[i] = new(sql.NullInt64)
		case video.FieldName, video.FieldURL, video.FieldDuration:
			values[i] = new(sql.NullString)
		case video.FieldCreatedAt, video.FieldDeletedAt, video.FieldUpdatedAt, video.FieldUploadedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video fields.
func (v *Video) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case video.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case video.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case video.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				v.CreatedBy = int(value.Int64)
			}
		case video.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				v.DeletedAt = new(time.Time)
				*v.DeletedAt = value.Time
			}
		case video.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				v.UpdatedBy = int(value.Int64)
			}
		case video.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case video.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case video.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				v.URL = value.String
			}
		case video.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				v.Size = value.Int64
			}
		case video.FieldDuration:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				v.Duration = value.String
			}
		case video.FieldUploadedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field uploaded_at", values[i])
			} else if value.Valid {
				v.UploadedAt = new(time.Time)
				*v.UploadedAt = value.Time
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Video.
// This includes values selected through modifiers, order, etc.
func (v *Video) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the Video entity.
func (v *Video) QueryCreator() *AdminQuery {
	return NewVideoClient(v.config).QueryCreator(v)
}

// QueryUpdater queries the "updater" edge of the Video entity.
func (v *Video) QueryUpdater() *AdminQuery {
	return NewVideoClient(v.config).QueryUpdater(v)
}

// QueryIpcReportEventVideo queries the "ipc_report_event_video" edge of the Video entity.
func (v *Video) QueryIpcReportEventVideo() *IPCReportEventQuery {
	return NewVideoClient(v.config).QueryIpcReportEventVideo(v)
}

// Update returns a builder for updating this Video.
// Note that you need to call Video.Unwrap() before calling this method if this Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Video) Update() *VideoUpdateOne {
	return NewVideoClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Video entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Video) Unwrap() *Video {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("dao: Video is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Video) String() string {
	var builder strings.Builder
	builder.WriteString("Video(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", v.CreatedBy))
	builder.WriteString(", ")
	if v := v.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", v.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(v.Name)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(v.URL)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", v.Size))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(v.Duration)
	builder.WriteString(", ")
	if v := v.UploadedAt; v != nil {
		builder.WriteString("uploaded_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Videos is a parsable slice of Video.
type Videos []*Video