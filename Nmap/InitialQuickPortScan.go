package Nmap

import (
	"Ariadne/ElasticLog"
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"time"
)

func (portScan *PortScan) InitialQuickPortScan(){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	// running 'nmap -p 21,80,443,8000,8080,8443 <target>', with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(portScan.target),nmap.WithPorts("21,80,443,8000,8080,8443"),nmap.WithContext(ctx),
	)
	if err != nil {
		panic(err)
	}
	result, warnings, err := scanner.Run()
	if err != nil {
		panic(err)
	}
	if warnings != nil {
		fmt.Println(warnings)
	}
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		for _, port := range host.Ports {
			if port.State.State == "open"{
				str := fmt.Sprintf("[%s] Found OPEN Port %d running service %s",portScan.ModuleName,port.ID,port.Service.String())
				portScan.logger.SendLog(ElasticLog.NewLog("DEBUG",str,portScan.ModuleName)) // one logs to debug , other to IMP
				fmt.Println(str)
				portScan.InitialScanPorts[port.Service.String()] = append(portScan.InitialScanPorts[port.Service.String()],port.ID)
				portScan.UniqueServices[port.Service.String()] = append(portScan.UniqueServices[port.Service.String()],port.ID)
				portScan.UniquePorts = append(portScan.UniquePorts,port.ID)
			}

		}
	}
	portScan.SendPortScanLogUpdate()
}