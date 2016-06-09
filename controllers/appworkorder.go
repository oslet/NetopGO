package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"strings"
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
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetAppOrderCount()
	appwos, _, err := models.GetAppOrders(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	for _, appwo := range appwos {
		appwo.Isapproved = models.IsApproved("app", dept.(string), appwo.Status, appwo.Upgradetype)
		// if "研发" == dept.(string) || "运维" == dept.(string) {
		// 	appwo.Isedit = appwo.Isapproved
		// } else {
		// 	appwo.Isedit = "false"
		// }

	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")
	this.Data["AppNameList"] = appNameList
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["AppWorkOrders"] = appwos
	this.Data["totals"] = total

	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "我的工单"
	this.Data["Href"] = ""
	this.Data["Category"] = "record/app"
	this.TplName = "appworkorder_list.html"
	return
}

func (this *AppWOController) AppOrder() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")

	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Path1"] = "系统发布"
	this.Data["Path2"] = "提交应用工单"
	this.Data["Href"] = "/workorder/my"
	this.Data["Category"] = "workorder/app"
	this.TplName = "appworkorder.html"
	return
}

func (this *AppWOController) AppOrderPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	apptype := this.Input().Get("apptype")
	appname := this.Input().Get("appname")
	version := this.Input().Get("version")
	jenkinsname := this.Input().Get("jenkinsname")
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
	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	var sqlfile string
	if fh != nil {
		attachment = fh.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("fileupload", attachment))
		if err != nil {
			beego.Error(err)
		}
	}
	if sql != nil {
		sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("fileupload", sqlfile))
		if err != nil {
			beego.Error(err)
		}
	}

	err = models.AddAppOrder(apptype, appname, version, jenkinsname, buildnum, featurelist, modifycfg, relayapp, upgradetype, sponsor, attachment, sqlfile, currDept)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Category"] = "workorder/app"
	this.TplName = "index.html"
	return
}

func (this *AppWOController) Approve() {
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
	this.Data["Href"] = "/workorder/my"
	this.Data["Category"] = "workorder/app"
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
	this.Data["Href"] = "/workorder/my"
	this.Data["Category"] = "workorder/app"
	this.TplName = "approllback.html"
	return
}

func (this *AppWOController) ApproveCommit() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")

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
	this.Redirect("/workorder/my", 302)
	return
}

func (this *AppWOController) ApproveRollback() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")

	beego.Info(id)
	appwo, err := models.GetAppwoById(id)
	if err != nil {
		beego.Error(err)
	}
	lastStatus, who, outcome := models.LastStatus("app", dept.(string), appwo.Status, appwo.Upgradetype)
	outcomevalue := this.Input().Get(outcome)
	err = models.ApproveRollback(id, lastStatus, outcome, outcomevalue, who, uname.(string), appwo.Status, appwo.Upgradetype, dept.(string))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/my", 302)
	return
}

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

	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Appwo"] = appwo
	//this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "工单详情"
	this.Data["Href"] = "/workorder/my"
	this.Data["Category"] = "workorder/app"
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

	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
	this.Data["Appwo"] = appwo
	this.Data["Auth"] = dept.(string)
	this.Data["Path1"] = "我的工单"
	this.Data["Path2"] = "重新发布"
	this.Data["Href"] = "/workorder/my"
	this.Data["Category"] = "workorder/app"
	this.TplName = "appworkorder_modify.html"
	return
}

func (this *AppWOController) ApproveModifyPost() {
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	id := this.Input().Get("id")
	apptype := this.Input().Get("apptype")
	appname := this.Input().Get("appname")
	upgradetype := this.Input().Get("upgradetype")
	version := this.Input().Get("version")
	jenkinsname := this.Input().Get("jenkinsname")
	buildnum := this.Input().Get("buildnum")
	featurelist := this.Input().Get("featurelist")
	modifycfg := this.Input().Get("modifycfg")
	relayapp := this.Input().Get("relayapp")
	old_attachment := this.Input().Get("old_attachment")
	old_sqlfile := this.Input().Get("old_sqlfile")

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	_, sql, err := this.GetFile("sqlfile")
	if err != nil {
		beego.Error(err)
	}
	var final_attachment string
	var final_sqlfile string
	if fh != nil {
		final_attachment = fh.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("fileupload", final_attachment))
		if err != nil {
			beego.Error(err)
		}
	} else {
		final_attachment = old_attachment
	}
	if sql != nil {
		final_sqlfile = sql.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("sqlfile", path.Join("fileupload", final_sqlfile))
		if err != nil {
			beego.Error(err)
		}
	} else {
		final_sqlfile = old_sqlfile
	}

	err = models.ApproveModify(id, apptype, appname, upgradetype, version, jenkinsname, buildnum, featurelist, modifycfg, relayapp, final_attachment, final_sqlfile, dept.(string))
	if err == nil {
		beego.Error(err)
	}
	this.Redirect("/workorder/my", 302)
	return
}
