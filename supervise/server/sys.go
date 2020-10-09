package server

import (
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"
	"supervise/models"
	"supervise/util"
)

/**
下发管控指令
 */
func Cmd(jsonStr string) map[string]string  {
  jsonData,err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		logs.Info("接收管控指令参数异常:"+err.Error())
	}
	txId, _ := jsonData.Get("txHash").String()
	op,_ := jsonData.Get("op").String()
	models.UpdateOp(txId,op)
	result := make(map[string]string)
	result["reviewType"] = "browser"
	result["reviewUrl"] = util.ReviewUrl
	return result
}
/**
心跳检测
 */
func Heartbeat(jsonStr string) map[string]interface{}{
	jsonData,err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		logs.Info("接收心跳参数异常:"+err.Error())
	}
	taskId,_ := jsonData.Get("taskId").String()
	checkpoint,_ := jsonData.Get("checkpoint").Int()
	if checkpoint == 0{
		checkpoint = 1
	}
	data := make(map[string]interface{})
	data["taskId"] = taskId
	//获取区块信息
	blocks := models.GetBlock(checkpoint)
	if blocks == nil {
		blocks = []models.Block{}
	}
	for i := 0; i < len(blocks); i++ {
	//根据区块查询交易信息
		num := blocks[i].Num
		txs := models.GetTransByNum(num)
		if txs == nil {
			txs = []models.Txs{}
		}
		blocks[i].Txs = txs
	}
	data["blocks"] = blocks
	if blocks != nil && len(blocks)>0 {
			checkpoint = checkpoint+len(blocks)
	}
	data["checkpoint"] = checkpoint
	return data
}
/**
保存巡检任务
 */
func SaveOrUpdateTask(taskId string,status string,hight int) map[string]interface{} {
	data := make(map[string]interface{})
	var task models.Task
	task.TaskId = taskId
	task.Height = hight
	task.Status = status
	task.InsertOrUpdate(&task)
	return data
}
/**
更新巡检任务
 */
func GetTask(taskId string) map[string]interface{} {
	data := make(map[string]interface{})
	task := models.GetTask(taskId)
	data["status"] = task.Status
	data["height"] = task.Height
	data["offset"] = task.Offset
	return data
}