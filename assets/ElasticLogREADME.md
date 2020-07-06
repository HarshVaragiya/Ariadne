# ElasticLog Module

- [X] Connecting to elasticsearch
    - Connects to `http://localhost:9200` to search for local elasticsearch instance
    - If not found looks for `ELASTICSEARCH_URL` in the environment variables and tries to connect

- [ ] Local logging 
    - If it is unable to connect to an elasticsearch instance, it shows you all the output
    - Log the output to a local file to be later imported by elasticsearch
