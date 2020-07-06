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
		time.Sleep(time.Second*40)
		tracker.KillTracker()
	}()

	fmt.Println("Getting Tracker Cred Channel .. ")
	c := tracker.GetTrackerCredChannel()

	fmt.Println("Cred Channel : " , c)
	for cred := range c{
		fmt.Println(cred)
	}


}