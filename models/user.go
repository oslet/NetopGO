package models

import (
	//	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"time"
)

type User struct {
	Id      int64
	Name    string `orm:size(100)`
	Passwd  string `orm:size(100)`
	Email   string `orm:size(50)`
	Dept    string `orm:size(20)`
	Created time.Time
	Auth    int64
	Tel     string `orm:size(11)`
}

func init() {
	orm.RegisterModel(new(User))
}

func Login(uname string) (*User, error) {
	o := orm.NewOrm()
	user := &User{}
	err := o.QueryTable("user").Filter("name", uname).One(user)
	return user, err
}

func GetUserCount() (int64, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	total, err := o.QueryTable("user").All(&users)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetUsers(currPage, pageSize int) ([]*User, int64, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	total, err := o.QueryTable("user").Limit(pageSize, (currPage-1)*pageSize).All(&users)
	if err != nil {
		return nil, 0, err
	}
	return users, total, err
}

func AddUser(name, passwd, email, tel, auth, dept string) error {
	o := orm.NewOrm()
	authInt, err := strconv.ParseInt(auth, 10, 64)
	if err != nil {
		return err
	}
	passwd = string(Base64Encode([]byte(passwd)))
	user := &User{
		Name:    name,
		Passwd:  passwd,
		Email:   email,
		Tel:     tel,
		Auth:    authInt,
		Dept:    dept,
		Created: time.Now(),
	}
	err = o.QueryTable("user").Filter("name", name).One(user)
	//fmt.Printf("+++++++is exists+++++++: %v\n", err)
	if err == nil {
		return nil

	}
	_, err = o.Insert(user)
	return err
}

func MofifyUser(id, name, passwd, email, tel, auth, dept string) error {
	o := orm.NewOrm()
	uid, err := strconv.ParseInt(id, 10, 64)
	authInt, err := strconv.ParseInt(auth, 10, 64)
	passwd = string(Base64Encode([]byte(passwd)))
	user := &User{
		Id: uid,
	}
	err = o.Read(user)
	if err == nil {
		user.Name = name
		user.Passwd = passwd
		user.Email = email
		user.Tel = tel
		user.Auth = authInt
		user.Dept = dept
	}
	o.Update(user)
	return err
}

func DeleteUser(id string) error {
	o := orm.NewOrm()
	uid, err := strconv.ParseInt(id, 10, 64)
	user := &User{
		Id: uid,
	}
	_, err = o.Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(id string) (*User, error) {
	o := orm.NewOrm()
	uid, err := strconv.ParseInt(id, 10, 64)
	user := &User{}
	err = o.QueryTable("user").Filter("id", uid).One(user)
	return user, err
}

func SearchUserCount(name string) (int64, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	total, err := o.QueryTable("user").Filter("name__icontains", name).All(&users)
	return total, err
}

func SearchUserByName(currPage, pageSize int, name string) ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	_, err := o.QueryTable("user").Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&users)
	return users, err
}

func ResetPasswd(id, passwd string) error {
	o := orm.NewOrm()
	uid, err := strconv.ParseInt(id, 10, 64)
	passwd = string(Base64Encode([]byte(passwd)))
	user := &User{
		Id: uid,
	}
	err = o.Read(user)
	if err == nil {
		user.Passwd = passwd
	}
	o.Update(user)
	return err
}

func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}
