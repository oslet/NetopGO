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

}
