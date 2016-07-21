package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Group struct {
	Id      int64
	Name    string `orm:size(50)`
	Conment string `orm:size(50)`
	Created string
}

func init() {
	orm.RegisterModel(new(Group))
}

func GetGroupCount() (int64, error) {
	o := orm.NewOrm()
	groups := make([]*Group, 0)
	total, err := o.QueryTable("group").All(&groups)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetGroups(currPage, pageSize int) ([]*Group, int64, error) {
	o := orm.NewOrm()
	groups := make([]*Group, 0)
	total, err := o.QueryTable("group").Limit(pageSize, (currPage-1)*pageSize).All(&groups)
	if err != nil {
		return nil, 0, err
	}
	return groups, total, err
}

func GetGroupById(id string) (*Group, error) {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	group := &Group{}
	err = o.QueryTable("group").Filter("id", gid).One(group)
	return group, err
}

func AddGroup(name, conment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	group := &Group{
		Name:    name,
		Conment: conment,
		Created: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := o.QueryTable("group").Filter("name", name).One(group)
	if err == nil {
		msg = "业务组" + name + " 已存在 "
		return nil, msg
	}
	_, err = o.Insert(group)
	msg = "添加业务组成功"
	return err, msg
}

func ModifyGroup(id, name, conment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	gid, err := strconv.ParseInt(id, 10, 64)
	group := &Group{
		Id: gid,
	}
	err = o.Read(group)
	if err == nil {
		group.Name = name
		group.Conment = conment
	}
	o.Update(group)
	return err, msg
}

func DeleteGroup(id string) error {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	group := &Group{
		Id: gid,
	}
	_, err = o.Delete(group)
	if err != nil {
		return err
	}
	return nil
}

func SearchGroupCount(name string) (int64, error) {
	o := orm.NewOrm()
	groups := make([]*Group, 0)
	total, err := o.QueryTable("group").Filter("name__icontains", name).All(&groups)
	return total, err
}

func SearchGroupByName(currPage, pageSize int, name string) ([]*Group, error) {
	o := orm.NewOrm()
	groups := make([]*Group, 0)
	_, err := o.QueryTable("group").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&groups)
	return groups, err
}

func GetNames() ([]*Group, error) {
	o := orm.NewOrm()
	groups := make([]*Group, 0)
	_, err := o.QueryTable("group").All(&groups)
	if err != nil {
		return nil, err
	}
	return groups, err
}
