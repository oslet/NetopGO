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

type DblistController struct {
	BaseController
}

func (this *DblistController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname.(string)
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "dblist列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/dblist/list"
	this.Data["Category"] = "asset/dblist"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetDblistCount()
	dblists, _, err := models.GetDblists(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["dblists"] = dblists
	this.Data["totals"] = total

	this.TplName = "dblist_list.html"
	return
}
func (this *DblistController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dblist"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	groups, err := models.GetNames()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Groups"] = groups

	id := this.Input().Get("id")
	if len(id) > 0 {
		dblist, err := models.GetDblistById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Dblist"] = dblist
		this.Data["DblistGroupName"] = dblist.Name
		this.Data["Path1"] = "dblist列表"
		this.Data["Path2"] = "修改dblist"
		this.Data["Href"] = "/asset/dblist/list"
		this.TplName = "dblist_modify.html"
		return
	}
	this.Data["Path1"] = "dblist列表"
	this.Data["Path2"] = "添加dblist"
	this.Data["Href"] = "/asset/dblist/list"
	this.TplName = "dblist_add.html"

}

func (this *DblistController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "dblist"

	Auth := role.(int64)
	this.Data["Auth"] = Auth
	id := this.Input().Get("id")
	ip := this.Input().Get("ip")
	port := this.Input().Get("port")
	dbinst := this.Input().Get("dbinst")
	dbname := this.Input().Get("dbname")
	isswitch := this.Input().Get("isswitch")
	attrteam := this.Input().Get("attrteam")
	name := this.Input().Get("name")
	if len(id) > 0 {
		err, msg := models.ModifyDblist(id, ip, port, dbinst, dbname, isswitch, attrteam, name)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddDblist(ip, port, dbinst, dbname, isswitch, attrteam, name)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "dblist列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/dblist/list"
	//this.Redirect("/asset/dblist/list", 302)
	this.TplName = "dblist_add.html"
}

func (this *DblistController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dblist"

	id := this.Input().Get("id")
	err := models.DeleteDblist(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "dblist列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/asset/dblist/list"
	this.Redirect("/asset/dblist/list", 302)
	return
}

func (this *DblistController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dblist"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteDblist(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *DblistController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "asset/dblist"

	name := this.Input().Get("keyword")
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchDblistCount(name)
	dblists, err := models.SearchDblistByName(int(currPage), int(pageSize), name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["dblists"] = dblists
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Path1"] = "dblist列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/asset/dblist/list"
	this.TplName = "dblist_list.html"
	return
}

func (this *DblistController) Detail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	id := this.Input().Get("id")
	dblistname, err := models.GetDblistById(id)
	if err != nil {
		beego.Error(err)
	}
	auth := role.(int64)
	this.Data["Auth"] = auth

	this.Data["dblist"] = dblistname
	this.Data["Path1"] = "db列表"
	this.Data["Path2"] = "db详情"
	this.Data["Href"] = "/asset/dblist/list"
	this.Data["Category"] = "dblist"
	this.TplName = "dblist_detail.html"
	return
}

func (this *DblistController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	Uname := uname.(string)
	this.Data["Uname"] = Uname
	this.Data["Role"] = role
	this.Data["Category"] = "dblist"
	values, columns, _ := models.QueryDblistExport()

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("db列表")
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
	filename := "all_dblist" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

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
