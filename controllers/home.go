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
	this.Data["TotalTimes"] = times
	this.Data["TotalSizes"] = sizes
	this.TplName = "index.html"
}
