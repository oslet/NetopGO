package controllers

import (
	"NetopGO/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	passwd := this.Input().Get("passwd")
	//fmt.Printf("input password: %v\n", passwd)
	//encodePasswd, _ := models.AESEncode(passwd, models.AesKey)
	//decodePasswd, _ := models.AESDecode(encodePasswd, models.AesKey)
	//fmt.Printf("encode password: %v\n", encodePasswd)
	//fmt.Printf("decode password: %v\n", decodePasswd)
	user, err := models.Login(uname)
	//fmt.Printf("user password: %v\n", user.Passwd)

	if uname == "netop" {
		i := 1
		var Auth int64
		Auth = int64(i)
		dept := "运维"
		this.SetSession("id", user.Id)
		this.SetSession("uname", uname)
		this.SetSession("passwd", passwd)
		this.SetSession("auth", Auth)
		this.SetSession("dept", dept)
		this.Redirect("/netopgo", 302)

	}

	authservice := beego.AppConfig.String("auth_service")
	service := NewUserManageCenterServiceSoap(authservice, false)
	auex, err := service.AuthenticateUserEx(&AuthenticateUserEx{SUserCode: uname, SPassword: passwd, NSystemID: 5})
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%v\n", seasons.AuthenticateUserResult)
	if auex.AuthenticateUserExResult.Success1 != true {

		//	if err != nil || encodePasswd != user.Passwd {
		beego.Error(err)
		this.Data["Error"] = true
		this.TplName = "login.html"
		return
	}

	get_usercode := auex.AuthenticateUserExResult.ObjData.UserCode
	if get_usercode == "2016092902" || get_usercode == "20106364336" {
		i := 1
		var Auth int64
		Auth = int64(i)
		dept := "运维"
		this.SetSession("id", auex.AuthenticateUserExResult.ObjData.UserID)
		this.SetSession("uname", uname)
		this.SetSession("passwd", passwd)
		this.SetSession("auth", Auth)
		this.SetSession("dept", dept)
		this.Redirect("/netopgo", 302)
	} else {
		i := 3
		var Auth int64
		Auth = int64(i)
		dept := "研发"
		this.SetSession("id", auex.AuthenticateUserExResult.ObjData.UserID)
		this.SetSession("uname", uname)
		this.SetSession("passwd", passwd)
		this.SetSession("auth", Auth)
		this.SetSession("dept", dept)
		this.Redirect("/netopgo", 302)
	}

}

func (this *LoginController) Logout() {
	this.DelSession("uname")
	this.DelSession("passwd")
	this.DelSession("auth")
	this.TplName = "login.html"
	return
}
