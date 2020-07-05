package main

import (
	"Ariadne/ElasticLog"
	"Ariadne/HTTP"
	"fmt"
	"sync"
)

func main(){
	logger := &ElasticLog.Logger{}
	logger.Init("gobuster")

	var wg sync.WaitGroup

	wordlist := "/home/harsh/Desktop/HackTheBox/Wordlist/directory-list-2.3-small.txt"
	scanner := HTTP.NewBasicGoBusterDir("http://192.168.1.1:80/","php,html,txt",wordlist,80,&wg,logger)
	report := scanner.Start()
	fmt.Println("Waiting for scans to finish")
	wg.Wait()
	fmt.Println(report)
	fmt.Println("Exiting!")
}