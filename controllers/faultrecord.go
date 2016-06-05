package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	//"path"
	"strconv"
	"strings"
)

type FaultRecordController struct {
	BaseController
}

func (this *FaultRecordController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	list := beego.AppConfig.String("app_name")
	arrList := strings.Split(list, ",")

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetFaultRecordCount()
	faultRecs, _, err := models.GetFaultRecords(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["List"] = arrList
	this.Data["paginator"] = res
	this.Data["FaultRecords"] = faultRecs
	this.Data["totals"] = total

	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/fault/list"
	this.Data["Category"] = "record/fault"
	this.TplName = "fault_record_list.html"
	return
}

func (this *FaultRecordController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/fault"
	list := beego.AppConfig.String("app_name")
	arrList := strings.Split(list, ",")

	this.Data["List"] = arrList
	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = "添加记录"
	this.Data["Href"] = "/record/fault/list"
	this.TplName = "fault_record_add.html"
	return
}

func (this *FaultRecordController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "record/fault"

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	level := this.Input().Get("level")
	system := this.Input().Get("system")
	appname := this.Input().Get("appname")
	category := this.Input().Get("category")
	issolu := this.Input().Get("issolu")
	operater := this.Input().Get("operater")
	starttime := this.Input().Get("starttime")
	endtime := this.Input().Get("endtime")
	solution := this.Input().Get("solution")
	effection := this.Input().Get("effection")
	analysis := this.Input().Get("analysis")
	nextstep := this.Input().Get("nextstep")

	if len(id) > 0 {
		faultRec, err := models.FaultRecordDetail(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["FaultRecord"] = faultRec
		this.Data["Path1"] = "故障记录"
		this.Data["Path2"] = "详细内容"
		this.Data["Href"] = "/record/fault/list"
		this.TplName = "fault_record_detail.html"
		return
	} else {
		err := models.AddFaultRecord(name, level, system, appname, category, issolu, operater, starttime, endtime, solution, effection, analysis, nextstep)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/fault/list"
	this.Redirect("/record/fault/list", 302)
	return
}

func (this *FaultRecordController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/fault"

	id := this.Input().Get("id")
	err := models.DeleteFaultRecord(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/fault/list"
	this.Redirect("/record/fault/list", 302)
	return
}

func (this *FaultRecordController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/fault"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteFaultRecord(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *FaultRecordController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/fault"

	cate := this.Input().Get("keyword")
	if "1" == cate {
		this.Data["Path1"] = "故障记录"
		this.Data["Path2"] = ""
		this.Data["Href"] = "/record/fault/list"
		this.Redirect("/record/fault/list", 302)
		return
	}
	//beego.Info(cate)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchFaultRecCount(cate)
	faultRecs, err := models.SearchFaultRecByName(int(currPage), int(pageSize), cate)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["FaultRecords"] = faultRecs
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = cate
	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/record/fault/list"
	this.TplName = "fault_record_list.html"
	return
}
