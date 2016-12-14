package controllers

import (
	"NetopGO/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type LineController struct {
	BaseController
}

func (this *LineController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "线路列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/line/list"
	this.Data["Category"] = "asset/line"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetLineCount()
	lines, _, err := models.GetLines(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Lines"] = lines
	this.Data["totals"] = total

	this.TplName = "line_list.html"
	return
}
func (this *LineController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "Line"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	if len(id) > 0 {
		line, err := models.GetLineById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Line"] = line
		this.Data["Path1"] = "线路列表"
		this.Data["Path2"] = "修改线路"
		this.Data["Href"] = "/asset/line/list"
		this.TplName = "line_modify.html"
		return
	}
	this.Data["Path1"] = "线路列表"
	this.Data["Path2"] = "添加线路"
	this.Data["Href"] = "/asset/line/list"
	this.TplName = "line_add.html"

}

func (this *LineController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "line"

	Auth := role.(int64)
	this.Data["Auth"] = Auth
	id := this.Input().Get("id")
	name := this.Input().Get("name")
	use := this.Input().Get("use")
	enable := this.Input().Get("enable")
	comment := this.Input().Get("comment")
	if len(id) > 0 {
		err, msg := models.ModifyLine(id, name, use, enable, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddLine(name, use, enable, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "线路列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/line/list"
	//this.Redirect("/asset/line/list", 302)
	this.TplName = "line_add.html"
}

func (this *LineController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "line"

	id := this.Input().Get("id")
	err := models.DeleteLine(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "线路列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/line/list"
	this.Redirect("/asset/line/list", 302)
	return
}

func (this *LineController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "line"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteLine(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *LineController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "asset/line"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchLineCount(name)
	lines, err := models.SearchLineByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Lines"] = lines
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "线路列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/asset/line/list"
	this.TplName = "line_list.html"
	return
}
