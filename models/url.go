package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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

func QueryUrlExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select name,comment from url")
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
