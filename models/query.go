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
	sqlTrun := sqlTrim[:8]
	sqlDrop := sqlTrim[:4]
	sqlAlter := sqlTrim[:5]

	beego.Info(sqlTrim)

	if "delete" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("您没有delete权限，请联系DBAs！")
	} else if "update" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("您没有update权限，请联系DBAs！")
	} else if "optimize" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("您没有optimize权限，请联系DBAs！")
	} else if "insert" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("您没有insert权限，请联系DBAs！")
	} else if "truncate" == strings.ToLower(sqlTrun) {
		return &result, columns, total, errors.New("您没有truncate权限，请联系DBAs！")
	} else if "drop" == strings.ToLower(sqlDrop) {
		return &result, columns, total, errors.New("您没有drop权限，请联系DBAs！")
	} else if "alter" == strings.ToLower(sqlAlter) {
		return &result, columns, total, errors.New("您没有alter权限，请联系DBAs！")
	} else {
		rows, err := conn.Query(sqlTrim)
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
	}
}

func QueryServer(alias, sqltext string) (*map[int64][]string, []string, int64, error, bool, int64) {
	result := make(map[int64][]string)
	var total int64
	var columns []string
	var isAffected bool = false
	o := orm.NewOrm()
	schema := &Schema{}
	o.QueryTable("schema").Filter("name", alias).All(schema)

	passwd, _ := AESDecode(schema.Passwd, AesKey)
	schemaUrl := schema.User + ":" + passwd + "@tcp(" + schema.Addr + ":" + schema.Port + ")/" + schema.DBName + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to %v server successfully !", schema.Name))

	conn, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		return &result, columns, total, err, isAffected, 0
	}

	defer conn.Close()
	sqlTrim := strings.Trim(sqltext, " ")
	sqlPrefix := sqlTrim[:6]
	sqlTrun := sqlTrim[:8]
	sqlDrop := sqlTrim[:4]
	sqlAlter := sqlTrim[:5]

	beego.Info(sqlTrim)

	if "delete" == strings.ToLower(sqlPrefix) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "update" == strings.ToLower(sqlPrefix) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "optimize" == strings.ToLower(sqlPrefix) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "insert" == strings.ToLower(sqlPrefix) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "truncate" == strings.ToLower(sqlTrun) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "drop" == strings.ToLower(sqlDrop) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "alter" == strings.ToLower(sqlAlter) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else if "create" == strings.ToLower(sqlAlter) {
		res, err := o.Raw(sqlTrim).Exec()
		if err != nil {
			return &result, columns, total, err, isAffected, 0
		}
		num, _ := res.RowsAffected()
		isAffected = true
		return &result, columns, total, nil, isAffected, num
	} else {
		rows, err := conn.Query(sqlTrim)
		if err != nil {
			return &result, columns, total, err, isAffected, 0
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
		return &result, columns, total, nil, isAffected, 0
	}
}

func QueryProxy(alias, sqltext string) (*map[int64][]string, []string, int64, error) {
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
	sqlTrun := sqlTrim[:8]
	sqlDrop := sqlTrim[:4]
	sqlAlter := sqlTrim[:5]

	beego.Info(sqlTrim)

	if "delete" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "update" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "optimize" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "insert" == strings.ToLower(sqlPrefix) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "truncate" == strings.ToLower(sqlTrun) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "drop" == strings.ToLower(sqlDrop) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "alter" == strings.ToLower(sqlAlter) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "create" == strings.ToLower(sqlAlter) {
		return &result, columns, total, errors.New("后端是代理，为保证数据一致性，暂不支持DML、DDL操作！")
	} else if "select" == strings.ToLower(sqlPrefix) {
		rows, err := conn.Query(sqlTrim)
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
	}
	return &result, columns, total, nil
}
