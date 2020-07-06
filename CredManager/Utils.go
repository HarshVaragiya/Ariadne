package CredManager

import (
	"bufio"
	"os"
	"strings"
)

func GetCredentialsFromFile(filename string)(usernames,passwords []string,err error){
	file, err := os.Open(filename)
	if err != nil {
		return nil,nil,err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		username := strings.TrimSpace(strings.Split(line,":")[0])
		password := strings.TrimSpace(strings.SplitAfter(line,":")[1])
		if username != "" {
			usernames = append(usernames, username)
		}
		if password != "" {
			passwords = append(passwords, password)
		}
	}
	// fmt.Println("Generated Credlist")
	return usernames,passwords,nil
}

func Contains(set []string,value string)int{
	if set == nil {
		return -1
	}
	for index , v := range set{
		if value == v {
			return index
		}
	}
	return -1
}

func ensureFileExistence(filename string){
	outFile, err := os.Open(filename)
	if err !=nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			file , err := os.Create(filename)
			if err != nil {
				panic(err)
			}
			file.Close()
			return
		}
		panic(err)
	}
	outFile.Close()
}