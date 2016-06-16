package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	//"strings"
	"github.com/tealeg/xlsx"
	"os"
	"time"
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
	var isViewItem bool
	if dept.(string) == "研发" || dept.(string) == "运维" || dept.(string) == "测试" {
		isViewItem = true
	} else {
		isViewItem = false
	}
	this.Data["IsViewItem"] = isViewItem
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
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	auth := role.(int64)
	this.Data["Auth"] = auth
	schemas, err := models.GetSchemaNames()
	if err != nil {
		beego.Error(err)
	}
	var flag bool
	if dept.(string) == "研发" || dept.(string) == "运维" || dept.(string) == "测试" {
		flag = true
	} else {
		flag = false
	}
	var isViewItem bool
	if dept.(string) == "研发" || dept.(string) == "运维" || dept.(string) == "测试" {
		isViewItem = true
	} else {
		isViewItem = false
	}
	this.Data["IsViewItem"] = isViewItem
	this.Data["Flag"] = flag
	this.Data["Schemas"] = schemas
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "提交DB工单"
	this.Data["Href"] = "/workorder/mydb/list"
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
	auth := role.(int64)
	this.Data["Auth"] = auth
	schema := this.Input().Get("schema")
	upgradeobj := this.Input().Get("upgradeobj")
	upgradetype := this.Input().Get("upgradetype")
	comment := this.Input().Get("comment")
	step := this.Input().Get("step")
	var sqlfile string
	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	if sql != nil {
		sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("attachment", sqlfile))
		if err != nil {
			beego.Error(err)
		}
	}
	err = models.AddDBOrder(schema, upgradeobj, upgradetype, comment, step, sqlfile, uname.(string), dept.(string))
	if err != nil {
		beego.Error(err)
	}

	this.Data["Category"] = "workorder/mydb"
	this.Redirect("/workorder/mydb/list", 302)
	return
}

func (this *DBWOController) DBInApp() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	auth := role.(int64)
	this.Data["Auth"] = auth
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
	this.Data["Href"] = "/workorder/mydb/list"
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
	auth := role.(int64)
	this.Data["Auth"] = auth
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
	this.Redirect("/workorder/mydb/list", 302)
	return
}

func (this *DBWOController) Detail() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	auth := role.(int64)
	this.Data["Auth"] = auth
	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}

	id := this.Input().Get("id")
	dbwo, err := models.GetDBwoById(id)
	if err != nil {
		beego.Error(err)
	}
	approver := models.GetUserByName(dbwo.Sponsor)
	var flag bool
	if approver.Dept == "研发" || approver.Dept == "运维" || approver.Dept == "测试" {
		flag = false
	} else {
		flag = true
	}
	this.Data["Flag"] = flag
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	//this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "工单详情"
	this.Data["Href"] = "/workorder/mydb/list"
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
	approver := models.GetUserByName(dbwo.Sponsor)
	dev, op, devReadonly, opReadonly := models.IsViewDBApprove(approver.Dept, dbwo.Status)
	this.Data["Dev"] = dev
	this.Data["Op"] = op
	this.Data["DevReadonly"] = devReadonly
	this.Data["OpReadonly"] = opReadonly
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "工单审批"
	this.Data["Href"] = "/workorder/mydb/list"
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
	this.Redirect("/workorder/mydb/list", 302)
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
	approver := models.GetUserByName(dbwo.Sponsor)
	dev, op, devReadonly, opReadonly := models.IsViewDBApprove(approver.Dept, dbwo.Status)
	this.Data["Dev"] = dev
	this.Data["Op"] = op
	this.Data["DevReadonly"] = devReadonly
	this.Data["OpReadonly"] = opReadonly
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "异常回滚"
	this.Data["Href"] = "/workorder/mydb/list"
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
	this.Redirect("/workorder/mydb/list", 302)
	return
}

func (this *DBWOController) DevApprove() {
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
	// approver := models.GetUserByName(dbwo.Sponsor)
	// dev, op, devReadonly, opReadonly := models.IsViewDBApprove(approver.Dept, dbwo.Status)
	// this.Data["Dev"] = dev
	// this.Data["Op"] = op
	// this.Data["DevReadonly"] = devReadonly
	// this.Data["OpReadonly"] = opReadonly
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "数据库工单"
	this.Data["Path2"] = "研发审批"
	this.Data["Href"] = "/workorder/mydb/list"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "devapprove.html"
	return
}

func (this *DBWOController) DevCommit() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")
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

	devOutcome := this.Input().Get("devoutcome")
	nextStatus := "正在实施"

	err = models.DevCommit(id, nextStatus, step, devOutcome, sqlfile, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/mydb/list", 302)
	return
}

func (this *DBWOController) DBApproveModify() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept

	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}

	id := this.Input().Get("id")
	dbwo, err := models.GetDBwoById(id)
	if err != nil {
		beego.Error(err)
	}
	approver := models.GetUserByName(dbwo.Sponsor)
	var flag bool
	if approver.Dept == "研发" || approver.Dept == "运维" || approver.Dept == "测试" {
		flag = false
	} else {
		flag = true
	}
	this.Data["Flag"] = flag
	this.Data["Schemas"] = schemas
	this.Data["Dbwo"] = dbwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "重新发布"
	this.Data["Href"] = "/workorder/mydb/list"
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbworkorder_modify.html"
	return
}

func (this *DBWOController) DBApproveModifyPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")
	schema := this.Input().Get("schema")
	upgradeobj := this.Input().Get("upgradeobj")
	upgradetype := this.Input().Get("upgradetype")
	comment := this.Input().Get("comment")
	step := this.Input().Get("step")
	devoutcome := this.Input().Get("devoutcome")

	old_sqlfile := this.Input().Get("old_sqlfile")

	_, sql, err := this.GetFile("new_sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var new_sqlfile string

	if sql != nil {
		new_sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("new_sqlfile", path.Join("attachment", new_sqlfile))
		if err != nil {
			beego.Error(err)
		}
	} else {
		new_sqlfile = old_sqlfile
	}

	err = models.DBApproveModify(id, schema, upgradeobj, upgradetype, comment, new_sqlfile, step, devoutcome, dept.(string))
	if err == nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/mydb/list", 302)
	return
}

func (this *DBWOController) Search() {
	var page string
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "workorder/mydb"

	schema := this.Input().Get("keyword")

	//beego.Info(appname)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchDbwoCount(schema)
	dbwos, err := models.SearchDbwo(int(currPage), int(pageSize), schema)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["Dept"] = dept.(string)
	this.Data["paginator"] = res
	this.Data["DBWorkOrders"] = dbwos
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = schema
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "数据库工单"
	this.Data["Href"] = ""
	this.Data["Category"] = "workorder/mydb"
	this.TplName = "dbworkorder_list.html"
	return
}

func (this *DBWOController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "workorder/mydb"
	method := this.Input().Get("method")
	values, columns, _ := models.QueryDBwosExport(method)

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	for _, val := range columns {
		cell := row.AddCell()
		cell.Value = val
	}
	//sheet.SetColWidth(1, len(columns), 100)
	for _, val := range *values {
		row = sheet.AddRow()
		for _, value := range val {
			cell := row.AddCell()
			cell.Value = value
		}
	}
	now := time.Now().String()
	var filename string
	if method == "month" {
		filename = "month_db" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
	} else if method == "all" {
		filename = "all_db" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
	}

	filepath := path.Join("export", filename)
	err := file.Save(filepath)
	if err != nil {
		beego.Error(err)
	}
	defer func() {
		os.Remove(filepath)
	}()
	this.Ctx.Output.Download(filepath, filename)
	return
}
