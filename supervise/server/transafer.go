package server

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"
	"supervise/models"
	"supervise/util"
)

/**
巡检检测  isReport  true:定时批量检查单笔上报  fasle:下单巡检任务批量巡检
 */
func CheckAll(serverPath string,taskId string,isReport bool)  error{
	invoks := models.GetInvoke()
	if len(invoks)<=100 && invoks !=nil {//接口每次最大支持100条
		logs.Info("单次数据提交小于100条")
		body, _ := json.Marshal(invoks)
		//检测
		jsonStr,err := util.SendRequestWidthKey(serverPath,string(body))
		if err != nil {
			logs.Info("敏感词检测异常:"+err.Error())
			return err
		}
		logs.Info("敏感词检测结果:"+jsonStr)
		jsonData, err := simplejson.NewJson([]byte(jsonStr))
		checkResult(jsonData,isReport,taskId)
	}else if len(invoks)>100 {//超过100条对数组进行分割
	invokArr := util.SplitArray(invoks,100)
		logs.Info("单次数据提交大于100条")
		for i := 0; i < len(invokArr); i++ {
			body, _ := json.Marshal(invokArr[i])
			jsonStr,err := util.SendRequestWidthKey(serverPath,string(body))
			if err != nil {
				logs.Info("敏感词检测异常:"+jsonStr)
				return err
			}
			logs.Info("敏感词检测结果:"+jsonStr)
			jsonData, err := simplejson.NewJson([]byte(jsonStr))
			checkResult(jsonData,isReport,taskId)
		}
	}else {
		logs.Info("巡检任务没有发现需要检测的交易")
	}
	//巡检上报
	if !isReport {
		InspectionReport("/v1/reg/inspection/report",taskId)
		models.UpdateTaskStatus(taskId,"complete",true)
	}
	return nil
}
/**
链前单笔交易检测
 */
func CheckOne(serverPath string,body string){
	jsonStr,err := util.SendRequestWidthKey(serverPath,body)
	if err != nil {
		fmt.Println(jsonStr)
	}
	jsonData, err := simplejson.NewJson([]byte(jsonStr))
	checkResult(jsonData,true,"")
}
/**
处理检测结果
 */
func checkResult(jsonData *simplejson.Json,isReport bool,taskId string)  {
	code,err := jsonData.Get("code").Int()
	if err != nil {
		logs.Info("获取返回值code失败:"+err.Error())
	}
	if code == 0 {
		var check =models.Check{}
			data := jsonData.Get("data")
			check.CheckId,_ = data.Get("txHash").String()
		if !isReport {
			check.TaskId = taskId
		}
			//解析hits
			hitsData(data,&check,isReport)
	}
}

/**
解析hits参数
 */
func hitsData(data *simplejson.Json,check *models.Check,isReport bool)  {
	hits,err := data.Get("hits").Array()
	if err != nil {
		logs.Info("hits解析失败:"+err.Error())
	}
	for i := 0; i < len(hits); i++ {
		hitsData := hits[i].(map[string]interface{})
		check.Txid = hitsData["txHash"].(string)
		trigger,_ := hitsData["trigger"].(bool)
		if trigger {
			check.Trigger = 1
		}else{
			check.Trigger = 0
		}
		check.Triggerlevel,_ = hitsData["level"].(json.Number).Int64()
		//保存解析结果
		check.SaveCheck(check)
		if trigger && isReport{//如果命中检测,进行上报
			Report(check.Txid,"/v1/reg/report")
		}
	}

}

/**
链前违规单笔上报
 */
func Report(txId string,serverPath string)  {
	invoke := models.GetReportInvoke(txId)
	if invoke != nil {
		for i := 0; i < len(invoke); i++ {
			body,_ := json.Marshal(invoke[i])
			jsonStr,err := util.SendRequestWidthKey(serverPath,string(body))
			logs.Info("单笔上报结果:"+jsonStr)
			if err != nil {
				fmt.Println(jsonStr)
			}
			txId = invoke[i]["txHash"].(string)
			 txidArr := []string{txId}
			reportResult(jsonStr,txidArr)
		}
	}
}
/**
巡检违规批量上报
 */
func InspectionReport(serverPath string,taskId string)  {
	invoke := models.GetReportInvokes()
	fmt.Printf("巡检违规批量上报数量:%d",len(invoke))
	if invoke != nil {
		params := make(map[string]interface{})
		params["taskId"] = taskId
		params["success"] = true
		params["message"] = "成功"
		var jsonStr string
		var err error
		if len(invoke) <=100{
			params["status"] = true
			params["result"] = invoke
			body, _ := json.Marshal(params)
			jsonStr,err = util.SendRequestWidthKey(serverPath,string(body))
			logs.Info("巡检上报小于100返回结果:"+jsonStr)
			if err != nil {
				logs.Info("巡检上报小于100异常:"+jsonStr)
			}
		}else {
			invokArr := util.SplitArray(invoke,100)
			for i := 0; i < len(invokArr); i++ {
				params["status"] = false
				if i == len(invokArr)-1 {
					params["status"] = true
				}
				params["result"] = invokArr[i]
				body, _ := json.Marshal(params)
				jsonStr,err = util.SendRequestWidthKey(serverPath,string(body))
				logs.Info("巡检上报大于100返回结果:"+jsonStr)
				if err != nil {
					logs.Info("巡检上报大于100异常:"+err.Error())
				}

			}

		}
		txidArr := []string{}
		for i := 0; i < len(invoke); i++ {
			txId := invoke[i]["txHash"].(string)
			txidArr = append(txidArr, txId)
		}
		reportResult(jsonStr,txidArr)
	}
}



/**
上报返回值解析方法
 */
func reportResult(jsonStr string,txId []string)  {
	jsonData, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		logs.Info("上报解析结果异常:"+err.Error())
	}
	code,_:=jsonData.Get("code").Int()
	var check =models.Check{}
	var checkArr []models.Check
	if code==0 {
		taskId, _ := jsonData.Get("data").Get("txHash").String()
		check.Isreport = 1
		check.TaskId = taskId
	}
	if code != 0{
		message, _ := jsonData.Get("message").String()
		/*r := strings.Index(message,"{")
		message = message[r:]*/
		check.Common = message
		requestId, _ := jsonData.Get("data").Get("requestId").String()
		check.TaskId = requestId
	}
	for i := 0; i < len(txId); i++ {
		check.Txid = txId[i]
		checkArr = append(checkArr, check)
	}
	check.UpdateCheck(checkArr)
}