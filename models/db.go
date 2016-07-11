package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

type Db struct {
	Id         int64
	Name       string `orm:size(50)`
	Uuid       string `orm:size(50)`
	Comment    string `orm:size(100)`
	Size       string `orm:size(10)`
	Role       string `orm:size(20)`
	User       string `orm:size(50)`
	Passwd     string `orm:size(50)`
	Port       string `orm:size(10)`
	Schemaname string `orm:size(20)`
	Created    time.Time
}

type SlowLog struct {
	Uuid      string
	Name      string
	Timestamp string
	AvgTime   float64
	Count     int64
	SqlText   string
}

type Explain struct {
	Id           int64
	SelectType   string
	Table        string
	Type         string
	PossibleKeys string
	Key          string
	KeyLen       int64
	Ref          string
	Rows         int64
	Extra        string
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

func GetDBByName(name string) (*Db, error) {
	o := orm.NewOrm()
	db := &Db{}
	err := o.QueryTable("db").Filter("name", name).One(db)
	return db, err
}

func AddDB(name, uuid, comment, size, role, user, password, port, schema string) error {
	o := orm.NewOrm()
	passwd, _ := AESEncode(password, AesKey)
	//fmt.Printf("***add passwd:%v\n", passwd)
	db := &Db{
		Name:       name,
		Uuid:       uuid,
		Comment:    comment,
		Created:    time.Now(),
		Size:       size,
		Role:       role,
		User:       user,
		Passwd:     passwd,
		Port:       port,
		Schemaname: schema,
	}
	err := o.QueryTable("db").Filter("name", name).One(db)
	if err == nil {
		return nil
	}
	_, err = o.Insert(db)
	return err
}

func ModifyDB(id, name, uuid, comment, size, role, user, password, port, schema string) error {
	o := orm.NewOrm()

	did, err := strconv.ParseInt(id, 10, 64)
	passwd, _ := AESEncode(password, AesKey)
	//fmt.Printf("***modify passwd:%v\n", passwd)
	db := &Db{
		Id: did,
	}
	err = o.Read(db)
	if err == nil {
		db.Name = name
		db.Uuid = uuid
		db.Comment = comment
		db.Size = size
		db.Role = role
		db.User = user
		db.Passwd = passwd
		db.Port = port
		db.Schemaname = schema
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

func GetSizeView(name string) ([]string, []float64, []int64, error) {
	o := orm.NewOrm()
	var time []string
	var currSize []float64
	var totalSizes []int64
	var totalSize int64
	o.Raw("select size from db where name=?", name).QueryRow(&totalSize)
	nums, err := o.Raw("select date_format(timestamp,'%Y-%m-%d') from inst_info where name=?", name).QueryRows(&time)
	o.Raw("select size from inst_info where name=?", name).QueryRows(&currSize)
	for i := 0; i < int(nums); i++ {
		totalSizes = append(totalSizes, totalSize)
	}
	return time, currSize, totalSizes, err
}

func GetSlowView(name string) ([]string, []int64, error) {
	o := orm.NewOrm()
	var time []string
	var count []int64
	_, err := o.Raw("select date_format(timestamp,'%Y-%m-%d') from slow_overview where name=?", name).QueryRows(&time)
	o.Raw("select count from slow_overview where name=?", name).QueryRows(&count)
	return time, count, err
}

func GetQpsView(name string) ([]string, []float64, []float64, error) {
	o := orm.NewOrm()
	var time []string
	var qps []float64
	var tps []float64
	_, err := o.Raw("select timestamp from qps_tps_overview where name=?", name).QueryRows(&time)
	o.Raw("select qps from qps_tps_overview where name=?", name).QueryRows(&qps)
	o.Raw("select tps from qps_tps_overview where name=?", name).QueryRows(&tps)
	return time, qps, tps, err
}

func GetSlowLogs(currPage, pageSize int, name string) ([]*SlowLog, error) {
	o := orm.NewOrm()
	slowLogs := make([]*SlowLog, 0)
	_, err := o.Raw("select uuid,name,timestamp,round(sum(query_time)/count(1),2) as avg_time,count(1) as count,sql_text from sql_info where name=? group by name,timestamp,uuid order by timestamp desc, avg_time desc,count desc limit ?,?;", name, (currPage-1)*pageSize, pageSize).QueryRows(&slowLogs)

	return slowLogs, err
}

func GetSlowCount(name string) (int64, error) {
	o := orm.NewOrm()
	var num []int64
	total, err := o.Raw("select count(1) from sql_info where name=? group by name,timestamp,uuid ;", name).QueryRows(&num)

	return total, err
}

func SqlExplain(name, sqltext string) ([]*Explain, int64, error) {
	result := make([]*Explain, 0)
	o := orm.NewOrm()
	db := &Db{}
	o.QueryTable("db").Filter("name", name).One(db)

	passwd, _ := AESDecode(db.Passwd, AesKey)
	schemaUrl := db.User + ":" + passwd + "@tcp(" + db.Name + ":" + db.Port + ")/" + db.Schemaname + "?charset=utf8"
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	sql_db, err := orm.GetDB(db.Name)
	beego.Info(schemaUrl)
	if sql_db == nil {
		orm.RegisterDataBase(db.Name, _DB_Driver, schemaUrl, 30)
		fmt.Println("=====> register database of slow explain")
	}

	sqlTrim := strings.Trim(sqltext, " ")
	explainSql := "explain " + sqlTrim
	fmt.Printf("**** explain sql:%v\n", explainSql)

	o.Using(db.Name)
	num, err := o.Raw(explainSql).QueryRows(&result)

	return result, num, err
}

func Query2Export(alias, sqltext string) (*map[int64][]string, []string, int64) {
	//fmt.Println("call queryserver function!")
	result := make(map[int64][]string)
	var columns []string
	var total int64
	o := orm.NewOrm()
	schema := &Schema{}
	o.QueryTable("schema").Filter("name", alias).All(schema)

	passwd, _ := AESDecode(schema.Passwd, AesKey)
	schemaUrl := schema.User + ":" + passwd + "@tcp(" + schema.Addr + ":" + schema.Port + ")/" + schema.DBName + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to %v server successfully !", schema.Name))

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}

	defer conn.Close()
	sqlTrim := strings.Trim(sqltext, " ")
	//beego.Info(sqlTrim)

	rows, err := conn.Query(sqlTrim)
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
