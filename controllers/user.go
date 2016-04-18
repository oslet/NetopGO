package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	var page string
	uname, role := this.IsLogined()
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 1
	total, err := models.GetUserCount()
	users, _, err := models.GetUsers(int(currPage), pageSize)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)
	this.Data["paginator"] = res
	this.Data["Users"] = users
	this.Data["totals"] = total
	this.Data["IsSearch"] = false
	this.TplName = "user_list.html"
	return
}

func (this *UserController) Add() {
	uname, role := this.IsLogined()
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	id := this.Input().Get("id")
	if len(id) > 0 {
		user, err := models.GetUserById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["User"] = user
		this.TplName = "user_modify.html"
		return
	}
	this.TplName = "user_add.html"

}

func (this *UserController) Delete() {
	uname, role := this.IsLogined()
	this.Data["Uname"] = uname
	this.Data["Role"] = role

	id := this.Input().Get("id")
	err := models.DeleteUser(id)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/user/list", 302)
	//this.TplName = "user_list.html"
	return
}

func (this *UserController) Post() {
	uname, role := this.IsLogined()
	this.Data["Uname"] = uname
	this.Data["Role"] = role

	id := this.Input().Get("id")
	name := this.Input().Get("uname")
	passwd := this.Input().Get("passwd")
	email := this.Input().Get("email")
	tel := this.Input().Get("tel")
	auth := this.Input().Get("auth")
	dept := this.Input().Get("dept")
	beego.Info(id)
	if len(id) > 0 {
		err := models.MofifyUser(id, name, passwd, email, tel, auth, dept)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddUser(name, passwd, email, tel, auth, dept)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Redirect("/user/list", 302)
	return
}

func (this *UserController) Search() {
	var page string
	uname, role := this.IsLogined()
	this.Data["Uname"] = uname
	this.Data["Role"] = role

	name := this.Input().Get("keyword")
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 1
	users, total, err := models.SearchUserByName(int(currPage), pageSize, name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)
	this.Data["paginator"] = res
	this.Data["Users"] = users
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.TplName = "user_list.html"
	return
}
