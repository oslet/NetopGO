package models

import (
	"github.com/astaxie/beego/orm"
	"math"
	"time"
)

type User struct {
	Id      int
	Name    string `orm:size(100)`
	Passwd  string `orm:size(100)`
	Email   string `orm:size(50)`
	Dept    string `orm:size(20)`
	Created time.Time
	Auth    int
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
	total, err := o.QueryTable("user").Limit(pageSize, (currPage-1)*1).All(&users)
	if err != nil {
		return nil, 0, err
	}
	return users, total, err
}

// func Paginator(page, pageSize int, total int64) map[string]interface{} {
// 	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
// 	if page > totalPages {
// 		page = totalPages
// 	}
// 	if page <= 0 {
// 		page = 1
// 	}

// }

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
