# Hydra Module

## Making a Module for Hydra
LibHydra requires these 4 methods to be implemented for a Module
```go
type Module interface {
	testCredential(string,string,string,*ElasticLog.Logger) bool
	preRunStart()
	kill()
	getModuleInfo()string
}
```
## Example Module 
**FTPCrack Module**
```go
type FTPCrack struct {
	ModuleName string
	isAlive bool
}
```
Module structure definition

## Function Implementation

### testCredentials function
- Function Signature
    ```go
    testCredential(target,username,password string,logger *ElasticLog.Logger) bool
    ```
- takes target URL/ip (depending on module this needs to be constructed),
    username and password and performs authentication check to see validity

- returns true if authentication is successful, logs the output in any case

### preRunStart,kill functions
Boilerplate code for controlling the module life
```go
func (ftpCrack *FTPCrack) preRunStart(){
    ftpCrack.isAlive = true
}

func (ftpCrack *FTPCrack) kill(){
    ftpCrack.isAlive = false
}
```

### getModuleInfo function
```go
func (ftpCrack *FTPCrack) getModuleInfo()string{
	ftpCrack.ModuleName = FtpModuleName
	return ftpCrack.ModuleName
}
```
returns Module name for use by libhydra for logging purposes.

## That is it.