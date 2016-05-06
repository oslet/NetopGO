package models

import (
	"github.com/astaxie/beego/orm"
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
	Created    time.Time
}

func init() {
	orm.RegisterModel(new(Dbrecord))
}

func AddDBRecord(schema, object, operation, backup, content, attachment, comment string) error {
	o := orm.NewOrm()
	record := &Dbrecord{
		Schema:     schema,
		Object:     object,
		Operation:  operation,
		Isbackup:   backup,
		Content:    content,
		Attachment: attachment,
		Comment:    comment,
		Created:    time.Now(),
	}
	_, err := o.Insert(record)
	return err
}
