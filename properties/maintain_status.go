package properties

type MaintainStatus int

const (
	Unknown MaintainStatus = iota
	Ready
	Processing
	WaitingCheck
	Completed
	Rejected
)

func (MaintainStatus) GetAll() []MaintainStatus {
	return []MaintainStatus{
		Unknown, Ready, Processing, WaitingCheck, Completed, Rejected,
	}
}

func (p MaintainStatus) Label() string {
	switch p {
	case Ready:
		return "待整改"
	case Processing:
		return "整改中"
	case WaitingCheck:
		return "验收中"
	case Completed:
		return "已完成"
	case Rejected:
		return "已驳回"
	default:
		return "Status Unknown"
	}
}

type LabeledEnum struct {
	Label string         `json:"label"`
	Value MaintainStatus `json:"value"`
}

func (p MaintainStatus) GetLabeledEnumList() []LabeledEnum {
	all := p.GetAll()
	labeledList := make([]LabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = LabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
