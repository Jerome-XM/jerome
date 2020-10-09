package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Invoke struct {
	TxId 	string `orm:"column(txId)" json:"txHash"`
	Remark  string	`json:"content"`
}

/**
获取未检测交易
 */
func GetInvoke() []Invoke {
	var o = orm.NewOrm()
	var invokes []Invoke
	_,err :=o.Raw("SELECT i.txId, i.remark FROM invokequeue i LEFT JOIN checkInfo c ON i.txId = c.txId WHERE c.id is null and i.remark <>''").QueryRows(&invokes)
	if err != nil {
		logs.Info("GetInvoke异常: "+err.Error())
	}
	return invokes
}

/**
 更新检测结果
 */
func UpdateInvoke(array []interface{}){
	o:= orm.NewOrm()
	p,err :=o.Raw("update invokequeue set remark = ? WHERE txid= ?","333","1108c1054035eb35a59333b9e955737df08ee1c312535c6b29fa3852305a40ff").Prepare()
	res,err := p.Exec("111","1108c1054035eb35a59333b9e955737df08ee1c312535c6b29fa3852305a40ff")
	res,err = p.Exec("222","40c6f7a06b98363aac61aeb25d977eb15541767c026ddc87feefa126f4de3d36")
	p.Close()
	if err == nil {
		num,_ := res.RowsAffected()
		fmt.Println("invokes update nums: ", num)
	}
}
/**
获取单个上报交易
*/
func GetReportInvoke(txId string) []orm.Params {
	var o = orm.NewOrm()
	var maps []orm.Params
	_,err :=o.Raw("select " +
		"c.txId txHash," +
		"i.facc fromAcct," +
		"i.tacc toAcct," +
		"i.remark content," +
		"IF(i.type=0,'normal','contract') type," +
		"'accept' op," +
		"UNIX_TIMESTAMP(invokeTime) createdAt " +
		"from invokequeue i,checkInfo c " +
		"where i.txId = c.txId " +
		"and c.isReport = 0 " +
		"and c.trigger = 1 " +
		"and c.txId = ?",txId).Values(&maps)
	if err != nil {
		logs.Info("GetReportInvoke异常: "+err.Error())
	}
	return maps
}

/**
获取所有上报交易
*/
func GetReportInvokes() []orm.Params {
	var o = orm.NewOrm()
	var maps []orm.Params
	_,err :=o.Raw("select " +
		"c.txId txHash," +
		"IFNULL(i.facc,'0x0') fromAcct," +
		"IFNULL(i.tacc,'0x0') toAcct," +
		"i.remark content," +
		"IF(i.type=0,'normal','contract') type," +
		"'accept' op," +
		"UNIX_TIMESTAMP(invokeTime) createdAt " +
		"from invokequeue i,checkInfo c " +
		"where i.txId = c.txId " +
		"and c.isReport = 0 " +
		"and c.`trigger` = 1").Values(&maps)
	if err != nil {
		logs.Info("GetReportInvokes异常: "+err.Error())
	}
	return maps
}