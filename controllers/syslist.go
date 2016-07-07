package controllers

import (
	"NetopGO/models"
	//"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type SyslistController struct {
	BaseController
}

func (this *SyslistController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "syslist"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/syslist/list"

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
	total, err := models.GetSysListCount()
	syslists, _, err := models.GetSysLists(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Syslists"] = syslists
	this.Data["totals"] = total

	this.TplName = "syslist_list.html"
	return
}
func (this *SyslistController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "syslist"
	/*
		groups, err := models.GetNames()
		if err != nil {
			beego.Error(err)
		}
		this.Data["Groups"] = groups
	*/
	id := this.Input().Get("id")
	if len(id) > 0 {
		syslist, err := models.GetSyslistById(id)
		if err != nil {
			beego.Error(err)
		}
		//syslist.Rootpwd, _ = models.AESDecode(syslist.Rootpwd, models.AesKey)
		//syslist.Readpwd, _ = models.AESDecode(syslist.Readpwd, models.AesKey)
		//fmt.Printf("***root :%v,read :%v\n", host.Rootpwd, host.Readpwd)
		this.Data["Syslist"] = syslist
		//	this.Data["HostGroupName"] = host.Group
		this.Data["Path1"] = "系统列表"
		this.Data["Path2"] = "修改系统应用"
		this.Data["Href"] = "/syslist/list"
		this.TplName = "syslist_modify.html"
		//this.TplName = "test.html"
		return
	}
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = "添加系统"
	this.Data["Href"] = "/syslist/list"
	this.TplName = "syslist_add.html"

}

func (this *SyslistController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "syslist"

	id := this.Input().Get("id")
	class := this.Input().Get("class")
	name := this.Input().Get("name")
	owner1 := this.Input().Get("owner1")
	owner2 := this.Input().Get("owner2")
	domain_name := this.Input().Get("domain_name")
	comment := this.Input().Get("comment")
	//beego.Info(idc)
	if len(id) > 0 {
		err := models.ModifySyslist(id, class, name, owner1, owner2, domain_name, comment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddSyslist(class, name, owner1, owner2, domain_name, comment)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/syslist/list"
	this.Redirect("/syslist/list", 302)
}

func (this *SyslistController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "syslist"

	id := this.Input().Get("id")
	err := models.DeleteSyslist(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/syslist/list"
	this.Redirect("/syslist/list", 302)
	return
}

func (this *SyslistController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "syslist"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteSyslist(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *SyslistController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "syslist"

	name := this.Input().Get("keyword")
	/*
		idc := this.Input().Get("idc")
		if idc == "1" {
			this.Data["Path1"] = "主机列表"
			this.Data["Path2"] = ""
			this.Data["Href"] = "/host/list"
			this.TplName = "host_list.html"
			return
		}
	*/
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchSyslistCount(name)
	syslists, err := models.SearchSyslistByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Syslists"] = syslists
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	//this.Data["Idc"] = idc
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/syslist/list"
	this.TplName = "syslist_list.html"
	return
}
