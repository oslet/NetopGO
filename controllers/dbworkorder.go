package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	//"strings"
)

type DBWOController struct {
	BaseController
}

func (this *DBWOController) Get() {
	var page string
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetDBOrderCount(dept.(string), uname.(string))
	dbwos, _, err := models.GetDBOrders(int(currPage), int(pageSize), dept.(string), uname.(string))
	if err != nil {
		beego.Error(err)
	}

	for _, dbwo := range dbwos {
		dbwo.Isapproved, dbwo.Isedit = models.IsDBApproved(dept.(string), dbwo.Status)
	}

	res := models.Paginator(int(currPage), int(pageSize), total)

	schemas, err := models.GetSchemaNames()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Schemas"] = schemas
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["DBWorkOrders"] = dbwos
	this.Data["totals"] = total

	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "数据库工单"
	this.Data["Href"] = ""
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbworkorder_list.html"
	return
}

func (this *DBWOController) DBOrder() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	schemas, err := models.GetSchemaNames()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Schemas"] = schemas
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "提交DB工单"
	this.Data["Href"] = "/workorder/mydb"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbworkorder.html"
	return
}

func (this *DBWOController) DBOrderPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	schema := this.Input().Get("schema")
	upgradeobj := this.Input().Get("upgradeobj")
	upgradetype := this.Input().Get("upgradetype")
	comment := this.Input().Get("comment")
	step := this.Input().Get("step")

	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var sqlfile string
	if sql != nil {
		sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("attachment", sqlfile))
		if err != nil {
			beego.Error(err)
		}
	}

	err = models.AddDBOrder(schema, upgradeobj, upgradetype, comment, sqlfile, step, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Data["Category"] = "workorder/mydb"
	this.Redirect("/workorder/mydb", 302)
	return
}

func (this *DBWOController) DBInApp() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	schemas, err := models.GetSchemaNames()
	if err != nil {
		beego.Error(err)
	}
	id := this.Input().Get("id")
	sqlfile := this.Input().Get("sqlfile")
	this.Data["Id"] = id
	this.Data["Sqlfile"] = sqlfile
	this.Data["Schemas"] = schemas
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "DB审批"
	this.Data["Href"] = "/workorder/mydb"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbinapprove.html"
}

func (this *DBWOController) DBInAppPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	id := this.Input().Get("id")
	sqlfile := this.Input().Get("sql")
	schema := this.Input().Get("schema")
	upgradeobj := this.Input().Get("upgradeobj")
	upgradetype := this.Input().Get("upgradetype")
	comment := this.Input().Get("comment")
	step := this.Input().Get("step")

	err := models.DBInAppCommit(id, schema, upgradeobj, upgradetype, comment, sqlfile, step, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Data["Category"] = "workorder/mydb"
	this.Redirect("/workorder/mydb", 302)
	return
}

func (this *DBWOController) Detail() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}

	id := this.Input().Get("id")
	dbwo, err := models.GetDBwoById(id)
	if err != nil {
		beego.Error(err)
	}

	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	//this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "工单详情"
	this.Data["Href"] = "/workorder/mydb"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbworkorder_detail.html"
	return
}

func (this *DBWOController) DBApprove() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}

	id := this.Input().Get("id")
	dbwo, err := models.GetDBwoById(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "工单审批"
	this.Data["Href"] = "/workorder/mydb"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbapprove.html"
	return
}

func (this *DBWOController) DBCommit() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")
	opoutcome := this.Input().Get("opoutcome")
	nextStatus := "实施完毕"

	err := models.DBCommit(id, nextStatus, opoutcome, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/mydb", 302)
	return
}

func (this *DBWOController) DBRollback() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}

	id := this.Input().Get("id")
	dbwo, err := models.GetDBwoById(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "异常回滚"
	this.Data["Href"] = "/workorder/mydb"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbrollback.html"
	return
}

func (this *DBWOController) DBRollbackPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")
	opoutcome := this.Input().Get("opoutcome")
	lastStatus := "异常回滚"

	err := models.DBRollback(id, lastStatus, opoutcome, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/mydb", 302)
	return
}
