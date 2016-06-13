package controllers

import (
	"NetopGO/models"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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
	quest := this.Input().Get("quest")
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	list := beego.AppConfig.String("app_name")
	arrList := strings.Split(list, ",")

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetFaultRecordCount(quest)
	faultRecs, _, err := models.GetFaultRecords(int(currPage), int(pageSize), quest)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["Question"] = quest
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
	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	appNameList := strings.Split(beego.AppConfig.String("AppName"), ",")
	questNames := models.GetQuestionNames()
	this.Data["QuestNames"] = questNames
	this.Data["AppTypeList"] = appTypeList
	this.Data["AppNameList"] = appNameList
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
	level := this.Input().Get("level")
	question := this.Input().Get("question")
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
		err := models.AddFaultRecord(question, level, system, appname, category, issolu, operater, starttime, endtime, solution, effection, analysis, nextstep)
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
	quest := this.Input().Get("quest")
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
	total, err := models.SearchFaultRecCount(cate, quest)
	faultRecs, err := models.SearchFaultRecByName(int(currPage), int(pageSize), cate, quest)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Question"] = quest
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

func (this *FaultRecordController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/fault"
	values, columns, _ := models.QueryFaultExport()

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	for _, val := range columns {
		cell := row.AddCell()
		cell.Value = val
	}
	//sheet.SetColWidth(1, len(columns), 100)
	for _, val := range *values {
		row = sheet.AddRow()
		for _, value := range val {
			cell := row.AddCell()
			cell.Value = value
		}
	}
	now := time.Now().String()
	filename := "all_fault" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

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
