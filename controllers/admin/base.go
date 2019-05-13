package admin

import (
	"database/sql"
	"log"
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

func (this *Base) Invoke(c *gin.Context) {

	ctls := map[string]interface{}{
		"index"		:		&Index{},
		"user"		:		&User{},
	}

	ctl := c.Param("ctl")

	if "" == ctl {
		c.Next()
	}

	controller, exist := ctls[ctl]

	if !exist {
		c.Next()
	}

	act := c.Param("act")

	if "" == act {
		c.Next()
	}

	first := strings.ToUpper(act[1:2])
	act = first + act[2:]

	refVal := reflect.ValueOf(controller)
	method := refVal.MethodByName(act)
	args := make([]reflect.Value, 0)
	args[1] = reflect.ValueOf("test")
	log.Println(args)
	method.Call(args)

}

func (this *Base) isPost(c *gin.Context) bool {
	return c.Request.Method == "POST"
}

func (this *Base) Display() {

}

func (this *Base) Assign() {

}