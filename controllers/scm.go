package controllers

import (
	"NetopGO/models"
	"strconv"
	"strings"

	"os"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

type ScmController struct {
	BaseController
}

func (this *ScmController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "SCM列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/scm/list"
	this.Data["Category"] = "scm"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetScmCount()
	scms, _, err := models.GetScms(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Scms"] = scms
	this.Data["totals"] = total

	this.TplName = "scm_list.html"
	return
}
func (this *ScmController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "Scm"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	if len(id) > 0 {
		scm, err := models.GetScmById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Scm"] = scm
		this.Data["Path1"] = "SCM列表"
		this.Data["Path2"] = "修改SCM"
		this.Data["Href"] = "/scm/list"
		this.TplName = "scm_modify.html"
		return
	}
	this.Data["Path1"] = "SCM列表"
	this.Data["Path2"] = "添加SCM"
	this.Data["Href"] = "/scm/list"
	this.TplName = "scm_add.html"

}

func (this *ScmController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "scm"

	Auth := role.(int64)
	this.Data["Auth"] = Auth
	id := this.Input().Get("id")
	name := this.Input().Get("name")
	isdeployment := this.Input().Get("isdeployment")
	ischeckin := this.Input().Get("ischeckin")
	owner := this.Input().Get("owner")
	company := this.Input().Get("company")
	scmaddr := this.Input().Get("scmaddr")
	comment := this.Input().Get("comment")
	if len(id) > 0 {
		err, msg := models.ModifyScm(id, name, isdeployment, ischeckin, owner, company, scmaddr, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddScm(name, isdeployment, ischeckin, owner, company, scmaddr, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "SCM列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/scm/list"
	//this.Redirect("/scm/list", 302)
	this.TplName = "scm_add.html"
}

func (this *ScmController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "scm"

	id := this.Input().Get("id")
	err := models.DeleteScm(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "SCM列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/scm/list"
	this.Redirect("/scm/list", 302)
	return
}

func (this *ScmController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "scm"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteScm(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *ScmController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "scm"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchScmCount(name)
	scms, err := models.SearchScmByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Scms"] = scms
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "SCM列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/scm/list"
	this.TplName = "scm_list.html"
	return
}

func (this *ScmController) Detail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	id := this.Input().Get("id")
	scmname, err := models.GetScmById(id)
	if err != nil {
		beego.Error(err)
	}
	auth := role.(int64)
	this.Data["Auth"] = auth

	this.Data["Scm"] = scmname
	this.Data["Path1"] = "系统列表"
	this.Data["Path2"] = "系统详情"
	this.Data["Href"] = "/scm/list"
	this.Data["Category"] = "scm"
	this.TplName = "scm_detail.html"
	return
}

func (this *ScmController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "scm"
	values, columns, _ := models.QueryScmExport()

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	for _, val := range columns {
		cell := row.AddCell()
		cell.Value = val
	}
	for _, val := range *values {
		row = sheet.AddRow()
		for _, value := range val {
			cell := row.AddCell()
			cell.Value = value
		}
	}
	now := time.Now().String()
	filename := "all_scm" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

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
