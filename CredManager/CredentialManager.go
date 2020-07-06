package CredManager


// List of credentials for testing for all submodules of Hydra or any other Modules
type CredList struct {
	usernames []string
	passwords []string

	crossConnect bool
	Credentials chan Cred
	TotalCreds int
}

func (credList *CredList) SetCredentials(usernames,passwords []string){
	credList.usernames = usernames
	credList.passwords = passwords
}

func (credList *CredList) SetCrossConnectStrategy(crossConnect bool){
	credList.crossConnect = crossConnect
}

func (credList *CredList) GetCredentialChannel() chan Cred {
	if credList.crossConnect {
		credList.crossConnectCreds()
	} else {
		credList.linearConnectCreds()
	}
	return credList.Credentials
}

func (credList *CredList) crossConnectCreds() {
	credList.TotalCreds = len(credList.usernames)*len(credList.passwords)
	credList.Credentials = make(chan Cred,credList.TotalCreds)
	for i := range credList.usernames{
		for j := range credList.passwords{
			credList.Credentials <- Cred{credList.usernames[i],credList.passwords[j]}
		}
	}
}

func (credList *CredList) linearConnectCreds() {
	credList.TotalCreds = len(credList.usernames)
	credList.Credentials = make(chan Cred,credList.TotalCreds)
	for i := range credList.usernames{
		credList.Credentials <- Cred{credList.usernames[i],credList.passwords[i]}
	}
}

func (credList *CredList) SetCredFile(filename string)error{
	var err error
	credList.usernames , credList.passwords , err = GetCredentialsFromFile(filename)
	return err
}