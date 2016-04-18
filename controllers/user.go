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
	this.TplName = "user_list.html"
}

func (this *UserController) Add() {
	uname, role := this.IsLogined()
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.TplName = "user_add.html"
}
