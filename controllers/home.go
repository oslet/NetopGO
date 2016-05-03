package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	times, sizes, err := models.GetAllSize()
	if err != nil {
		beego.Error(err)
	}
	userNums, err := models.GetUserCount()
	hostNums, err := models.GetHostCount()
	dbNums, err := models.GetDBCount()
	nowSize, err := models.GetNowSize()
	this.Data["NowSize"] = nowSize
	this.Data["UserNums"] = userNums
	this.Data["HostNums"] = hostNums
	this.Data["DbNums"] = dbNums
	this.Data["TotalTimes"] = times
	this.Data["TotalSizes"] = sizes
	this.TplName = "index.html"
}
