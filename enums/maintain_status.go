package enums

type MaintainStatus int

const (
	MSUnknown MaintainStatus = iota
	MS1
	MS2
	MS3
	MS4
	MS5
)

func (MaintainStatus) GetAll() []MaintainStatus {
	return []MaintainStatus{
		MSUnknown,
		MS1,
		MS2,
		MS3,
		MS4,
		MS5,
	}
}

func (p MaintainStatus) Label() string {
	switch p {
	case MS1:
		return "待整改"
	case MS2:
		return "整改中"
	case MS3:
		return "验收中"
	case MS4:
		return "已完成"
	case MS5:
		return "已驳回"
	default:
		return "Status Unknown"
	}
}

type MSLabeledEnum struct {
	Label string         `json:"label"`
	Value MaintainStatus `json:"value"`
}

func (p MaintainStatus) GetLabeledEnumList() []MSLabeledEnum {
	all := p.GetAll()
	labeledList := make([]MSLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = MSLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
