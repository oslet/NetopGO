package controllers

import (
	"NetopGO/models"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

type DailyreportController struct {
	BaseController
}

func (this *DailyreportController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "workorder/dailyreport"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "程序更新报表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/workorder/dailyreport/list"

	/*
		groups, err := models.GetNames()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Groups"] = groups
	*/
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetDailyreportCount()
	Dailyreportlists, _, err := models.GetDailyreports(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Dailyreportlists"] = Dailyreportlists
	this.Data["totals"] = total

	this.TplName = "dailyreport_list.html"
	return
}
func (this *DailyreportController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dailyreport"
	Auth := role.(int64)
	this.Data["Auth"] = Auth
	/*
		groups, err := models.GetNames()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Groups"] = groups
	*/
	id := this.Input().Get("id")
	if len(id) > 0 {
		dailyreportlist, err := models.GetDailyreportlistById(id)
		if err != nil {
			beego.Error(err)
		}
		//dailyreportlist.Rootpwd, _ = models.AESDecode(dailyreportlist.Rootpwd, models.AesKey)
		//dailyreportlist.Readpwd, _ = models.AESDecode(dailyreportlist.Readpwd, models.AesKey)
		//fmt.Printf("***root :%v,read :%v\n", host.Rootpwd, host.Readpwd)
		this.Data["Dailyreportlist"] = dailyreportlist
		//	this.Data["HostGroupName"] = host.Group
		this.Data["Path1"] = "程序更新报表"
		this.Data["Path2"] = "修改更新报表"
		this.Data["Href"] = "/workorder/dailyreport/list"
		this.TplName = "dailyreport_modify.html"
		//this.TplName = "test.html"
		return
	}
	this.Data["Path1"] = "程序更新报表"
	this.Data["Path2"] = "添加报表"
	this.Data["Href"] = "/workorder/dailyreport/list"
	this.TplName = "dailyreport_add.html"

}

func (this *DailyreportController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "dailyreport"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	appsys := this.Input().Get("appsys")
	appname := this.Input().Get("appname")
	appcontent := this.Input().Get("appcontent")
	applicgrp := this.Input().Get("applicgrp")
	applicant := this.Input().Get("applicant")
	publisher := this.Input().Get("publisher")
	department := this.Input().Get("department")
	publishtime := this.Input().Get("publishtime")
	followstatus := this.Input().Get("followstatus")
	followman := this.Input().Get("followman")
	isinitial := this.Input().Get("isinitial")
	//beego.Info(idc)
	if len(id) > 0 {
		err, msg := models.ModifyDailyreportlist(id, appsys, appname, appcontent, applicgrp, applicant, publisher, department, publishtime, followstatus, followman, isinitial)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddDailyreportlist(appsys, appname, appcontent, applicgrp, applicant, publisher, department, publishtime, followstatus, followman, isinitial)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "程序更新报表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/workorder/dailyreport/list"
	//this.Redirect("/workorder/dailyreport/list", 302)
	this.TplName = "dailyreport_add.html"
}

func (this *DailyreportController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dailyreport"

	id := this.Input().Get("id")
	err := models.DeleteDailyreportlist(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "程序更新报表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/workorder/dailyreport/list"
	this.Redirect("/workorder/dailyreport/list", 302)
	return
}

func (this *DailyreportController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dailyreport"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteDailyreportlist(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *DailyreportController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "workorder/dailyreport"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchDailyreportlistCount(name)
	dailyreports, err := models.SearchDailyreportlistByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Dailyreportlists"] = dailyreports
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "dailyreports列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/workorder/dailyreport/list"
	this.TplName = "dailyreport_list.html"
	return
}

func (this *DailyreportController) Detail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	id := this.Input().Get("id")
	dailyreport, err := models.GetDailyreportById(id)
	if err != nil {
		beego.Error(err)
	}
	auth := role.(int64)
	this.Data["Auth"] = auth

	this.Data["Dailyreport"] = dailyreport
	this.Data["Path1"] = "程序更新报表"
	this.Data["Path2"] = "详情"
	this.Data["Href"] = "/workorder/dailyreport/list"
	this.Data["Category"] = "dailyreport"
	this.TplName = "dailyreport_detail.html"
	return
}

func (this *DailyreportController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dailyreport"
	values, columns, _ := models.QueryDailyreportExport()

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	for _, val := range columns {
		cell := row.AddCell()
		cell.Value = val
	}
	for _, val := range *values {
		row = sheet.AddRow()
		for _, value := range val {
			cell := row.AddCell()
			cell.Value = value
		}
	}
	now := time.Now().String()
	filename := "all_dailyreport" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

	filepath := path.Join("export", filename)
	err := file.Save(filepath)
	if err != nil {
		beego.Error(err)
	}
	defer func() {
		os.Remove(filepath)
	}()
	this.Ctx.Output.Download(filepath, filename)
	return
}
