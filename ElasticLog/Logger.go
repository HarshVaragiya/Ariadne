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
	isElastic 	bool
}

func NewElasticLogger(loggingIndex string)*Logger{
	newLogger := &Logger{}
	newLogger.Init(loggingIndex)
	return newLogger
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
	logger.Lock.Lock()
	logger.index += 1
	defer logger.Lock.Unlock()
	if !logger.isElastic{
		logger.localLogger(jsonLog)
		return
	}
	if err == nil {
		req := esapi.IndexRequest{
			Index:      logger.LoggingIndex,
			DocumentID: strconv.Itoa(logger.index),
			Body:       strings.NewReader(jsonLog),
			Refresh:    "true",
		}
		res, err := req.Do(logger.ctx, logger.ElasticClient)
		if err != nil {
			logger.isElastic = false
			// process this one log here
			logger.localLogger(jsonLog)
		}
		defer res.Body.Close()
	}
}

func (logger *Logger)localLogger(v interface{}){
	// write to a file somewhere
	fmt.Println(v)
}