package identity

type AccountType int8

var (
	IsAdmin                = "IsAdmin"
	Operator   AccountType = 1
	NormalUser AccountType = 2
)