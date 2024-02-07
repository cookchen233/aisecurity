package enums

type EventStatus int

const (
	ESUnknown EventStatus = iota
	ES1
	ES2
	ES3
)

func (EventStatus) GetAll() []EventStatus {
	return []EventStatus{
		ESUnknown, ES1, ES2, ES3,
	}
}

func (p EventStatus) Label() string {
	switch p {
	case ES1:
		return "待处理"
	case ES2:
		return "处理中"
	case ES3:
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
