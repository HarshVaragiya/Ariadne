package main

import (
	"Ariadne/CredManager"
	"fmt"
	"time"
)

func main(){
	tracker := CredManager.NewCredFileTracker("/home/harsh/Desktop/HackTheBox/test/credentials.txt")
	go tracker.Track()

	go func(){
		time.Sleep(time.Second*60)
		tracker.KillTracker()
	}()

	c := tracker.GetTrackerCredChannel()

	for cred := range c{
		fmt.Println(cred)
	}


}