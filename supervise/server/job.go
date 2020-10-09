package server

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func InitTask()  {
	tk := toolbox.NewTask("transaferCheck","0 */2 * * * *",TransaferCheck)
	toolbox.AddTask("transaferCheck",tk)
}

func TransaferCheck() error {
	logs.Info("定时查询检测交易任务开始")
	err :=CheckAll("/v1/reg/kw","",true)
	logs.Info("定时查询检测交易任务结束")
	return err
}