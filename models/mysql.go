package models

import (
	"database/sql"
	_ "web/packs/gin/plugins/mysql"
	"web/packs/util"
)

type Mysql struct {
	instance *sql.DB
}

var this Mysql

func init() {
	var err error

	db_user := util.Gmysql_user
	db_pwd  := util.Gmysql_pwd
	db_host := util.Gmysql_host
	db_port := util.Gmysql_port
	db_name := util.Gmysql_dbname
	db_char := util.Gmysql_charset

	dns := db_user + ":" + db_pwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=" + db_char + "&loc=Asia%2FShanghai"

	this.instance, err = sql.Open("mysql", dns)

	err_dping := this.instance.Ping()

	util.CheckErr(err_dping)

	db_idel	:= util.Str2int(util.Gmysql_max_idle_conns)
	util.CheckErr(err)

	db_open := util.Str2int(util.Gmysql_max_open_conns)
	util.CheckErr(err)

	this.instance.SetMaxIdleConns(db_idel)

	this.instance.SetMaxOpenConns(db_open)
}

func (_ *Mysql) GetInstance() *sql.DB{
	return this.instance
}

func (_ *Mysql) GetRow(querySql string, record map[string]string) error {
	row, err := this.instance.Query(querySql)
	defer row.Close()

	if nil != err {
		return err
	}

	columns, err := row.Columns()
	if nil != err {
		return err
	}

	scanArgs := make([]interface{}, len(columns))
	values   := make([]interface{}, len(columns))

	for j := range values {
		scanArgs[j] = &values[j]
	}

	row.Next()
	err = row.Scan(scanArgs...)
	for i, col := range values {
		if nil != col {
			record[columns[i]] = string(col.([]byte))
		}
	}

	return  nil
}

func (_ *Mysql) GetAll(querySql string, records *[]map[string]string) error {
	rows, err := this.instance.Query(querySql)
	defer rows.Close()

	if nil != err {
		return err
	}

	columns, err := rows.Columns()
	if nil != err {
		return err
	}

	scanArgs 	:= make([]interface{}, len(columns))
	values 		:= make([]interface{}, len(columns))
	record      := make(map[string]string)

	for j := range values {
		scanArgs[j] = &values[j]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		*records = append(*records, record)
	}

	return nil
}