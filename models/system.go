package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type System struct {
	Id            int64
	Class         string `orm:size(50)`
	Status        string `orm:size(20)`
	Name          string `orm:size(50)`
	Buss_owner    string `orm:size(20)`
	Buss_attr     string `orm:size(20)`
	Domain_name   string `orm:size(50)`
	Risk_category string `orm:size(50)`
	Attr_dev	string `"column(attr_dev)"`
	Attr_test	string `"column(attr_test)"`
	Attr_arch	string `"column(attr_arch)"`
	Attr_ops	string `"column(attr_ops)"`
	Controller    string `orm:size(50)`
	Responsible   string `orm:size(20)`
	Team          string `orm:size(20)`
	Company       string `orm:size(50)`
	Support_level string `orm:size(20)`
	Numbers       string `orm:size(20)`
	Total_core    string `orm:size(20)`
	Total_mem     string `orm:size(20)`
	Total_disk    string `orm:size(20)`
	Area          string `orm:size(20)`
	Windows       string `orm:size(50)`
	Comment       string `orm:size(100)`
	Created       time.Time
}

// register host model
func init() {
	orm.RegisterModel(new(System))
}

func GetSystemCount() (int64, error) {
	o := orm.NewOrm()
	Systemlist := make([]*System, 0)
	total, err := o.QueryTable("system").All(&Systemlist)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetSystems(currPage, pageSize int) ([]*System, int64, error) {
	o := orm.NewOrm()
	Systemlist := make([]*System, 0)
	total, err := o.QueryTable("system").Limit(pageSize, (currPage-1)*pageSize).All(&Systemlist)
	if err != nil {
		return nil, 0, err
	}
	return Systemlist, total, err
}

func GetSystemlistById(id string) (*System, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	Systemlist := &System{}
	err = o.QueryTable("system").Filter("id", hid).One(Systemlist)
	return Systemlist, err
}

func AddSystemlist(class, status, name, buss_owner, buss_attr, domain_name, risk_category, attr_dev, attr_test, attr_arch, attr_ops, controller, responsible, team, company, support_level, numbers, total_core, total_mem, total_disk, area, windows, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	Systemlist := &System{
		Class:         class,
		Status:        status,
		Name:          name,
		Buss_owner:    buss_owner,
		Buss_attr:     buss_attr,
		Domain_name:   domain_name,
		Risk_category: risk_category,
		Attr_dev: attr_dev,
		Attr_test: attr_test,
		Attr_arch: attr_arch,
		Attr_ops: attr_ops,
		Controller:    controller,
		Responsible:   responsible,
		Team:          team,
		Company:       company,
		Support_level: support_level,
		Numbers:       numbers,
		Total_core:    total_core,
		Total_mem:     total_mem,
		Total_disk:    total_disk,
		Area:          area,
		Windows:       windows,
		Comment:       comment,
		Created:       time.Now(),
	}
	err := o.QueryTable("system").Filter("name", name).One(Systemlist)
	if err == nil {
		msg = "系统名已存在"
		return nil, msg
	}
	_, err = o.Insert(Systemlist)
	msg = "添加系统成功"
	return err, msg
}

func ModifySystemlist(id, class, status, name, buss_owner, buss_attr, domain_name, risk_category, attr_dev, attr_test, attr_arch, attr_ops, controller, responsible, team, company, support_level, numbers, total_core, total_mem, total_disk, area, windows, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	hid, err := strconv.ParseInt(id, 10, 64)
	Systemlist := &System{
		Id: hid,
	}
	err = o.Read(Systemlist)
	if err == nil {
		Systemlist.Class = class
		Systemlist.Status = status
		Systemlist.Name = name
		Systemlist.Buss_owner = buss_owner
		Systemlist.Buss_attr = buss_attr
		Systemlist.Domain_name = domain_name
		Systemlist.Risk_category = risk_category
		Systemlist.Attr_dev = attr_dev
		Systemlist.Attr_test = attr_test
		Systemlist.Attr_arch = attr_arch
		Systemlist.Attr_ops = attr_ops
		Systemlist.Controller = controller
		Systemlist.Comment = comment
		Systemlist.Responsible = responsible
		Systemlist.Team = team
		Systemlist.Company = company
		Systemlist.Support_level = support_level
		Systemlist.Numbers = numbers
		Systemlist.Total_core = total_core
		Systemlist.Total_mem = total_mem
		Systemlist.Total_disk = total_disk
		Systemlist.Area = area
		Systemlist.Windows = windows
	}
	o.Update(Systemlist)
	msg = "修改成功"
	return err, msg
}

func DeleteSystemlist(id string) error {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	Systemlist := &System{
		Id: hid,
	}
	_, err = o.Delete(Systemlist)
	if err != nil {
		return err
	}
	return nil
}

func SearchSystemlistCount(class, name string) (int64, error) {
	o := orm.NewOrm()
	Systemlists := make([]*System, 0)
	total, err := o.QueryTable("system").Filter("class", class).Filter("name__icontains", name).All(&Systemlists)
	return total, err
}

func SearchSystemlistByName(currPage, pageSize int, class, name string) ([]*System, error) {
	o := orm.NewOrm()
	Systemlists := make([]*System, 0)
	_, err := o.QueryTable("system").Filter("class", class).Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&Systemlists)
	return Systemlists, err
}

func GetSystemById(id string) (*System, error) {
	o := orm.NewOrm()
	sid, err := strconv.ParseInt(id, 10, 64)
	system := &System{}
	err = o.QueryTable("system").Filter("id", sid).One(system)
	return system, err
}

func QuerySystemExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select class,status,name,buss_owner,buss_attr, domain_name,risk_category, attr_dev, attr_test, attr_arch, attr_ops, controller,responsible,team,company,support_level,numbers,total_core,total_mem,total_disk,area,windows,comment from system")
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
