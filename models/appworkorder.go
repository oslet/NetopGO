package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

/*
工单状态
1-测试流程中
2-审批流程中
3-实施流程中（不可修改）
4-验证流程中
5-工单已关闭
6-异常已回滚（回滚后工单状态可编辑）
附带DB状态
1-无DB变更
2-研发审批
3-正在实施
4-实施完毕
5-异常回滚
*/
type Appworkorder struct {
	Id           int64
	Appname      string `orm:size(30)`
	Version      string `orm:size(10)`
	Apptype      string `orm:size(20)`
	Upgradetype  string `orm:size(20)`
	FeatureList  string `orm:size(2048)`
	ModifyCfg    string `orm:size(2048)`
	RelayApp     string `orm:size(255)`
	Sqlfile      string `orm:size(100)`
	Attachment   string `orm:size(100)`
	JenkinsName  string `orm:size(100)`
	BuildNum     string `orm:size(10)`
	Sponsor      string `orm:size(50)`
	Tester       string `orm:size(50)`
	Approver     string `orm:size(50)`
	Operater     string `orm:size(50)`
	Finalchker   string `orm:size(50)`
	TestOutcome  string `orm:size(1024)`
	PrdtOutcome  string `orm:size(1024)`
	OpOutcome    string `orm:size(1024)`
	FinalOutcome string `orm:size(1024)`
	Status       string `orm:size(50)`
	DbStatus     string `orm:size(50)`
	Isapproved   string `orm:size(50)`
	Isedit       string `orm:size(5)`
	RequestCount int64
	Created      string `orm:size(20)`
}

func init() {
	orm.RegisterModel(new(Appworkorder))
}

func AddAppOrder(apptype, appname, version, jenkinsname, buildnum, featurelist, modifycfg, relayapp, upgradetype, sponsor, attachment, sqlfile, currDept string) error {
	o := orm.NewOrm()
	var dbstatus string
	var status string
	if len(strings.TrimSpace(sqlfile)) > 0 {
		dbstatus = "正在实施"
	} else {
		dbstatus = "无DB变更"
	}
	if currDept == "运维" {
		status = "实施流程中"
	} else {
		status = "测试流程中"
	}
	appwo := &Appworkorder{
		Appname:      appname,
		Version:      version,
		Apptype:      apptype,
		Upgradetype:  upgradetype,
		FeatureList:  featurelist,
		ModifyCfg:    modifycfg,
		RelayApp:     relayapp,
		Sqlfile:      sqlfile,
		Attachment:   attachment,
		JenkinsName:  jenkinsname,
		BuildNum:     buildnum,
		Sponsor:      sponsor,
		Status:       status,
		Isedit:       "false",
		DbStatus:     dbstatus,
		RequestCount: 1,
		Created:      time.Now().String()[:18],
	}
	_, err := o.Insert(appwo)
	return err
}

func GetAppOrderCount(auth int64) (int64, error) {
	var total int64
	var err error
	o := orm.NewOrm()
	appwo := make([]*Appworkorder, 0)
	if auth == 2 {
		total, err = o.QueryTable("appworkorder").Filter("status", "实施流程中").Filter("db_status", "正在实施").All(&appwo)
		if err != nil {
			return 0, err
		}
	} else {
		total, err = o.QueryTable("appworkorder").All(&appwo)
		if err != nil {
			return 0, err
		}
	}

	return total, err
}

func GetAppOrders(currPage, pageSize int, auth int64) ([]*Appworkorder, int64, error) {
	var total int64
	var err error
	o := orm.NewOrm()
	appwo := make([]*Appworkorder, 0)
	if auth == 2 {
		total, err = o.QueryTable("appworkorder").Filter("status", "实施流程中").Filter("db_status", "正在实施").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appwo)
		if err != nil {
			return nil, 0, err
		}
	} else {
		total, err = o.QueryTable("appworkorder").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appwo)
		if err != nil {
			return nil, 0, err
		}
	}

	return appwo, total, err
}

func GetAppwoById(id string) (*Appworkorder, error) {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appwo := &Appworkorder{}
	err = o.QueryTable("appworkorder").Filter("id", aid).One(appwo)
	return appwo, err
}

func ApproveCommit(id, nextStatus, outcome, outcomevalue, who, uname string) error {
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appwo := &Appworkorder{
		Id: aid,
	}

	err = o.Read(appwo)
	if err == nil {
		appwo.Status = nextStatus
		if who == "tester" {
			appwo.Tester = uname
		} else if who == "approver" {
			appwo.Approver = uname
		} else if who == "finalchker" {
			appwo.Finalchker = uname
		} else if who == "operater" {
			appwo.Operater = uname
		}
		if outcome == "testoutcome" {
			appwo.TestOutcome = outcomevalue
		} else if outcome == "prdtoutcome" {
			appwo.PrdtOutcome = outcomevalue
		} else if outcome == "opoutcome" {
			appwo.OpOutcome = outcomevalue
		} else if outcome == "finaloutcome" {
			appwo.FinalOutcome = outcomevalue
		}
	}
	o.Update(appwo)
	return err
}

func ApproveRollback(id, nextStatus, outcome, outcomevalue, who, uname, status, upgradeType, dept string) error {
	o := orm.NewOrm()
	var isEdit string
	aid, err := strconv.ParseInt(id, 10, 64)
	appwo := &Appworkorder{
		Id: aid,
	}

	if dept == "测试" && upgradeType == "修复bug" && status == "测试流程中" {
		isEdit = "true"
	} else if dept == "测试" && upgradeType == "产品发布" && status == "测试流程中" {
		isEdit = "true"
	} else if dept == "运维" && upgradeType == "系统运维" && status == "实施流程中" {
		isEdit = "true"
	} else {
		isEdit = "false"
	}

	err = o.Read(appwo)
	if err == nil {
		appwo.Status = nextStatus
		if who == "tester" {
			appwo.Tester = uname
		} else if who == "approver" {
			appwo.Approver = uname
		} else if who == "finalchker" {
			appwo.Finalchker = uname
		} else if who == "operater" {
			appwo.Operater = uname
		}
		if outcome == "testoutcome" {
			appwo.TestOutcome = outcomevalue
		} else if outcome == "prdtoutcome" {
			appwo.PrdtOutcome = outcomevalue
		} else if outcome == "opoutcome" {
			appwo.OpOutcome = outcomevalue
		} else if outcome == "finaloutcome" {
			appwo.FinalOutcome = outcomevalue
		}
		appwo.Isedit = isEdit
	}
	o.Update(appwo)
	return err
}

func ApproveModify(id, apptype, appname, upgradetype, version, jenkinsname, buildnum, featurelist, modifycfg, relayapp, final_attachment, final_sqlfile, dept string) error {
	var status string
	o := orm.NewOrm()
	aid, err := strconv.ParseInt(id, 10, 64)
	appwo := &Appworkorder{
		Id: aid,
	}
	err = o.Read(appwo)
	if upgradetype == "系统运维" && dept == "运维" && appwo.Status == "异常已回滚" {
		status = "实施流程中"
	} else {
		status = "测试流程中"
	}
	if err == nil {
		appwo.Apptype = apptype
		appwo.Appname = appname
		appwo.Upgradetype = upgradetype
		appwo.Version = version
		appwo.JenkinsName = jenkinsname
		appwo.BuildNum = buildnum
		appwo.FeatureList = featurelist
		appwo.ModifyCfg = modifycfg
		appwo.RelayApp = relayapp
		appwo.Attachment = final_attachment
		appwo.Sqlfile = final_sqlfile
		appwo.Status = status
		appwo.RequestCount = appwo.RequestCount + 1
		appwo.TestOutcome = ""
		appwo.PrdtOutcome = ""
		appwo.OpOutcome = ""
		appwo.FinalOutcome = ""
		appwo.Tester = ""
		appwo.Approver = ""
		appwo.Operater = ""
		appwo.Finalchker = ""
		appwo.Isedit = "false"
	}
	o.Update(appwo)
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

func NextStatus(cate, dept, status, upgradeType string) (string, string, string) {
	var nextStatus string
	var who string
	var outcome string
	if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "测试流程中" {
		nextStatus = "实施流程中"
		who = "tester"
		outcome = "testoutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "审批流程中" {
		nextStatus = "实施流程中"
		who = "approver"
		outcome = "testoutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "测试流程中" {
		nextStatus = "审批流程中"
		who = "tester"
		outcome = "testoutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "验证流程中" {
		nextStatus = "工单已关闭"
		who = "finalchker"
		outcome = "finaloutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "验证流程中" {
		nextStatus = "工单已关闭"
		who = "finalchker"
		outcome = "finaloutcome"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "审批流程中" {
		nextStatus = "实施流程中"
		who = "approver"
		outcome = "prdtoutcome"
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "实施流程中" {
		nextStatus = "验证流程中"
		who = "operater"
		outcome = "opoutcome"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "实施流程中" {
		nextStatus = "验证流程中"
		who = "operater"
		outcome = "opoutcome"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "实施流程中" {
		nextStatus = "工单已关闭"
		who = "operater"
		outcome = "opoutcome"
	}
	return nextStatus, who, outcome
}

func LastStatus(cate, dept, status, upgradeType string) (string, string, string) {
	var lastStatus string
	var who string
	var outcome string
	if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "测试流程中" {
		lastStatus = "异常已回滚"
		who = "tester"
		outcome = "testoutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "测试流程中" {
		lastStatus = "异常已回滚"
		who = "tester"
		outcome = "testoutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "验证流程中" {
		lastStatus = "实施流程中"
		who = "finalchker"
		outcome = "finaloutcome"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "验证流程中" {
		lastStatus = "实施流程中"
		who = "finalchker"
		outcome = "finaloutcome"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "审批流程中" {
		lastStatus = "测试流程中"
		who = "approver"
		outcome = "prdtoutcome"
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "实施流程中" {
		lastStatus = "测试流程中"
		who = "operater"
		outcome = "opoutcome"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "实施流程中" {
		lastStatus = "审批流程中"
		who = "operater"
		outcome = "opoutcome"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "实施流程中" {
		lastStatus = "异常已回滚"
		who = "operater"
		outcome = "opoutcome"
	}
	return lastStatus, who, outcome
}

func IsApproved(cate, dept, status, upgradeType, dbStatus string) string {
	fmt.Printf("*****this :%v", dbStatus)
	var flag string
	if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "测试流程中" {
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "测试流程中" {
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "系统运维" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "审批流程中" {
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "系统运维" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "系统运维" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "验证流程中" {
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "验证流程中" {
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "系统运维" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "系统运维" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "异常已回滚" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "false"
	} else if cate == "app" && dept == "测试" && upgradeType == "系统运维" && status == "异常已回滚" {
		flag = "false"
	}

	if cate == "app" && dept == "研发" && upgradeType == "修复bug" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "产品发布" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "系统运维" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "修复bug" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "产品发布" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "系统运维" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "修复bug" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "产品发布" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "系统运维" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "修复bug" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "产品发布" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "系统运维" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "修复bug" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "产品发布" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "系统运维" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "研发" && upgradeType == "修复bug" && status == "异常已回滚" {
		flag = "true"
	} else if cate == "app" && dept == "研发" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "true"
	} else if cate == "app" && dept == "研发" && upgradeType == "系统运维" && status == "异常已回滚" {
		flag = "false"
	}

	if cate == "app" && dept == "产品" && upgradeType == "修复bug" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "系统运维" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "修复bug" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "审批流程中" {
		flag = "true"
	} else if cate == "app" && dept == "产品" && upgradeType == "系统运维" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "修复bug" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "系统运维" && status == "实施流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "修复bug" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "系统运维" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "修复bug" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "系统运维" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "修复bug" && status == "异常已回滚" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "false"
	} else if cate == "app" && dept == "产品" && upgradeType == "系统运维" && status == "异常已回滚" {
		flag = "false"
	}

	if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "测试流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "审批流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "审批流程中" {
		flag = "true"
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "实施流程中" {
		if dbStatus == "正在实施" {
			flag = "false"
		} else {
			flag = "true"
		}
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "实施流程中" {
		if dbStatus == "正在实施" {
			flag = "false"
		} else {
			flag = "true"
		}
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "实施流程中" {
		if dbStatus == "正在实施" {
			flag = "false"
		} else {
			flag = "true"
		}
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "验证流程中" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "工单已关闭" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "修复bug" && status == "异常已回滚" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "false"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "异常已回滚" {
		flag = "false"
	}
	return flag
}

func IsViewDiv(dept, status, upgradeType string) (testOutcome, prdtOutcome, opOutcome, finalOutcome, testRead, productRead, opRead, finalRead string) {
	var test, product, op, final string
	var testReadonly, productReadonly, opReadonly, finalReadonly string
	if dept == "测试" && status == "测试流程中" && upgradeType == "修复bug" {
		test = "true"
		product = "false"
		op = "false"
		final = "false"
		testReadonly = "false"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "测试" && status == "测试流程中" && upgradeType == "产品发布" {
		test = "true"
		product = "false"
		op = "false"
		final = "false"
		testReadonly = "false"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "测试" && status == "验证流程中" && upgradeType == "修复bug" {
		test = "true"
		product = "false"
		op = "true"
		final = "true"
		testReadonly = "true"
		productReadonly = "false"
		opReadonly = "true"
		finalReadonly = "false"
	} else if dept == "测试" && status == "验证流程中" && upgradeType == "产品发布" {
		test = "true"
		product = "true"
		op = "true"
		final = "true"
		testReadonly = "true"
		productReadonly = "true"
		opReadonly = "true"
		finalReadonly = "false"
	} else if dept == "测试" && status == "审批流程中" && upgradeType == "修复bug" {
		test = "true"
		product = "false"
		op = "false"
		final = "false"
		testReadonly = "false"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "测试" && status == "审批流程中" && upgradeType == "产品发布" {
		test = "true"
		product = "false"
		op = "false"
		final = "false"
		testReadonly = "true"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "产品" && status == "审批流程中" && upgradeType == "产品发布" {
		test = "true"
		product = "true"
		op = "false"
		final = "false"
		testReadonly = "true"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "运维" && status == "实施流程中" && upgradeType == "修复bug" {
		test = "true"
		product = "false"
		op = "true"
		final = "false"
		testReadonly = "true"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "运维" && status == "实施流程中" && upgradeType == "产品发布" {
		test = "true"
		product = "true"
		op = "true"
		final = "false"
		testReadonly = "true"
		productReadonly = "true"
		opReadonly = "false"
		finalReadonly = "false"
	} else if dept == "运维" && status == "实施流程中" && upgradeType == "系统运维" {
		test = "false"
		product = "false"
		op = "true"
		final = "false"
		testReadonly = "false"
		productReadonly = "false"
		opReadonly = "false"
		finalReadonly = "false"
	}

	return test, product, op, final, testReadonly, productReadonly, opReadonly, finalReadonly
}
