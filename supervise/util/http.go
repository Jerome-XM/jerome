package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)


/**
发送请求到接口
*/
func SendRequestWidthKey(serverPath string, parama string) (string,error) {
	key := GetKey(serverPath,parama)
	client := &http.Client{}
	req, err := http.NewRequest("POST", Server+serverPath, strings.NewReader(parama))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization_v2",key)
	resp, err := client.Do(req)


	if err != nil {
		return "",err
	}
	if resp.StatusCode!= 200{
		return "",errors.New("错误码:"+ strconv.Itoa(resp.StatusCode) )
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(body)

	return jsonStr,nil
}