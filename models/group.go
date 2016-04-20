package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Group struct {
	Id      int64
	Name    string `orm:size(50)`
	Conment string `orm:size(50)`
	Created time.Time
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

func AddGroup(name, conment string) error {
	o := orm.NewOrm()
	group := &Group{
		Name:    name,
		Conment: conment,
		Created: time.Now(),
	}
	err := o.QueryTable("group").Filter("name", name).One(group)
	if err == nil {
		return nil
	}
	_, err = o.Insert(group)
	return err
}

func ModifyGroup(id, name, conment string) error {
	o := orm.NewOrm()
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
	return err
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
