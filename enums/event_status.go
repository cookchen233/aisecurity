package enums

type EventStatus int

const (
	ESUnknown EventStatus = iota
	ESPending
	ESProcessing
	ESProcessed
)

func (EventStatus) GetAll() []EventStatus {
	return []EventStatus{
		ESUnknown, ESPending, ESProcessing, ESProcessed,
	}
}

func (p EventStatus) Label() string {
	switch p {
	case ESPending:
		return "待处理"
	case ESProcessing:
		return "处理中"
	case ESProcessed:
		return "已处理"
	default:
		return "Unknown"
	}
}

type ESLabeledEnum struct {
	Label string      `json:"label"`
	Value EventStatus `json:"value"`
}

func (p EventStatus) GetLabeledEnumList() []ESLabeledEnum {
	all := p.GetAll()
	labeledList := make([]ESLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = ESLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
