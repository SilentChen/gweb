package admin

import (
	"fmt"
	"log"
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

	log.Println(users)

	c.HTML(200, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (this *User) Add(c *gin.Context) {
	log.Println("testtest")
	/*c.HTML(200, "admin/user/add", map[string]interface{}{
		"test"		:		"a",
	})*/
}