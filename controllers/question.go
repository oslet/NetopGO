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

type QuestController struct {
	BaseController
}

func (this *QuestController) Get() {
	var page string
	uid, uname, role, dept := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Dept"] = dept
	this.Data["IsSearch"] = false
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	// var questRecs []*Question
	// var total int64
	// var err error
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetQuestRecordCount()
	questRecs, _, err := models.GetQuestRecords(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}

	res := models.Paginator(int(currPage), int(pageSize), total)

	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")
	this.Data["AppTypeList"] = appTypeList
	this.Data["QuestRecords"] = questRecs
	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["totals"] = total
	this.Data["Path1"] = "问题记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/quest/list"
	this.Data["Category"] = "record/quest"
	this.TplName = "question_list.html"
	return
}

func (this *QuestController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/quest"
	appTypeList := strings.Split(beego.AppConfig.String("AppType"), ",")

	id := this.Input().Get("id")
	if len(id) > 0 {
		quest, err := models.GetQuestionById(id)
		if err != nil {
			beego.Error(err)
		}
		this.Data["AppTypeList"] = appTypeList
		this.Data["QuestRecord"] = quest
		this.Data["Path1"] = "问题记录"
		this.Data["Path2"] = "修改问题"
		this.Data["Href"] = "/record/quest/list"
		this.TplName = "quest_modify.html"
		return
	}
	this.Data["AppTypeList"] = appTypeList
	this.Data["Path1"] = "问题记录"
	this.Data["Path2"] = "添加记录"
	this.Data["Href"] = "/record/quest/list"
	this.TplName = "quest_add.html"
	return
}

func (this *QuestController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "record/quest"

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	influce := this.Input().Get("influce")
	owner := this.Input().Get("owner")
	status := this.Input().Get("status")
	comment := this.Input().Get("comment")
	beego.Info(id)
	if len(id) > 0 {
		err := models.ModifyQuestion(id, name, influce, owner, status, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/record/quest/list", 302)
		return
	} else {
		err := models.AddQuestRecord(name, influce, owner, status, comment)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "问题记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/quest/list"
	this.Redirect("/record/quest/list", 302)
	return
}

func (this *QuestController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/quest"

	id := this.Input().Get("id")
	err := models.DeleteQuestRecord(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/record/quest/list"
	this.Redirect("/record/quest/list", 302)
	return
}

func (this *QuestController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/quest"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteQuestRecord(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *QuestController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/quest"

	apptype := this.Input().Get("keyword")
	if "1" == apptype {
		this.Data["Path1"] = "故障记录"
		this.Data["Path2"] = ""
		this.Data["Href"] = "/record/quest/list"
		this.Redirect("/record/quest/list", 302)
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
	total, err := models.SearchQuestRecCount(apptype)
	questRecs, err := models.SearchQuestRecByAppType(int(currPage), int(pageSize), apptype)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["QuestRecords"] = questRecs
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = apptype
	this.Data["Path1"] = "故障记录"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/record/quest/list"
	this.TplName = "question_list.html"
	return
}

func (this *QuestController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "record/quest"
	values, columns, _ := models.QueryQuestionExport()

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
	filename := "all_question" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"

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
