package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	//"strings"
	"time"
)

type Dbworkorder struct {
	Id           int64
	Schemaname   string `orm:size(50)`
	Upgradeobj   string `orm:size(15)`
	Upgradetype  string `orm:size(50)`
	Step         string `orm:size(1024)`
	Comment      string `orm:size(50)`
	Sqlfile      string `orm:size(50)`
	Status       string `orm:size(30)`
	Sponsor      string `orm:size(50)`
	Operater     string `orm:size(50)`
	Finalchker   string `orm:size(50)`
	OpOutcome    string `orm:size(1024)`
	RequestCount int64
	Isedit       string `orm:size(5)`
	Isapp        string `orm:size(5)`
	Isapproved   string `orm:size(50)`
	Created      string `orm:size(20)`
}

func init() {
	orm.RegisterModel(new(Dbworkorder))
}

func AddDBOrder(schema, upgradeobj, upgradetype, comment, sqlfile, step, sponsor string) error {
	o := orm.NewOrm()

	dbwo := &Dbworkorder{
		Schemaname:   schema,
		Upgradeobj:   upgradeobj,
		Upgradetype:  upgradetype,
		Step:         step,
		Comment:      comment,
		Sqlfile:      sqlfile,
		Status:       "正在实施",
		Sponsor:      sponsor,
		Isedit:       "false",
		Isapp:        "flase",
		RequestCount: 1,
		Created:      time.Now().String()[:18],
	}
	_, err := o.Insert(dbwo)
	return err
}

func GetDBOrderCount(dept, sponsor string) (int64, error) {
	var total int64
	var err error
	o := orm.NewOrm()
	dbwo := make([]*Dbworkorder, 0)
	if "运维" == dept {
		total, err = o.QueryTable("dbworkorder").All(&dbwo)
		if err != nil {
			return 0, err
		}
	} else {
		total, err = o.QueryTable("dbworkorder").Filter("sponsor", sponsor).All(&dbwo)
		if err != nil {
			return 0, err
		}
	}

	return total, err
}

func GetDBOrders(currPage, pageSize int, dept, sponsor string) ([]*Dbworkorder, int64, error) {
	var total int64
	var err error
	o := orm.NewOrm()
	dbwo := make([]*Dbworkorder, 0)
	if "运维" == dept {
		total, err = o.QueryTable("dbworkorder").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbwo)
		if err != nil {
			return nil, 0, err
		}
	} else {
		total, err = o.QueryTable("dbworkorder").Filter("sponsor", sponsor).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbwo)
		if err != nil {
			return nil, 0, err
		}
	}
	return dbwo, total, err
}

func GetDBwoById(id string) (*Dbworkorder, error) {
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	dbwo := &Dbworkorder{}
	err = o.QueryTable("dbworkorder").Filter("id", did).One(dbwo)
	return dbwo, err
}

func DBInAppCommit(id, schema, upgradeobj, upgradetype, comment, sqlfile, step, sponsor string) error {
	aid, _ := strconv.ParseInt(id, 10, 64)
	o := orm.NewOrm()
	err := o.Begin()
	appwo := &Appworkorder{
		Id: aid,
	}
	err = o.Read(appwo)
	if err == nil {
		appwo.DbStatus = "实施完毕"
	}
	_, err = o.Update(appwo)
	dbwo := &Dbworkorder{
		Schemaname:   schema,
		Upgradeobj:   upgradeobj,
		Upgradetype:  upgradetype,
		Step:         step,
		Comment:      comment,
		Sqlfile:      sqlfile,
		Status:       "实施完毕",
		Sponsor:      sponsor,
		Isedit:       "false",
		Isapp:        "true",
		RequestCount: 1,
		Created:      time.Now().String()[:18],
	}
	_, err = o.Insert(dbwo)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

func IsDBApproved(dept, status string) (string, string) {
	var flag string
	var isEdit string
	if dept == "运维" {
		if status == "实施完毕" {
			flag = "false"
			isEdit = "false"
		} else if status == "异常回滚" {
			flag = "false"
			isEdit = "false"
		} else if status == "正在实施" {
			flag = "true"
			isEdit = "false"
		}
	} else {
		if status == "实施完毕" {
			flag = "false"
			isEdit = "false"
		} else if status == "异常回滚" {
			flag = "true"
			isEdit = "true"
		} else if status == "正在实施" {
			flag = "true"
			isEdit = "false"
		}
	}
	return flag, isEdit
}

func GetSchemaNamesArray() ([]string, error) {
	o := orm.NewOrm()
	var schemaNames []string
	_, err := o.Raw("select name from  `schema`").QueryRows(&schemaNames)
	return schemaNames, err
}

func DBCommit(id, nextStatus, opoutcome, uname string) error {
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	dbwo := &Dbworkorder{
		Id: did,
	}

	err = o.Read(dbwo)
	if err == nil {
		dbwo.Status = nextStatus
		dbwo.OpOutcome = opoutcome
		dbwo.Operater = uname
	}
	o.Update(dbwo)
	return err
}

func DBRollback(id, lastStatus, opoutcome, uname string) error {
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	dbwo := &Dbworkorder{
		Id: did,
	}

	err = o.Read(dbwo)
	if err == nil {
		dbwo.Status = lastStatus
		dbwo.OpOutcome = opoutcome
		dbwo.Operater = uname
	}
	o.Update(dbwo)
	return err
}
