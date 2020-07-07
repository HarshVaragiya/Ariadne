# ElasticLog Module

## Example Usage

1. Declaring a logger
    ```go
    logger := ElasticLog.NewElasticLogger("defaultIndex")
    ```
   Note : the function NewElasticLogger returns a pointer to the logger (type `*ElasticLog.Logger`)
   
2. logging to the logger

    - simple logs of type "DEBUG", "INFO" etc can be easily logged by using
        ```go
        logger.SendLog(ElasticLog.NewLog("ERROR",Error,ModuleName))
        ```
    - progress logs can be logged by using
        ```go
        hydra.logger.SendLog(ElasticLog.NewProgressLog(ModuleName, Target, Done, Total))
        ```
3. passing the logger around

    - functions can receive the exact logger when passed as reference 
        ```go
        func myFunction(logger *ElasticLog.Logger){}   
        ```
    