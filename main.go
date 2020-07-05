package main

import (
	"Ariadne/ElasticLog"
	"Ariadne/Nmap"
	"fmt"
	"sync"
)

func main(){
	logger := &ElasticLog.Logger{}
	logger.Init("nmap")

	var wg sync.WaitGroup
	scanner := Nmap.NewPortScanner("192.168.1.1",&wg,logger)
	scanner.DefaultScan()
	fmt.Println("Waiting for scans to finish")
	wg.Wait()
	fmt.Println(scanner.PortsFoundLog)
	fmt.Println("Exiting!")
}