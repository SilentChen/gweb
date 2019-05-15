package admin

import (
	"database/sql"
	"net/http"
	"reflect"
	"strings"
	"web/models"
	"web/packs/gin"
	"web/packs/util"
)

type Base struct {
	db *models.Mysql	//mysql instance
	pz int				//pagesize
}

var this Base

func (this *Base) pageSize() int {
	if 0 == this.pz {
		this.pz = util.Str2int(util.Gwebsetting.Get("webPageSize"))
	}

	return this.pz
}

func (this *Base) pageOffset(page int) int {
	if page < 1 {
		page = 1
	}

	return (page - 1) * this.pz
}

func (this *Base) mysqlInstance() *models.Mysql{
	if nil == this.db {
		this.db = new(models.Mysql)
	}

	return this.db
}

func (this *Base) dbInstance() *sql.DB {
	if nil == this.db {
		this.mysqlInstance()
	}

	return this.db.GetInstance()
}

func (this *Base) isPost(c *gin.Context) bool {
	return c.Request.Method == "POST"
}

func (this *Base) Invoke(c *gin.Context) {

	ctls := map[string]interface{}{
		"index"		:		&Index{},
		"user"		:		&User{},
	}

	ctl := c.Param("ctl")
	act := c.Param("act")

	if "" == act && "" == ctl {
		ctl = "index"
		act = "/index"
	}else if "/" == act && "" != ctl {
		act = "/" + ctl
		ctl = "index"
	}

	controller, exist := ctls[ctl]
	if !exist {
		c.HTML(http.StatusNotFound, "admin/default", map[string]interface{}{"message"	:	"bad ctl",})
		return
	}

	first := strings.ToUpper(act[1:2])		//change the second char into upper
	act = first + act[2:]					//cut the string begin from the third char, first is '/', the second will be replace by it's upper own

	refVal := reflect.ValueOf(controller)
	method := refVal.MethodByName(act)

	if method.Kind() == reflect.Invalid {
		c.HTML(http.StatusNotFound, "admin/default", map[string]interface{}{"message"	:	"bad act",})
		return
	}

	c.Set("ctl", ctl)
	c.Set("act", act)

	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(c)
	method.Call(args)
}