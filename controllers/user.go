package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetUserCount()
	users, _, err := models.GetUsers(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Users"] = users
	this.Data["totals"] = total
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "用户列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/user/list"
	this.TplName = "user_list.html"
	return
}

func (this *UserController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	id := this.Input().Get("id")
	if len(id) > 0 {
		user, err := models.GetUserById(id)
		if err != nil {
			beego.Error(err)
		}
		user.Passwd, _ = models.AESDecode(user.Passwd, models.AesKey)
		//beego.Info(user.Passwd)
		this.Data["User"] = user
		this.Data["Path1"] = "用户列表"
		this.Data["Path2"] = "修改用户"
		this.Data["Href"] = "/user/list"
		this.TplName = "user_modify.html"
		return
	}
	this.Data["Path1"] = "用户列表"
	this.Data["Path2"] = "添加用户"
	this.Data["Href"] = "/user/list"
	this.TplName = "user_add.html"

}

func (this *UserController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	id := this.Input().Get("id")
	err := models.DeleteUser(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "用户列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/user/list"
	this.Redirect("/user/list", 302)
	return
}

func (this *UserController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteUser(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *UserController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	id := this.Input().Get("id")
	name := this.Input().Get("uname")
	passwd := this.Input().Get("passwd")
	email := this.Input().Get("email")
	tel := this.Input().Get("tel")
	auth := this.Input().Get("auth")
	dept := this.Input().Get("dept")
	//beego.Info(id)
	//beego.Info(passwd)
	if len(id) > 0 {
		err := models.ModifyUser(id, name, passwd, email, tel, auth, dept)
		if err != nil {
			//		beego.Info("call")
			beego.Error(err)
		}
	} else {
		err := models.AddUser(name, passwd, email, tel, auth, dept)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "用户列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/user/list"
	this.Redirect("/user/list", 302)
	return
}

func (this *UserController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchUserCount(name)
	users, err := models.SearchUserByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Users"] = users
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "用户列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/user/list"
	this.TplName = "user_list.html"
	return
}

func (this *UserController) Detail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	id := this.Input().Get("id")
	user, err := models.GetUserById(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["User"] = user
	this.Data["Path1"] = "用户列表"
	this.Data["Path2"] = "个人信息"
	this.Data["Href"] = "/user/list"
	this.TplName = "user_detail.html"
	return

}

func (this *UserController) ResetPasswd() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["Category"] = "user"

	var flag int
	id := this.Input().Get("id")
	action := this.Input().Get("action")
	if action == "view" {
		flag = 0
		this.Data["Flag"] = flag
		this.Data["Id"] = id
		this.Data["Path1"] = "用户列表"
		this.Data["Path2"] = "修改密码"
		this.Data["Href"] = "/user/list"
		this.TplName = "reset_password.html"
		return
	} else {
		passwd0 := this.Input().Get("passwd0")
		passwd1 := this.Input().Get("passwd1")
		passwd2 := this.Input().Get("passwd2")
		user, err := models.GetUserById(id)
		if err != nil {
			beego.Error(err)
		}
		enpasswd, _ := models.AESEncode(passwd0, models.AesKey)
		if enpasswd != user.Passwd {
			flag = 1
			this.Data["Flag"] = flag
			this.Data["Path1"] = "用户列表"
			this.Data["Path2"] = "修改密码"
			this.Data["Href"] = "/user/list"
			this.TplName = "reset_password.html"
			return
		}
		if passwd1 != passwd2 {
			flag = 2
			this.Data["Flag"] = flag
			this.Data["Path1"] = "用户列表"
			this.Data["Path2"] = "修改密码"
			this.Data["Href"] = "/user/list"
			this.TplName = "reset_password.html"
			return
		}
		err = models.ResetPasswd(id, passwd1)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/netopgo", 302)
	}
}
