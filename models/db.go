package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type DB struct {
	Id      int64
	Name    string `orm:size(50)`
	Uuid    string `orm:size(50)`
	Created time.Time
}

func init() {
	orm.RegisterModel(new(DB))
}
