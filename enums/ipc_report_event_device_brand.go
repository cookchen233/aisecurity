package enums

type IPCReportEventDeviceBrand IntEnum

const (
	IREDBUnknown IPCReportEventDeviceBrand = iota
	IREDB1
)

func (IPCReportEventDeviceBrand) GetAll() []IPCReportEventDeviceBrand {
	return []IPCReportEventDeviceBrand{
		IREDBUnknown, IREDB1,
	}
}

func (p IPCReportEventDeviceBrand) Label() string {
	switch p {
	case IREDB1:
		return "图为"
	default:
		return "Unknown"
	}
}

type IREDBLabeledEnum struct {
	Label string                    `json:"label"`
	Value IPCReportEventDeviceBrand `json:"value"`
}

func (p IPCReportEventDeviceBrand) GetLabeledEnumList() []IREDBLabeledEnum {
	all := p.GetAll()
	labeledList := make([]IREDBLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = IREDBLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
