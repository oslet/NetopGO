package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Question struct {
	Id            int64
	InterNum      string `orm:size(30)`
	Name          string `orm:size(100)`
	InfluceBusine string `orm:size(20)`
	Owner         string `orm:size(30)`
	FaultCount    int64
	Status        string `orm:size(20)`
	Comment       string `orm:size(512)`
	Created       string `orm:size(20)`
}

func init() {
	orm.RegisterModel(new(Question))
}

func GetQuestRecordCount() (int64, error) {
	o := orm.NewOrm()
	questRecs := make([]*Question, 0)
	total, err := o.QueryTable("question").All(&questRecs)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetQuestRecords(currPage, pageSize int) ([]*Question, int64, error) {
	o := orm.NewOrm()
	questRecs := make([]*Question, 0)
	total, err := o.QueryTable("question").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&questRecs)
	if err != nil {
		return nil, 0, err
	}
	return questRecs, total, err
}

func AddQuestRecord(name, influce, owner, status, comment string) error {
	o := orm.NewOrm()
	now := time.Now().String()
	interNum := "IP" + now[:4] + now[5:7] + now[8:10] + now[11:13] + now[14:16] + now[17:19]

	record := &Question{
		InterNum:      interNum,
		Name:          name,
		InfluceBusine: influce,
		Owner:         owner,
		FaultCount:    0,
		Status:        status,
		Comment:       comment,
		Created:       time.Now().String()[:19],
	}
	err := o.QueryTable("question").Filter("name", name).One(record)
	if err == nil {
		return nil
	}
	_, err = o.Insert(record)
	return err
}

func ModifyQuestion(id, name, influce, owner, status, comment string) error {
	o := orm.NewOrm()
	qid, err := strconv.ParseInt(id, 10, 64)
	quest := &Question{
		Id: qid,
	}
	err = o.Read(quest)
	if err == nil {
		quest.Name = name
		quest.InfluceBusine = influce
		quest.Owner = owner
		quest.Comment = comment
	}
	o.Update(quest)
	return err
}

func GetQuestionById(id string) (*Question, error) {
	o := orm.NewOrm()
	qid, err := strconv.ParseInt(id, 10, 64)
	quest := &Question{}
	err = o.QueryTable("question").Filter("id", qid).One(quest)
	return quest, err
}

func DeleteQuestRecord(id string) error {
	o := orm.NewOrm()
	qid, err := strconv.ParseInt(id, 10, 64)
	quest := &Question{
		Id: qid,
	}
	_, err = o.Delete(quest)
	if err != nil {
		return err
	}
	return nil
}

func SearchQuestRecCount(apptype string) (int64, error) {
	o := orm.NewOrm()
	questRecs := make([]*Question, 0)
	total, err := o.QueryTable("question").Filter("influce_busine__icontains", apptype).All(&questRecs)
	return total, err
}

func SearchQuestRecByAppType(currPage, pageSize int, apptype string) ([]*Question, error) {
	o := orm.NewOrm()
	questRecs := make([]*Question, 0)
	_, err := o.QueryTable("question").Filter("influce_busine__icontains", apptype).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&questRecs)
	return questRecs, err
}

func QueryQuestionExport() (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}
	defer conn.Close()

	rows, err := conn.Query("select inter_num,name,influce_busine,owner,fault_count,status,comment from  question")
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

func GetQuestionNames() []string {
	o := orm.NewOrm()
	var name []string
	o.Raw("select name from  question").QueryRows(&name)
	return name
}
