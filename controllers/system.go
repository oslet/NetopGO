package controllers

import (
	"NetopGO/models"
	//"fmt"
	"strconv"
	"strings"

	"os"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

type SystemController struct {
	BaseController
}

func (this *SystemController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "asset/system"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/system/list"

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
	total, err := models.GetSystemCount()
	Systemlists, _, err := models.GetSystems(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Systemlists"] = Systemlists
	this.Data["totals"] = total

	this.TplName = "system_list.html"
	return
}
func (this *SystemController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "system"
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
		systemlist, err := models.GetSystemlistById(id)
		if err != nil {
			beego.Error(err)
		}
		//systemlist.Rootpwd, _ = models.AESDecode(systemlist.Rootpwd, models.AesKey)
		//systemlist.Readpwd, _ = models.AESDecode(systemlist.Readpwd, models.AesKey)
		//fmt.Printf("***root :%v,read :%v\n", host.Rootpwd, host.Readpwd)
		this.Data["Systemlist"] = systemlist
		//	this.Data["HostGroupName"] = host.Group
		this.Data["Path1"] = "系统列表"
		this.Data["Path2"] = "修改系统应用"
		this.Data["Href"] = "/asset/system/list"
		this.TplName = "system_modify.html"
		//this.TplName = "test.html"
		return
	}
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = "添加系统"
	this.Data["Href"] = "/asset/system/list"
	this.TplName = "system_add.html"

}

func (this *SystemController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "system"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	class := this.Input().Get("class")
	name := this.Input().Get("name")
	owner1 := this.Input().Get("owner1")
	owner2 := this.Input().Get("owner2")
	domain_name := this.Input().Get("domain_name")
	controller := this.Input().Get("controller")
	responsible := this.Input().Get("responsible")
	team := this.Input().Get("team")
	company := this.Input().Get("company")
	support_level := this.Input().Get("support_level")
	comment := this.Input().Get("comment")
	//beego.Info(idc)
	if len(id) > 0 {
		err, msg := models.ModifySystemlist(id, class, name, owner1, owner2, domain_name, controller, responsible, team, company, support_level, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddSystemlist(class, name, owner1, owner2, domain_name, controller, responsible, team, company, support_level, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/system/list"
	//this.Redirect("/asset/system/list", 302)
	this.TplName = "system_add.html"
}

func (this *SystemController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "system"

	id := this.Input().Get("id")
	err := models.DeleteSystemlist(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/system/list"
	this.Redirect("/asset/system/list", 302)
	return
}

func (this *SystemController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "system"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteSystemlist(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *SystemController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "asset/system"

	name := this.Input().Get("keyword")

	class := this.Input().Get("class")
	if class == "1" {
		this.Data["Path1"] = "系统列表"
		this.Data["Path2"] = ""
		this.Data["Href"] = "/asset/system/list"
		this.TplName = "system_list.html"
		return
	}

	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchSystemlistCount(class, name)
	Systemlists, err := models.SearchSystemlistByName(int(currPage), int(pageSize), class, name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Systemlists"] = Systemlists
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	//this.Data["Idc"] = idc
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/asset/system/list"
	this.TplName = "system_list.html"
	return
}

func (this *SystemController) Detail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	id := this.Input().Get("id")
	sysname, err := models.GetSystemById(id)
	if err != nil {
		beego.Error(err)
	}
	auth := role.(int64)
	this.Data["Auth"] = auth

	this.Data["System"] = sysname
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = "系统详情"
	this.Data["Href"] = "/asset/system/list"
	this.Data["Category"] = "system"
	this.TplName = "system_detail.html"
	return
}

func (this *SystemController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "system"
	values, columns, _ := models.QuerySystemExport()

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
	filename := "all_system" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

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
