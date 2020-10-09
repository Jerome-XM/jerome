package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main(){
	/*body := "{\"checkpoint\":1,\"taskId\":\"bd593933b91040fcb06739da3e057922\"}"
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	nonce := strconv.Itoa(rand.Intn(1000000000))
	fmt.Println(nonce+"长度:"+strconv.Itoa(len(nonce)))
	request := util.GetRequest("/v1/sys/heartbeat",body,timestamp,nonce)
	fmt.Println("request:"+request)
	signature := util.HmacSHA512(util.AppKey,request)
	fmt.Println("signature:"+signature)
	fmt.Println("JG-"+signature+"-"+timestamp+"-"+nonce)*/

	/*message:=[]byte("hello")
	hashCode:=GetSHA256HashCode(message)
	fmt.Println(hashCode)*/

	/*str := "不存在敏感词交易-txHash:{2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824}"
	ss := strings.Index(str,"{")
	str = str[ss:]
	fmt.Println(str[1:len(str)-1])*/
	a:=make(map[string]interface{})
	a["taskId"] = "1qaz2wsx"
	a["checkpoint"] = 1255
	r,err :=json.Marshal(a)
	parama := string(r)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://49.73.154.68:8090/v1/sys/heartbeat", strings.NewReader(parama))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode!= 200{
		logs.Info("错误码:"+ strconv.Itoa(resp.StatusCode) )
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(body)

	fmt.Println(jsonStr)

}

func GetSHA256HashCode(message []byte)string{
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}
