package Hydra

import (
	"Ariadne/CredManager"
	"Ariadne/ElasticLog"
	"fmt"
	"github.com/jlaffaye/ftp"
	"strings"
	"sync"
	"time"
)

const FtpModuleName = "FTPCrack"

type FTPCrack struct {
	ModuleName string
	isAlive bool
}

func (ftpCrack *FTPCrack) preRunStart(){
	ftpCrack.isAlive = true
}

func (ftpCrack *FTPCrack) kill(){
	ftpCrack.isAlive = false
}

func (ftpCrack *FTPCrack) getModuleInfo()string{
	ftpCrack.ModuleName = FtpModuleName
	return ftpCrack.ModuleName
}

func (ftpCrack *FTPCrack) testCredential(target,username,password string,logger *ElasticLog.Logger) bool {
	timeoutError := fmt.Sprintf("dial tcp %s: i/o timeout",target)
	for ;ftpCrack.isAlive;{
		conn, err := ftp.Dial(target, ftp.DialWithTimeout(20*time.Second))
		if err != nil{
			if err.Error() == timeoutError {
				logger.SendLog(ElasticLog.NewLog("DEBUG","Taking a long break due to i/o timeout",ftpCrack.ModuleName))
				time.Sleep(time.Minute*2)
			}
			if strings.Contains(err.Error(),"421") {
				logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),ftpCrack.ModuleName))
				time.Sleep(time.Minute*2)
			}
			logger.SendLog(ElasticLog.NewLog("ERROR",err.Error(),ftpCrack.ModuleName))
			continue
		}
		err = conn.Login(username, password)
		if err == nil {
			logger.SendLog(CredManager.NewCredentialLog(username,password,true,6,ftpCrack.ModuleName,target))
			fmt.Printf("[%s] Possible Valid Credentials for %s => %s : %s \n",ftpCrack.ModuleName,target,username,password)
			return true
		} else if strings.Contains(err.Error(),"530") {
			logger.SendLog(CredManager.NewCredentialLog(username,password,false,0,ftpCrack.ModuleName,target))
			time.Sleep(time.Second*4)
			return false
		}
	}
	return false
}

func DefaultFTPCredentialChecker(target,filename string,threads int,logger *ElasticLog.Logger,parentWaitGroup *sync.WaitGroup) *LibHydra {
	newDefaultCredList := &CredManager.CredList{}
	err := newDefaultCredList.SetCredFile(filename)
	newDefaultCredList.SetCrossConnectStrategy(false)
	if err!=nil {
		panic(err)
	}
	newFTPCracker := &LibHydra{
		target:          target,
		Total:           newDefaultCredList.TotalCreds,
		findOneOnly:     true,
		logger:          logger,
		threads:         threads,
		Credentials:     newDefaultCredList.GetCredentialChannel(),
		parentWaitGroup: parentWaitGroup,
	}
	newFTPCracker.AttachModule(&FTPCrack{})
	return newFTPCracker
}