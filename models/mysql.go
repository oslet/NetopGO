package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_Driver = mysql
)

func RegisterDB() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_schema := beego.AppConfig.String("db_schema")
	db_user := beego.AppConfig.String("db_user")
	db_passwd := beego.AppConfig.String("db_passwd")

	jdbcUrl := 
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	orm.RegisterDataBase("default", _DB_Driver, , 30)
}
