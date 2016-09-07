package controllers

import (
	"NetopGO/models"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

type AppWOController struct {
	BaseController
}

func (this *AppWOController) Get() {
	var page string
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role

	this.Data["IsSearch"] = false
	auth := role.(int64)

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	pageAuth, _ := strconv.ParseInt(this.Input().Get("pageAuth"), 10, 64)
	pageDept := this.Input().Get("pageDept")
	this.Data["PageAuth"] = pageAuth
	this.Data["PageDept"] = pageDept
	// beego.Info(pageAuth)
	// beego.Info(pageDept)

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetAppOrderCount(pageAuth, pageDept, uname.(string))
	appwos, _, err := models.GetAppOrders(int(currPage), int(pageSize), pageAuth, pageDept, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	for _, appwo := range appwos {
		appwo.Isapproved = models.IsApproved("app", dept.(string), appwo.Status, appwo.Upgradetype, appwo.DbStatus)
		// if "研发" == dept.(string) || "运维" == dept.(string) {
		// 	appwo.Isedit = appwo.Isapproved
		// } else {
		// 	appwo.Isedit = "false"
		// }

	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Schemas"] = schemas
	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Auth"] = auth
	this.Data["Dept"] = dept.(string)
	this.Data["paginator"] = res
	this.Data["AppWorkOrders"] = appwos
	this.Data["totals"] = total
	var isViewItem bool
	if dept.(string) == "研发" || dept.(string) == "运维" || dept.(string) == "测试" || dept.(string) == "产品" {
		isViewItem = true
	} else {
		isViewItem = false
	}

	this.Data["IsViewItem"] = isViewItem
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "我的工单"
	this.Data["Href"] = ""
	this.Data["Category"] = "workorder/my"
	this.TplName = "appworkorder_list.html"
	return
}

func (this *AppWOController) AppOrder() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	var isViewItem bool
	if dept.(string) == "研发" || dept.(string) == "运维" {
		isViewItem = true
	} else {
		isViewItem = false
	}
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["IsViewItem"] = isViewItem
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "提交应用工单"
	this.Data["Href"] = "/workorder/my/list"
	this.Data["Category"] = "workorder/my"
	this.TplName = "appworkorder.html"
	return
}

func (this *AppWOController) AppOrderPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Auth"] = role.(int64)
	apptype := this.Input().Get("apptype")
	appname := this.Input().Get("appname")
	version := this.Input().Get("version")
	sourcecodename := this.Input().Get("sourcecodename")
	buildnum := this.Input().Get("buildnum")
	featurelist := this.Input().Get("featurelist")
	modifycfg := this.Input().Get("modifycfg")
	relayapp := this.Input().Get("relayapp")
	upgradetype := this.Input().Get("upgradetype")
	sponsor := uname.(string)
	currDept := dept.(string)
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	_, report, err := this.GetFile("testreport")
	if err != nil {
		beego.Error(err)
	}
	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var testreport string
	var sqlfile string
	if fh != nil {
		attachment = fh.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}
	if report != nil {
		timestamp := time.Now().UnixNano()
		testreport = strconv.FormatInt(timestamp, 10) + report.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("testreport", path.Join("attachment", testreport))
		if err != nil {
			beego.Error(err)
		}
	}
	if sql != nil {
		sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("attachment", sqlfile))
		if err != nil {
			beego.Error(err)
		}
	}

	err = models.AddAppOrder(apptype, appname, version, sourcecodename, buildnum, featurelist, modifycfg, relayapp, upgradetype, sponsor, attachment, testreport, sqlfile, currDept)
	if err != nil {
		beego.Error(err)
	}
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Redirect("/workorder/my/list", 302)
	return
}

func (this *AppWOController) Approve() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	auth := role.(int64)
	this.Data["Auth"] = auth

	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	id := this.Input().Get("id")
	appwo, err := models.GetAppwoById(id)
	if err != nil {
		beego.Error(err)
	}
	test, product, op, final, testReadonly, productReadonly, opReadonly, finalReadonly := models.IsViewDiv(dept.(string), appwo.Status, appwo.Upgradetype)
	this.Data["Test"] = test
	this.Data["Product"] = product
	this.Data["Op"] = op
	this.Data["Final"] = final
	this.Data["TestReadonly"] = testReadonly
	this.Data["ProductReadonly"] = productReadonly
	this.Data["OpReadonly"] = opReadonly
	this.Data["FinalReadonly"] = finalReadonly
	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Appwo"] = appwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "工单审批"
	this.Data["Href"] = "/workorder/my/list"
	this.Data["Category"] = "workorder/my"
	this.TplName = "approve.html"
	return
}

func (this *AppWOController) Rollback() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	id := this.Input().Get("id")
	appwo, err := models.GetAppwoById(id)
	if err != nil {
		beego.Error(err)
	}
	test, product, op, final, testReadonly, productReadonly, opReadonly, finalReadonly := models.IsViewDiv(dept.(string), appwo.Status, appwo.Upgradetype)
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["Test"] = test
	this.Data["Product"] = product
	this.Data["Op"] = op
	this.Data["Final"] = final
	this.Data["TestReadonly"] = testReadonly
	this.Data["ProductReadonly"] = productReadonly
	this.Data["OpReadonly"] = opReadonly
	this.Data["FinalReadonly"] = finalReadonly
	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Appwo"] = appwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "异常回滚"
	this.Data["Href"] = "/workorder/my/list"
	this.Data["Category"] = "workorder/my"
	this.TplName = "approllback.html"
	return
}

func (this *AppWOController) ApproveCommit() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["Auth"] = role.(int64)
	id := this.Input().Get("id")
	auth := role.(int64)
	this.Data["Auth"] = auth

	beego.Info(id)
	appwo, err := models.GetAppwoById(id)
	if err != nil {
		beego.Error(err)
	}
	nextStatus, who, outcome := models.NextStatus("app", dept.(string), appwo.Status, appwo.Upgradetype)
	outcomevalue := this.Input().Get(outcome)
	err = models.ApproveCommit(id, nextStatus, outcome, outcomevalue, who, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/my/list", 302)
	return
}

func (this *AppWOController) ApproveRollback() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["Auth"] = role.(int64)
	id := this.Input().Get("id")
	auth := role.(int64)
	this.Data["Auth"] = auth
	opoutcome := this.Input().Get("opoutcome")
	finaloutcome := this.Input().Get("finaloutcome")
	err := models.ApproveRollback(id, "异常已回滚", opoutcome, finaloutcome, uname.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/my/list", 302)
	return
}

// func (this *AppWOController) ApproveRollback() {
// 	uid, uname, role, dept := this.IsLogined()
// 	this.Data["Id"] = uid
// 	this.Data["Uname"] = uname
// 	this.Data["Role"] = role
// 	this.Data["Dept"] = dept
// 	this.Data["Auth"] = role.(int64)
// 	id := this.Input().Get("id")

// 	beego.Info(id)
// 	appwo, err := models.GetAppwoById(id)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	lastStatus, who, outcome := models.LastStatus("app", dept.(string), appwo.Status, appwo.Upgradetype)
// 	outcomevalue := this.Input().Get(outcome)
// 	err = models.ApproveRollback(id, lastStatus, outcome, outcomevalue, who, uname.(string), appwo.Status, appwo.Upgradetype, dept.(string))
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	this.Redirect("/workorder/my/list", 302)
// 	return
// }

func (this *AppWOController) Detail() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	id := this.Input().Get("id")
	appwo, err := models.GetAppwoById(id)
	if err != nil {
		beego.Error(err)
	}
	auth := role.(int64)
	this.Data["Auth"] = auth

	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Appwo"] = appwo
	//this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "工单详情"
	this.Data["Href"] = "/workorder/my/list"
	this.Data["Category"] = "workorder/my"
	this.TplName = "appworkorder_detail.html"
	return
}

func (this *AppWOController) ApproveModify() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")
	appwo, err := models.GetAppwoById(id)
	if err != nil {
		beego.Error(err)
	}
	// err = models.ApproveCommit(id, nextStatus, outcome, outcomevalue, who, uname.(string))
	// if err != nil {
	// 	beego.Error(err)
	// }
	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Appwo"] = appwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "重新发布"
	this.Data["Href"] = "/workorder/my/list"
	this.Data["Category"] = "workorder/my"
	this.TplName = "appworkorder_modify.html"
	return
}

func (this *AppWOController) ApproveModifyPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	auth := role.(int64)
	this.Data["Auth"] = auth

	id := this.Input().Get("id")
	apptype := this.Input().Get("apptype")
	appname := this.Input().Get("appname")
	upgradetype := this.Input().Get("upgradetype")
	version := this.Input().Get("version")
	sourcecodename := this.Input().Get("sourcecodename")
	buildnum := this.Input().Get("buildnum")
	featurelist := this.Input().Get("featurelist")
	modifycfg := this.Input().Get("modifycfg")
	relayapp := this.Input().Get("relayapp")
	old_attachment := this.Input().Get("old_attachment")
	old_testreport := this.Input().Get("old_testreport")
	old_sqlfile := this.Input().Get("old_sqlfile")

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	_, report, err := this.GetFile("testreport")
	if err != nil {
		beego.Error(err)
	}
	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var final_attachment string
	var final_testreport string
	var final_sqlfile string
	if fh != nil {
		final_attachment = fh.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("attachment", final_attachment))
		if err != nil {
			beego.Error(err)
		}
	} else {
		final_attachment = old_attachment
	}
	if report != nil {
		final_testreport = report.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("testreport", path.Join("attachment", final_testreport))
		if err != nil {
			beego.Error(err)
		}
	} else {
		final_testreport = old_testreport
	}
	if sql != nil {
		final_sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("attachment", final_sqlfile))
		if err != nil {
			beego.Error(err)
		}
	} else {
		final_sqlfile = old_sqlfile
	}

	err = models.ApproveModify(id, apptype, appname, upgradetype, version, sourcecodename, buildnum, featurelist, modifycfg, relayapp, final_attachment, final_testreport, final_sqlfile, dept.(string))
	if err == nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/my/list", 302)
	return
}

func (this *AppWOController) Search() {
	var page string
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Auth"] = role.(int64)
	this.Data["Category"] = "workorder/my"
	auth := role.(int64)
	this.Data["Auth"] = auth

	pageAuth, _ := strconv.ParseInt(this.Input().Get("pageAuth"), 10, 64)
	pageDept := this.Input().Get("pageDept")
	this.Data["PageAuth"] = pageAuth
	this.Data["PageDept"] = pageDept
	// beego.Info(pageAuth)
	// beego.Info(pageDept)
	apptype := this.Input().Get("apptype")
	appname := this.Input().Get("appname")
	//auth := this.Input().Get("auth")
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchAppwoCount(apptype, appname, pageDept, uname.(string), pageAuth)
	appwos, err := models.SearchAppwo(int(currPage), int(pageSize), apptype, appname, pageDept, uname.(string), pageAuth)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	schemas, err := models.GetSchemaNamesArray()
	if err != nil {
		beego.Error(err)
	}
	var isViewItem bool
	if dept.(string) == "研发" || dept.(string) == "运维" || dept.(string) == "测试" || dept.(string) == "产品" {
		isViewItem = true
	} else {
		isViewItem = false
	}
	this.Data["IsViewItem"] = isViewItem
	this.Data["Schemas"] = schemas
	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Dept"] = dept.(string)
	this.Data["paginator"] = res
	this.Data["AppWorkOrders"] = appwos
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["AppType"] = apptype
	this.Data["AppName"] = appname
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "我的工单"
	this.Data["Href"] = ""
	this.Data["Category"] = "workorder/my"
	this.TplName = "appworkorder_list.html"
	return
}

func (this *AppWOController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "workorder/my"
	method := this.Input().Get("method")
	values, columns, _ := models.QueryAppwosExport(method)
	auth := role.(int64)
	this.Data["Auth"] = auth

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
		filename = "month_app" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
	} else if method == "all" {
		filename = "all_app" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
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

func (this *AppWOController) CloseOrder() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "workorder/my"
	auth := role.(int64)
	this.Data["Auth"] = auth

	id := this.Input().Get("id")
	err := models.CloseOrder(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "我的工单"
	this.Data["Href"] = ""
	this.Redirect("/workorder/my/list", 302)
	return
}
