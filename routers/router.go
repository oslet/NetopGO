package routers

import (
	"NetopGO/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/user/list", &controllers.UserController{})
	//beego.AutoRouter(&controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:Add")
	beego.Router("/user/add", &controllers.UserController{}, "post:Post")
	beego.Router("/user/modify", &controllers.UserController{}, "post:Post")
	beego.Router("/user/del", &controllers.UserController{}, "get:Delete")
	beego.Router("/user/search", &controllers.UserController{}, "get:Search")
	beego.Router("/user/detail", &controllers.UserController{}, "get:Detail")
	beego.Router("/user/bitchDel", &controllers.UserController{}, "post:BitchDelete")
	beego.Router("/user/reset_password", &controllers.UserController{}, "get:ResetPasswd")
	beego.Router("/user/reset_password", &controllers.UserController{}, "post:ResetPasswd")

	beego.Router("/group/list", &controllers.GroupController{})
	beego.Router("/group/add", &controllers.GroupController{}, "get:Add")
	beego.Router("/group/add", &controllers.GroupController{}, "post:Post")
	beego.Router("/group/modify", &controllers.GroupController{}, "post:Post")
	beego.Router("/group/del", &controllers.GroupController{}, "get:Delete")
	beego.Router("/group/bitchDel", &controllers.GroupController{}, "post:BitchDelete")
	beego.Router("/group/search", &controllers.GroupController{}, "get:Search")

	beego.Router("/host/list", &controllers.HostController{})
	beego.Router("/host/add", &controllers.HostController{}, "get:Add")
	beego.Router("/host/add", &controllers.HostController{}, "post:Post")
	beego.Router("/host/modify", &controllers.HostController{}, "post:Post")
	beego.Router("/host/del", &controllers.HostController{}, "get:Delete")
	beego.Router("/host/bitchDel", &controllers.HostController{}, "post:BitchDelete")
	beego.Router("/host/search", &controllers.HostController{}, "get:Search")
}
