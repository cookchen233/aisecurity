package enums

type AdminStatus IntEnum

const (
	ASUnknown AdminStatus = iota
	AS1
	AS2
)

func (AdminStatus) GetAll() []AdminStatus {
	return []AdminStatus{
		ASUnknown, AS1, AS2,
	}
}

func (p AdminStatus) Label() string {
	switch p {
	case AS1:
		return "正常"
	case AS2:
		return "禁用"
	default:
		return "Unknown"
	}
}

type ASLabeledEnum struct {
	Label string      `json:"label"`
	Value AdminStatus `json:"value"`
}

func (p AdminStatus) GetLabeledEnumList() []ASLabeledEnum {
	all := p.GetAll()
	labeledList := make([]ASLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = ASLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
