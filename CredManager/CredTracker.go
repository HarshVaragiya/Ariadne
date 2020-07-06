package CredManager

import (
	"fmt"
	"os"
	"time"
)

// tracks a file of credentials where each line is of type
// username1:password1
// username2:password2
// ...
// automatically checks cross credential usage
type CredFileTracker struct {
	trackedFile string
	checkedUsernames []string
	checkedPasswords []string

	possibleCredentials chan Cred
	isAlive bool
	sleepDuration time.Duration
}

func NewCredFileTracker(filename string)*CredFileTracker{
	ensureFileExistence(filename)
	tracker := &CredFileTracker{
		trackedFile: filename,
		checkedUsernames: nil,
		checkedPasswords: nil,
		possibleCredentials: make(chan Cred,100),
		isAlive: true,
		sleepDuration: time.Second*5,
	}
	return tracker
}

func (tracker *CredFileTracker) Track() {
	for ;tracker.isAlive; {
		usernameQueue, passwordQueue, err := GetCredentialsFromFile(tracker.trackedFile)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
			// panic(err)
		}
		filteredUsernames, filteredPasswords, foundNew := tracker.filterAllCreds(usernameQueue, passwordQueue)
		if !foundNew {
			time.Sleep(tracker.sleepDuration)
		}
		tracker.processNewCreds(filteredUsernames, filteredPasswords)
	}
}

func (tracker *CredFileTracker) processNewCreds (usernames,passwords []string){
	// check each new username with known passwords + new passwords
	var myMap map[string][]string = make(map[string][]string)
	for _,username := range usernames{
		possiblePasswords := append(tracker.checkedPasswords,passwords...)
		myMap[username] = append(myMap[username],possiblePasswords...)
	}
	// check each of the old usernames with new password
	for _,password := range passwords{
		for _,username := range tracker.checkedUsernames{
			myMap[username] = append(myMap[username],password)
		}
	}
	// flatten the map to two arrays []username , []password
	for username,passphrases := range myMap{
		for _ , passphrase := range passphrases{
			tracker.possibleCredentials <- Cred{username,passphrase}
		}
	}
	tracker.checkedUsernames = append(tracker.checkedUsernames,usernames...)
	tracker.checkedPasswords = append(tracker.checkedPasswords,passwords...)
}

func (tracker *CredFileTracker) filterAllCreds (usernames,passwords []string)(filteredUsernames,filteredPasswords []string,foundNew bool){
	foundNew = true
	for _ , username := range usernames {
		if Contains(tracker.checkedUsernames,username)==-1{
			filteredUsernames = append(filteredUsernames,username)
		}
	}

	for _ , password := range passwords{
		if Contains(tracker.checkedPasswords,password)==-1{
			filteredPasswords = append(filteredPasswords,password)
		}
	}

	if filteredUsernames == nil && filteredPasswords == nil {
		foundNew = false
	}

	return
}

func (tracker *CredFileTracker) GetTrackerCredChannel()chan Cred{
	return tracker.possibleCredentials
}

func (tracker *CredFileTracker) KillTracker(){
	close(tracker.possibleCredentials)
	tracker.isAlive = false
}