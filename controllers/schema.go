package controllers

import (
	"NetopGO/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type SchemaController struct {
	BaseController
}

func (this *SchemaController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
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
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetSchemaCount()
	schemas, _, err := models.GetSchemas(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	for _, val := range schemas {
		val.Size, _ = models.GetSizeBySchema(val.Name)
	}

	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Schemas"] = schemas
	this.Data["totals"] = total

	this.TplName = "schema_list.html"
	return
}

func (this *SchemaController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

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
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "schema"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	dbname := this.Input().Get("dbname")
	user := this.Input().Get("user")
	passwd := this.Input().Get("passwd")
	comment := this.Input().Get("comment")
	addr := this.Input().Get("addr")
	port := this.Input().Get("port")
	partition := this.Input().Get("partition")
	status := this.Input().Get("status")
	if len(id) > 0 {
		err := models.ModifySchema(id, name, dbname, partition, user, passwd, status, comment, addr, port)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddSchema(name, dbname, partition, user, passwd, status, comment, addr, port)
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
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

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
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

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
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchSchemaCount(name)
	schemas, err := models.SearchSchemaByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
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

func (this *SchemaController) Partition() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "schema"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	schmea := this.Input().Get("schema")
	flag := this.Input().Get("flag")
	total, _ := strconv.ParseInt(this.Input().Get("num"), 10, 64)

	parts, num, err := models.GetPartDetail(flag, schmea)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Partitions"] = parts
	this.Data["Num"] = num
	this.Data["Total"] = total
	this.TplName = "part_detail.html"
	return
}

func (this *SchemaController) View() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["Category"] = "db"
	schema := this.Input().Get("schema")
	time, size, total, err := models.GetTotalSizeView(schema)
	slowTime, count, err := models.GetTotalSlowView(schema)
	qpsTiem, qps, tps, err := models.GetTotalQpsView(schema)
	if err != nil {
		beego.Error(err)
	}
	this.Data["SizeTimes"] = time
	this.Data["CurrSizes"] = size
	this.Data["TotalSizes"] = total
	this.Data["SlowTimes"] = slowTime
	this.Data["SlowCounts"] = count
	this.Data["QpsTimes"] = qpsTiem
	this.Data["Qps"] = qps
	this.Data["Tps"] = tps
	this.Data["Path1"] = "Schema列表"
	this.Data["Path2"] = "图表展示"
	this.Data["Href"] = "/schema/list"
	this.TplName = "schema_view.html"
}
