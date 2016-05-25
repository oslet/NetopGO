package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"strings"
)

type RecordController struct {
	BaseController
}

func (this *RecordController) Get() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	// list := beego.AppConfig.String("db_record_schema_list")
	// arrList := strings.Split(list, ",")
	schemas, err := models.GetSchemaNames()
	if err != nil {
		beego.Error(err)
	}

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetDBRecordCount()
	dbRecs, _, err := models.GetDBRecords(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["Schemas"] = schemas
	this.Data["paginator"] = res
	this.Data["DBRecords"] = dbRecs
	this.Data["totals"] = total

	this.Data["Path1"] = "DB升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/db/list"
	this.Data["Category"] = "record/db"
	this.TplName = "db_record_list.html"
	return
}

func (this *RecordController) Add() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/db"
	list := beego.AppConfig.String("db_record_schema_list")
	arrList := strings.Split(list, ",")

	this.Data["List"] = arrList
	this.Data["Path1"] = "DB升级记录"
	this.Data["Path2"] = "添加记录"
	this.Data["Href"] = "/record/db/list"
	this.TplName = "db_record_add.html"
	return
}

func (this *RecordController) Post() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "record/db"

	id := this.Input().Get("id")
	schema := this.Input().Get("schema")
	object := this.Input().Get("object")
	operation := this.Input().Get("operation")
	backup := this.Input().Get("backup")
	content := this.Input().Get("content")
	comment := this.Input().Get("comment")
	operater := uname.(string)

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if fh != nil {
		attachment = fh.Filename
		//beego.Info(attachment)
		err := this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(id) > 0 {
		dbRec, err := models.DBRecordDetail(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["DBRecord"] = dbRec
		this.Data["Path1"] = "DB升级记录"
		this.Data["Path2"] = "操作内容"
		this.Data["Href"] = "/record/db/list"
		this.TplName = "db_record_detail.html"
		return
	} else {
		err := models.AddDBRecord(schema, object, operation, backup, content, attachment, comment, operater)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Data["Path1"] = "DB升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/db/list"
	this.Redirect("/record/db/list", 302)
	return
}

func (this *RecordController) Delete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/db"

	id := this.Input().Get("id")
	err := models.DeleteDBRecord(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "DB升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/db/list"
	this.Redirect("/record/db/list", 302)
	return
}

func (this *RecordController) BitchDelete() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/db"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteDBRecord(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *RecordController) Search() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/db"

	schema := this.Input().Get("keyword")
	//beego.Info(schema)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchDBRecCount(schema)
	dbRecs, err := models.SearchDBRecBySchema(int(currPage), int(pageSize), schema)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["DBRecords"] = dbRecs
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = schema
	this.Data["Path1"] = "DB升级记录"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/record/db/list"
	this.TplName = "db_record_list.html"
	return
}
