package enums

type IPCReportEventStatus int

const (
	IRESUnknown IPCReportEventStatus = iota
	IRESPending
	IRESProcessing
	IRESProcessed
)

func (IPCReportEventStatus) GetAll() []IPCReportEventStatus {
	return []IPCReportEventStatus{
		IRESUnknown, IRESPending, IRESProcessing, IRESProcessed,
	}
}

func (p IPCReportEventStatus) Label() string {
	switch p {
	case IRESPending:
		return "待处理"
	case IRESProcessing:
		return "处理中"
	case IRESProcessed:
		return "已处理"
	default:
		return "Unknown"
	}
}

type IRESLabeledEnum struct {
	Label string               `json:"label"`
	Value IPCReportEventStatus `json:"value"`
}

func (p IPCReportEventStatus) GetLabeledEnumList() []IRESLabeledEnum {
	all := p.GetAll()
	labeledList := make([]IRESLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = IRESLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
