package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Schema struct {
	Id      int64
	Addr    string `orm:size(50)`
	Port    string `orm:size(10)`
	Name    string `orm:size(50)`
	Comment string `orm:size(50)`
	User    string `orm:size(50)`
	Passwd  string `orm:size(50)`
	DBName  string `orm:size(50)`
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Schema))
}

func GetSchemaCount() (int64, error) {
	o := orm.NewOrm()
	schemas := make([]*Schema, 0)
	total, err := o.QueryTable("schema").All(&schemas)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetSchemas(currPage, pageSize int) ([]*Schema, int64, error) {
	o := orm.NewOrm()
	schemas := make([]*Schema, 0)
	total, err := o.QueryTable("schema").Limit(pageSize, (currPage-1)*pageSize).All(&schemas)
	if err != nil {
		return nil, 0, err
	}
	return schemas, total, err
}

func GetSchemaById(id string) (*Schema, error) {
	o := orm.NewOrm()
	sid, err := strconv.ParseInt(id, 10, 64)
	schema := &Schema{}
	err = o.QueryTable("schema").Filter("id", sid).One(schema)
	return schema, err
}

func AddSchema(name, dbname, user, passwd, comment, addr, port string) error {
	o := orm.NewOrm()
	passwd, _ = AESEncode(passwd, AesKey)
	schema := &Schema{
		Name:    name,
		User:    user,
		DBName:  dbname,
		Passwd:  passwd,
		Comment: comment,
		Addr:    addr,
		Port:    port,
		Created: time.Now(),
	}
	err := o.QueryTable("schema").Filter("name", name).One(schema)
	if err == nil {
		return nil
	}
	_, err = o.Insert(schema)
	return err
}

func ModifySchema(id, name, dbname, user, passwd, comment, addr, port string) error {
	o := orm.NewOrm()
	sid, err := strconv.ParseInt(id, 10, 64)
	passwd, _ = AESEncode(passwd, AesKey)
	schema := &Schema{
		Id: sid,
	}
	err = o.Read(schema)
	if err == nil {
		schema.Name = name
		schema.DBName = dbname
		schema.User = user
		schema.Passwd = passwd
		schema.Comment = comment
		schema.Addr = addr
		schema.Port = port
	}
	o.Update(schema)
	return err
}

func DeleteSchema(id string) error {
	o := orm.NewOrm()
	sid, err := strconv.ParseInt(id, 10, 64)
	schema := &Schema{
		Id: sid,
	}
	_, err = o.Delete(schema)
	if err != nil {
		return err
	}
	return nil
}

func SearchSchemaCount(name string) (int64, error) {
	o := orm.NewOrm()
	schemas := make([]*Schema, 0)
	total, err := o.QueryTable("schema").Filter("name__icontains", name).All(&schemas)
	return total, err
}

func SearchSchemaByName(currPage, pageSize int, name string) ([]*Schema, error) {
	o := orm.NewOrm()
	schemas := make([]*Schema, 0)
	_, err := o.QueryTable("schema").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&schemas)
	return schemas, err
}

func GetSchemaNames() ([]*Schema, error) {
	o := orm.NewOrm()
	schemas := make([]*Schema, 0)
	_, err := o.QueryTable("schema").All(&schemas)
	if err != nil {
		return nil, err
	}
	return schemas, err
}
