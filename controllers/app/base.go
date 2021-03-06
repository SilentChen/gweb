package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"web/models"
	"web/packs/gin"
	"web/packs/util"
)

const (
	MSG_SUCCESS				=	"success"
	MSG_ERROR				=	"error"
	APISTATUS_OK			=	"0"			//	success
	APISTATUS_FAIL			=	"1"			//	fail, maybe network, try again
	APISTATUS_ERR			=	"-1"		//	error, program go wrong
	APISTATUS_ERRPARAMS		=	"-2"		//	bad params
	APISTATUS_ERRINVOLID	=	"-3"		//	involid opt
	APISTATUS_MYSQLQUERY	=	"-4"		//	mysql query error
	APISTATUS_MYSQLAROW		=	"-5"		//	mysql affected row error
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

	return (page - 1) * this.pageSize()
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

func (this *Base) getAndCheckParams(c *gin.Context, params []string) (map[string]string, []string) {
	ret := make(map[string]string)
	var errMsg []string

	if len(params) < 1 {
		return ret, errMsg
	}

	tmp := ""
	for _, v := range params {
		tmp = c.Query(v)
		if "" == tmp {
			tmp = c.PostForm(v)
		}

		if "" == tmp {
			errMsg = append(errMsg, fmt.Sprintf("bad param: %s", v))
		}else{
			ret[v] = tmp
		}

	}

	return ret, errMsg
}

func (this *Base) Invoke(c *gin.Context) {
	ctls := map[string]interface{}{
		"index"		:		&Index{},
	}

	ctl := c.Param("ctl")
	act := c.Param("act")

	if "" == act && "" == ctl {
		ctl = "index"
		act = "index"
	}else if "" == ctl && "" != act {
		ctl = "index"
	}

	controller, exist := ctls[ctl]
	if !exist {
		this.errorShow(c, []string{"bad ctl"})
		return
	}

	first  := strings.ToUpper(act[0:1])
	action := first + act[1:]
	refVal := reflect.ValueOf(controller)
	method := refVal.MethodByName(action)
	if method.Kind() == reflect.Invalid {
		this.errorShow(c, []string{"bad act"})
		return
	}
	c.Set("ctl", ctl)
	c.Set("act", act)

	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(c)
	method.Call(args)
}

func (this *Base) errorShow(c *gin.Context, errMsg []string) {
	var msg string
	if len(errMsg) > 0 {
		for _, v := range errMsg {
			msg += v + "<br/>"
		}
	}
	if "" == msg {
		msg = "Oh God ! Something Went Wrong !"
	}
	c.HTML(http.StatusForbidden, "app/layout/default", map[string]interface{}{"msg":msg})
	return
}

func (this *Base) display(c *gin.Context, params map[string]interface{}) {
	template := "app/"
	ctl := c.GetString("ctl")
	act := c.GetString("act")
	if "" == ctl || "" == act {
		template += "layout/default"
	}else{
		template += ctl + "/" + act
	}

	c.HTML(http.StatusOK, template, params)
}