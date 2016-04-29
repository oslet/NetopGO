package main

import (
	"NetopGO/models"
	_ "NetopGO/routers"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	//"time"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true

	//o := orm.NewOrm()
	// psw := "nbs2010"
	// encode1 := models.Base64Encode([]byte(psw))
	// beego.Info(string(encode1))
	// // decode, _ := models.Base64Decode(encode)
	// // beego.Info(string(decode))
	// aesKey := beego.AppConfig.String("aesKey")
	// beego.Info(aesKey)
	// encode, _ := models.AESEncode(psw, aesKey)
	// beego.Info(encode)
	// decode, _ := models.AESDecode(encode, aesKey)
	// beego.Info(decode)

	// passwd := string(models.Base64Encode([]byte("nbs2010")))
	// admin := &models.User{
	// 	Name:    "admin",
	// 	Passwd:  passwd,
	// 	Email:   "admin@tingyun.com",
	// 	Dept:    "op",
	// 	Created: time.Now(),
	// 	Auth:    1,
	// 	Tel:     "18202808939",
	// }
	// o.Insert(admin)
	// dba := &models.User{
	// 	Name:    "dba",
	// 	Passwd:  passwd,
	// 	Email:   "dba@tingyun.com",
	// 	Dept:    "op",
	// 	Created: time.Now(),
	// 	Auth:    2,
	// 	Tel:     "18202808939",
	// }
	// o.Insert(dba)
	// guest := &models.User{
	// 	Name:    "guest",
	// 	Passwd:  passwd,
	// 	Email:   "guest@tingyun.com",
	// 	Dept:    "op",
	// 	Created: time.Now(),
	// 	Auth:    3,
	// 	Tel:     "18202808939",
	// }
	// o.Insert(guest)
	// encode, _ := models.AESEncode("upbjsxt", models.AesKey)
	// host := &models.Host{
	// 	Name:    "localhost",
	// 	Ip:      "127.0.0.1",
	// 	Cpu:     "4核",
	// 	Mem:     "8GB",
	// 	Disk:    "1TB",
	// 	Idc:     "Ucloud",
	// 	Root:    "quenlang",
	// 	Read:    "quenlang",
	// 	Rootpwd: encode,
	// 	Readpwd: encode,
	// 	Group:   "flume",
	// 	Created: time.Now(),
	// }
	// o.Insert(host)
	// schemaPwd, _ := models.AESEncode("upbjsxt", models.AesKey)
	// schema1 := &models.Schema{
	// 	Name:    "lens_conf",
	// 	Comment: "lens_conf",
	// 	Created: time.Now(),
	// 	DBName:  "lens_conf",
	// 	User:    "lens",
	// 	Passwd:  schemaPwd,
	// }
	// o.Insert(schema1)
	// schema2 := &models.Schema{
	// 	Name:    "lens_server_data",
	// 	Comment: "lens_server_data",
	// 	Created: time.Now(),
	// 	DBName:  "lens_server_data",
	// 	User:    "lens",
	// 	Passwd:  schemaPwd,
	// }
	// o.Insert(schema2)
	// group := &models.Group{
	// 	Name:    "flume",
	// 	Conment: "flume",
	// 	Created: time.Now(),
	// }
	// o.Insert(group)

	// db := &models.Db{
	// 	Name:    "dbmaster_conf",
	// 	Uuid:    "udb-sahjd",
	// 	Comment: "核心配置",
	// 	Created: time.Now(),
	// }
	// o.Insert(db)
	// str := "nbs2010"
	// beego.Info(models.Md5Encode([]byte(str)))
	beego.Run()
}
