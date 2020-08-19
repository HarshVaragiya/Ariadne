package Core

import "strings"


func portServiceParser(services map[string][]uint16)(ret map[string][]uint16){
	ret = make(map[string][]uint16)
	var http []uint16
	var ftp []uint16
	for key,value := range services{
		if strings.Contains(key,"http"){
			http = append(http,value...)
		} else if strings.Contains(key,"ftp"){
			ftp = append(ftp,value...)
		} else {
			ret[key] = value
		}
	}
	ret["http"] = http
	ret["ftp"] = ftp
	return ret
}