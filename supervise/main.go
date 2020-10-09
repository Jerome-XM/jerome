package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "supervise/routers"
	"supervise/server"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	server.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()

	logs.SetLogger(logs.AdapterFile,`{"filename":"superviseLog/supervise.log"}`)
	logs.EnableFuncCallDepth(true)
	beego.Run()
}

