package Core

import (
	"Ariadne/ElasticLog"
	"Ariadne/HTTP"
	"Ariadne/Nmap"
	"context"
	"fmt"
	"sync"
)

const MaxHttpThreads = 50
const HttpDirWordlist = "/home/harsh/Desktop/HackTheBox/Wordlist/directory-list-2.3-medium.txt"

type AriadneTarget struct {
	rootWaitGroup		*sync.WaitGroup
	rootContext 	     context.Context
	rootCancelFunc 	     context.CancelFunc
	rootTarget			 string

	logger 				*ElasticLog.Logger
	httpJobs 			*[]HTTP.GobusterDir

	portScanWaitGroup,httpWaitGroup sync.WaitGroup
}

func NewAriadneTarget(target,loggingIndex string,ctx context.Context)*AriadneTarget{
	var waitGroup sync.WaitGroup
	logger := ElasticLog.NewElasticLogger(loggingIndex)
	newAriadneTarget := AriadneTarget{
		rootWaitGroup:  &waitGroup,
		rootTarget:     target,
		logger:         logger,
	}
	newAriadneTarget.rootContext, newAriadneTarget.rootCancelFunc = context.WithCancel(ctx)
	return &newAriadneTarget
}

func (target *AriadneTarget) StartEnumerating(httpExtensions string){
	portScanner := Nmap.NewPortScanner(target.rootTarget,&target.portScanWaitGroup,target.logger)
	portScanner.DefaultScan()
	target.portScanWaitGroup.Wait() // wait for the port scan to finish...
	fmt.Println(portScanner.DisplayHumanReadablePorts()) // print nmap scan report

	// write a parser that parses portScanner.UniqueServices result and starts required modules
	services := portServiceParser(portScanner.UniqueServices)
	fmt.Println("Following services appear to be running : ",services)
	fmt.Println("Starting HTTP enumeration on ports ",services["http"])
	target.HttpHandler(httpExtensions,services["http"],&target.httpWaitGroup)

	// do other stuff

	target.httpWaitGroup.Wait() // wait in the end .. continue before this

	fmt.Println("\nHttp Report/s :")
	for _ , job := range *target.httpJobs{
		fmt.Println(job.GetReport().DisplayHumanReadableEndpoints())
	}
}

func (target *AriadneTarget) HttpHandler(httpExtensions string, ports []uint16,httpWaitGroup *sync.WaitGroup){
	threads := MaxHttpThreads / len(ports)
	var httpJobs []HTTP.GobusterDir
	for _,port := range ports{
		targetURL := fmt.Sprintf("http://%s:%d/",target.rootTarget,port)
		newHttpJob := HTTP.NewBasicGoBusterDir(targetURL,httpExtensions,HttpDirWordlist,threads,target.rootContext,httpWaitGroup,target.logger)
		newHttpJob.Start()
		httpJobs = append(httpJobs,*newHttpJob)
	}
	target.httpJobs = &httpJobs
}