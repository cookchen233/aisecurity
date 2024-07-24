package enums

type EnabledStatus IntEnum

const (
	ENSUnknown EnabledStatus = iota
	ENS1
	ENS2
)

func (EnabledStatus) GetAll() []EnabledStatus {
	return []EnabledStatus{
		ENSUnknown, ENS1, ENS2,
	}
}

func (p EnabledStatus) Label() string {
	switch p {
	case ENS1:
		return "启用"
	case ENS2:
		return "停用"
	default:
		return "Unknown"
	}
}

type ENSLabeledEnum struct {
	Label string        `json:"label"`
	Value EnabledStatus `json:"value"`
}

func (p EnabledStatus) GetLabeledEnumList() []ENSLabeledEnum {
	all := p.GetAll()
	labeledList := make([]ENSLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = ENSLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
