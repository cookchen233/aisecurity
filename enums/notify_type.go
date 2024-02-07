package enums

type NotifyType int

const (
	NTUnknown NotifyType = iota
	NT1
	NT2
	NT3
)

func (NotifyType) GetAll() []NotifyType {
	return []NotifyType{
		NTUnknown,
		NT1,
		NT2,
		NT3,
	}
}

func (p NotifyType) Label() string {
	switch p {
	case NT1:
		return "公众号"
	case NT2:
		return "短信"
	default:
		return "Status Unknown"
	}
}

type NTLabeledEnum struct {
	Label string     `json:"label"`
	Value NotifyType `json:"value"`
}

func (p NotifyType) GetLabeledEnumList() []NTLabeledEnum {
	all := p.GetAll()
	labeledList := make([]NTLabeledEnum, len(all))
	for i, enum := range all {
		labeledList[i] = NTLabeledEnum{
			Label: enum.Label(),
			Value: enum,
		}
	}
	return labeledList
}
