package admin

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
		"user"		:		&User{},
		"article"	:		&Article{},
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
		this.errorShow(c, []string{"bad ctl"})
		return
	}

	first := strings.ToUpper(act[1:2])		//change the second char into upper
	act = first + act[2:]					//cut the string begin from the third char, first is '/', the second will be replace by it's upper own

	refVal := reflect.ValueOf(controller)
	method := refVal.MethodByName(act)
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
	c.HTML(http.StatusForbidden, "admin/default", map[string]interface{}{"msg":msg})
	return
}