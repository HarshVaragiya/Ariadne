package Hydra

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func (credList *CredList) GetCredentialChannel() chan Cred{
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
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		credList.usernames = append(credList.usernames, strings.Split(line,":")[0])
		credList.passwords = append(credList.passwords, strings.SplitAfter(line,":")[1])
	}
	fmt.Println("Generated Credlist")
	return nil
}

