package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Scm struct {
	Id           int64
	Name         string `orm:size(50)`
	Isdeployment string `orm:size(2)`
	Ischeckin    string `orm:size(2)`
	Owner        string `orm:size(20)`
	Company      string `orm:size(50)`
	Scmaddr      string `orm:size(50)`
	Comment      string `orm:size(50)`
	Created      time.Time
}

func init() {
	orm.RegisterModel(new(Scm))
}

func GetScmCount() (int64, error) {
	o := orm.NewOrm()
	scms := make([]*Scm, 0)
	total, err := o.QueryTable("scm").All(&scms)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetScms(currPage, pageSize int) ([]*Scm, int64, error) {
	o := orm.NewOrm()
	scms := make([]*Scm, 0)
	total, err := o.QueryTable("scm").Limit(pageSize, (currPage-1)*pageSize).All(&scms)
	if err != nil {
		return nil, 0, err
	}
	return scms, total, err
}

func GetScmById(id string) (*Scm, error) {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	scm := &Scm{}
	err = o.QueryTable("scm").Filter("id", gid).One(scm)
	return scm, err
}

func AddScm(name, isdeployment, ischeckin, owner, company, scmaddr, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	scm := &Scm{
		Name:         name,
		Isdeployment: isdeployment,
		Ischeckin:    ischeckin,
		Owner:        owner,
		Company:      company,
		Scmaddr:      scmaddr,
		Comment:      comment,
		Created:      time.Now(),
	}
	err := o.QueryTable("scm").Filter("name", name).One(scm)
	if err == nil {
		msg = "线路" + name + "已存在"
		return nil, msg
	}
	_, err = o.Insert(scm)
	msg = "添加线路成功"
	return err, msg
}

func ModifyScm(id, name, isdeployment, ischeckin, owner, company, scmaddr, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	gid, err := strconv.ParseInt(id, 10, 64)
	scm := &Scm{
		Id: gid,
	}
	err = o.Read(scm)
	if err == nil {
		scm.Name = name
		scm.Isdeployment = isdeployment
		scm.Ischeckin = ischeckin
		scm.Owner = owner
		scm.Company = company
		scm.Scmaddr = scmaddr
		scm.Comment = comment
	}
	o.Update(scm)
	msg = "修改成功"
	return err, msg
}

func DeleteScm(id string) error {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	scm := &Scm{
		Id: gid,
	}
	_, err = o.Delete(scm)
	if err != nil {
		return err
	}
	return nil
}

func SearchScmCount(name string) (int64, error) {
	o := orm.NewOrm()
	scms := make([]*Scm, 0)
	total, err := o.QueryTable("scm").Filter("name__icontains", name).All(&scms)
	return total, err
}

func SearchScmByName(currPage, pageSize int, name string) ([]*Scm, error) {
	o := orm.NewOrm()
	scms := make([]*Scm, 0)
	_, err := o.QueryTable("scm").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&scms)
	return scms, err
}

func GetScmNames() ([]*Scm, error) {
	o := orm.NewOrm()
	scms := make([]*Scm, 0)
	_, err := o.QueryTable("scm").All(&scms)
	if err != nil {
		return nil, err
	}
	return scms, err
}

func QueryScmExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select name,isdeployment,ischeckin,owner,company,scmaddr,comment from scm")
	if err != nil {
		return &result, columns, total
	}
	defer rows.Close()
	columns, err = rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		var row []string
		_ = rows.Scan(scans...)
		for _, col := range values {
			row = append(row, string(col))
		}
		total = total + 1
		result[total] = row
	}

	return &result, columns, total
}
