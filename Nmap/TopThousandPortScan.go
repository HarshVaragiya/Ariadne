package Nmap

import (
	"Ariadne/ElasticLog"
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"time"
)

func (portScan *PortScan) TopThousandPortScan(){
	defer portScan.ParentWaitGroup.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	scanner, err := nmap.NewScanner(nmap.WithTargets(portScan.target),nmap.WithContext(ctx))
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
				str := fmt.Sprintf("Found OPEN Port %d running service %s",port.ID,port.Service.String())
				portScan.logger.SendLog(ElasticLog.NewLog("IMP",str,portScan.ModuleName))
				fmt.Println(str)
				portScan.DefaultScanPorts[port.Service.String()] = append(portScan.DefaultScanPorts[port.Service.String()],port.ID)
				counter := 0x00
				for _,p := range portScan.UniquePorts{
					if p != port.ID{
						counter +=1
					}
				}
				if counter == len(portScan.UniquePorts){
					// unique result
					portScan.UniquePorts = append(portScan.UniquePorts,port.ID)
					portScan.UniqueServices[port.Service.String()] = append(portScan.UniqueServices[port.Service.String()],port.ID)
				}
			}
		}
	}
	portScan.SendPortScanLogUpdate()
}