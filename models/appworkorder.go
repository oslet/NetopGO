package models

import (
	"github.com/astaxie/beego/orm"
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
	Id          int64
	Appname     string `orm:size(30)`
	Verion      string `orm:size(10)`
	Apptype     string `orm:size(20)`
	Upgradetype string `orm:size(20)`
	FeatureList string `orm:size(2048)`
	ModifyCfg   string `orm:size(2048)`
	RelayApp    string `orm:size(255)`
	Sqlfile     string `orm:size(100)`
	Attachment  string `orm:size(100)`
	JenkinsName string `orm:size(100)`
	BuildNum    string `orm:size(10)`
	Sponsor     string `orm:size(50)`
	Tester      string `orm:size(50)`
	Approver    string `orm:size(50)`
	Operater    string `orm:size(50)`
	Status      string `orm:size(50)`
	DbStatus    string `orm:size(50)`
	Isedit      string `orm:size(5)`
	Created     string `orm:size(20)`
}

func init() {
	orm.RegisterModel(new(Appworkorder))
}

func AddAppOrder(apptype, appname, version, jenkinsname, buildnum, featurelist, modifycfg, relayapp, upgradetype, sponsor, attachment, sqlfile string) error {
	o := orm.NewOrm()
	var dbstatus string
	if len(strings.TrimSpace(sqlfile)) > 0 {
		dbstatus = "正在实施"
	} else {
		dbstatus = "无DB变更"
	}
	appwo := &Appworkorder{
		Appname:     appname,
		Verion:      version,
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
		Status:      "测试流程中",
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
