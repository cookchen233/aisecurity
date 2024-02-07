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
	ET24
	ET25
	ET26
	ET27
	ET28
	ET29
	ET30
	ET31
	ET32
	ET33
)

func (EventType) GetAll() []EventType {
	return []EventType{
		ETUnknown, ET1, ET2, ET3, ET4, ET5, ET6, ET7, ET8, ET9, ET10, ET11, ET12, ET13, ET14,
		ET15, ET16, ET17, ET18, ET19, ET20, ET21, ET22, ET23, ET24, ET25, ET26, ET27, ET28,
		ET29, ET30, ET31, ET32, ET33,
	}
}

func (p EventType) Label() string {
	switch p {
	case ET1:
		return "明烟明火"
	case ET2:
		return "人脸识别"
	case ET3:
		return "车辆禁停"
	case ET4:
		return "车牌识别"
	case ET5:
		return "小动物检测"
	case ET6:
		return "未戴口罩"
	case ET7:
		return "人数超员"
	case ET8:
		return "人数监控"
	case ET9:
		return "离岗"
	case ET10:
		return "越线"
	case ET11:
		return "区域入侵"
	case ET12:
		return "人员徘徊"
	case ET13:
		return "翻越围栏"
	case ET14:
		return "攀爬"
	case ET15:
		return "客流计数"
	case ET16:
		return "打架"
	case ET17:
		return "电动车识别"
	case ET18:
		return "未佩戴安全帽"
	case ET19:
		return "未穿反光衣"
	case ET20:
		return "未穿长袖"
	case ET21:
		return "未穿工服"
	case ET22:
		return "倒地"
	case ET23:
		return "抽烟"
	case ET24:
		return "打电话"
	case ET25:
		return "通道占用"
	case ET26:
		return "未戴安全帽"
	case ET27:
		return "倒地V2"
	case ET28:
		return "睡岗"
	case ET29:
		return "玩手机"
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

func (p EventType) GetEventTypeByFuyouType(label string) EventType {
	eventType, exists := TwoTypeToEventTypeMap[label]
	if !exists {
		return ETUnknown
	}
	return eventType
}

var TwoTypeToEventTypeMap = map[string]EventType{
	"Playing":          ET29,
	"Smoking":          ET23,
	"Calling":          ET24,
	"CurrentPeopleNum": ET8,
	"NoVest":           ET19,
	"NoHelmet":         ET18,
	"NoMask":           ET6,
	"Falldown":         ET22,
	"NoSleeve":         ET20,
	"Occupied":         ET25,
	// Add other labels here
}

var FuyouTypeToEventTypeMap = map[string]EventType{
	"Playing":          ET29,
	"Smoking":          ET23,
	"Calling":          ET24,
	"CurrentPeopleNum": ET8,
	"NoVest":           ET19,
	"NoHelmet":         ET18,
	"NoMask":           ET6,
	"Falldown":         ET22,
	"NoSleeve":         ET20,
	"Occupied":         ET25,
	// Add other labels here
}
