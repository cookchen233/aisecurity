// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/sweep"
	"aisecurity/ent/dao/sweepschedule"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SweepSchedule is the model entity for the SweepSchedule schema.
type SweepSchedule struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time"`
	// 创建者
	CreatorID int `json:"creator_id"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time"`
	// 最后更新者
	UpdaterID int `json:"updater_id"`
	// 最后更新时间
	UpdateTime time.Time `json:"update_time"`
	// 名称
	Name string `json:"name" validate:"required"`
	// 排查ID
	SweepID int `json:"sweep_id"`
	// 任务状态
	ScheduleStatus enums.AdminStatus `json:"schedule_status"`
	// 任务开始时间
	ActionTime time.Time `json:"action_time"`
	// 提醒
	Remind types.ScheduleRemind `json:"remind"`
	// 重复
	Repeat types.ScheduleRepeat `json:"repeat" validate:"required"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SweepScheduleQuery when eager-loading is set.
	Edges        SweepScheduleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SweepScheduleEdges holds the relations/edges for other nodes in the graph.
type SweepScheduleEdges struct {
	// Creator holds the value of the creator edge.
	Creator *Admin `json:"creator,omitempty"`
	// Updater holds the value of the updater edge.
	Updater *Admin `json:"updater,omitempty"`
	// Sweep holds the value of the sweep edge.
	Sweep *Sweep `json:"sweep,omitempty"`
	// Workers holds the value of the workers edge.
	Workers []*Admin `json:"workers,omitempty"`
	// SweepResult holds the value of the sweep_result edge.
	SweepResult []*SweepResult `json:"sweep_result,omitempty"`
	// SweepResultDetails holds the value of the sweep_result_details edge.
	SweepResultDetails []*SweepResultDetails `json:"sweep_result_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SweepScheduleEdges) CreatorOrErr() (*Admin, error) {
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
func (e SweepScheduleEdges) UpdaterOrErr() (*Admin, error) {
	if e.loadedTypes[1] {
		if e.Updater == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Updater, nil
	}
	return nil, &NotLoadedError{edge: "updater"}
}

// SweepOrErr returns the Sweep value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SweepScheduleEdges) SweepOrErr() (*Sweep, error) {
	if e.loadedTypes[2] {
		if e.Sweep == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sweep.Label}
		}
		return e.Sweep, nil
	}
	return nil, &NotLoadedError{edge: "sweep"}
}

// WorkersOrErr returns the Workers value or an error if the edge
// was not loaded in eager-loading.
func (e SweepScheduleEdges) WorkersOrErr() ([]*Admin, error) {
	if e.loadedTypes[3] {
		return e.Workers, nil
	}
	return nil, &NotLoadedError{edge: "workers"}
}

// SweepResultOrErr returns the SweepResult value or an error if the edge
// was not loaded in eager-loading.
func (e SweepScheduleEdges) SweepResultOrErr() ([]*SweepResult, error) {
	if e.loadedTypes[4] {
		return e.SweepResult, nil
	}
	return nil, &NotLoadedError{edge: "sweep_result"}
}

// SweepResultDetailsOrErr returns the SweepResultDetails value or an error if the edge
// was not loaded in eager-loading.
func (e SweepScheduleEdges) SweepResultDetailsOrErr() ([]*SweepResultDetails, error) {
	if e.loadedTypes[5] {
		return e.SweepResultDetails, nil
	}
	return nil, &NotLoadedError{edge: "sweep_result_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SweepSchedule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sweepschedule.FieldRemind, sweepschedule.FieldRepeat:
			values[i] = new([]byte)
		case sweepschedule.FieldID, sweepschedule.FieldCreatorID, sweepschedule.FieldUpdaterID, sweepschedule.FieldSweepID, sweepschedule.FieldScheduleStatus:
			values[i] = new(sql.NullInt64)
		case sweepschedule.FieldName:
			values[i] = new(sql.NullString)
		case sweepschedule.FieldCreateTime, sweepschedule.FieldDeleteTime, sweepschedule.FieldUpdateTime, sweepschedule.FieldActionTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SweepSchedule fields.
func (ss *SweepSchedule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sweepschedule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ss.ID = int(value.Int64)
		case sweepschedule.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ss.CreateTime = value.Time
			}
		case sweepschedule.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				ss.CreatorID = int(value.Int64)
			}
		case sweepschedule.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				ss.DeleteTime = new(time.Time)
				*ss.DeleteTime = value.Time
			}
		case sweepschedule.FieldUpdaterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updater_id", values[i])
			} else if value.Valid {
				ss.UpdaterID = int(value.Int64)
			}
		case sweepschedule.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ss.UpdateTime = value.Time
			}
		case sweepschedule.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ss.Name = value.String
			}
		case sweepschedule.FieldSweepID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sweep_id", values[i])
			} else if value.Valid {
				ss.SweepID = int(value.Int64)
			}
		case sweepschedule.FieldScheduleStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_status", values[i])
			} else if value.Valid {
				ss.ScheduleStatus = enums.AdminStatus(value.Int64)
			}
		case sweepschedule.FieldActionTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field action_time", values[i])
			} else if value.Valid {
				ss.ActionTime = value.Time
			}
		case sweepschedule.FieldRemind:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field remind", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ss.Remind); err != nil {
					return fmt.Errorf("unmarshal field remind: %w", err)
				}
			}
		case sweepschedule.FieldRepeat:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field repeat", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ss.Repeat); err != nil {
					return fmt.Errorf("unmarshal field repeat: %w", err)
				}
			}
		default:
			ss.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SweepSchedule.
// This includes values selected through modifiers, order, etc.
func (ss *SweepSchedule) Value(name string) (ent.Value, error) {
	return ss.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the SweepSchedule entity.
func (ss *SweepSchedule) QueryCreator() *AdminQuery {
	return NewSweepScheduleClient(ss.config).QueryCreator(ss)
}

// QueryUpdater queries the "updater" edge of the SweepSchedule entity.
func (ss *SweepSchedule) QueryUpdater() *AdminQuery {
	return NewSweepScheduleClient(ss.config).QueryUpdater(ss)
}

// QuerySweep queries the "sweep" edge of the SweepSchedule entity.
func (ss *SweepSchedule) QuerySweep() *SweepQuery {
	return NewSweepScheduleClient(ss.config).QuerySweep(ss)
}

// QueryWorkers queries the "workers" edge of the SweepSchedule entity.
func (ss *SweepSchedule) QueryWorkers() *AdminQuery {
	return NewSweepScheduleClient(ss.config).QueryWorkers(ss)
}

// QuerySweepResult queries the "sweep_result" edge of the SweepSchedule entity.
func (ss *SweepSchedule) QuerySweepResult() *SweepResultQuery {
	return NewSweepScheduleClient(ss.config).QuerySweepResult(ss)
}

// QuerySweepResultDetails queries the "sweep_result_details" edge of the SweepSchedule entity.
func (ss *SweepSchedule) QuerySweepResultDetails() *SweepResultDetailsQuery {
	return NewSweepScheduleClient(ss.config).QuerySweepResultDetails(ss)
}

// Update returns a builder for updating this SweepSchedule.
// Note that you need to call SweepSchedule.Unwrap() before calling this method if this SweepSchedule
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *SweepSchedule) Update() *SweepScheduleUpdateOne {
	return NewSweepScheduleClient(ss.config).UpdateOne(ss)
}

// Unwrap unwraps the SweepSchedule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ss *SweepSchedule) Unwrap() *SweepSchedule {
	_tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("dao: SweepSchedule is not a transactional entity")
	}
	ss.config.driver = _tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *SweepSchedule) String() string {
	var builder strings.Builder
	builder.WriteString("SweepSchedule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ss.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ss.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.CreatorID))
	builder.WriteString(", ")
	if v := ss.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("updater_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.UpdaterID))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ss.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ss.Name)
	builder.WriteString(", ")
	builder.WriteString("sweep_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.SweepID))
	builder.WriteString(", ")
	builder.WriteString("schedule_status=")
	builder.WriteString(fmt.Sprintf("%v", ss.ScheduleStatus))
	builder.WriteString(", ")
	builder.WriteString("action_time=")
	builder.WriteString(ss.ActionTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("remind=")
	builder.WriteString(fmt.Sprintf("%v", ss.Remind))
	builder.WriteString(", ")
	builder.WriteString("repeat=")
	builder.WriteString(fmt.Sprintf("%v", ss.Repeat))
	builder.WriteByte(')')
	return builder.String()
}

// SweepSchedules is a parsable slice of SweepSchedule.
type SweepSchedules []*SweepSchedule