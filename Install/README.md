# Installing Ariadne

1. Run install script located at [Install/install.sh](install.sh) from `Install` folder.
    ```bash
    ~/go/src/Ariadne/Install $ ./install.sh
    ```
2. Building Ariadne
    ```bash
    go build main.go -o <binary-name>
    ```

## Explanation for install script
1. Install necessary packages `golang,nmap,git`
    ```bash
    $ sudo apt install golang nmap git -y
    ```
2. Install necessary golang dependencies 
    ```bash
    go get github.com/sirupsen/logrus
    go get github.com/jlaffaye/ftp
    go get github.com/akamensky/argparse
    go get github.com/Ullaakut/nmap
    go get github.com/elastic/go-elasticsearch
    ```
3. Apply the libgobuster patch to gobuster that you just downloaded
   ```bash
   patch -u -b ~/go/src/github.com/OJ/gobuster/libgobuster/libgobuster.go -i "$PWD/libgobuster.patch" 
   ```
   This patch exports a few variables `Mu (sync.RWMutex)`, `RequestsIssues and RequestsExpected (int)`
   which are necessary for finding progress made by libgobuster
   