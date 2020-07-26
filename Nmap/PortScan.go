package Nmap

import (
	"Ariadne/ElasticLog"
	"sync"
)

type PortScan struct {
	InitialScanPorts map[string][]uint16
	DefaultScanPorts map[string][]uint16

	UniquePorts 	 []uint16
	UniqueServices   map[string][]uint16

	target 			 string
	ParentWaitGroup  *sync.WaitGroup
	logger 			 *ElasticLog.Logger
	ModuleName		 string
	
	PortsFoundLog    PortScanLog
	done			 int
	total		     int
}

func NewPortScanner(target string,parentWaitGroup *sync.WaitGroup,logger *ElasticLog.Logger) *PortScan{
	portScan := PortScan{
		InitialScanPorts: make(map[string][]uint16),
		DefaultScanPorts: make(map[string][]uint16),
		UniqueServices:   make(map[string][]uint16),
		target:           target,
		ParentWaitGroup:  parentWaitGroup,
		logger:           logger,
		ModuleName:  	  "NMAP",
	}
	return &portScan
}

func (portScan *PortScan) DefaultScan(){
	portScan.total = 2
	portScan.ParentWaitGroup.Add(1)
	go portScan.AllPortScan()
	portScan.InitialQuickPortScan()
}

func (portScan *PortScan)SendPortScanLogUpdate(){
	portScan.PortsFoundLog.UniquePorts = portScan.UniquePorts
	portScan.PortsFoundLog.UniqueServices = portScan.UniqueServices
	portScan.logger.SendLog(portScan.PortsFoundLog)
}

type PortScanLog struct {
	UniquePorts []uint16
	UniqueServices map[string][]uint16
}
