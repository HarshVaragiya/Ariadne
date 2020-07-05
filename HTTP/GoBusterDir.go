package HTTP

import (
	"Ariadne/ElasticLog"
	"context"
	"fmt"
	"github.com/OJ/gobuster/gobusterdir"
	"github.com/OJ/gobuster/helper"
	"github.com/OJ/gobuster/libgobuster"
	"sync"
	"time"
)

type GobusterDir struct{
	parentWaitGroup *sync.WaitGroup
	context         context.Context
	cancel          context.CancelFunc
	timeoutMinutes  int
	dirbuster       *libgobuster.Gobuster
	Entities        []Endpoint
	targetBaseURL   string
	threadCount     int
	wordlist        string
	extensions      string
	Done            bool
	ReportOutput 	*DirSearchReport
	ModuleName		string
	logger 			*ElasticLog.Logger
}

func NewBasicGoBusterDir(targetURL,extensions,wordlist string,threadCount int,parentWaitGroup *sync.WaitGroup,logger *ElasticLog.Logger)*GobusterDir{
	dirbuster := GobusterDir{}
	dirbuster.ModuleName = "GOBUSTER-DIR"
	dirbuster.Done = false
	dirbuster.parentWaitGroup = parentWaitGroup
	dirbuster.context, dirbuster.cancel = context.WithTimeout(context.Background(), time.Duration(5)*time.Minute)
	dirbuster.targetBaseURL = targetURL
	dirbuster.wordlist = wordlist
	dirbuster.threadCount = threadCount
	dirbuster.extensions = extensions
	dirbuster.logger = logger
	err := dirbuster.defaults()
	if err != nil{
		logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),dirbuster.ModuleName))
		return nil
	}
	return &dirbuster
}
func (dirbuster *GobusterDir) defaults() error {
	httpOpts := libgobuster.OptionsHTTP{}
	httpOpts.URL = dirbuster.targetBaseURL
	httpOpts.UserAgent = "pwn/0.1"
	httpOpts.InsecureSSL = true
	pluginOpts := gobusterdir.NewOptionsDir()
	pluginOpts.Password = httpOpts.Password
	pluginOpts.URL = httpOpts.URL
	pluginOpts.UserAgent = httpOpts.UserAgent
	pluginOpts.Username = httpOpts.Username
	pluginOpts.Proxy = httpOpts.Proxy
	pluginOpts.Cookies = httpOpts.Cookies
	pluginOpts.Timeout = httpOpts.Timeout
	pluginOpts.FollowRedirect = httpOpts.FollowRedirect
	pluginOpts.InsecureSSL = httpOpts.InsecureSSL
	pluginOpts.Headers = httpOpts.Headers
	pluginOpts.StatusCodes = "200,204,301,302,307,401,403"
	pluginOpts.StatusCodesBlacklist = "404"

	opts := libgobuster.NewOptions()
	opts.Threads = dirbuster.threadCount
	opts.Wordlist = dirbuster.wordlist

	pluginOpts.StatusCodesParsed,_ = helper.ParseStatusCodes(pluginOpts.StatusCodes)
	pluginOpts.ExtensionsParsed,_ = helper.ParseExtensions(dirbuster.extensions)
	plugin, err := gobusterdir.NewGobusterDir(dirbuster.context, opts, pluginOpts)

	if err != nil{
		dirbuster.logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),dirbuster.ModuleName))
		return err
	}
	ctx, _ := context.WithCancel(dirbuster.context)
	dirbuster.dirbuster, err = libgobuster.NewGobuster(ctx, opts, plugin)
	if err!=nil{
		dirbuster.logger.SendLog(ElasticLog.NewLog("ERROR","Libgobuster Error" + err.Error(),dirbuster.ModuleName))
		return err
	}
	return nil
}

func(dirbuster *GobusterDir) updateEntities(){
	for r := range dirbuster.dirbuster.Results(){
		s, _ := r.ToString(dirbuster.dirbuster)
		if s != "" {
			endpoint := Endpoint{
				Entity:     r.Entity,
				StatusCode: r.StatusCode,
			}
			dirbuster.Entities = append(dirbuster.Entities,endpoint)
			if dirbuster.ReportOutput != nil{
				dirbuster.ReportOutput.Endpoints = append(dirbuster.ReportOutput.Endpoints,endpoint)

				dirbuster.logger.SendLog(dirbuster.ReportOutput)
				dirbuster.logger.SendLog(ElasticLog.NewLog("IMP","Endpoint found : "+dirbuster.targetBaseURL+r.Entity,dirbuster.ModuleName))
				fmt.Printf("[%s] Endpoint found : %s%s  [StatusCode %d] \n",dirbuster.ModuleName,dirbuster.targetBaseURL,r.Entity,r.StatusCode)
				// log output to elasticsearch
			}
		}
	}
}

func (dirbuster *GobusterDir) GetProgress()(int,int){
	dirbuster.dirbuster.Mu.RLock()
	defer dirbuster.dirbuster.Mu.RUnlock()
	return dirbuster.dirbuster.RequestsIssued+1, dirbuster.dirbuster.RequestsExpected
}

func (dirbuster *GobusterDir)Start() *DirSearchReport{
	dirbuster.ReportOutput = &DirSearchReport{TargetURL: dirbuster.targetBaseURL,Priority: 6}
	dirbuster.parentWaitGroup.Add(1)
	go dirbuster.run()
	return dirbuster.ReportOutput
}

func (dirbuster *GobusterDir)run(){
	defer dirbuster.parentWaitGroup.Done()
	go dirbuster.updateEntities()
	go dirbuster.logProgress()
	err := dirbuster.dirbuster.Start()
	if err != nil {
		dirbuster.logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),dirbuster.ModuleName))
		fmt.Println(err) // should actually be panic - but YOLO
	}
	dirbuster.Done = true
}

func (dirbuster *GobusterDir) logProgress(){
	for ;!dirbuster.Done;{
		time.Sleep(time.Second)
		done,total := dirbuster.GetProgress()
		dirbuster.logger.SendLog(ElasticLog.NewProgressLog(dirbuster.ModuleName,dirbuster.targetBaseURL,done,total))
	}
}