package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type GroupController struct {
	BaseController
}

func (this *GroupController) Get() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "业务组列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/group/list"
	this.Data["Category"] = "group"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetGroupCount()
	groups, _, err := models.GetGroups(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Groups"] = groups
	this.Data["totals"] = total

	this.TplName = "group_list.html"
	return
}
func (this *GroupController) Add() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "group"

	id := this.Input().Get("id")
	if len(id) > 0 {
		group, err := models.GetGroupById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Group"] = group
		this.Data["Path1"] = "业务组列表"
		this.Data["Path2"] = "修改业务组"
		this.Data["Href"] = "/group/list"
		this.TplName = "group_modify.html"
		return
	}
	this.Data["Path1"] = "业务组列表"
	this.Data["Path2"] = "添加业务组"
	this.Data["Href"] = "/group/list"
	this.TplName = "group_add.html"

}

func (this *GroupController) Post() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "group"

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	conment := this.Input().Get("conment")
	if len(id) > 0 {
		err := models.ModifyGroup(id, name, conment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddGroup(name, conment)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "业务组列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/group/list"
	this.Redirect("/group/list", 302)
}

func (this *GroupController) Delete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "group"

	id := this.Input().Get("id")
	err := models.DeleteGroup(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "业务组列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/group/list"
	this.Redirect("/group/list", 302)
	return
}

func (this *GroupController) BitchDelete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "group"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteGroup(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *GroupController) Search() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "group"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchGroupCount(name)
	groups, err := models.SearchGroupByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Groups"] = groups
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "业务组列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/group/list"
	this.TplName = "group_list.html"
	return
}
