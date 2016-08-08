package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RecycleHost struct {
	Id           int64
	Class        string `orm:size(50)`
	Service_name string `orm:size(50)`
	Name         string `orm:size(50)`
	Ip           string `orm:size(15)`
	Port         string `orm:size(15)`
	Os_type      string `orm:size(50)`
	Owner        string `orm:size(50)`
	Cpu          string `orm:size(50)`
	Mem          string `orm:size(50)`
	Disk         string `orm:size(50)`
	Idc          string `orm:size(50)`
	Group        string `orm:size(50)`
	Comment      string `orm:size(100)`
	Created      time.Time
}

// register recyclehost model
func init() {
	orm.RegisterModel(new(RecycleHost))
}

func GetRecycleHostCount() (int64, error) {
	o := orm.NewOrm()
	recyclehosts := make([]*RecycleHost, 0)
	total, err := o.QueryTable("recycle_host").All(&recyclehosts)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetRecycleHosts(currPage, pageSize int) ([]*RecycleHost, int64, error) {
	o := orm.NewOrm()
	recyclehosts := make([]*RecycleHost, 0)
	total, err := o.QueryTable("recycle_host").Limit(pageSize, (currPage-1)*pageSize).All(&recyclehosts)
	if err != nil {
		return nil, 0, err
	}
	return recyclehosts, total, err
}

func GetRecycleHostById(id string) (*RecycleHost, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	recyclehost := &RecycleHost{}
	err = o.QueryTable("recycle_host").Filter("id", hid).One(recyclehost)
	return recyclehost, err
}

func SearchRecycleHostWeekCount() (int64, error) {
	o := orm.NewOrm()
	//recyclehosts := make([]*RecycleHost, 0)
	var count int64
	//total, err := o.QueryTable("recyclehost").Filter("name__icontains", name).All(&recyclehosts)
	err := o.Raw("select count(*) as Count from recycle_host where date_sub(curdate(), INTERVAL ? DAY) <= date(`created`)", 7).QueryRow(&count)
	return count, err
}

func SearchRecycleHostByWeek(currPage, pageSize int, created string) ([]*RecycleHost, error) {
	o := orm.NewOrm()
	recyclehosts := make([]*RecycleHost, 0)
	/*
			var cond *orm.Condition
			cond = orm.NewCondition()
			cond = cond.Or("name__icontains", name)
			cond = cond.Or("ip__icontains", "ip")
			var qs orm.QuerySeter
			qs = o.QueryTable("recyclehost").Filter("idc", idc).Limit(pageSize, (currPage-1)*pageSize).SetCond(cond)
			_, err := qs.All(&recyclehosts)

		//_, err := o.QueryTable("recyclehost").Filter("created", created).Limit(pageSize, (currPage-1)*pageSize).All(&recyclehosts)
		//_, err := o.Raw("select * from recyclehost where date_sub(curdate(), INTERVAL 7 DAY) <= date(?)", created).QueryRows(&recyclehosts)
		_, err := orm.NewOrm().QueryTable("recyclehost").Filter("created").
	*/
	_, err := o.Raw("select * from recycle_host where date_sub(curdate(), INTERVAL 7 DAY) <= date(`created`)").QueryRows(&recyclehosts)
	return recyclehosts, err
}

func QueryRecycleHostWeekExport(method string) (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_recyclehost") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}

	defer conn.Close()
	if method == "week" {
		rows, err := conn.Query("select name,ip,service_name,cpu,mem,disk,idc,`group`,created from `recycle_host` where DATE_SUB(CURDATE(), INTERVAL 7 day) <= date(`created`)")
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
	} else if method == "all" {
		rows, err := conn.Query("select name,ip,service_name,cpu,mem,disk,idc,`group`,created from `recycle_host`")
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
	}

	return &result, columns, total
}
