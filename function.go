package main

import(
	"fmt"
	"crypto/sha256"
	"encoding/json"
)

func GetSHA256HashCode(msg string) string {
	hashcode := sha256.New()
	hashcode.Write([]byte(msg))
	return fmt.Sprintf("%x",hashcode.Sum(nil))
}

func GetJsonarraytoString(arr []Transaction) (res string) {
	if len(arr) == 0 {
		return
	}
	for _,i :=range arr{
		jsonval , err := json.Marshal(i)
		if err != nil {
        	        fmt.Println("账单格式错误",err)
        	}
		res += string(jsonval)+","
	}
	res = res[:len(res)-1]	//去最后一个,
	return
}
