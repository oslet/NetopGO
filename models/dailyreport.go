package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Dailyreport struct {
	Id           int64
	Appsys       string `orm:size(50)`
	Appname      string `orm:size(50)`
	Appcontent   string `orm:size(1000)`
	Applicgrp    string `orm:size(50)`
	Applicant    string `orm:size(50)`
	Publisher    string `orm:size(50)`
	Department   string `orm:size(50)`
	Publishtime  string `orm:size(50)`
	Followstatus string `orm:size(50)`
	Followman    string `orm:size(50)`
	Isinitial    string `orm:size(50)`
	Created      time.Time
}

// register host model
func init() {
	orm.RegisterModel(new(Dailyreport))
}

func GetDailyreportCount() (int64, error) {
	o := orm.NewOrm()
	Dailyreportlist := make([]*Dailyreport, 0)
	total, err := o.QueryTable("dailyreport").All(&Dailyreportlist)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetDailyreports(currPage, pageSize int) ([]*Dailyreport, int64, error) {
	o := orm.NewOrm()
	Dailyreportlist := make([]*Dailyreport, 0)
	total, err := o.QueryTable("dailyreport").Limit(pageSize, (currPage-1)*pageSize).All(&Dailyreportlist)
	if err != nil {
		return nil, 0, err
	}
	return Dailyreportlist, total, err
}

func GetDailyreportlistById(id string) (*Dailyreport, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	Dailyreportlist := &Dailyreport{}
	err = o.QueryTable("dailyreport").Filter("id", hid).One(Dailyreportlist)
	return Dailyreportlist, err
}

func AddDailyreportlist(appsys, appname, appcontent, applicgrp, applicant, publisher, department, publishtime, followstatus, followman, isinitial string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	Dailyreportlist := &Dailyreport{
		Appsys:       appsys,
		Appname:      appname,
		Applicgrp:    applicgrp,
		Appcontent:   appcontent,
		Applicant:    applicant,
		Publisher:    publisher,
		Department:   department,
		Publishtime:  publishtime,
		Followstatus: followstatus,
		Followman:    followman,
		Isinitial:    isinitial,
		Created:      time.Now(),
	}
	err := o.QueryTable("dailyreport").Filter("appname", appname).Filter("publishtime", publishtime).One(Dailyreportlist)
	if err == nil {
		msg = "程序更新报表存在重复记录"
		return nil, msg
	}
	_, err = o.Insert(Dailyreportlist)
	msg = "添加程序更新报表成功"
	return err, msg
}

func ModifyDailyreportlist(id, appsys, appname, appcontent, applicgrp, applicant, publisher, department, publishtime, followstatus, followman, isinitial string) (error, string) {
	o := orm.NewOrm()
	var msg string
	//rootpwd, _ = AESEncode(rootpwd, AesKey)
	//readpwd, _ = AESEncode(readpwd, AesKey)
	hid, err := strconv.ParseInt(id, 10, 64)
	Dailyreportlist := &Dailyreport{
		Id: hid,
	}
	err = o.Read(Dailyreportlist)
	if err == nil {
		Dailyreportlist.Appsys = appsys
		Dailyreportlist.Appname = appname
		Dailyreportlist.Appcontent = appcontent
		Dailyreportlist.Applicgrp = applicgrp
		Dailyreportlist.Applicant = applicant
		Dailyreportlist.Publisher = publisher
		Dailyreportlist.Department = department
		Dailyreportlist.Publishtime = publishtime
		Dailyreportlist.Followstatus = followstatus
		Dailyreportlist.Followman = followman
		Dailyreportlist.Isinitial = isinitial
	}
	o.Update(Dailyreportlist)
	msg = "修改成功"
	return err, msg
}

func DeleteDailyreportlist(id string) error {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	Dailyreportlist := &Dailyreport{
		Id: hid,
	}
	_, err = o.Delete(Dailyreportlist)
	if err != nil {
		return err
	}
	return nil
}

/*
func SearchDailyreportlistCount(appname string) (int64, error) {
	o := orm.NewOrm()
	dailyreportlists := make([]*Dailyreport, 0)
	total, err := o.QueryTable("dailyreport").Filter("appname__icontains", appname).All(&dailyreportlists)
	return total, err
}
*/
func SearchDailyreportlistCount(name string) (int64, error) {
	o := orm.NewOrm()
	//dailyreportlists := make([]*Dailyreport, 0)
	var num int64
	var err error
	ids := []string{"%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%"}
	o.Raw("select count(*) from dailyreport where appsys like binary ? or appname like binary ? or appcontent like binary ? or applicant like binary ? or publisher like binary ? or publishtime like binary ? or followstatus like binary ? or followman like binary ? or isinitial like binary ?", ids).QueryRow(&num)
	return num, err
}

func SearchDailyreportlistByName(currPage, pageSize int, name string) ([]*Dailyreport, error) {
	o := orm.NewOrm()
	Dailyreportlists := make([]*Dailyreport, 0)
	ids := []string{"%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%"}
	if len(name) == 0 {
		_, err := o.QueryTable("dailyreport").Limit(pageSize, (currPage-1)*pageSize).All(&Dailyreportlists)
		return Dailyreportlists, err
	} else {
		_, err := o.Raw("select * from dailyreport where appsys like binary ? or appname like binary ? or appcontent like binary ? or applicant like binary ? or publisher like binary ? or publishtime like binary ? or followstatus like binary ? or followman like binary ? or isinitial like binary ? limit ?,?", ids, (currPage-1)*pageSize, pageSize).QueryRows(&Dailyreportlists)
		return Dailyreportlists, err
	}
}

func GetDailyreportById(id string) (*Dailyreport, error) {
	o := orm.NewOrm()
	sid, err := strconv.ParseInt(id, 10, 64)
	dailyreport := &Dailyreport{}
	err = o.QueryTable("dailyreport").Filter("id", sid).One(dailyreport)
	return dailyreport, err
}

func QueryDailyreportExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select appsys, appname, appcontent, applicgrp, applicant, publisher, department, publishtime, followstatus, followman, isinitial from dailyreport")
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
