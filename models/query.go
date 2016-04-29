package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func Query(alias, sqltext string) (*map[int64][]string, []string, int64, error) {
	result := make(map[int64][]string)
	var total int64
	var columns []string
	o := orm.NewOrm()
	schema := &Schema{}
	o.QueryTable("schema").Filter("name", alias).All(schema)

	passwd, _ := AESDecode(schema.Passwd, AesKey)
	schemaUrl := schema.User + ":" + passwd + "@tcp(" + schema.Addr + ":" + schema.Port + ")/" + schema.DBName + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to %v server successfully !", schema.Name))

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total, err
	}

	defer conn.Close()
	sqlTrim := strings.Trim(sqltext, " ")
	sqlPrefix := sqlTrim[:6]
	if "select" == strings.ToLower(sqlPrefix) {
		rows, err := conn.Query(sqlTrim)
		beego.Info(sqltext)
		if err != nil {
			return &result, columns, total, err
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
		beego.Info(result)
		return &result, columns, total, nil
	} else {
		return &result, columns, total, errors.New("没有执行权限，请联系DBAs！")
	}
}
