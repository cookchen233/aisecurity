package enums

type RepeatType int

const (
	RTUnknown RepeatType = iota
	RT1
	RT2
	RT3
	RT4
	RT5
	RT6
	RT7
)

func (RepeatType) GetAll() []RepeatType {
	return []RepeatType{
		RTUnknown,
		RT1,
		RT2,
		RT3,
		RT4,
		RT5,
		RT6,
		RT7,
	}
}

func (p RepeatType) Label() string {
	switch p {
	case RT1:
		return "不重复"
	case RT2:
		return "每日"
	case RT3:
		return "每周"
	case RT4:
		return "每月"
	case RT5:
		return "每隔?日"
	case RT6:
		return "每隔?周"
	case RT7:
		return "每隔?月"
	default:
		return "Status Unknown"
	}
}

type RTLabeledEnum struct {
	Label string     `json:"label"`
	Value RepeatType `json:"value"`
}

func (p RepeatType) GetLabeledEnumList() []RTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]RTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = RTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
