package enums

type CheckInType int

const (
	CITUnknown CheckInType = iota
	CIT1
	CIT2
	CIT3
)

func (CheckInType) GetAll() []CheckInType {
	return []CheckInType{
		CITUnknown,
		CIT1,
		CIT2,
		CIT3,
	}
}

func (p CheckInType) Label() string {
	switch p {
	case CIT1:
		return "系统打卡"
	case CIT2:
		return "无需签到"
	default:
		return "Status Unknown"
	}
}

type CITLabeledEnum struct {
	Label string      `json:"label"`
	Value CheckInType `json:"value"`
}

func (p CheckInType) GetLabeledEnumList() []CITLabeledEnum {
	all := p.GetAll()
	labeledList := make([]CITLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = CITLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
