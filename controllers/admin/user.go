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

	count, users, _ := this.mysqlInstance().GetAll(fmt.Sprintf("select ? from user limit %d,%d", this.pageOffset(page), this.pageSize()), "*")

	pagebar := util.NewPager(page, count, this.pageSize(), "/admin/user/list", true).ToString()

	c.HTML(http.StatusOK, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (this *User) Edit(c *gin.Context) {
	uinfo := map[string]string {
		"username"		:		"",
		"password"		:		"",
		"password2"		:		"",
		"email"			:		"",
		"active"		:		"0",

	}

	id := c.Query("id")
	if "" != id {
		tmp, err :=  this.mysqlInstance().GetRow(fmt.Sprint("select * from user where id = %d", id))
		util.CheckErr(err)
		log.Println(tmp)
	}


	c.HTML(http.StatusOK, "admin/user/edit", map[string]interface{}{
		"uinfo"		:		uinfo,
	})
}