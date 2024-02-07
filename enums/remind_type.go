package enums

type RemindType int

const (
	RMDTUnknown RemindType = iota
	RMDT1
	RMDT2
	RMDT3
)

func (RemindType) GetAll() []RemindType {
	return []RemindType{
		RMDTUnknown,
		RMDT1,
		RMDT2,
		RMDT3,
	}
}

func (p RemindType) Label() string {
	switch p {
	case RMDT1:
		return "公众号"
	case RMDT2:
		return "短信"
	default:
		return "Status Unknown"
	}
}

type RMDTLabeledEnum struct {
	Label string     `json:"label"`
	Value RemindType `json:"value"`
}

func (p RemindType) GetLabeledEnumList() []RMDTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]RMDTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = RMDTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
