package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Host struct {
	Id           int64
	Class        string `orm:size(50)`
	Service_name string `orm:size(50)`
	Name         string `orm:size(50)`
	Ip           string `orm:size(15)`
	Pubip        string `orm:size(15)`
	Port         string `orm:size(15)`
	Os_type      string `orm:size(50)`
	Owner        string `orm:size(50)`
	Department   string `orm:size(50)`
	Cpu          string `orm:size(50)`
	Mem          string `orm:size(50)`
	Disk         string `orm:size(50)`
	Idc          string `orm:size(50)`
	Root         string `orm:size(10)`
	Rootpwd      string `orm:size(50)`
	Read         string `orm:size(10)`
	Readpwd      string `orm:size(50)`
	Group        string `orm:size(50)`
	Comment      string `orm:size(100)`
	Created      time.Time
}

// register host model
func init() {
	orm.RegisterModel(new(Host))
}

func GetHostCount() (int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	//total, err := o.QueryTable("host").All(&hosts)
	total, err := o.Raw("select distinct(ip) from host").QueryRows(&hosts)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetAppCount() (int64, error) {
	o := orm.NewOrm()
	apps := make([]*Host, 0)
	total, err := o.QueryTable("host").All(&apps)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetHosts(currPage, pageSize int) ([]*Host, int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
	if err != nil {
		return nil, 0, err
	}
	return hosts, total, err
}

func GetHostById(id string) (*Host, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{}
	err = o.QueryTable("host").Filter("id", hid).One(host)
	return host, err
}

func AddHost(class, service_name, name, ip, pubip, port, os_type, owner, department, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	rootpwd, _ = AESEncode(rootpwd, AesKey)
	readpwd, _ = AESEncode(readpwd, AesKey)
	host := &Host{
		Class:        class,
		Service_name: service_name,
		Name:         name,
		Ip:           ip,
		Pubip:        pubip,
		Port:         port,
		Os_type:      os_type,
		Owner:        owner,
		Department:   department,
		Root:         root,
		Read:         read,
		Rootpwd:      rootpwd,
		Readpwd:      readpwd,
		Cpu:          cpu,
		Mem:          mem,
		Disk:         disk,
		Group:        group,
		Idc:          idc,
		Comment:      comment,
		Created:      time.Now(),
	}
	err := o.QueryTable("host").Filter("name", name).Filter("service_name", service_name).Filter("ip", ip).One(host)
	if err == nil {
		msg = "主机名或服务名已存在"
		return nil, msg
	}
	_, err = o.Insert(host)
	msg = "添加主机成功"
	return err, msg
}

func ModifyHost(id, class, service_name, name, ip, pubip, port, os_type, owner, department, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	rootpwd, _ = AESEncode(rootpwd, AesKey)
	readpwd, _ = AESEncode(readpwd, AesKey)
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{
		Id: hid,
	}
	err = o.Read(host)
	if err == nil {
		host.Class = class
		host.Service_name = service_name
		host.Name = name
		host.Ip = ip
		host.Pubip = pubip
		host.Port = port
		host.Os_type = os_type
		host.Owner = owner
		host.Department = department
		host.Root = root
		host.Read = read
		host.Rootpwd = rootpwd
		host.Readpwd = readpwd
		host.Cpu = cpu
		host.Mem = mem
		host.Disk = disk
		host.Group = group
		host.Idc = idc
		host.Comment = comment
	}
	o.Update(host)
	msg = "修改成功"
	return err, msg
}

func DeleteHost(id string) error {
	o := orm.NewOrm()
	var err1 error
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{
		Id: hid,
	}
	if _, err1 = o.Raw("insert into recycle_host(class,service_name,name,ip,pubip,port,os_type,owner,department,cpu,mem,disk,`group`,idc,COMMENT,created) select class,service_name,name,ip,pubip,port,os_type,owner,department,cpu,mem,disk,`group`,idc,COMMENT,CURRENT_TIMESTAMP() as created from `host` where id=?", hid).Exec(); err1 == nil {

		if _, err = o.Delete(host); err == nil {
			return nil
		}
	}
	return nil

}

func SearchHostCount(idc, name string) (int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	ids := []string{"%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%"}
	//total, err := o.QueryTable("host").Filter("idc", idc).Filter("name__icontains", name).All(&hosts)
	if len(name) == 0 {
		total, err := o.QueryTable("host").Filter("idc", idc).All(&hosts)
		return total, err
	} else {
		total, err := o.Raw("select * from host where name like ? or service_name like ? or ip like ? or comment like ? and idc LIKE ?", ids, idc).QueryRows(&hosts)
		return total, err
	}
	//return total, err
}

func SearchHostByName(currPage, pageSize int, idc, name string) ([]*Host, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	ids := []string{"%" + name + "%", "%" + name + "%", "%" + name + "%", "%" + name + "%"}
	//_, err := o.QueryTable("host").Filter("idc", idc).Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
	//var err error
	if len(name) == 0 {
		_, err := o.QueryTable("host").Filter("idc", idc).Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
		return hosts, err
	} else {
		_, err := o.Raw("select * from host where name like ? or service_name like ? or ip like ? or comment like ? and idc = ? limit ?,?", ids, idc, (currPage-1)*pageSize, pageSize).QueryRows(&hosts)
		return hosts, err
	}
	//return hosts, err
}

/*
func SearchHostWeekCount() (int64, error) {
	o := orm.NewOrm()
	//hosts := make([]*Host, 0)
	var hosts int
	//total, err := o.QueryTable("host").Filter("name__icontains", name).All(&hosts)
	total, err := o.Raw("select count(*) as Count from host where date_sub(curdate(), INTERVAL ? DAY) <= date(`created`)", 7).QueryRows(&hosts)
	return total, err
}
*/

func SearchHostWeekCount() (int64, error) {
	o := orm.NewOrm()
	//hosts := make([]*Host, 0)
	var count int64
	//total, err := o.QueryTable("host").Filter("name__icontains", name).All(&hosts)
	err := o.Raw("select count(*) as Count from host where date_sub(curdate(), INTERVAL ? DAY) <= date(`created`)", 7).QueryRow(&count)
	return count, err
}

func SearchHostByWeek(currPage, pageSize int, created string) ([]*Host, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	/*
			var cond *orm.Condition
			cond = orm.NewCondition()
			cond = cond.Or("name__icontains", name)
			cond = cond.Or("ip__icontains", "ip")
			var qs orm.QuerySeter
			qs = o.QueryTable("host").Filter("idc", idc).Limit(pageSize, (currPage-1)*pageSize).SetCond(cond)
			_, err := qs.All(&hosts)

		//_, err := o.QueryTable("host").Filter("created", created).Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
		//_, err := o.Raw("select * from host where date_sub(curdate(), INTERVAL 7 DAY) <= date(?)", created).QueryRows(&hosts)
		_, err := orm.NewOrm().QueryTable("host").Filter("created").
	*/
	_, err := o.Raw("select * from host where date_sub(curdate(), INTERVAL 7 DAY) <= date(`created`)").QueryRows(&hosts)
	return hosts, err
}

func QueryHostWeekExport(method string) (*map[int64][]string, []string, int64) {
	result := make(map[int64][]string)
	var columns []string
	var total int64
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total
	}

	defer conn.Close()
	if method == "week" {
		rows, err := conn.Query("select id,class,name,service_name,ip,pubip,port,os_type,owner,department,cpu,mem,disk,idc,`group`,comment,created from `host` where DATE_SUB(CURDATE(), INTERVAL 7 day) <= date(`created`) order by id")
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
		rows, err := conn.Query("select id,class,name,service_name,ip,pubip,port,os_type,owner,department,cpu,mem,disk,idc,`group`,comment,created from `host` order by id")
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
