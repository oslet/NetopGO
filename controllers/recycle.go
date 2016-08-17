package controllers

import (
	"NetopGO/models"
	//"fmt"
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

type RecycleHostController struct {
	BaseController
}

func (this *RecycleHostController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/recycle"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "主机回收列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/report/recycle/list"

	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetRecycleHostCount()
	hosts, _, err := models.GetRecycleHosts(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total

	this.TplName = "report_host_recycle.html"
	return
}

func (this *RecycleHostController) ReportWeek() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/recycle"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/report/recycle/list"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetRecycleHostCount()
	hosts, _, err := models.GetRecycleHosts(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total

	this.TplName = "report_host_recycle.html"
	return
}

func (this *RecycleHostController) SearchWeek() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/recycle"

	week := this.Input().Get("keyword")

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchRecycleHostWeekCount()
	hosts, err := models.SearchRecycleHostByWeek(int(currPage), int(pageSize), week)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	//this.Data["Keyword"] = name
	//this.Data["Idc"] = idc
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/report/recycle/list"
	this.TplName = "report_host_recycle.html"
	return
}

func (this *RecycleHostController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/recycle"
	method := this.Input().Get("method")
	values, columns, _ := models.QueryRecycleHostWeekExport(method)
	auth := role.(int64)
	this.Data["Auth"] = auth

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
	var filename string
	if method == "week" {
		filename = "week_recyclehost" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
	} else if method == "all" {
		filename = "all_recyclehost" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
	}

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

func (this *RecycleHostController) ReportSendMail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/host"
	Auth := role.(int64)
	this.Data["Auth"] = Auth
	//date := time.Now().Format("2006-01-02 15:04:05")

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

	host := "smtp.sina.com"
	email := "falconmail@sina.com"
	password := "NAlv-W73wZNdHJ8i"
	toEmail := "yun.li@platenogroup.com"

	from := mail.Address{"主机报表", email}
	to := mail.Address{"", toEmail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("主机列表")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	body := `
	  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>HOST Result</title>
      </head>
      <body>
        <h2>HOST Result</h2>
        <p>Your email has been sent to: </p>
        <pre>{{html .}}</pre>
        <p><a href="/">Start again!</a></p>
        <div>
          <p><b>© 2014 RubyLearning. All rights reserved.</b></p>
        </div>
      </body>
    </html>
`

	//	var mailTemplate = template.Must(template.New("mail").Parse(mailTemplateHTML))
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)

	err := smtp.SendMail(
		host+":25",
		auth,
		email,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		panic(err)
	}
	this.Redirect("/report/recycle/list", 302)
	return
}
