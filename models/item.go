package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type SlowOverview struct {
	Name  string
	Count int64
}

type SizeChange struct {
	Name string
	Size float64
}

func GetAllSize() ([]string, []float64, error) {
	o := orm.NewOrm()
	firstDay := time.Now().String()[:8] + "01 00:00:00"
	var time []string
	var size []float64
	_, err := o.Raw("select `timestamp` from all_size where `timestamp`>=? ", firstDay).QueryRows(&time)
	if err != nil {
		beego.Error(err)
	}
	_, err = o.Raw("select size from all_size where timestamp>=? ", firstDay).QueryRows(&size)
	if err != nil {
		beego.Error(err)
	}
	for i, value := range time {
		time[i] = string(value)[:10]
	}
	return time, size, err
}

func GetNowSize() (float64, error) {
	var size float64
	today := time.Now().String()[:10] + " 00:00:00"
	o := orm.NewOrm()
	err := o.Raw("select size from all_size where timestamp>=? limit 1", today).QueryRow(&size)
	//beego.Info(size)
	return size, err
}

func GetSlowOverview() ([]*SlowOverview, error) {
	o := orm.NewOrm()
	slows := make([]*SlowOverview, 0)
	today := time.Now().String()[:10] + " 00:00:00"
	_, err := o.Raw("select name,count from  slow_overview where timestamp>=? order by count desc limit 12", today).QueryRows(&slows)
	return slows, err
}

func GetSizeChange() ([]*SizeChange, error) {
	o := orm.NewOrm()
	sizeChange := make([]*SizeChange, 0)
	firstDay := time.Now().String()[:8] + "01 00:00:00"
	//beego.Info(firstDay)
	today := time.Now().String()[:10] + " 00:00:00"
	//beego.Info(today)
	_, err := o.Raw("select  a.name,(b.size-a.size) as size from  (select schemaname name,sum(size) size from inst_info where timestamp=? and name like '%master%' group by schemaname) a join  (select schemaname name,sum(size) size from inst_info where timestamp=? and name like '%master%' group by schemaname) b on a.name=b.name order by size desc", firstDay, today).QueryRows(&sizeChange)
	return sizeChange, err
}

func GetDBRecordMonth() (float64, error) {
	var num float64
	firstDay := time.Now().String()[:8] + "01 00:00:00"
	o := orm.NewOrm()
	err := o.Raw("select count(*) from dbworkorder where created>=?", firstDay).QueryRow(&num)
	return num, err
}

func GetAppRecordMonth() (float64, error) {
	var num float64
	firstDay := time.Now().String()[:8] + "01 00:00:00"
	o := orm.NewOrm()
	err := o.Raw("select count(*) from appworkorder where created>=?", firstDay).QueryRow(&num)
	return num, err
}

func GetQuestionRecordMonth() (float64, error) {
	var num float64
	firstDay := time.Now().String()[:8] + "01 00:00:00"
	o := orm.NewOrm()
	err := o.Raw("select count(*) from question where created>=? and status=?", firstDay, "挂起").QueryRow(&num)
	return num, err
}

func FaultRecordMonth() (float64, error) {
	var num float64
	firstDay := time.Now().String()[:8] + "01 00:00:00"
	o := orm.NewOrm()
	err := o.Raw("select count(*) from faultrecord where starttime>=?", firstDay).QueryRow(&num)
	return num, err
}

func GetUnoverOrderNums(auth int64, dept, uname string) (float64, int64, string, error) {
	o := orm.NewOrm()
	var num float64
	var err error
	var pageAuth int64
	var pageDept string
	if auth == 2 {
		err = o.Raw("select count(*) from  dbworkorder where status=?", "正在实施").QueryRow(&num)
		pageAuth = 2
		pageDept = "运维"
	} else if auth == 1 {
		err = o.Raw("select count(*) from  appworkorder where status=?", "实施流程中").QueryRow(&num)
		pageAuth = 1
		pageDept = "运维"
	} else if dept == "测试" {
		err = o.Raw("select count(*) from  appworkorder where status in (?,?)", "测试流程中", "验证流程中").QueryRow(&num)
		pageAuth = 3
		pageDept = "测试"
	} else if dept == "研发" {
		err = o.Raw("select count(*) from  appworkorder where status<>?", "工单已关闭").QueryRow(&num)
		pageAuth = 3
		pageDept = "研发"
	} else if dept == "产品" {
		err = o.Raw("select count(*) from  appworkorder where status=?", "审批流程中").QueryRow(&num)
		pageAuth = 3
		pageDept = "产品"
	} else {
		err = nil
		num = 0
		pageAuth = 3
		pageDept = ""
	}
	return num, pageAuth, pageDept, err
}
