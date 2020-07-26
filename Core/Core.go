package Core

import (
	"Ariadne/ElasticLog"
	"Ariadne/Nmap"
	"context"
	"log"
	"sync"
)

type AriadneTarget struct {
	rootWaitGroup		*sync.WaitGroup
	rootContext 	     context.Context
	rootCancelFunc 	     context.CancelFunc
	rootTarget			 string

	logger 				*ElasticLog.Logger
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

func (target *AriadneTarget) StartEnumerating(){
	var portScanWaitGroup sync.WaitGroup
	portScanner := Nmap.NewPortScanner(target.rootTarget,&portScanWaitGroup,target.logger)
	portScanner.DefaultScan()
	portScanWaitGroup.Wait() // wait for the port scan to finish...
	log.Println("Services Found : ",portScanner.UniqueServices)

	// write a parser that parses portScanner.UniqueServices result and starts required modules
}

