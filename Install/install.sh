#!/usr/bin/env bash
echo "[+] Install Script for Ariadne"

echo '------------------------------------------------'
echo "[+] Installing golang,git,nmap"
sudo apt install golang git nmap

echo '------------------------------------------------'
echo "[+] Getting necessary dependencies"
read -rsp $'[+] Press any key to continue (or ctrl+c to quit)...\n' -n1 key
go get github.com/sirupsen/logrus
go get github.com/jlaffaye/ftp
go get github.com/akamensky/argparse
go get github.com/Ullaakut/nmap
go get github.com/elastic/go-elasticsearch

echo '------------------------------------------------'
echo "[+] Patching Gobuster"
read -rsp $'[+] Press any key to continue (or ctrl+c to quit)...\n' -n1 key
patch -u -b ~/go/src/github.com/OJ/gobuster/libgobuster/libgobuster.go -i "$PWD/libgobuster.patch"
echo "[+] Patched libgobuster/libgobuster.go"

echo '------------------------------------------------'
echo "[+] To build Ariadne:"
echo "$ make build"
echo "[+] from the Ariadne directory"

echo '------------------------------------------------'
echo "[+] To run Ariadne:"
echo "$ make run"
echo "[+] from the Ariadne directory"

