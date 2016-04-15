package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id      int8
	Name    string `orm:size(100)`
	Passwd  string `orm:size(100)`
	Email   string `orm:size(50)`
	Dept    string `orm:size(20)`
	Created time.Time
	Auth    int
	Tel     string `orm:size(11)`
}

func init() {
	orm.RegisterModel(new(User))
}
