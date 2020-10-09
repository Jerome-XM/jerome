package util

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"sort"
	"strconv"
	"time"
)
/**
HmacSHA512加密
 */
func HmacSHA512(key string, data string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
/**
拼接请求
 */
func GetRequest(url string,body string,timestamp string,nonce string)  string{
	arrayString := []string{url,body,timestamp,nonce}
	sort.Strings(arrayString)
	var str = ""
	for i := 0; i < len(arrayString); i++ {
		str = str + arrayString[i]
	}
	return str
}

func GetKey(serverPath string,body string) string {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	nonce := strconv.Itoa(rand.Intn(1000000000))
	request := GetRequest(serverPath,body,timestamp,nonce)
	signature := HmacSHA512(AppKey,request)
	authorization_v2 := "JG-"+signature+"-"+timestamp+"-"+nonce
	return  authorization_v2
}


