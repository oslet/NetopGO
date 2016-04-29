package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type DBController struct {
	BaseController
}

func (this *DBController) Get() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "db"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "DB列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/db/list"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 2
	total, err := models.GetDBCount()
	dbs, _, err := models.GetDBs(int(currPage), pageSize)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)

	this.Data["paginator"] = res
	this.Data["DBs"] = dbs
	this.Data["totals"] = total

	this.TplName = "db_list.html"
	return
}
func (this *DBController) Add() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "db"

	id := this.Input().Get("id")
	if len(id) > 0 {
		db, err := models.GetDBById(id)
		if err != nil {
			beego.Error(err)
		}

		this.Data["DB"] = db
		this.Data["Path1"] = "DB列表"
		this.Data["Path2"] = "修改DB"
		this.Data["Href"] = "/db/list"
		this.TplName = "db_modify.html"
		//this.TplName = "test.html"
		return
	}
	this.Data["Path1"] = "DB列表"
	this.Data["Path2"] = "添加DB"
	this.Data["Href"] = "/db/list"
	this.TplName = "db_add.html"

}

func (this *DBController) Post() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "db"

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	uuid := this.Input().Get("uuid")
	comment := this.Input().Get("comment")
	if len(id) > 0 {
		err := models.ModifyDB(id, name, uuid, comment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddDB(name, uuid, comment)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "DB列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/db/list"
	this.Redirect("/db/list", 302)
}

func (this *DBController) Delete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "db"

	id := this.Input().Get("id")
	err := models.DeleteDB(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "DB列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/db/list"
	this.Redirect("/db/list", 302)
	return
}

func (this *DBController) BitchDelete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "db"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteDB(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *DBController) Search() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "db"

	name := this.Input().Get("keyword")
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 1
	total, err := models.SearchDBCount(name)
	dbs, err := models.SearchDBByName(int(currPage), pageSize, name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)
	this.Data["paginator"] = res
	this.Data["DBs"] = dbs
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "DB列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/db/list"
	this.TplName = "db_list.html"
	return
}

func (this *DBController) Query() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "db"
	var error int
	schema := this.Input().Get("schema")
	flag := this.Input().Get("flag")
	sqltext := this.Input().Get("sql")
	if "result" == flag {
		values, columns, total, msg := models.Query(schema, sqltext)
		if msg != nil {
			error = 1
		}
		this.Data["Values"] = values
		this.Data["Columns"] = columns
		this.Data["Total"] = total
		this.Data["Sqltext"] = sqltext
		this.Data["Error"] = error
		this.Data["Msg"] = msg
		this.Data["Path1"] = "查询窗口"
		this.Data["Path2"] = "查询结果"
		this.Data["Href"] = "/db/query?sql=" + sqltext
		this.TplName = "query_result.html"
		return
	}

	schemas, err := models.GetSchemaNames()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Sqltext"] = sqltext
	this.Data["Schemas"] = schemas
	this.Data["Path1"] = "查询窗口"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/db/query"
	this.TplName = "query.html"
	return
}
