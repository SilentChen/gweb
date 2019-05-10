package admin

import (
	"log"
	"web/packs/gin"
	"web/packs/util"
)

type User struct {
	Base
}

func (this *User) Ulist(c *gin.Context) {
	var page int

	var count int

	count = 100

	var users []map[string]string

	this.dbInstance().GetAll("select * from user", &users)

	log.Println(users, page, count, this.pageSize())

	pagebar := util.NewPager(page, int(count), this.pageSize(), "/admin/user/list", true).ToString()
	log.Println(pagebar)
	/*c.HTML(200, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})*/
}

func (_ *User) Uadd(c *gin.Context) {

}