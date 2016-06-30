package main

import (
	"NetopGO/models"
	_ "NetopGO/routers"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	//"time"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", false, false)
	//orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.Debug = true
	beego.Run()
}

//test
