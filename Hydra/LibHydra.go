package Hydra

import (
	"Ariadne/CredManager"
	"Ariadne/ElasticLog"
	"context"
	"sync"
)

const ModuleName = "HYDRA"

type Module interface {
	testCredential(string,string,string,*ElasticLog.Logger,context.Context) bool
	preRunStart()
	kill()
	getModuleInfo()string
}

type LibHydra struct {
	target string    // ip:port type
	done   int       // done and total creds to be tested
	Total  int

	findOneOnly bool
	foundCred bool

	logger      	*ElasticLog.Logger

	threads          int
	parentModuleName string
	kill             bool
	parentWaitGroup  *sync.WaitGroup
	lock             sync.Mutex

	Credentials chan CredManager.Cred
	ctx 			 context.Context
	cancelFunc 		 context.CancelFunc

	module			 Module
}

func NewLibHydraModule(target string,credChannel chan CredManager.Cred,threads int,logger *ElasticLog.Logger,parentWaitGroup *sync.WaitGroup,ctx context.Context) *LibHydra {
	newLibHydraModule := &LibHydra{
		target:          target,
		Credentials:     credChannel,
		logger:          logger,
		threads:         threads,
		parentWaitGroup: parentWaitGroup,
	}
	newLibHydraModule.ctx, newLibHydraModule.cancelFunc = context.WithCancel(ctx)
	return newLibHydraModule
}

func (hydra *LibHydra) AttachModule(module Module){
	hydra.module = module
}

func (hydra *LibHydra) StartCracking(){
	hydra.parentModuleName = hydra.module.getModuleInfo()
	hydra.module.preRunStart()
	hydra.parentWaitGroup.Add(hydra.threads)
	for i:=0;i<hydra.threads;i++{
		go hydra.checkCredentials(hydra.Credentials,hydra.parentWaitGroup)
	}
}

func (hydra *LibHydra) KillCrackingSession(){
	if hydra.kill == true{
		return
	}
	hydra.cancelFunc()
	hydra.module.kill()
	hydra.kill = true
	close(hydra.Credentials) // to kill all threads
}

func (hydra *LibHydra) checkCredentials(credentials chan CredManager.Cred,group *sync.WaitGroup){
	defer group.Done()
	select {
	case credential := <- credentials:
		if !hydra.kill{
			isValid := hydra.module.testCredential(hydra.target, credential.Username, credential.Password,hydra.logger,hydra.ctx)
			if !hydra.kill{
				hydra.lock.Lock()
				hydra.done += 1
				if hydra.done == hydra.Total {
					hydra.KillCrackingSession() // when total is known - else (dynamic cases ex credtracker) channel needs to be closed
				}
				hydra.lock.Unlock()
				hydra.logger.SendLog(ElasticLog.NewProgressLog(ModuleName, hydra.target,hydra.parentModuleName, hydra.done, hydra.Total))
			}
			if isValid && !hydra.kill{
				hydra.foundCred = true
				if hydra.findOneOnly {
					hydra.KillCrackingSession()
					hydra.kill = true // Update 2 - seems to work with different function to update the value in struct
					hydra.logger.SendLog(ElasticLog.NewProgressLog(ModuleName, hydra.target,hydra.parentModuleName, hydra.Total, hydra.Total))
				}
			}
		}
		case <- hydra.ctx.Done():
			return
	}

}

