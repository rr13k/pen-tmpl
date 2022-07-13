package enum

const (
	CommonUser = iota
	AdminUser
	SuperUser
)

type CaseRunStatus = int

const (
	CaseRunStatus_Wait CaseRunStatus = iota
	CaseRunStatus_Error
	CaseRunStatus_Run
	CaseRunStatus_Success
)
