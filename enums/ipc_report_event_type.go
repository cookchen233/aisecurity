package enums

type IPCReportEventType IntEnum

const (
	IRETUnknown IPCReportEventType = iota
	IRET1
	IRET2
	IRET3
	IRET4
	IRET5
	IRET6
	IRET7
	IRET8
	IRET9
	IRET10
	IRET11
	IRET12
	IRET13
	IRET14
	IRET15
	IRET16
	IRET17
	IRET18
	IRET19
	IRET20
	IRET21
	IRET22
	IRET23
)

func (IPCReportEventType) GetAll() []IPCReportEventType {
	return []IPCReportEventType{
		IRETUnknown, IRET1, IRET2, IRET3, IRET4, IRET5, IRET6, IRET7, IRET8, IRET9, IRET10, IRET11, IRET12, IRET13, IRET14, IRET15, IRET16, IRET17, IRET18, IRET19, IRET20, IRET21, IRET22,
	}
}

func (p IPCReportEventType) Label() string {
	switch p {
	case IRET1:
		return "明烟明火"
	case IRET2:
		return "人脸识别"
	case IRET17:
		return "未穿工服"
	case IRET18:
		return "打电话"
	case IRET19:
		return "抽烟"
	case IRET20:
		return "驾驶员左顾右盼"
	default:
		return "Unknown"
	}
}

type IRETLabeledEnum struct {
	Label string             `json:"label"`
	Value IPCReportEventType `json:"value"`
}

func (p IPCReportEventType) GetLabeledEnumList() []IRETLabeledEnum {
	all := p.GetAll()
	labeledList := make([]IRETLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = IRETLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}

func (p IPCReportEventType) GetEventTypeByTwoType(label string) IPCReportEventType {
	eventType, exists := TwoTypeToEventTypeMap[label]
	if !exists {
		return IRETUnknown
	}
	return eventType
}

var TwoTypeToEventTypeMap = map[string]IPCReportEventType{
	"Calling": IRET18,
	"Smoking": IRET19,
	// Add other labels here
}
