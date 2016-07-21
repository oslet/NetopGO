package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysList struct {
	Id          int64
	Class       string `orm:size(50)`
	Name        string `orm:size(50)`
	Owner1      string `orm:size(50)`
	Owner2      string `orm:size(10)`
	Domain_name string `orm:size(50)`
	Comment     string `orm:size(100)`
	Created     time.Time
}

// register host model
func init() {
	orm.RegisterModel(new(SysList))
}

func GetSysListCount() (int64, error) {
	o := orm.NewOrm()
	Syslist := make([]*SysList, 0)
	total, err := o.QueryTable("sys_list").All(&Syslist)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetSysLists(currPage, pageSize int) ([]*SysList, int64, error) {
	o := orm.NewOrm()
	Syslist := make([]*SysList, 0)
	total, err := o.QueryTable("sys_list").Limit(pageSize, (currPage-1)*pageSize).All(&Syslist)
	if err != nil {
		return nil, 0, err
	}
	return Syslist, total, err
}

func GetSyslistById(id string) (*SysList, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	Syslist := &SysList{}
	err = o.QueryTable("sys_list").Filter("id", hid).One(Syslist)
	return Syslist, err
}

func AddSyslist(class, name, owner1, owner2, domain_name, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	Syslist := &SysList{
		Class:       class,
		Name:        name,
		Owner1:      owner1,
		Owner2:      owner2,
		Domain_name: domain_name,
		Comment:     comment,
		Created:     time.Now(),
	}
	err := o.QueryTable("sys_list").Filter("name", name).One(Syslist)
	if err == nil {
		msg = "系统名已存在"
		return nil, msg
	}
	_, err = o.Insert(Syslist)
	msg = "添加系统成功"
	return err, msg
}

func ModifySyslist(id, class, name, owner1, owner2, domain_name, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	hid, err := strconv.ParseInt(id, 10, 64)
	Syslist := &SysList{
		Id: hid,
	}
	err = o.Read(Syslist)
	if err == nil {
		Syslist.Class = class
		Syslist.Name = name
		Syslist.Owner1 = owner1
		Syslist.Owner2 = owner2
		Syslist.Domain_name = domain_name
		Syslist.Comment = comment
	}
	o.Update(Syslist)
	msg = "修改成功"
	return err, msg
}

func DeleteSyslist(id string) error {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	Syslist := &SysList{
		Id: hid,
	}
	_, err = o.Delete(Syslist)
	if err != nil {
		return err
	}
	return nil
}

func SearchSyslistCount(class, name string) (int64, error) {
	o := orm.NewOrm()
	Syslists := make([]*SysList, 0)
	total, err := o.QueryTable("sys_list").Filter("class", class).Filter("name__icontains", name).All(&Syslists)
	return total, err
}

func SearchSyslistByName(currPage, pageSize int, class, name string) ([]*SysList, error) {
	o := orm.NewOrm()
	Syslists := make([]*SysList, 0)
	/*
		var cond *orm.Condition
		cond = orm.NewCondition()
		cond = cond.Or("name__icontains", name)
		//cond = cond.Or("ip__icontains", "ip")
		var qs orm.QuerySeter
		qs = o.QueryTable("sys_list").Limit(pageSize, (currPage-1)*pageSize).SetCond(cond)
		_, err := qs.All(&Syslists)
	*/
	_, err := o.QueryTable("sys_list").Filter("class", class).Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&Syslists)
	return Syslists, err
}
