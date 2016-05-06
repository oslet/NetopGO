package controllers

import (
	"github.com/astaxie/beego"
)

type RecordController struct {
	BaseController
}

func (this *RecordController) Get() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
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

	//id := this.Input().Get("id")
	// if len(id) > 0 {
	// 	// schema, err := models.GetSchemaById(id)
	// 	// if err != nil {
	// 	// 	beego.Error(err)
	// 	// }
	// 	this.Data["Path1"] = "DB升级记录"
	// 	this.Data["Path2"] = "添加记录"
	// 	this.Data["Href"] = "/record/db/list"
	// 	this.TplName = "schema_modify.html"
	// 	return
	// }
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
	attachment := this.Input().Get("attachment")

	if len(id) > 0 {
		// err := models.ModifySchema(id, name, dbname, partition, user, passwd, status, comment, addr, port)
		// if err != nil {
		// 	beego.Error(err)
		// }
	} else {
		err := models.AddDBRecord(schema, object, operation, backup, content, attachment, comment)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Data["Path1"] = "DB升级记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/db/list"
	this.Redirect("/schema/list", 302)
}
