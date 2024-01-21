package enums

type DeviceModel IntEnum

const (
	DMUnknown DeviceModel = iota
	DM1
)

func (DeviceModel) GetAll() []DeviceModel {
	return []DeviceModel{
		DMUnknown, DM1,
	}
}

func (p DeviceModel) Label() string {
	switch p {
	case DM1:
		return "图为TW-T1206"
	default:
		return "Unknown"
	}
}

type DMLabeledEnum struct {
	Label string      `json:"label"`
	Value DeviceModel `json:"value"`
}

func (p DeviceModel) GetLabeledEnumList() []DMLabeledEnum {
	all := p.GetAll()
	labeledList := make([]DMLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = DMLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
