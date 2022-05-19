package utils

import (
	"encoding/json"
	"fmt"
)

//print struct json format to console
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

//check if list constains element
func ContainsString(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

//check duplicates
func DupesCheck(list []string) map[string]int {
	dupesCount := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the dupesCount map
		_, exist := dupesCount[item]

		if exist {
			dupesCount[item] += 1 // increase counter by 1 if already in the map
		} else {
			dupesCount[item] = 1 // else start counting from 1
		}
	}

	return dupesCount
}

func PrintAppInfo(appName string, env string, dbname string, dbuser string, dbhost string, dbport string, idleConn int, maxConn int, host string, port string, brokersList []string, topicsList []string) {
	fmt.Printf("\nAPP NAME: %s\nENV: %s\nDB NAME: %s\nDB USER: %s\nDB HOST: %s\nDB PORT: %s\nDB IDLE CONN: %d\nDB MAX CONN: %d\nHOST: %s\nPORT: %s\nBROKERS:%v\nTOPICS:%v", appName, env, dbname, dbuser, dbhost, dbport, idleConn, maxConn, host, port, brokersList, topicsList)
}
