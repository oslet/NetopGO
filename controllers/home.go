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
	if err != nil {
		beego.Error(err)
	}
	hostNums, err := models.GetHostCount()
	if err != nil {
		beego.Error(err)
	}
	dbNums, err := models.GetDBCount()
	if err != nil {
		beego.Error(err)
	}
	nowSize, err := models.GetNowSize()
	if err != nil {
		beego.Error(err)
	}
	slows, err := models.GetSlowOverview()
	if err != nil {
		beego.Error(err)
	}
	sizeChange, err := models.GetSizeChange()
	if err != nil {
		beego.Error(err)
	}
	dbRecordNums, err := models.GetDBRecordMonth()
	if err != nil {
		beego.Error(err)
	}
	appRecordNums, err := models.GetAppRecordMonth()
	if err != nil {
		beego.Error(err)
	}
	this.Data["AppRecordNums"] = appRecordNums
	this.Data["DBRecordNums"] = dbRecordNums
	this.Data["SizeChange"] = sizeChange
	this.Data["Slows"] = slows
	this.Data["NowSize"] = nowSize
	this.Data["UserNums"] = userNums
	this.Data["HostNums"] = hostNums
	this.Data["DbNums"] = dbNums
	this.Data["TotalTimes"] = times
	this.Data["TotalSizes"] = sizes
	this.TplName = "index.html"
}
