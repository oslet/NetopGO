package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Query(alias, sqltext string) (*map[int64][]string, []string, int64) {

	o := orm.NewOrm()
	schema := &Schema{}
	o.QueryTable("schema").Filter("name", alias).All(schema)

	passwd, _ := AESDecode(schema.Passwd, AesKey)
	schemaUrl := schema.User + ":" + passwd + "@tcp(" + schema.Addr + ":" + schema.Port + ")/" + schema.DBName + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to %v server successfully !", schema.Name))

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		fmt.Println("mysql connect error")
		return nil, nil, 0
	}

	defer conn.Close()
	rows, err := conn.Query(sqltext)
	if err != nil {
		fmt.Println("mysql query error", err.Error())
	}
	defer rows.Close()
	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	result := make(map[int64][]string)
	var total int64
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
	return &result, columns, total
}
