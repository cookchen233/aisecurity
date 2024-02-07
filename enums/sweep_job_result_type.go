package enums

type SweepJobResultType int

const (
	SRTUnknown SweepJobResultType = iota
	SRT1
	SRT2
	SRT3
	SRT4
)

func (SweepJobResultType) GetAll() []SweepJobResultType {
	return []SweepJobResultType{
		SRTUnknown,
		SRT1,
		SRT2,
		SRT3,
		SRT4,
	}
}

func (p SweepJobResultType) Label() string {
	switch p {
	case SRT1:
		return "合格"
	case SRT2:
		return "不合格"
	case SRT3:
		return "整改中"
	case SRT4:
		return "暂时无法评估"

	default:
		return "Status Unknown"
	}
}

type SRTLabeledEnum struct {
	Label string             `json:"label"`
	Value SweepJobResultType `json:"value"`
}

func (p SweepJobResultType) GetLabeledEnumList() []SRTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]SRTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = SRTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
