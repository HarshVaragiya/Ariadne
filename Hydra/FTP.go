package Hydra

import (
	"Ariadne/ElasticLog"
	"fmt"
	"github.com/jlaffaye/ftp"
	"sync"
	"time"
)

type FTP struct {
	target string  //ip:port type
	done int       //done and remaining creds to be tested
	remaining int

	CredList CredList

	findOneOnly bool
	foundCred bool

	logger      *ElasticLog.Logger

	threads int
	ModuleName string
	kill *bool
	parentWaitGroup *sync.WaitGroup
}

func FTPDefaultCredentialCheck(target,filename string,logger *ElasticLog.Logger,kill *bool,parentWaitGroup *sync.WaitGroup)FTP{
	newDefaultCredList := CredList{}
	err := newDefaultCredList.SetCredFile(filename)
	newDefaultCredList.SetCrossConnectStrategy(false)
	if err!=nil {
		panic(err)
	}
	newFTPCracker := FTP{
		target:          target,
		done:            0,
		remaining:       newDefaultCredList.TotalCreds,
		CredList:        newDefaultCredList,
		findOneOnly:     true,
		foundCred:       false,
		logger:          logger,
		threads:         4,
		ModuleName:      "FTPCrack",
		kill:            kill,
		parentWaitGroup: parentWaitGroup,
	}
	return newFTPCracker
}

func (ftpCrack *FTP) StartCracking(){
	credentials := ftpCrack.CredList.GetCredentialChannel()
	ftpCrack.parentWaitGroup.Add(ftpCrack.threads)
	for i:=0;i<ftpCrack.threads;i++{
		go ftpCrack.CheckCredentials(credentials,ftpCrack.parentWaitGroup)
	}
}

func (ftpCrack *FTP)CheckCredentials(credentials chan Cred,group *sync.WaitGroup){
	defer group.Done()
	for credential := range credentials{
		isValid := ftpCrack.CheckFTPCredential(ftpCrack.target,credential.Username,credential.Password)
		if isValid {
			ftpCrack.foundCred = true
			if ftpCrack.findOneOnly {
				done := true
				ftpCrack.kill = &done  // no idea if this will work or not
			}
		}
		ftpCrack.done +=1
		ftpCrack.logger.SendLog(ElasticLog.NewProgressLog(ftpCrack.ModuleName,ftpCrack.target,ftpCrack.done,ftpCrack.remaining))
	}
}

func (ftpCrack *FTP) CheckFTPCredential (target,username,password string) bool {
	timeoutError := fmt.Sprintf("dial tcp %s: i/o timeout",target)
	invalidPassphraseError := fmt.Sprintf("530 User %s cannot log in.",username)
	for ;!*ftpCrack.kill;{
		conn, err := ftp.Dial(target, ftp.DialWithTimeout(20*time.Second))
		if err != nil{
			if err.Error() == timeoutError {
				ftpCrack.logger.SendLog(ElasticLog.NewLog("DEBUG","Taking a long break due to i/o timeout",ftpCrack.ModuleName))
				time.Sleep(time.Minute*2)
				continue
			}
			ftpCrack.logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),ftpCrack.ModuleName))
		}
		err = conn.Login(username, password)
		if err.Error() == invalidPassphraseError {
			ftpCrack.logger.SendLog(NewCredential(username,password,false,0,ftpCrack.ModuleName,target))
			time.Sleep(time.Second*4)
			return false
		}
		if err == nil {
			ftpCrack.logger.SendLog(NewCredential(username,password,true,6,ftpCrack.ModuleName,target))
			return true
		}
	}
	return false
}