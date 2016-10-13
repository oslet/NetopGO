package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Url struct {
	Id      int64
	Name    string `orm:size(50)`
	Comment string `orm:size(50)`
	Created string
}

func init() {
	orm.RegisterModel(new(Url))
}

func GetUrlCount() (int64, error) {
	o := orm.NewOrm()
	urls := make([]*Url, 0)
	total, err := o.QueryTable("url").All(&urls)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetUrls(currPage, pageSize int) ([]*Url, int64, error) {
	o := orm.NewOrm()
	urls := make([]*Url, 0)
	total, err := o.QueryTable("url").Limit(pageSize, (currPage-1)*pageSize).All(&urls)
	if err != nil {
		return nil, 0, err
	}
	return urls, total, err
}

func GetUrlById(id string) (*Url, error) {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	url := &Url{}
	err = o.QueryTable("url").Filter("id", gid).One(url)
	return url, err
}

func AddUrl(name, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	url := &Url{
		Name:    name,
		Comment: comment,
		Created: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := o.QueryTable("url").Filter("name", name).One(url)
	if err == nil {
		msg = "url" + name + "已存在"
		return nil, msg
	}
	_, err = o.Insert(url)
	msg = "添加url成功"
	return err, msg
}

func ModifyUrl(id, name, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	gid, err := strconv.ParseInt(id, 10, 64)
	url := &Url{
		Id: gid,
	}
	err = o.Read(url)
	if err == nil {
		url.Name = name
		url.Comment = comment
	}
	o.Update(url)
	msg = "修改成功"
	return err, msg
}

func DeleteUrl(id string) error {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	url := &Url{
		Id: gid,
	}
	_, err = o.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func SearchUrlCount(name string) (int64, error) {
	o := orm.NewOrm()
	urls := make([]*Url, 0)
	total, err := o.QueryTable("url").Filter("name__icontains", name).All(&urls)
	return total, err
}

func SearchUrlByName(currPage, pageSize int, name string) ([]*Url, error) {
	o := orm.NewOrm()
	urls := make([]*Url, 0)
	_, err := o.QueryTable("url").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&urls)
	return urls, err
}

func GetUrlNames() ([]*Url, error) {
	o := orm.NewOrm()
	urls := make([]*Url, 0)
	_, err := o.QueryTable("url").All(&urls)
	if err != nil {
		return nil, err
	}
	return urls, err
}
