package models

import (
	//"fmt"
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
1-准备就绪
2-正在实施
3-无DB变更
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
	TestOutcome  string `orm:size(1024)`
	PrdtOutcome  string `orm:size(1024)`
	OpOutcome    string `orm:size(1024)`
	FinalOutcome string `orm:size(1024)`
	Status       string `orm:size(50)`
	DbStatus     string `orm:size(50)`
	Isapproved   string `orm:size(50)`
	Isedit       string `orm:size(5)`
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
		Appname:     appname,
		Version:     version,
		Apptype:     apptype,
		Upgradetype: upgradetype,
		FeatureList: featurelist,
		ModifyCfg:   modifycfg,
		RelayApp:    relayapp,
		Sqlfile:     sqlfile,
		Attachment:  attachment,
		JenkinsName: jenkinsname,
		BuildNum:    buildnum,
		Sponsor:     sponsor,
		Status:      status,
		Isedit:      "true",
		DbStatus:    dbstatus,
		Created:     time.Now().String()[:18],
	}
	_, err := o.Insert(appwo)
	return err
}

func GetAppOrderCount() (int64, error) {
	o := orm.NewOrm()
	appwo := make([]*Appworkorder, 0)
	total, err := o.QueryTable("appworkorder").All(&appwo)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetAppOrders(currPage, pageSize int) ([]*Appworkorder, int64, error) {
	o := orm.NewOrm()
	appwo := make([]*Appworkorder, 0)
	total, err := o.QueryTable("appworkorder").OrderBy("-created").Limit(pageSize, (currPage-1)*pageSize).All(&appwo)
	if err != nil {
		return nil, 0, err
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

func IsApproved(cate, dept, status, upgradeType string) string {
	var flag string
	if cate == "app" && dept == "测试" && upgradeType == "修复bug" && status == "测试流程中" {
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "测试流程中" {
		flag = "false"
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
		flag = "true"
	} else if cate == "app" && dept == "测试" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "true"
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
		flag = "true"
	} else if cate == "app" && dept == "产品" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "true"
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
		flag = "true"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "实施流程中" {
		flag = "true"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "实施流程中" {
		flag = "true"
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
		flag = "true"
	} else if cate == "app" && dept == "运维" && upgradeType == "产品发布" && status == "异常已回滚" {
		flag = "true"
	} else if cate == "app" && dept == "运维" && upgradeType == "系统运维" && status == "异常已回滚" {
		flag = "true"
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
		product = "true"
		op = "true"
		final = "true"
		testReadonly = "true"
		productReadonly = "true"
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
		product = "true"
		op = "true"
		final = "false"
		testReadonly = "true"
		productReadonly = "true"
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
		testReadonly = "flase"
		productReadonly = "flase"
		opReadonly = "false"
		finalReadonly = "false"
	}

	return test, product, op, final, testReadonly, productReadonly, opReadonly, finalReadonly
}
