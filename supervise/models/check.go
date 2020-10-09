package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Check struct {
	Id   		 int
	Txid 		 string	`orm:"column(txId)"`
	Dealtype 	 int	`orm:"column(dealType)"`
	Isreport 	 int	`orm:"column(isReport)"`
	Trigger 	 int
	Triggerlevel int64	`orm:"column(triggerLevel)"`
	Op 			 string
	TaskId		 string `orm:"column(taskId)"`
	CheckId		 string `orm:"column(checkId)"`
	Common  	 string
}

func init(){
	orm.RegisterModel(new(Check))
}

func (c Check) TableName() string {
	return "checkInfo"
}

func (c *Check) SaveCheck(ch *Check)  {
	var o = orm.NewOrm()
	_,err := o.InsertOrUpdate(ch,"txid")
	if err != nil {
		logs.Info("检测信息插入失败:"+err.Error())
	}
}

func UpdateOp(txId string,op string)  {
	var o = orm.NewOrm()
	_,err := o.Raw("UPDATE checkInfo set op = ? WHERE txId = ?",op,txId).Exec()
	if err != nil {
		logs.Info("更新op异常:"+err.Error())
	}
}

func (c *Check) UpdateCheck(ch []Check)  {
	var o = orm.NewOrm()
	p,err :=o.Raw("UPDATE checkInfo SET isReport = ?,taskId = ?,common = ? WHERE txId = ?").Prepare()
	if err != nil {
		logs.Info("UpdateCheck异常:"+err.Error())
	}
	for i := 0; i < len(ch); i++ {
		isReport := ch[i].Isreport
		taskId := ch[i].TaskId
		common := ch[i].Common
		txid := ch[i].Txid
		_,err := p.Exec(isReport,taskId,common,txid)
		if err != nil {
			logs.Info("报告信息更新失败:"+err.Error())
		}
	}

	p.Close()
}