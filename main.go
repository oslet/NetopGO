package main

import (
	"NetopGO/models"
	_ "NetopGO/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", true, true)
}

func main() {
	orm.Debug = true

	o := orm.NewOrm()
	// psw := "nbs2010"
	// encode := models.Base64Encode([]byte(psw))
	// beego.Info(string(encode))
	// decode, _ := models.Base64Decode(encode)
	// beego.Info(string(decode))
	passwd := string(models.Base64Encode([]byte("nbs2010")))
	admin := &models.User{
		Name:    "admin",
		Passwd:  passwd,
		Email:   "admin@tingyun.com",
		Dept:    "op",
		Created: time.Now(),
		Auth:    1,
		Tel:     "18202808939",
	}
	o.Insert(admin)
	dba := &models.User{
		Name:    "dba",
		Passwd:  passwd,
		Email:   "dba@tingyun.com",
		Dept:    "op",
		Created: time.Now(),
		Auth:    2,
		Tel:     "18202808939",
	}
	o.Insert(dba)
	guest := &models.User{
		Name:    "guest",
		Passwd:  passwd,
		Email:   "guest@tingyun.com",
		Dept:    "op",
		Created: time.Now(),
		Auth:    3,
		Tel:     "18202808939",
	}
	o.Insert(guest)

	host := &models.Host{
		Name:    "ucd-ty-ice-log-1",
		Ip:      "192.168.2.17",
		Cpu:     "4核",
		Mem:     "8GB",
		Disk:    "1TB",
		Idc:     "Ucloud",
		Rootpwd: passwd,
		Readpwd: passwd,
		Group:   "flume",
		Created: time.Now(),
	}
	o.Insert(host)

	group := &models.Group{
		Name:    "flume",
		Conment: "flume",
		Created: time.Now(),
	}
	o.Insert(group)
	// str := "nbs2010"
	// beego.Info(models.Md5Encode([]byte(str)))
	beego.Run()
}
