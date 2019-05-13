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

	count, _ := this.mysqlInstance().GetOne("select count(*) from user")

	var users []map[string]string

	this.mysqlInstance().GetAll(fmt.Sprintf("select * from user limit %d,%d", this.pageOffset(page), this.pageSize()), &users)

	pagebar := util.NewPager(page, util.Str2int(count), this.pageSize(), "/admin/user/list", true).ToString()

	c.HTML(200, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (_ *User) Uadd(c *gin.Context) {

}