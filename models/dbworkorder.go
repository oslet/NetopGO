package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	//"strings"
	//"fmt"
	"time"
)

/**
1、除运维、研发、测试之外的提数，研发审批上传sql，运维执行，如果回滚，研发可编辑
2、研发、测试提运维直接执行，异常可回滚
3、运维自提直接走流程
*/
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
	Approver     string `orm:size(50)`
	Operater     string `orm:size(50)`
	DevOutcome   string `orm:size(1024)`
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

func AddDBOrder(schema, upgradeobj, upgradetype, comment, step, sqlfile, sponsor, dept string) error {
	o := orm.NewOrm()
	var status string
	dbwo := &Dbworkorder{}
	if dept == "运维" || dept == "测试" || dept == "研发" {
		status = "正在实施"
		dbwo = &Dbworkorder{
			Schemaname:   schema,
			Upgradeobj:   upgradeobj,
			Upgradetype:  upgradetype,
			Comment:      comment,
			Status:       status,
			Sponsor:      sponsor,
			Step:         step,
			Sqlfile:      sqlfile,
			Isedit:       "false",
			Isapp:        "false",
			RequestCount: 1,
			Created:      time.Now().String()[:18],
		}
	} else {
		status = "研发审批"
		dbwo = &Dbworkorder{
			Schemaname:   schema,
			Upgradeobj:   upgradeobj,
			Upgradetype:  upgradetype,
			Comment:      comment,
			Status:       status,
			Sponsor:      sponsor,
			Isedit:       "false",
			Isapp:        "false",
			RequestCount: 1,
			Created:      time.Now().String()[:18],
		}
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
	} else if "研发" == dept {
		total, err = o.Raw("select db.* from dbworkorder db join user on db.sponsor=user.name where sponsor=? or user.dept not in(?,?,?)", "dev", "研发", "测试", "运维").QueryRows(&dbwo)
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
	} else if "研发" == dept {
		// total, err = o.QueryTable("dbworkorder").Filter("sponsor", sponsor).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbwo)
		total, err = o.Raw("select db.* from dbworkorder db join user on db.sponsor=user.name where sponsor=? or user.dept not in(?,?,?) limit ?,?", "dev", "研发", "测试", "运维", (currPage-1)*pageSize, pageSize).QueryRows(&dbwo)
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
			isEdit = "true"
		} else if status == "正在实施" {
			flag = "true"
			isEdit = "false"
		} else {
			flag = "false"
			isEdit = "false"
		}
	} else if dept == "研发" {
		if status == "实施完毕" {
			flag = "false"
			isEdit = "false"
		} else if status == "异常回滚" {
			flag = "false"
			isEdit = "true"
		} else if status == "正在实施" {
			flag = "false"
			isEdit = "false"
		} else if status == "研发审批" {
			flag = "true"
			isEdit = "false"
		} else {
			flag = "false"
			isEdit = "false"
		}

	} else if dept == "测试" {
		if status == "实施完毕" {
			flag = "false"
			isEdit = "false"
		} else if status == "异常回滚" {
			flag = "false"
			isEdit = "true"
		} else if status == "正在实施" {
			flag = "false"
			isEdit = "false"
		} else if status == "研发审批" {
			flag = "false"
			isEdit = "false"
		} else {
			flag = "false"
			isEdit = "false"
		}

	} else {
		flag = "false"
		isEdit = "false"
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

func IsViewDBApprove(sponsorDept, status string) (devOutcome, opOutcome, devRead, opRead string) {
	var dev, op string
	var devReadonly, opReadonly string
	if sponsorDept == "测试" && status == "正在实施" {
		dev = "false"
		op = "true"
		devReadonly = "false"
		opReadonly = "true"
	} else if sponsorDept == "运维" && status == "正在实施" {
		dev = "false"
		op = "true"
		devReadonly = "false"
		opReadonly = "true"
	} else if sponsorDept == "研发" && status == "正在实施" {
		dev = "false"
		op = "true"
		devReadonly = "false"
		opReadonly = "true"
	} else {
		dev = "true"
		op = "true"
		devReadonly = "false"
		opReadonly = "true"
	}
	return dev, op, devReadonly, opReadonly
}

func DevCommit(id, nextStatus, step, devOutcome, sqlfile, uname string) error {
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	dbwo := &Dbworkorder{
		Id: did,
	}

	err = o.Read(dbwo)
	if err == nil {
		dbwo.Status = nextStatus
		dbwo.DevOutcome = devOutcome
		dbwo.Approver = uname
		dbwo.Step = step
		dbwo.Sqlfile = sqlfile
	}
	o.Update(dbwo)
	return err
}

func DBApproveModify(id, schema, upgradeobj, upgradetype, comment, new_sqlfile, step, devoutcome, dept string) error {
	var status string
	o := orm.NewOrm()
	did, err := strconv.ParseInt(id, 10, 64)
	dbwo := &Dbworkorder{
		Id: did,
	}
	err = o.Read(dbwo)
	// approveDept := GetUserDeptByApprover(dbwo.Approver)
	// if approveDept == "" && dbwo.Status == "异常回滚" {
	// 	status = "实施流程中"
	// } else {
	// 	status = "测试流程中"
	// }
	status = "正在实施"
	if err == nil {
		dbwo.Schemaname = schema
		dbwo.Upgradeobj = upgradeobj
		dbwo.Upgradetype = upgradetype
		dbwo.Comment = comment
		dbwo.Sqlfile = new_sqlfile
		dbwo.Step = step
		dbwo.DevOutcome = devoutcome
		dbwo.RequestCount = dbwo.RequestCount + 1
		dbwo.Operater = ""
		dbwo.OpOutcome = ""
		dbwo.Isedit = "false"
		dbwo.Status = status
	}
	o.Update(dbwo)
	return err
}

func SearchCount(schema string) (int64, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	total, err := o.QueryTable("dbrecord").Filter("schema__icontains", schema).All(&dbRecs)
	return total, err
}

func Search(currPage, pageSize int, schema string) ([]*Dbrecord, error) {
	o := orm.NewOrm()
	dbRecs := make([]*Dbrecord, 0)
	_, err := o.QueryTable("dbrecord").Filter("schema__icontains", schema).OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&dbRecs)
	return dbRecs, err
}
