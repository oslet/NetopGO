package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type SchemaController struct {
	BaseController
}

func (this *SchemaController) Get() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "Schema列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/schema/list"
	this.Data["Category"] = "schema"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 2
	total, err := models.GetSchemaCount()
	schemas, _, err := models.GetSchemas(int(currPage), pageSize)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)

	this.Data["paginator"] = res
	this.Data["Schemas"] = schemas
	this.Data["totals"] = total

	this.TplName = "schema_list.html"
	return
}

func (this *SchemaController) Add() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"

	id := this.Input().Get("id")
	if len(id) > 0 {
		schema, err := models.GetSchemaById(id)
		if err != nil {
			beego.Error(err)
		}
		schema.Passwd, _ = models.AESDecode(schema.Passwd, models.AesKey)
		this.Data["Schema"] = schema
		this.Data["Path1"] = "Schema列表"
		this.Data["Path2"] = "修改Schema"
		this.Data["Href"] = "/schema/list"
		this.TplName = "schema_modify.html"
		return
	}
	this.Data["Path1"] = "Schema列表"
	this.Data["Path2"] = "添加Schema"
	this.Data["Href"] = "/schema/list"
	this.TplName = "schema_add.html"

}

func (this *SchemaController) Post() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "schema"

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	dbname := this.Input().Get("dbname")
	user := this.Input().Get("user")
	passwd := this.Input().Get("passwd")
	comment := this.Input().Get("comment")
	if len(id) > 0 {
		err := models.ModifySchema(id, name, dbname, user, passwd, comment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddSchema(name, dbname, user, passwd, comment)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "Schema列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/schema/list"
	this.Redirect("/schema/list", 302)
}

func (this *SchemaController) Delete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"

	id := this.Input().Get("id")
	err := models.DeleteSchema(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "Schema列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/schema/list"
	this.Redirect("/schema/list", 302)
	return
}

func (this *SchemaController) BitchDelete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteSchema(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *SchemaController) Search() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"

	name := this.Input().Get("keyword")
	beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 1
	total, err := models.SearchSchemaCount(name)
	schemas, err := models.SearchSchemaByName(int(currPage), pageSize, name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)
	this.Data["paginator"] = res
	this.Data["Schemas"] = schemas
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "Schema列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/schema/list"
	this.TplName = "schema_list.html"
	return
}
