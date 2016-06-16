package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type AuditController struct {
	BaseController
}

func (this *AuditController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	auth := role.(int64)
	this.Data["Auth"] = auth

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	schemas, err := models.GetSchemaNames()
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetAuditCount()
	audits, _, err := models.GetAudits(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	this.Data["Auth"] = auth
	this.Data["Schemas"] = schemas
	this.Data["paginator"] = res
	this.Data["Audits"] = audits
	this.Data["totals"] = total

	this.Data["Path1"] = "日志审计"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/audit/list"
	this.Data["Category"] = "audit"
	this.TplName = "audit_list.html"
	return
}

func (this *AuditController) Detail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "audit"
	auth := role.(int64)
	this.Data["Auth"] = auth
	id := this.Input().Get("id")
	audit, err := models.AuditDetail(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Audit"] = audit
	this.Data["Path1"] = "日志审计"
	this.Data["Path2"] = "sqltext"
	this.Data["Href"] = "/audit/list"
	this.Data["Category"] = "audit"
	this.TplName = "audit_detail.html"
	return
}

func (this *AuditController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "audit"
	auth := role.(int64)
	this.Data["Auth"] = auth
	id := this.Input().Get("id")
	err := models.DeleteAudit(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "日志审计"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/audit/list"
	this.Redirect("/audit/list", 302)
	return
}

func (this *AuditController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "audit"
	auth := role.(int64)
	this.Data["Auth"] = auth
	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteAudit(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *AuditController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "audit"
	auth := role.(int64)
	this.Data["Auth"] = auth
	schema := this.Input().Get("keyword")
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchAuditCount(schema)
	audits, err := models.SearchAuditBySchema(int(currPage), int(pageSize), schema)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Audits"] = audits
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = schema
	this.Data["Path1"] = "日志审计"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/audit/list"
	this.TplName = "audit_list.html"
	return
}
