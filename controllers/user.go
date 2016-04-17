package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	uname := this.GetSession("uname")
	auth := this.GetSession("auth")
	beego.Info(uname)
	beego.Info(auth)
	if uname == nil {
		this.Redirect("/login", 302)
		return
	}
	this.Data["Uname"] = uname
	switch auth {
	case 1:
		this.Data["Auth"] = "超级管理员"
	case 2:
		this.Data["Auth"] = "数据库管理员"
	case 3:
		this.Data["Auth"] = "来宾用户"
	}
	if uname == "admin" {
		this.Data["Admin"] = true
	} else {
		this.Data["Admin"] = false
	}

	this.TplName = "user_list.html"
}
