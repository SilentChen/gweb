package admin

import (
	"log"
	"reflect"
	"web/models"
	"web/packs/gin"
)

type Base struct {
	dbinstance *models.Mysql
}

func (this *Base) dbInstance() *models.Mysql{
	if nil == this.dbinstance {
		this.dbinstance = new(models.Mysql)
	}

	return this.dbinstance
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