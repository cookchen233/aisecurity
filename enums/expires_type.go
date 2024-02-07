package enums

type ExpiresType int

const (
	EXTUnknown ExpiresType = iota
	EXT1
	EXT2
)

func (ExpiresType) GetAll() []ExpiresType {
	return []ExpiresType{
		EXTUnknown,
		EXT1,
		EXT2,
	}
}

func (p ExpiresType) Label() string {
	switch p {
	case EXT1:
		return "一直有效"
	case EXT2:
		return "截止到"
	default:
		return "Status Unknown"
	}
}

type EXTLabeledEnum struct {
	Label string      `json:"label"`
	Value ExpiresType `json:"value"`
}

func (p ExpiresType) GetLabeledEnumList() []EXTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]EXTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = EXTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
