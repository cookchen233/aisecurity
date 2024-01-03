package properties

type MaintainStatus int

const (
	Ready         MaintainStatus = 1
	Processing    MaintainStatus = 2
	AwaitingCheck MaintainStatus = 3
	Completed     MaintainStatus = 4
	Rejected      MaintainStatus = 5
)
