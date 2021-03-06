package ElasticLog

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	Type  	  string
	Text 	  string
	Module 	  string
	Priority  int
}

type Progress struct {
	Type     string
	Target   string
	Done     int
	Total    int
	Module   string
	Work	 string
	Priority int
}

func NewLog(Type,Text,Module string)Log{
	var priority int
	switch Type{
	case "TRACE":
		priority = 0
	case "DEBUG":
		priority = 1
	case "INFO":
		priority = 2
	case "WARN":
		priority = 3
	case "ERROR":
		priority = 4
	case "IMP":
		priority = 5
	}
	newLog := Log{Type,Text,Module,priority}
	return newLog
}

func NewProgressLog(Module,Target,Work string,Done,Remaining int)Progress{
	return Progress{"PROGRESS",Target,Done,Remaining,Module,Work,1}
}

func LogToJson(log interface{}) (string,error) {
	b, err := json.Marshal(log)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return err.Error(),err
	}
	return string(b),nil
}

