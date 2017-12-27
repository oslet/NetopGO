package main

import (
	"NetopGO/models"
	_ "NetopGO/routers"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	//"time"
	"github.com/robfig/cron"
)

func init() {
	models.RegisterDB()
	//orm.RunSyncdb("default", false, false)
	orm.RunSyncdb("default", false, true)

	c := cron.New()
	spec := "0 */10 * * * 1-5"

	c.AddFunc(spec, func() {
		models.TasksForDailyReport()
	})

	c.Start()
	//select {} //阻塞主线程不退出

}

func main() {
	//orm.Debug = true
	beego.Run()

}
