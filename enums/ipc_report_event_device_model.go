package enums

type IPCReportEventDeviceModel IntEnum

const (
	IREDMUnknown IPCReportEventDeviceModel = iota
	IREDM1
)

func (IPCReportEventDeviceModel) GetAll() []IPCReportEventDeviceModel {
	return []IPCReportEventDeviceModel{
		IREDMUnknown, IREDM1,
	}
}

func (p IPCReportEventDeviceModel) Label() string {
	switch p {
	case IREDM1:
		return "图为TW-T1206"
	default:
		return "Unknown"
	}
}

type IREDMLabeledEnum struct {
	Label string                    `json:"label"`
	Value IPCReportEventDeviceModel `json:"value"`
}

func (p IPCReportEventDeviceModel) GetLabeledEnumList() []IREDMLabeledEnum {
	all := p.GetAll()
	labeledList := make([]IREDMLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = IREDMLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
