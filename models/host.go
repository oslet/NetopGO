package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Host struct {
	Id      int64
	Name    string `orm:size(50)`
	Ip      string `orm:size(15)`
	Cpu     string `orm:size(50)`
	Mem     string `orm:size(50)`
	Disk    string `orm:size(50)`
	Idc     string `orm:size(50)`
	Rootpwd string `orm:size(50)`
	Readpwd string `orm:size(50)`
	Group   string `orm:size(50)`
	Created time.Time
}

// register host model
func init() {
	orm.RegisterModel(new(Host))
}

func GetHostCount() (int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").All(&hosts)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetHosts(currPage, pageSize int) ([]*Host, int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
	if err != nil {
		return nil, 0, err
	}
	return hosts, total, err
}

func GetHostById(id string) (*Host, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{}
	err = o.QueryTable("host").Filter("id", hid).One(host)
	return host, err
}

func AddHost(name, ip, rootpwd, readpwd, cpu, mem, disk, group, idc string) error {
	o := orm.NewOrm()
	rootpwd = Md5Encode([]byte(rootpwd))
	readpwd = Md5Encode([]byte(readpwd))
	host := &Host{
		Name:    name,
		Ip:      ip,
		Rootpwd: rootpwd,
		Readpwd: readpwd,
		Cpu:     cpu,
		Mem:     mem,
		Disk:    disk,
		Group:   group,
		Idc:     idc,
		Created: time.Now(),
	}
	err := o.QueryTable("host").Filter("name", name).One(host)
	if err == nil {
		return nil
	}
	_, err = o.Insert(host)
	return err
}

func ModifyHost(id, name, ip, rootpwd, readpwd, cpu, mem, disk, group, idc string) error {
	o := orm.NewOrm()
	rootpwd = Md5Encode([]byte(rootpwd))
	readpwd = Md5Encode([]byte(readpwd))
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{
		Id: hid,
	}
	err = o.Read(host)
	if err == nil {
		host.Name = name
		host.Ip = ip
		host.Rootpwd = rootpwd
		host.Readpwd = readpwd
		host.Cpu = cpu
		host.Mem = mem
		host.Disk = disk
		host.Group = group
		host.Idc = idc
	}
	o.Update(host)
	return err
}

func DeleteHost(id string) error {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{
		Id: hid,
	}
	_, err = o.Delete(host)
	if err != nil {
		return err
	}
	return nil
}

func SearchHostCount(idc, group, name string) (int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").Filter("idc", idc).Filter("group", group).Filter("name__icontains", name).All(&hosts)
	return total, err
}

func SearchHostByName(currPage, pageSize int, idc, group, name string) ([]*Host, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	_, err := o.QueryTable("host").Filter("idc", idc).Filter("group", group).Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
	return hosts, err
}
