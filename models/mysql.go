package models

import (
	"database/sql"
	"strings"
	"sync"
	_ "web/packs/gin/plugins/mysql"
	"web/packs/util"
)

type Mysql struct {
	instance *sql.DB
	lock	sync.Mutex
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

	this.instance, err = sql.Open(util.Gapp_db, dns)

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

func (_ *Mysql) GetAll (querySql string, columns string) (int, *[]map[string]string, error) {
	var total_num int
	var err error
	var ret []map[string]string

	err = this.instance.QueryRow(strings.Replace(querySql,"?", "count(*)", 1)).Scan(&total_num)


	if(nil != err || total_num < 1) {
		return total_num,&ret, err
	}

	var rows *sql.Rows
	rows, err  = this.instance.Query(strings.Replace(querySql, "?", columns, 1))
	defer rows.Close()

	if nil != err {
		return total_num,&ret, err
	}

	var rcol  []string
	rcol, err = rows.Columns()
	if err != nil {
		return total_num,&ret, err
	}
	cnum := len(rcol)

	ret = make([]map[string]string, total_num)
	scaner := make([]interface{}, cnum)
	values := make([]interface{}, cnum)
	for j := range values {
		scaner[j] = &values[j]
	}

	index := 0
	for rows.Next() {
		err = rows.Scan(scaner...)
		ret[index] = make(map[string]string)
		for i, col := range values {
			if nil != col {
				ret[index][rcol[i]] = string(col.([]byte))
			}
		}
		index += 1
	}

	return total_num,&ret, nil
}

func (_ *Mysql) GetRow(querySql string) (*map[string]string, error) {
	ret := make(map[string]string)

	row, err := this.instance.Query(querySql)
	defer row.Close()
	if nil != err {
		return &ret, err
	}

	columns, err := row.Columns()
	if nil != err {
		return &ret, err
	}

	cnum := len(columns)
	scaner := make([]interface{}, cnum)
	values := make([]interface{}, cnum)

	for j := range values {
		scaner[j] = &values[j]
	}

	row.Next()
	err = row.Scan(scaner...)
	for i, col := range values {
		if nil != col {
			ret[columns[i]] = string(col.([]byte))
		}
	}

	return  &ret, nil
}

func (_ *Mysql) GetOne(querySql string) (string, error) {
	var tmp interface{}
	ret := ""

	err := this.instance.QueryRow(querySql).Scan(&tmp)

	if nil != err {
		return ret, err
	}

	ret = string(tmp.([]byte))

	return ret, err
}

func (_ *Mysql) Exec(querySql string) (sql.Result, error) {
	return this.instance.Exec(querySql)
}
