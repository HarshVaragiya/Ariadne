package Hydra

type CredentialLog struct {
	Type 	 string
	Username string
	Password string
	IsValid bool
	Priority int

	Module string
	Target string
}

func NewCredential(Username,Password string,IsValid bool,Priority int,Module,Target string) CredentialLog {
	return CredentialLog{"CRED",Username,Password,IsValid,Priority,Module,Target}
}

type Cred struct {
	Username string
	Password string
}