package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Audit struct {
	Id       int64
	Schema   string `orm:size(50)`
	Operater string `orm:size(50)`
	Status   string `orm:size(50)`
	Sqltext  string `orm:size(1000)`
	Created  string `orm:size(19)`
}

func init() {
	orm.RegisterModel(new(Audit))
}

func GetAuditCount() (int64, error) {
	o := orm.NewOrm()
	audits := make([]*Audit, 0)
	nums, err := o.QueryTable("audit").All(&audits)
	return nums, err
}

func GetAudits(currPage, pageSize int) ([]*Audit, int64, error) {
	o := orm.NewOrm()
	audits := make([]*Audit, 0)
	total, err := o.QueryTable("audit").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&audits)
	if err != nil {
		return nil, 0, err
	}
	return audits, total, err
}

func DeleteAudit(id string) error {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	audit := &Audit{
		Id: aid,
	}
	_, err = o.Delete(audit)
	if err != nil {
		return err
	}
	return nil
}

func AuditDetail(id string) (*Audit, error) {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	audit := &Audit{}
	err = o.QueryTable("audit").Filter("id", aid).One(audit)
	return audit, err
}

func SearchAuditCount(schema string) (int64, error) {
	o := orm.NewOrm()
	audit := make([]*Audit, 0)
	total, err := o.QueryTable("audit").Filter("schema__icontains", schema).OrderBy("-created").All(&audit)
	return total, err
}

func SearchAuditBySchema(currPage, pageSize int, schema string) ([]*Audit, error) {
	o := orm.NewOrm()
	audits := make([]*Audit, 0)
	_, err := o.QueryTable("audit").Filter("schema__icontains", schema).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&audits)
	return audits, err
}

func WriteAuditLog(schema, operater, sqltext, status string) error {
	o := orm.NewOrm()
	audit := &Audit{
		Schema:   schema,
		Operater: operater,
		Status:   status,
		Sqltext:  sqltext,
		Created:  time.Now().String()[:19],
	}
	_, err := o.Insert(audit)
	return err
}
