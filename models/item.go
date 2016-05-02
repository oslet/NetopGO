package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
