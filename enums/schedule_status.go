package enums

type ScheduleStatus IntEnum

const (
	SSUnknown ScheduleStatus = iota
	SS1
	SS2
)

func (ScheduleStatus) GetAll() []ScheduleStatus {
	return []ScheduleStatus{
		SSUnknown, SS1, SS2,
	}
}

func (p ScheduleStatus) Label() string {
	switch p {
	case SS1:
		return "正常"
	case SS2:
		return "停用"
	default:
		return "Unknown"
	}
}

type SSLabeledEnum struct {
	Label string         `json:"label"`
	Value ScheduleStatus `json:"value"`
}

func (p ScheduleStatus) GetLabeledEnumList() []SSLabeledEnum {
	all := p.GetAll()
	labeledList := make([]SSLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = SSLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
