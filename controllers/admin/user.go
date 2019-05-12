package admin

import (
	"fmt"
	"web/packs/gin"
	"web/packs/util"
)

type User struct {
	Base
}

func (this *User) Ulist(c *gin.Context) {
	page := util.Str2int(c.Query("page"))

	count := this.rowsCount("select count(*) as `count` from user", "count")

	offset := this.pageOffset(page)

	var users []map[string]string

	this.mysqlInstance().GetAll(fmt.Sprintf("select * from user limit %d,%d", offset, this.pageSize()), &users)

	pagebar := util.NewPager(page, count, this.pageSize(), "/admin/user/list", true).ToString()

	c.HTML(200, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (_ *User) Uadd(c *gin.Context) {

}