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
	//orm.RunSyncdb("default", false, false)
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	models.AddUser("netop", "nbs2010", "netop@tingyun.com", "15201481200", "1", "运维")
	//models.IsViewDBApprove("产品", "研发审批")
	beego.Run()
}
