package main

import (
	"Ariadne/CredManager"
	"Ariadne/ElasticLog"
	"Ariadne/Hydra"
	"fmt"
	"sync"
	"time"
)

func main(){

	//wordlist := "/home/harsh/Desktop/HackTheBox/Wordlist/ftp-betterdefaultpasslist.txt"
	logger := ElasticLog.NewElasticLogger("defaultIndex")

	var wg sync.WaitGroup

	//credList := CredManager.NewCredListFromFile(wordlist,false)
	tracker := CredManager.NewCredFileTracker("/home/harsh/Desktop/HackTheBox/test/credentials.txt")
	go tracker.Track()
	newCracker := Hydra.NewLibHydraModule("192.168.1.1:21",tracker.GetTrackerCredChannel(),1,logger,&wg)
	newCracker.AttachModule(&Hydra.FTPCrack{})
	newCracker.StartCracking()
	go func(){
		// time to stop
		time.Sleep(time.Minute)
		fmt.Println("\nKilling tracker")
		tracker.KillTracker()
	}()
	wg.Wait()
}