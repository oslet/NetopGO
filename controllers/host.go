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
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

type HostController struct {
	BaseController
}

func (this *HostController) Get() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/host/list"

	groups, err := models.GetNames()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Groups"] = groups

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetHostCount()
	hosts, _, err := models.GetHosts(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total

	this.TplName = "host_list.html"
	return
}
func (this *HostController) Add() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"
	auth := role.(int64)
	this.Data["Auth"] = auth

	groups, err := models.GetNames()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Groups"] = groups

	id := this.Input().Get("id")
	if len(id) > 0 {
		host, err := models.GetHostById(id)
		if err != nil {
			beego.Error(err)
		}
		host.Rootpwd, _ = models.AESDecode(host.Rootpwd, models.AesKey)
		host.Readpwd, _ = models.AESDecode(host.Readpwd, models.AesKey)
		//fmt.Printf("***root :%v,read :%v\n", host.Rootpwd, host.Readpwd)
		this.Data["Host"] = host
		this.Data["HostGroupName"] = host.Group
		this.Data["Path1"] = "主机列表"
		this.Data["Path2"] = "修改主机"
		this.Data["Href"] = "/host/list"
		this.TplName = "host_modify.html"
		//this.TplName = "test.html"
		return
	}
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = "添加主机"
	this.Data["Href"] = "/host/list"
	this.TplName = "host_add.html"

}

func (this *HostController) Post() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "host"
	Auth := role.(int64)
	this.Data["Auth"] = Auth

	id := this.Input().Get("id")
	class := this.Input().Get("class")
	service_name := this.Input().Get("service_name")
	name := this.Input().Get("name")
	ip := this.Input().Get("ip")
	port := this.Input().Get("port")
	os_type := this.Input().Get("os_type")
	owner := this.Input().Get("owner")
	root := this.Input().Get("root")
	read := this.Input().Get("read")
	rootpwd := this.Input().Get("rootpwd")
	readpwd := this.Input().Get("readpwd")
	cpu := this.Input().Get("cpu")
	mem := this.Input().Get("mem")
	disk := this.Input().Get("disk")
	group := this.Input().Get("group")
	idc := this.Input().Get("idc")
	comment := this.Input().Get("comment")
	//beego.Info(idc)
	if len(id) > 0 {
		err, msg := models.ModifyHost(id, class, service_name, name, ip, port, os_type, owner, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	} else {
		err, msg := models.AddHost(class, service_name, name, ip, port, os_type, owner, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Message"] = msg
	}
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/host/list"
	//this.("/host/list", 302)
	this.TplName = "host_add.html"
}

func (this *HostController) Delete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"

	id := this.Input().Get("id")
	err := models.DeleteHost(id)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/host/list"
	this.Redirect("/host/list", 302)
	return
}

func (this *HostController) BitchDelete() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"

	ids := strings.Split(this.Input().Get("ids"), ",")
	for _, id := range ids {
		err := models.DeleteHost(id)
		if err != nil {
			this.Ctx.WriteString("删除失败！")
		}
	}
	//this.Redirect("/user/list", 302)
	this.Ctx.WriteString("删除成功！")
	return
}

func (this *HostController) ReportWeek() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/host"
	this.Data["IsSearch"] = false
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/report/host/list"

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.GetHostCount()
	hosts, _, err := models.GetHosts(int(currPage), int(pageSize))
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), int(pageSize), total)

	auth := role.(int64)
	this.Data["Auth"] = auth
	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total

	this.TplName = "report_host_list.html"
	return
}

func (this *HostController) SearchWeek() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/host"

	week := this.Input().Get("keyword")

	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchHostWeekCount()
	hosts, err := models.SearchHostByWeek(int(currPage), int(pageSize), week)
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
	this.Data["Href"] = "/report/host/list"
	this.TplName = "report_host_list.html"
	return
}

func (this *HostController) Export() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/host"
	method := this.Input().Get("method")
	values, columns, _ := models.QueryHostWeekExport(method)
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
		filename = "week_host" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
	} else if method == "all" {
		filename = "all_host" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19] + ".xlsx"
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

func (this *HostController) Search() {
	var page string
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"

	name := this.Input().Get("keyword")
	idc := this.Input().Get("idc")
	if idc == "1" {
		this.Data["Path1"] = "主机列表"
		this.Data["Path2"] = ""
		this.Data["Href"] = "/host/list"
		this.TplName = "host_list.html"
		return
	}
	//beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize, _ := strconv.ParseInt(beego.AppConfig.String("pageSize"), 10, 64)
	total, err := models.SearchHostCount(idc, name)
	hosts, err := models.SearchHostByName(int(currPage), int(pageSize), idc, name)
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
	this.Data["Keyword"] = name
	this.Data["Idc"] = idc
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/host/list"
	this.TplName = "host_list.html"
	return
}

func (this *HostController) ReportSendMail() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "report/host"
	Auth := role.(int64)
	this.Data["Auth"] = Auth
	//date := time.Now().Format("2006-01-02 15:04:05")

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

	host := "mail.in.7daysinn.cn"
	email := "falconalert@in.7daysinn.cn"
	password := "RaAgpfAbcM8ubfNU5vnJr__rGMp5gOCP"
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
	this.Redirect("/report/host/list", 302)
	return
}

func (this *HostController) WebConsole() {
	uid, uname, role, _ := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"

	id := this.Input().Get("id")
	ip := this.Input().Get("ip")
	user := this.Input().Get("user")
	account := this.Input().Get("role")

	host, err := models.GetHostById(id)
	if err != nil {
		beego.Error(err)
	}
	vmAddr := ip + ":22"
	var passwd string
	if "1" == account {
		passwd, _ = models.AESDecode(host.Rootpwd, models.AesKey)
	} else {
		passwd, _ = models.AESDecode(host.Readpwd, models.AesKey)
	}

	ssh_info := make([]string, 0, 0)
	ssh_info = append(ssh_info, user)
	ssh_info = append(ssh_info, passwd)
	ssh_info = append(ssh_info, vmAddr)
	//beego.Info(user)
	//beego.Info(passwd)
	//beego.Info(vmAddr)

	b64_ssh_info, err := models.AESEncode(strings.Join(ssh_info, "\n"), models.AesKey)
	//beego.Info(b64_ssh_info)
	wsAddr := "ws://" + this.Ctx.Request.Host + "/console/sshws" + "?vm_info=" + b64_ssh_info
	//beego.Info(wsAddr)
	this.Data["Uname"] = uname
	this.Data["WsAddr"] = wsAddr
	this.TplName = "console.html"
	return
}
