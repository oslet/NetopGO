package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) IsLogined() (uid, uname, auth, dept interface{}) {
	uid = this.GetSession("id")
	uname = this.GetSession("uname")
	auth = this.GetSession("auth")
	dept = this.GetSession("dept")
	if uname == nil {
		this.Redirect("/login", 302)
		return
	}
	return

}
