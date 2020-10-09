package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Block struct {
	Num int `json:"height"`
	PrevisourHash string `orm:"column(previousHash)" json:"parentHash"`
	DataHash string `orm:"column(dataHash)" json:"hash"`
	InsertTime int `orm:"column(insertTime)" json:"createdAt"`
	Txs []Txs `json:"txs"`
}

type Txs struct{
	TxId string `orm:"column(txId)" json:"hash"`
	Facc string `json:"fromAcct"`
	Tacc string	`json:"toAcct"`
}

func GetBlock(blockNum int) []Block{
	var o = orm.NewOrm()
	var blocks []Block
	_,err :=o.Raw("SELECT num,previousHash,dataHash,UNIX_TIMESTAMP(insertTime) insertTime from blockbasicinfo WHERE num >= ? ORDER BY num desc",blockNum).QueryRows(&blocks)
	if err != nil {
		logs.Info("心跳查询区块信息nums: ", err.Error())
	}
	return blocks
}

func GetBlockMax() Block{
	var o = orm.NewOrm()
	var blocks Block
	err := o.Raw("SELECT num,previousHash,dataHash,UNIX_TIMESTAMP(insertTime) insertTime from blockbasicinfo ORDER BY num desc limit 1").QueryRow(&blocks)
	if err != nil {
		logs.Info("查询区块最大高度信息异常: ", err.Error())
	}
	return blocks
}

func GetTransByNum(blockNum int) []Txs {
	var o = orm.NewOrm()
	var txs []Txs
	_,err :=o.Raw("SELECT txId,IFNULL(facc,'0x0') facc,IFNULL(tacc,'0x0') tacc from invokequeue  WHERE blockNum = ?",blockNum).QueryRows(&txs)
	if err != nil {
		logs.Info("心跳交易信息nums: ", err.Error())
	}
	return txs

}