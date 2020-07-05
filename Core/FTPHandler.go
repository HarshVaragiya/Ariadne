package Core

import (
	"Ariadne/ElasticLog"
	"Ariadne/Hydra"
	"fmt"
	"sync"
)

func TestFTP(logger *ElasticLog.Logger,parentWaitGroup *sync.WaitGroup){
	ftpTarget := fmt.Sprintf("%s:%d","192.168.1.1",21)
	kill := false
	defaultCredChecker := Hydra.FTPDefaultCredentialCheck(ftpTarget,"/home/harsh/Desktop/HackTheBox/Wordlist/ftp-betterdefaultpasslist.txt",logger,&kill,parentWaitGroup)
	fmt.Println("Starting ...")
	defaultCredChecker.StartCracking()
	fmt.Println("Waiting for threads to end ...")
	parentWaitGroup.Wait()
}
