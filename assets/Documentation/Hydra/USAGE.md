# Hydra Module

## Example usage of Hydra Module
1. Using Hydra to test for default credentials
    ```go
    wordlist := "/home/harsh/Desktop/HackTheBox/Wordlist/ftp-betterdefaultpasslist.txt"
    logger := ElasticLog.NewElasticLogger("defaultIndex")
    
    var wg sync.WaitGroup
    
    credList := CredManager.NewCredListFromFile(wordlist,false) // reads credentials from file
    newCracker := Hydra.NewLibHydraModule("192.168.1.1:21",credList,4,logger,&wg)
    newCracker.AttachModule(&Hydra.FTPCrack{})                  // attach to FTPCracker module
    newCracker.StartCracking()
    
    wg.Wait() // wait for newCracker threads to exit
    ```
   
2. Using Hydra with CredTracker to test credentials dynamically
    - See [CredManager - CredTracker](../CredManager/README.md)
   
   