package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Task struct {
	TaskId string `orm:"column(taskId);pk"`
	Status string
	Height int
	Offset int
}
func init(){
	orm.RegisterModel(new(Task))
}
func (t Task) TableName() string {
	return "checkTask"
}
func (t *Task) InsertOrUpdate(task *Task) {
	var o = orm.NewOrm()
	_,err := o.InsertOrUpdate(task,"taskId")
	if err != nil {
		logs.Info("任务信息插入失败:"+err.Error())
	}
}

func GetTask(taskId string) Task  {
	var o = orm.NewOrm()
	var task Task
	o.Raw("SELECT taskId,`status`,height,`offset` FROM checkTask WHERE taskId = ?",taskId).QueryRow(&task)
	return task
}
func UpdateTaskStatus(taskId string,status string,offset bool){
	var o = orm.NewOrm()
	var err error
	if offset {
		_,err = o.Raw("UPDATE checkTask set `status` = ? ,`offset` = height WHERE taskId = ?",status,taskId).Exec()
	}else {
		_,err = o.Raw("UPDATE checkTask set `status` = ?  WHERE taskId = ?",status,taskId).Exec()
	}

	if err != nil {
		logs.Info("更新巡检任务异常:"+err.Error())
	}
}
