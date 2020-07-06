# Ariadne
Shows path to root

## Modules 

Each Module for Ariadne adds more functionality for enumeration of the target

- [X] ElasticLog
    - [Documentation](assets/ElasticLogREADME.md)
    - Logging module that logs everything to elasticsearch for easy indexing and search with kibana
    - Failsover to logging locally if elasticsearch is not configured

- [X] CredManager
    - [Documentation](assets/CredManagerREADME.md)
    - Module for managing credentials for Hydra and other modules
    - Spawn Goroutines that track "credentials.txt" or any such file for new credentials
    and supply them to other modules to test if they work

- [X] HTTP 
    - [Documentation](assets/HTTPREADME.md)
    - GoBusterDir mode for finding files on http ports
    - GoBusterVhost mode to be added later

- [X] Hydra
    - [Documentation](assets/HydraREADME.md)
    - Service Communication Module
    - Test Default credentials,anonymous logins for services
    - Can be integrated with CredTrackers that track new credentials and try them on the target

- [X] NMAP
    - [Documentation](assets/NmapREADME.md)
    - Scans target ports systematically to find services
    - Multi-stage operation
