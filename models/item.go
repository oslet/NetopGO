package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func GetAllSize() ([]string, []float64, error) {
	o := orm.NewOrm()
	var time []string
	var size []float64
	_, err := o.Raw("select `timestamp` from all_size").QueryRows(&time)
	if err != nil {
		beego.Error(err)
	}
	_, err = o.Raw("select size from all_size").QueryRows(&size)
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
	beego.Info(size)
	return size, err
}
