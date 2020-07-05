package main

import (
	"Ariadne/Core"
	"Ariadne/ElasticLog"
	"fmt"
)

func main(){
	logger := &ElasticLog.Logger{}
	logger.Init("self")
	
	log := ElasticLog.NewLog("DEBUG","Starting Ariadne","root")
	logger.SendLog(log)

	Core.TestFTP(logger)
	fmt.Println("Exiting")
}