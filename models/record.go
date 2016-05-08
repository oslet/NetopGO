package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Dbrecord struct {
	Id         int64
	Schema     string `orm:size(50)`
	Object     string `orm:size(50)`
	Operation  string `orm:size(50)`
	Isbackup   string `orm:size(10)`
	Content    string `orm:size(2000)`
	Attachment string `orm:size(100)`
	Comment    string `orm:szie(100)`
	Operater   string `orm:size(20)`
	Created    string `orm:size(19)`
}

type Apprecord struct {
	Id        int64
	Group     string `orm:size(50)`
	Operation string `orm:size(50)`
	Appname   string `orm:size(100)`
	Disthost  string `orm:size(100)`
	Isauto    string `orm:size(20)`
	Applicant string `orm:size(50)`
	Operater  string `orm:size(20)`
	Content   string `orm:size(2000)`
	Created   string `orm:size(19)`
}

func init() {
	orm.RegisterModel(new(Dbrecord), new(Apprecord))
}

func GetDBRecordCount() (int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").All(&dbRecs)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetDBRecords(currPage, pageSize int) ([]*Dbrecord, int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbRecs)
	if err != nil {
		return nil, 0, err
	}
	return dbRecs, total, err
}

func AddDBRecord(schema, object, operation, backup, content, attachment, comment, operater string) error {
	o := orm.NewOrm()
	record := &Dbrecord{
		Schema:     schema,
		Object:     object,
		Operation:  operation,
		Isbackup:   backup,
		Content:    content,
		Attachment: attachment,
		Comment:    comment,
		Operater:   operater,
		Created:    time.Now().String()[:19],
	}
	_, err := o.Insert(record)
	return err
}

func DeleteDBRecord(id string) error {
	o := orm.NewOrm()
	rid, err := strconv.ParseInt(id, 10, 64)
	dbRec := &Dbrecord{
		Id: rid,
	}
	_, err = o.Delete(dbRec)
	if err != nil {
		return err
	}
	return nil
}

func DBRecordDetail(id string) (*Dbrecord, error) {
	o := orm.NewOrm()
	rid, err := strconv.ParseInt(id, 10, 64)
	dbRec := &Dbrecord{}
	err = o.QueryTable("dbrecord").Filter("id", rid).One(dbRec)
	return dbRec, err
}

func SearchDBRecCount(schema string) (int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").Filter("schema__icontains", schema).All(&dbRecs)
	return total, err
}

func SearchDBRecBySchema(currPage, pageSize int, schema string) ([]*Dbrecord, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	_, err := o.QueryTable("dbrecord").Filter("schema__icontains", schema).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbRecs)
	return dbRecs, err
}

func GetAppRecordCount() (int64, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	total, err := o.QueryTable("apprecord").All(&appRecs)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetAppRecords(currPage, pageSize int) ([]*Apprecord, int64, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	total, err := o.QueryTable("apprecord").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appRecs)
	if err != nil {
		return nil, 0, err
	}
	return appRecs, total, err
}

func AddAppRecord(group, operation, appname, disthost, isauto, applicant, content, operater string) error {
	o := orm.NewOrm()
	record := &Apprecord{
		Group:     group,
		Operation: operation,
		Appname:   appname,
		Disthost:  disthost,
		Isauto:    isauto,
		Applicant: applicant,
		Content:   content,
		Operater:  operater,
		Created:   time.Now().String()[:19],
	}
	_, err := o.Insert(record)
	return err
}

func DeleteAppRecord(id string) error {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appRec := &Apprecord{
		Id: aid,
	}
	_, err = o.Delete(appRec)
	if err != nil {
		return err
	}
	return nil
}

func AppRecordDetail(id string) (*Apprecord, error) {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appRec := &Apprecord{}
	err = o.QueryTable("apprecord").Filter("id", aid).One(appRec)
	return appRec, err
}

func SearchAppRecCount(appname string) (int64, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	total, err := o.QueryTable("apprecord").Filter("appname__icontains", appname).All(&appRecs)
	return total, err
}

func SearchAppRecByName(currPage, pageSize int, appname string) ([]*Apprecord, error) {
	o := orm.NewOrm()
	appRecs := make([]*Apprecord, 0)
	_, err := o.QueryTable("apprecord").Filter("appname__icontains", appname).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appRecs)
	return appRecs, err
}
