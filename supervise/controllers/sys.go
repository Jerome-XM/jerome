package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"
	"supervise/models"
	"supervise/server"
	"supervise/util"
)

type SysController struct {
	beego.Controller
}
func (s *SysController) URLMapping()  {
	s.Mapping("Inspection",s.Inspection)
	s.Mapping("Cmd",s.Cmd)
	s.Mapping("Heartbeat",s.Heartbeat)
	s.Mapping("InspectionGet",s.InspectionGet)
	s.Mapping("InspectionDelete",s.InspectionDelete)
}
//下达巡检指令
// @router /inspection [post]
func (s *SysController)Inspection(){
	jsonStr :=string(s.Ctx.Input.RequestBody)
	logs.Info("--------------------监管指令---------------------")
	logs.Info("接收到参数:"+jsonStr)
	jsonData, _ := simplejson.NewJson([]byte(jsonStr))
	taskId,_ := jsonData.Get("taskId").String()
	var task models.Task
	task.TaskId = taskId
	task.Status = "processing"
	block := models.GetBlockMax()
	task.Height = block.Num
	task.InsertOrUpdate(&task)
	server.CheckAll("/v1/reg/kw",taskId,false)
	s.Data["json"] = util.SuccResp("成功",nil)
	s.ServeJSON()
}

//下发管控指令
// @router /cmd  [post]
func (s *SysController)Cmd(){
	jsonStr :=string(s.Ctx.Input.RequestBody)
	logs.Info("--------------------管控指令---------------------")
	logs.Info("接收到参数:"+jsonStr)
	data := server.Cmd(jsonStr)
	s.Data["json"] = util.SuccResp("ok",data)
	s.ServeJSON()
}

//心跳检测
// @router /heartbeat  [post]
func (s *SysController)Heartbeat(){
	jsonStr :=string(s.Ctx.Input.RequestBody)
	logs.Info("--------------------心跳检测---------------------")
	logs.Info("接收到参数:"+jsonStr)
	data := server.Heartbeat(jsonStr)
	s.Data["json"] = util.SuccResp("成功",data)
	s.ServeJSON()
}

//查询巡检状态
// @router /inspection/:taskId [get]
func (s *SysController)InspectionGet(){
	taskId :=s.Ctx.Input.Param(":taskId")
	logs.Info("--------------------查看巡检状态---------------------")
	logs.Info("接收到参数taskId:"+taskId)
	data := server.GetTask(taskId)
	if data["status"] == "" {
		s.Data["json"] = util.FalseResp("没有查询到任务")
	}else{
		s.Data["json"] = util.SuccResp("ok",data)
	}
	s.ServeJSON()
}
//取消巡检指令
// @router /inspection/:taskId [delete]
func (s *SysController)InspectionDelete(){
	taskId :=s.Ctx.Input.Param(":taskId")
	logs.Info("--------------------取消巡检状态---------------------")
	logs.Info("接收到参数taskId:"+taskId)
	models.UpdateTaskStatus(taskId,"none",false)
	s.Data["json"] = util.SuccResp("ok",nil)
	s.ServeJSON()
}