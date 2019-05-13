package admin

import (
	"database/sql"
	"log"
	"reflect"
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
	m := c.Param("module")
	a := c.Param("action")
	mt := reflect.TypeOf(m)
	at := reflect.TypeOf(a)

	log.Printf("mt is: %s and at is: %s", mt, at)

	mtv := reflect.ValueOf(m)
	atv := reflect.ValueOf(a)

	log.Printf("mtv is : %s and atv is: %s", mtv, atv)

	//method := mtv.MethodByName("info")
	//rgs := make([]reflect.Value, 0)

	//method.Call(args)
}

func (this *Base) isPost(c *gin.Context) bool {
	return c.Request.Method == "POST"
}

func (this *Base) Display() {

}

func (this *Base) Assign() {

}