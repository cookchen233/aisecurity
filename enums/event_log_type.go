package enums

type EventLogType IntEnum

const (
	ELTUnknown EventLogType = iota
	ELT1
	ELT2
	ELT3
	ELT4
	ELT5
)

func (EventLogType) GetAll() []EventLogType {
	return []EventLogType{
		ELTUnknown, ELT1, ELT2, ELT3, ELT4, ELT5,
	}
}

func (p EventLogType) Label() string {
	switch p {
	case ELT1:
		return "告警上报"
	case ELT2:
		return "分派处理人"
	case ELT3:
		return "认领处理"
	case ELT4:
		return "标记已完成"
	case ELT5:
		return "重新分派处理人"
	default:
		return "Unknown"
	}
}

type ELTLabeledEnum struct {
	Label string       `json:"label"`
	Value EventLogType `json:"value"`
}

func (p EventLogType) GetLabeledEnumList() []ELTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]ELTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = ELTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
