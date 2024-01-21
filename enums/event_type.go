package enums

type EventType IntEnum

const (
	ETUnknown EventType = iota
	ET1
	ET2
	ET3
	ET4
	ET5
	ET6
	ET7
	ET8
	ET9
	ET10
	ET11
	ET12
	ET13
	ET14
	ET15
	ET16
	ET17
	ET18
	ET19
	ET20
	ET21
	ET22
	ET23
)

func (EventType) GetAll() []EventType {
	return []EventType{
		ETUnknown, ET1, ET2, ET3, ET4, ET5, ET6, ET7, ET8, ET9, ET10, ET11, ET12, ET13, ET14, ET15, ET16, ET17, ET18, ET19, ET20, ET21, ET22,
	}
}

func (p EventType) Label() string {
	switch p {
	case ET1:
		return "明烟明火"
	case ET2:
		return "人脸识别"
	case ET17:
		return "未穿工服"
	case ET18:
		return "打电话"
	case ET19:
		return "抽烟"
	case ET20:
		return "驾驶员左顾右盼"
	default:
		return "Unknown"
	}
}

type ETLabeledEnum struct {
	Label string    `json:"label"`
	Value EventType `json:"value"`
}

func (p EventType) GetLabeledEnumList() []ETLabeledEnum {
	all := p.GetAll()
	labeledList := make([]ETLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = ETLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}

func (p EventType) GetEventTypeByTwoType(label string) EventType {
	eventType, exists := TwoTypeToEventTypeMap[label]
	if !exists {
		return ETUnknown
	}
	return eventType
}

var TwoTypeToEventTypeMap = map[string]EventType{
	"Calling": ET18,
	"Smoking": ET19,
	// Add other labels here
}
