package CredManager

import "fmt"

// tracks a file of credentials where each line is of type
// username1:password1
// username2:password2
// ...
// automatically checks cross credential usage
type CredFileTracker struct {
	trackedFile string
	checkedLines []string
	checkedUsernames []string
	checkedPasswords []string
	previousFileHash [32]byte
}

func NewCredFileTracker(filename string)*CredFileTracker{
	tracker := &CredFileTracker{
		trackedFile: filename,
		checkedLines:     nil,
		checkedUsernames: nil,
		checkedPasswords: nil,
		previousFileHash: [32]byte{},
	}
	tracker.init()
	return tracker
}

func (tracker *CredFileTracker)init(){
	usernameQueue,passwordQueue,err := GetCredentialsFromFile(tracker.trackedFile)
	if err !=nil{
		panic(err)
	}
	tracker.processNewCreds(usernameQueue,passwordQueue)
}

func (tracker *CredFileTracker) processNewCreds(usernames,passwords []string) {
	// check each username with known passwords + new passwords
	var newUsernames,newPasswords []string
	for _,username := range usernames{
		possiblePasswords := append(tracker.checkedPasswords,passwords...)
		for i:=0;i<len(possiblePasswords);i++ {
			newUsernames = append(newUsernames, username)
		}
		newPasswords = append(newPasswords,possiblePasswords...)
	}
	fmt.Println(newUsernames,"\n",newPasswords)
	// check each password with known usernames + new usernames
}

func (tracker *CredFileTracker) filterAllCreds (usernames,passwords []string)(filteredUsernames,filteredPasswords []string){
	for _ , username := range usernames {
		if !Contains(tracker.checkedUsernames,username){
			filteredUsernames = append(filteredUsernames,username)
		}
	}

	for _ , password := range passwords{
		if !Contains(tracker.checkedPasswords,password){
			filteredPasswords = append(filteredPasswords,password)
		}
	}
	return
}

























