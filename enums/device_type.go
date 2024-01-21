package enums

type DeviceType IntEnum

const (
	DTUnknown DeviceType = iota
	DT1
)

func (DeviceType) GetAll() []DeviceType {
	return []DeviceType{
		DTUnknown, DT1,
	}
}

func (p DeviceType) Label() string {
	switch p {
	case DT1:
		return "网络摄像机"
	default:
		return "Unknown"
	}
}

type DTLabeledEnum struct {
	Label string     `json:"label"`
	Value DeviceType `json:"value"`
}

func (p DeviceType) GetLabeledEnumList() []DTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]DTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = DTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
