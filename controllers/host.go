package controllers

import (
	"NetopGO/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type HostController struct {
	BaseController
}

func (this *HostController) Get() {
	var page string
	uid, uname, role := this.IsLogined()
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
	pageSize := 2
	total, err := models.GetHostCount()
	hosts, _, err := models.GetHosts(int(currPage), pageSize)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)

	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total

	this.TplName = "host_list.html"
	return
}
func (this *HostController) Add() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"

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
		fmt.Printf("***root :%v,read :%v\n", host.Rootpwd, host.Readpwd)
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
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["IsSearch"] = false
	this.Data["Category"] = "host"

	id := this.Input().Get("id")
	name := this.Input().Get("name")
	ip := this.Input().Get("ip")
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
	beego.Info(idc)
	if len(id) > 0 {
		err := models.ModifyHost(id, name, ip, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.AddHost(name, ip, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = ""
	this.Data["Href"] = "/host/list"
	this.Redirect("/host/list", 302)
}

func (this *HostController) Delete() {
	uid, uname, role := this.IsLogined()
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
	uid, uname, role := this.IsLogined()
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

func (this *HostController) Search() {
	var page string
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.Data["Category"] = "host"

	name := this.Input().Get("keyword")
	group := this.Input().Get("group")
	idc := this.Input().Get("idc")
	if group == "1" || idc == "1" {
		this.Data["Path1"] = "主机列表"
		this.Data["Path2"] = ""
		this.Data["Href"] = "/host/list"
		this.TplName = "host_list.html"
		return
	}
	beego.Info(name)
	if len(this.Input().Get("page")) == 0 {
		page = "1"
	} else {
		page = this.Input().Get("page")
	}
	currPage, _ := strconv.ParseInt(page, 10, 64)
	pageSize := 1
	total, err := models.SearchHostCount(idc, group, name)
	hosts, err := models.SearchHostByName(int(currPage), pageSize, idc, group, name)
	if err != nil {
		beego.Error(err)
	}
	res := models.Paginator(int(currPage), pageSize, total)
	this.Data["paginator"] = res
	this.Data["Hosts"] = hosts
	this.Data["totals"] = total
	this.Data["IsSearch"] = true
	this.Data["Keyword"] = name
	this.Data["Idc"] = idc
	this.Data["Group"] = group
	this.Data["Path1"] = "主机列表"
	this.Data["Path2"] = "搜索结果"
	this.Data["Href"] = "/host/list"
	this.TplName = "host_list.html"
	return
}

func (this *HostController) WebConsole() {
	uid, uname, role := this.IsLogined()
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
	beego.Info(user)
	beego.Info(passwd)
	beego.Info(vmAddr)

	b64_ssh_info, err := models.AESEncode(strings.Join(ssh_info, "\n"), models.AesKey)
	beego.Info(b64_ssh_info)
	wsAddr := "ws://" + this.Ctx.Request.Host + "/console/sshws" + "?vm_info=" + b64_ssh_info
	beego.Info(wsAddr)
	this.Data["Uname"] = uname
	this.Data["WsAddr"] = wsAddr
	this.TplName = "console.html"
	return
}
