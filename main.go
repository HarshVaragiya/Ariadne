package main

import (
	"Ariadne/Core"
	"Ariadne/ElasticLog"
	"fmt"
	"sync"
)

func main(){
	logger := &ElasticLog.Logger{}
	logger.Init("self")
	
	log := ElasticLog.NewLog("DEBUG","Starting Ariadne","root")
	logger.SendLog(log)

	var wg sync.WaitGroup
	Core.TestFTP(logger,&wg)
	fmt.Println("Waiting for threads to exit")
	wg.Wait()
	fmt.Println("Exiting")
}