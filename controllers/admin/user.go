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
	page := util.Str2int(c.DefaultQuery("page", "0"))

	count, users, _ := this.mysqlInstance().GetAll(fmt.Sprintf("select ? from user limit %d,%d", this.pageOffset(page), this.pageSize()), "*")

	pagebar := util.NewPager(page, count, this.pageSize(), "/admin/user/list", true).ToString()

	c.HTML(http.StatusOK, "admin/user/list", map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (this *User) Edit(c *gin.Context) {
	uinfo := map[string]string {
		"id"			:		"",
		"username"		:		"",
		"email"			:		"",
		"active"		:		"0",
	}

	desc := "add"

	id := c.Query("id")
	if "" != id {
		tmp, err := this.mysqlInstance().GetRow(fmt.Sprintf("select id,user_name as username, email, active from user where id = %d", util.Str2int(id)))
		util.CheckErr(err)
		uinfo = *tmp

		if "" != uinfo["id"] {
			desc = "update"
		}
	}

	if this.isPost(c) {
		var querySql,username,password,password2,email,active,id string

		this.paramCheckExist(c,"username", &username, "bad username")
		this.paramCheckExist(c,"password", &password, "bad password")
		this.paramCheckExist(c,"password2", &password2, "bad password2")
		this.paramCheckExist(c,"email", &email, "bad email")
		this.paramCheckExist(c,"active", &active, "bad active")

		id = c.DefaultPostForm("id", "")

		if "" != id {
			querySql = fmt.Sprintf("update user set user_name = %s, password = password(%s), email = %s, active = %s whhere id = %s", username, password, email, active, id)
		}else{
			querySql = fmt.Sprintf("insert into user (user_name, password, email, active) values(%s,password(%s),%s,%s)", username, password, email, active)
		}
		log.Println("sql is: ", querySql)
		affRow, _ := this.mysqlInstance().Exec(querySql)
		log.Println("affective row is: ", affRow)
	}

	c.HTML(http.StatusOK, "admin/user/edit", map[string]interface{}{
		"uinfo"		:		uinfo,
		"desc"		:		desc,
	})
}