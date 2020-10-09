package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"supervise/server"
	"supervise/util"
)

type TransaferController struct {
	beego.Controller
}

func (c *TransaferController) URLMapping()  {
	c.Mapping("sensitivWords",c.SensitivWords)
}
//敏感词检测接口
// @router /sensitivWords [post]
func (c *TransaferController)SensitivWords(){
	jsonData :=string(c.Ctx.Input.RequestBody)
	fmt.Println("--------------------敏感词检测接口---------------------")
	fmt.Println("接收到参数:"+jsonData)
	serverUrl := "v1/reg/kw"
	server.CheckOne(serverUrl,jsonData)
	c.Data["json"] = util.SuccResp("成功",nil)
	c.ServeJSON()
}
