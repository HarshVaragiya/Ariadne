package ElasticLog

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"strconv"
	"strings"
	"sync"
)

type Logger struct {
	LoggingIndex string
	ctx context.Context
	ElasticClient *elasticsearch.Client
	index int
	Lock *sync.Mutex
}

func (logger *Logger) Init(loggingIndex string){
	logger.ctx = context.Background()
	var err error
	logger.ElasticClient, err = elasticsearch.NewDefaultClient()
	if err !=nil{
		panic(err)
	}
	logger.Lock = &sync.Mutex{}
	logger.index = 0x00
	logger.LoggingIndex = loggingIndex
}

func (logger *Logger)SendLog(log interface{}){
	jsonLog , err := LogToJson(log)
	fmt.Println(jsonLog)
	logger.Lock.Lock()
	logger.index += 1
	defer logger.Lock.Unlock()
	if err == nil {
		req := esapi.IndexRequest{
			Index:      logger.LoggingIndex,
			DocumentID: strconv.Itoa(logger.index),
			Body:       strings.NewReader(jsonLog),
			Refresh:    "true",
		}
		res, err := req.Do(logger.ctx, logger.ElasticClient)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
	}
}
