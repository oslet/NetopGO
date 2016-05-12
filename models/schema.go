package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Schema struct {
	Id        int64
	Addr      string `orm:size(50)`
	Port      string `orm:size(10)`
	Name      string `orm:size(50)`
	Comment   string `orm:size(50)`
	User      string `orm:size(50)`
	Passwd    string `orm:size(50)`
	DBName    string `orm:size(50)`
	Partition int64
	Status    int64
	Size      float64
	Created   time.Time
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

func GetSchemaByName(name string) (*Schema, error) {
	o := orm.NewOrm()
	schema := &Schema{}
	err := o.QueryTable("schema").Filter("name", name).One(schema)
	return schema, err
}

func AddSchema(name, dbname, partition, user, passwd, status, comment, addr, port string) error {
	o := orm.NewOrm()
	parNum, _ := strconv.ParseInt(partition, 10, 64)
	statusInt, _ := strconv.ParseInt(status, 10, 64)
	passwd, _ = AESEncode(passwd, AesKey)
	schema := &Schema{
		Name:      name,
		User:      user,
		DBName:    dbname,
		Passwd:    passwd,
		Comment:   comment,
		Addr:      addr,
		Port:      port,
		Partition: parNum,
		Status:    statusInt,
		Created:   time.Now(),
	}
	err := o.QueryTable("schema").Filter("name", name).One(schema)
	if err == nil {
		return nil
	}
	_, err = o.Insert(schema)
	return err
}

func ModifySchema(id, name, dbname, partition, user, passwd, status, comment, addr, port string) error {
	o := orm.NewOrm()
	parNum, _ := strconv.ParseInt(partition, 10, 64)
	statusInt, _ := strconv.ParseInt(status, 10, 64)
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
		schema.Partition = parNum
		schema.Status = statusInt
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

type Partition struct {
	Instance  string
	Timestamp string
	Count     int64
}

func GetPartDetail(flag, schema string) ([]*Partition, int64, error) {
	o := orm.NewOrm()
	parts := make([]*Partition, 0)

	var num int64
	var err error
	if "min" == flag {
		yesterday := time.Now().AddDate(0, 0, -1).String()[:10] + "23:59:59"
		num, err = o.Raw("SELECT instance, timestamp,count FROM `partition_info` WHERE `schemaname` = ? and `timestamp` > ? and `type` = ?", schema, yesterday, flag).QueryRows(&parts)
	} else if "hour" == flag {
		lastweek := time.Now().AddDate(0, 0, -7).String()[:10] + "23:59:59"
		num, err = o.Raw("SELECT instance, timestamp,count FROM `partition_info` WHERE `schemaname` = ? and `timestamp` > ? and `type` = ?", schema, lastweek, flag).QueryRows(&parts)
	} else {
		var lastquarter string
		year := time.Now().String()[:5]
		month := time.Now().String()[6:7]
		if month >= "1" && month < "4" {
			lastquarter = year + "01-01"
		} else if month >= "4" && month < "7" {
			lastquarter = year + "04-01"
		} else if month >= "7" && month < "10" {
			lastquarter = year + "07-01"
		} else if month >= "10" {
			lastquarter = year + "10-01"
		}
		num, err = o.Raw("SELECT instance, timestamp,count FROM `partition_info` WHERE `schemaname`  = ? and `timestamp` > ? and `type` = ?", schema, lastquarter, flag).QueryRows(&parts)
	}

	return parts, num, err
}
func GetSizeBySchema(schema string) (float64, error) {
	o := orm.NewOrm()
	var size float64
	today := time.Now().String()[:11] + "00:00:00"
	err := o.Raw("select sum(size) from inst_info where schemaname=? and timestamp=? and name like '%master%'", schema, today).QueryRow(&size)
	return size, err
}
