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

type UrlController struct {
	BaseController
}

func (this *UrlController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "url列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/url/list"
	this.Data["Category"] = "asset/url"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetUrlCount()
	urls, _, err := models.GetUrls(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Urls"] = urls
	this.Data["totals"] = total

	this.TplName = "url_list.html"
	return
}
func (this *UrlController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "Url"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	if len(id) > 0 {
		url, err := models.GetUrlById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Url"] = url
		this.Data["Path1"] = "url列表"
		this.Data["Path2"] = "修改url"
		this.Data["Href"] = "/asset/url/list"
		this.TplName = "url_modify.html"
		return
	}
	this.Data["Path1"] = "url列表"
	this.Data["Path2"] = "添加url"
	this.Data["Href"] = "/asset/url/list"
	this.TplName = "url_add.html"

}

func (this *UrlController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "url"

	Auth := role.(int64)
	this.Data["Auth"] = Auth
	id := this.Input().Get("id")
	name := this.Input().Get("name")
	comment := this.Input().Get("comment")
	if len(id) > 0 {
		err, msg := models.ModifyUrl(id, name, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddUrl(name, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "url列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/url/list"
	//this.Redirect("/asset/url/list", 302)
	this.TplName = "url_add.html"
}

func (this *UrlController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "url"

	id := this.Input().Get("id")
	err := models.DeleteUrl(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "url列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/url/list"
	this.Redirect("/asset/url/list", 302)
	return
}

func (this *UrlController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "url"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteUrl(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *UrlController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "asset/url"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchUrlCount(name)
	urls, err := models.SearchUrlByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Urls"] = urls
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "url列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/asset/url/list"
	this.TplName = "url_list.html"
	return
}

func (this *UrlController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "line"
	values, columns, _ := models.QueryUrlExport()

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
	filename := "all_url" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

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
