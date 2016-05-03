package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Db struct {
	Id      int64
	Name    string `orm:size(50)`
	Uuid    string `orm:size(50)`
	Comment string `orm:size(100)`
	Size    int64
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Db))
}

func GetDBCount() (int64, error) {
	o := orm.NewOrm()
	dbs := make([]*Db, 0)
	total, err := o.QueryTable("db").All(&dbs)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetDBs(currPage, pageSize int) ([]*Db, int64, error) {
	o := orm.NewOrm()
	dbs := make([]*Db, 0)
	total, err := o.QueryTable("db").Limit(pageSize, (currPage-1)*pageSize).All(&dbs)
	if err != nil {
		return nil, 0, err
	}
	return dbs, total, err
}

func GetDBById(id string) (*Db, error) {
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	db := &Db{}
	err = o.QueryTable("db").Filter("id", did).One(db)
	return db, err
}

func AddDB(name, uuid, comment, size string) error {
	o := orm.NewOrm()
	sizeInt, _ := strconv.ParseInt(size, 10, 64)
	db := &Db{
		Name:    name,
		Uuid:    uuid,
		Comment: comment,
		Created: time.Now(),
		Size:    sizeInt,
	}
	err := o.QueryTable("db").Filter("name", name).One(db)
	if err == nil {
		return nil
	}
	_, err = o.Insert(db)
	return err
}

func ModifyDB(id, name, uuid, comment, size string) error {
	o := orm.NewOrm()

	did, err := strconv.ParseInt(id, 10, 64)
	sizeInt, _ := strconv.ParseInt(size, 10, 64)
	db := &Db{
		Id: did,
	}
	err = o.Read(db)
	if err == nil {
		db.Name = name
		db.Uuid = uuid
		db.Comment = comment
		db.Size = sizeInt
	}
	o.Update(db)
	return err
}

func DeleteDB(id string) error {
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	db := &Db{
		Id: did,
	}
	_, err = o.Delete(db)
	if err != nil {
		return err
	}
	return nil
}

func SearchDBCount(name string) (int64, error) {
	o := orm.NewOrm()
	dbs := make([]*Db, 0)
	total, err := o.QueryTable("db").Filter("name__icontains", name).All(&dbs)
	return total, err
}

func SearchDBByName(currPage, pageSize int, name string) ([]*Db, error) {
	o := orm.NewOrm()
	dbs := make([]*Db, 0)
	_, err := o.QueryTable("db").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&dbs)
	return dbs, err
}
