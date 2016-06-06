package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"strings"
)

type AppWOController struct {
	BaseController
}

func (this *AppWOController) Get() {
	var page string
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetAppOrderCount()
	appwos, _, err := models.GetAppOrders(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	for _, appwo := range appwos {
		appwo.Isapproved = models.IsApproved("app", dept.(string), appwo.Status, appwo.Upgradetype)
		if "研发" == dept.(string) || "运维" == dept.(string) {
			appwo.Isedit = appwo.Isapproved
		} else {
			appwo.Isedit = "false"
		}

	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")
	this.Data["AppNameList"] = appNameList
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["AppWorkOrders"] = appwos
	this.Data["totals"] = total

	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "我的工单"
	this.Data["Href"] = ""
	this.Data["Category"] = "record/app"
	this.TplName = "appworkorder_list.html"
	return
}

func (this *AppWOController) AppOrder() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "提交应用工单"
	this.Data["Href"] = "/workorder/my"
	this.Data["Category"] = "workorder/app"
	this.TplName = "appworkorder.html"
	return
}

func (this *AppWOController) AppOrderPost() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	apptype := this.Input().Get("apptype")
	appname := this.Input().Get("appname")
	version := this.Input().Get("version")
	jenkinsname := this.Input().Get("jenkinsname")
	buildnum := this.Input().Get("buildnum")
	featurelist := this.Input().Get("featurelist")
	modifycfg := this.Input().Get("modifycfg")
	relayapp := this.Input().Get("relayapp")
	upgradetype := this.Input().Get("upgradetype")
	sponsor := uname.(string)

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var sqlfile string
	if fh != nil {
		attachment = fh.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("fileupload", attachment))
		if err != nil {
			beego.Error(err)
		}
	}
	if sql != nil {
		sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("fileupload", sqlfile))
		if err != nil {
			beego.Error(err)
		}
	}

	err = models.AddAppOrder(apptype, appname, version, jenkinsname, buildnum, featurelist, modifycfg, relayapp, upgradetype, sponsor, attachment, sqlfile)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Category"] = "workorder/app"
	this.TplName = "index.html"
	return
}

func (this *AppWOController) Approve() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	id := this.Input().Get("id")
	appwo, err := models.GetAppwoById(id)

	this.Data["Appwo"] = appwo
	this.Data["Auth"] = role.(string)
	this.TplName = "approve.html"
	return
}
