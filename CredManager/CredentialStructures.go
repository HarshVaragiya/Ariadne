package CredManager

// For general purpose
type Cred struct {
	Username string
	Password string
}

// Logging to elasticsearch and full trace logging purposes
type CredentialLog struct {
	Type 	 string
	Username string
	Password string
	IsValid bool
	Priority int

	Module string
	Target string
}

// Constructor for log
func NewCredentialLog(Username,Password string,IsValid bool,Priority int,Module,Target string) CredentialLog {
	return CredentialLog{"CRED",Username,Password,IsValid,Priority,Module,Target}
}