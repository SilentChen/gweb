package admin

import (
	"fmt"
	"log"
	"net/http"
	"web/packs/gin"
	"web/packs/util"
)

type User struct {
	Base
}

func (this *User) List(c *gin.Context) {
	page := util.Str2int(c.Query("page"))

	count, _ := this.mysqlInstance().GetOne("select count(*) from user")

	var users []map[string]string

	this.mysqlInstance().GetAll(fmt.Sprintf("select * from user limit %d,%d", this.pageOffset(page), this.pageSize()), &users)

	pagebar := util.NewPager(page, util.Str2int(count), this.pageSize(), "/admin/user/list", true).ToString()

	c.HTML(http.StatusOK, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (this *User) Edit(c *gin.Context) {
	uinfo := map[string]interface{} {
		"username"		:		"",
		"password"		:		"",
		"password2"		:		"",
		"email"			:		"",
		"active"		:		0,

	}

	id := c.Query("id")
	log.Println(id)


	c.HTML(http.StatusOK, "admin/user/edit", map[string]interface{}{
		"uinfo"		:		uinfo,
	})
}