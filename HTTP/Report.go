package HTTP

import (
	"fmt"
	"strings"
)

type DirSearchReport struct {
	TargetURL string
	Endpoints []Endpoint

	Priority  int
}

type Endpoint struct{
	Entity string
	StatusCode int
}

func(report *DirSearchReport) DisplayHumanReadableEndpoints()string{
	ret := strings.Repeat("=",60)
	ret += fmt.Sprintf("\nGobuster dir search report for URL: %s \n",report.TargetURL)
	ret += strings.Repeat("-",60) + "\n"
	for _ , endpoint := range report.Endpoints{
		url := fmt.Sprintf("%s%s",report.TargetURL,endpoint.Entity)
		ret += fmt.Sprintf(" %-40s - STATUS [%3d] \n",url,endpoint.StatusCode)
	}
	ret += strings.Repeat("-",60) + "\n"
	ret += "End of report.\n"
	ret += strings.Repeat("=",60) + "\n"
	return ret
}
