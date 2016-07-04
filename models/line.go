package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Line struct {
	Id      int64
	Name    string `orm:size(50)`
	Conment string `orm:size(50)`
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Line))
}

func GetLineCount() (int64, error) {
	o := orm.NewOrm()
	lines := make([]*Line, 0)
	total, err := o.QueryTable("line").All(&lines)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetLines(currPage, pageSize int) ([]*Line, int64, error) {
	o := orm.NewOrm()
	lines := make([]*Line, 0)
	total, err := o.QueryTable("line").Limit(pageSize, (currPage-1)*pageSize).All(&lines)
	if err != nil {
		return nil, 0, err
	}
	return lines, total, err
}

func GetLineById(id string) (*Line, error) {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	line := &Line{}
	err = o.QueryTable("line").Filter("id", gid).One(line)
	return line, err
}

func AddLine(name, conment string) error {
	o := orm.NewOrm()
	line := &Line{
		Name:    name,
		Conment: conment,
		Created: time.Now(),
	}
	err := o.QueryTable("line").Filter("name", name).One(line)
	if err == nil {
		return nil
	}
	_, err = o.Insert(line)
	return err
}

func ModifyLine(id, name, conment string) error {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	line := &Line{
		Id: gid,
	}
	err = o.Read(line)
	if err == nil {
		line.Name = name
		line.Conment = conment
	}
	o.Update(line)
	return err
}

func DeleteLine(id string) error {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	line := &Line{
		Id: gid,
	}
	_, err = o.Delete(line)
	if err != nil {
		return err
	}
	return nil
}

func SearchLineCount(name string) (int64, error) {
	o := orm.NewOrm()
	lines := make([]*Line, 0)
	total, err := o.QueryTable("line").Filter("name__icontains", name).All(&lines)
	return total, err
}

func SearchLineByName(currPage, pageSize int, name string) ([]*Line, error) {
	o := orm.NewOrm()
	lines := make([]*Line, 0)
	_, err := o.QueryTable("line").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&lines)
	return lines, err
}

func GetLineNames() ([]*Line, error) {
	o := orm.NewOrm()
	lines := make([]*Line, 0)
	_, err := o.QueryTable("line").All(&lines)
	if err != nil {
		return nil, err
	}
	return lines, err
}
