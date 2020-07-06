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
		usernames = append(usernames, strings.Split(line,":")[0])
		passwords = append(passwords, strings.SplitAfter(line,":")[1])
	}
	// fmt.Println("Generated Credlist")
	return usernames,passwords,nil
}

func Contains(set []string,value string)bool{
	if set == nil {
		return false
	}
	for _ , v := range set{
		if value == v {
			return true
		}
	}
	return false
}