package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	//"path"
	"strconv"
	"strings"
)

type AppRecordController struct {
	BaseController
}

func (this *AppRecordController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	list := beego.AppConfig.String("app_name")
	arrList := strings.Split(list, ",")

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetAppRecordCount()
	appRecs, _, err := models.GetAppRecords(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["List"] = arrList
	this.Data["paginator"] = res
	this.Data["AppRecords"] = appRecs
	this.Data["totals"] = total

	this.Data["Path1"] = "应用升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/app/list"
	this.Data["Category"] = "record/app"
	this.TplName = "app_record_list.html"
	return
}

func (this *AppRecordController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/app"
	list := beego.AppConfig.String("app_name")
	arrList := strings.Split(list, ",")

	this.Data["List"] = arrList
	this.Data["Path1"] = "应用升级记录"
	this.Data["Path2"] = "添加记录"
	this.Data["Href"] = "/record/app/list"
	this.TplName = "app_record_add.html"
	return
}

func (this *AppRecordController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "record/app"

	id := this.Input().Get("id")
	group := this.Input().Get("group")
	operation := this.Input().Get("operation")
	appname := this.Input().Get("appname")
	disthost := this.Input().Get("disthost")
	isauto := this.Input().Get("isauto")
	applicant := this.Input().Get("applicant")
	content := this.Input().Get("content")
	operater := uname.(string)

	if len(id) > 0 {
		appRec, err := models.AppRecordDetail(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["AppRecord"] = appRec
		this.Data["Path1"] = "应用升级记录"
		this.Data["Path2"] = "操作内容"
		this.Data["Href"] = "/record/app/list"
		this.TplName = "app_record_detail.html"
		return
	} else {
		err := models.AddAppRecord(group, operation, appname, disthost, isauto, applicant, content, operater)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Data["Path1"] = "应用升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/app/list"
	this.Redirect("/record/app/list", 302)
	return
}

func (this *AppRecordController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/app"

	id := this.Input().Get("id")
	err := models.DeleteAppRecord(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "应用升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/app/list"
	this.Redirect("/record/app/list", 302)
	return
}

func (this *AppRecordController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/app"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteAppRecord(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *AppRecordController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/app"

	appname := this.Input().Get("keyword")
	if "1" == appname {
		this.Data["Path1"] = "应用升级记录"
		this.Data["Path2"] = ""
		this.Data["Href"] = "/record/app/list"
		this.Redirect("/record/app/list", 302)
		return
	}
	//beego.Info(appname)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchAppRecCount(appname)
	appRecs, err := models.SearchAppRecByName(int(currPage), int(pageSize), appname)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["AppRecords"] = appRecs
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = appname
	this.Data["Path1"] = "应用升级记录"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/record/app/list"
	this.TplName = "app_record_list.html"
	return
}
