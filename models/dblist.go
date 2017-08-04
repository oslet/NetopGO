package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Dblist struct {
	Id       int64
	IP       string `orm:"column(ip)"`
	Port     string `orm:"column(port)"`
	DBInst   string `orm:"column(dbinst)"`
	DBName   string `orm:"column(dbname)"`
	IsSwitch string `orm:"column(isswitch)"`
	AttrTeam string `orm:"column(attrteam)"`
	Name	 string `"orm:column(name)"`
	Created  time.Time
}

func init() {
	orm.RegisterModel(new(Dblist))
}

func GetDblistCount() (int64, error) {
	o := orm.NewOrm()
	dblists := make([]*Dblist, 0)
	total, err := o.QueryTable("dblist").All(&dblists)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetDblists(currPage, pageSize int) ([]*Dblist, int64, error) {
	o := orm.NewOrm()
	dblists := make([]*Dblist, 0)
	total, err := o.QueryTable("dblist").Limit(pageSize, (currPage-1)*pageSize).All(&dblists)
	if err != nil {
		return nil, 0, err
	}
	return dblists, total, err
}

func GetDblistById(id string) (*Dblist, error) {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	dblist := &Dblist{}
	err = o.QueryTable("dblist").Filter("id", gid).One(dblist)
	return dblist, err
}

func AddDblist(ip, port, dbinst, dbname, isswitch, attrteam, name string) (error, string) {
	o := orm.NewOrm()
	var msg string
	dblist := &Dblist{
		IP:         ip,
		Port: port,
		DBInst:    dbinst,
		DBName:        dbname,
		IsSwitch:      isswitch,
		AttrTeam:   attrteam,
		Name:      name,
		Created:      time.Now(),
	}
	err := o.QueryTable("dblist").Filter("name", dbname).One(dblist)
	if err == nil {
		msg = "数据库名称" + dbname + "已存在"
		return nil, msg
	}
	_, err = o.Insert(dblist)
	msg = "添加数据库成功"
	fmt.Println("add Dblist : ", dblist)
	return err, msg
}

func ModifyDblist(id, ip, port, dbinst, dbname, isswitch, attrteam, name string) (error, string) {
	o := orm.NewOrm()
	var msg string
	gid, err := strconv.ParseInt(id, 10, 64)
	dblist := &Dblist{
		Id: gid,
	}
	err = o.Read(dblist)
	if err == nil {
		dblist.IP = ip
		dblist.Port = port
		dblist.DBInst = dbinst
		dblist.DBName = dbname
		dblist.IsSwitch = isswitch
		dblist.AttrTeam = attrteam
		dblist.Name = name
	}
	o.Update(dblist)
	msg = "修改成功"
	return err, msg
}

func DeleteDblist(id string) error {
	o := orm.NewOrm()
	gid, err := strconv.ParseInt(id, 10, 64)
	dblist := &Dblist{
		Id: gid,
	}
	_, err = o.Delete(dblist)
	if err != nil {
		return err
	}
	return nil
}

func SearchDblistCount(dbname string) (int64, error) {
	o := orm.NewOrm()
	dblists := make([]*Dblist, 0)
	total, err := o.QueryTable("dblist").Filter("name__icontains", dbname).All(&dblists)
	return total, err
}
/*
func SearchDblistByName(currPage, pageSize int, dbname string) ([]*Dblist, error) {
	o := orm.NewOrm()
	dblists := make([]*Dblist, 0)
	_, err := o.QueryTable("dblist").Filter("name__icontains", dbname).Limit(pageSize, (currPage-1)*pageSize).All(&dblists)
	return dblists, err
}
*/

func SearchDblistByName(currPage, pageSize int, name string) ([]*Dblist, error) {
	o := orm.NewOrm()
	dblists := make([]*Dblist, 0)
	ids := []string{"%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%"}
	if len(name) == 0 {
		_, err := o.QueryTable("dblist").Limit(pageSize, (currPage-1)*pageSize).All(&dblists)
		return dblists, err
	} else {
		_, err := o.Raw("select * from dblist where ip like ? or port like ? or dbinst like ? or dbname like ? or isswitch like ? or attrteam like ? or name like ? limit ?,?", ids, (currPage-1)*pageSize, pageSize).QueryRows(&dblists)
		return dblists, err
	}
}

func GetDblistNames() ([]*Dblist, error) {
	o := orm.NewOrm()
	dblists := make([]*Dblist, 0)
	_, err := o.QueryTable("dblist").All(&dblists)
	if err != nil {
		return nil, err
	}
	return dblists, err
}

func QueryDblistExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select ip, port, dbinst, dbname, isswitch, attrteam, name from dblist")
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
