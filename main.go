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

	wordlist := "/home/harsh/Desktop/HackTheBox/Wordlist/small.txt"
	scanner := HTTP.NewBasicGoBusterDir("http://127.0.0.1:8000/","php,html,txt",wordlist,50,&wg,logger)
	report := scanner.Start()
	fmt.Println("Waiting for scans to finish")
	wg.Wait()
	fmt.Println(report)
	fmt.Println("Exiting!")
}