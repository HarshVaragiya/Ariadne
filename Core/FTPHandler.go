package Core

import (
	"Ariadne/ElasticLog"
	"Ariadne/Hydra"
	"fmt"
	"sync"
)

func TestFTP(logger *ElasticLog.Logger){
	var myWaitGroup sync.WaitGroup
	ftpTarget := fmt.Sprintf("%s:%d","127.0.0.1",21)
	threads := 2
	defaultCredChecker := Hydra.FTPDefaultCredentialCheck(ftpTarget,
		"/home/harsh/Desktop/HackTheBox/Wordlist/ftp-betterdefaultpasslist.txt",threads,logger,&myWaitGroup)
	fmt.Println("Starting ...")
	defaultCredChecker.StartCracking()
	fmt.Println("Waiting for threads to end ...")
	myWaitGroup.Wait()
}
