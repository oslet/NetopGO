package routers

import (
	"NetopGO/controllers"
	"NetopGO/models"
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
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
	beego.Router("/host/webconsole", &controllers.HostController{}, "get:WebConsole")
	beego.Handler("/console/sshws", websocket.Handler(models.SSHWebSocketHandler))

	beego.Router("/schema/list", &controllers.SchemaController{})
	beego.Router("/schema/add", &controllers.SchemaController{}, "get:Add")
	beego.Router("/schema/add", &controllers.SchemaController{}, "post:Post")
	beego.Router("/schema/modify", &controllers.SchemaController{}, "post:Post")
	beego.Router("/schema/del", &controllers.SchemaController{}, "get:Delete")
	beego.Router("/schema/bitchDel", &controllers.SchemaController{}, "post:BitchDelete")
	beego.Router("/schema/partition", &controllers.SchemaController{}, "get:Partition")

	beego.Router("/db/list", &controllers.DBController{})
	beego.Router("/db/add", &controllers.DBController{}, "get:Add")
	beego.Router("/db/add", &controllers.DBController{}, "post:Post")
	beego.Router("/db/modify", &controllers.DBController{}, "post:Post")
	beego.Router("/db/del", &controllers.DBController{}, "get:Delete")
	beego.Router("/db/bitchDel", &controllers.DBController{}, "post:BitchDelete")
	beego.Router("/db/search", &controllers.DBController{}, "get:Search")
	beego.Router("/db/query", &controllers.DBController{}, "get:Query")
	beego.Router("/db/detail", &controllers.DBController{}, "get:Detail")
}
