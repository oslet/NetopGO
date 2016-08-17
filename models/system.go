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
	Id          int64
	Class       string `orm:size(50)`
	Name        string `orm:size(50)`
	Owner1      string `orm:size(20)`
	Owner2      string `orm:size(20)`
	Domain_name string `orm:size(50)`
	Controller  string `orm:size(50)`
	Responsible string `orm:size(20)`
	Team        string `orm:size(20)`
	Company     string `orm:size(50)`
	Comment     string `orm:size(100)`
	Created     time.Time
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

func AddSystemlist(class, name, owner1, owner2, domain_name, controller, responsible, team, company, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	Systemlist := &System{
		Class:       class,
		Name:        name,
		Owner1:      owner1,
		Owner2:      owner2,
		Domain_name: domain_name,
		Controller:  controller,
		Responsible: responsible,
		Team:        team,
		Company:     company,
		Comment:     comment,
		Created:     time.Now(),
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

func ModifySystemlist(id, class, name, owner1, owner2, domain_name, controller, responsible, team, company, comment string) (error, string) {
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
		Systemlist.Name = name
		Systemlist.Owner1 = owner1
		Systemlist.Owner2 = owner2
		Systemlist.Domain_name = domain_name
		Systemlist.Controller = controller
		Systemlist.Comment = comment
		Systemlist.Responsible = responsible
		Systemlist.Team = team
		Systemlist.Company = company
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
	/*
		var cond *orm.Condition
		cond = orm.NewCondition()
		cond = cond.Or("name__icontains", name)
		//cond = cond.Or("ip__icontains", "ip")
		var qs orm.QuerySeter
		qs = o.QueryTable("system").Limit(pageSize, (currPage-1)*pageSize).SetCond(cond)
		_, err := qs.All(&Systemlists)
	*/
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

	rows, err := conn.Query("select class,name,owner1,owner2, domain_name,controller,responsible,team,company,comment from system")
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
