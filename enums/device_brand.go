package enums

type DeviceBrand IntEnum

const (
	DBUnknown DeviceBrand = iota
	DB1
)

func (DeviceBrand) GetAll() []DeviceBrand {
	return []DeviceBrand{
		DBUnknown, DB1,
	}
}

func (p DeviceBrand) Label() string {
	switch p {
	case DB1:
		return "图为"
	default:
		return "Unknown"
	}
}

type DBLabeledEnum struct {
	Label string      `json:"label"`
	Value DeviceBrand `json:"value"`
}

func (p DeviceBrand) GetLabeledEnumList() []DBLabeledEnum {
	all := p.GetAll()
	labeledList := make([]DBLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = DBLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
