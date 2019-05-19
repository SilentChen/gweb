package admin

import (
	"fmt"
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

		params, errMsg := this.getAndCheckParams(c, []string{"username", "password", "password2", "email", "active"})
		if len(errMsg) > 0 {
			this.errorShow(c, errMsg)
			return
		}

		username 	= params["username"]
		password 	= params["password"]
		password2 	= params["password2"]
		email 		= params["email"]
		active 		= params["active"]

		id = c.DefaultPostForm("id", "")

		if password != password2 {
			this.errorShow(c, []string{"bad two params"})
			return
		}

		var arow int64
		if "" != id {
			querySql = fmt.Sprintf("update user set user_name = '%s', password = password('%s'), email = '%s', active = '%s' where id = %s", username, password, email, active, id)
			dbRet, err := this.mysqlInstance().Exec(querySql)
			util.CheckErr(err)
			arow, _ = dbRet.RowsAffected()

		}else{
			querySql = fmt.Sprintf("insert into user (user_name, password, email, active) values('%s',password('%s'),'%s','%s')", username, password, email, active)
			dbRet, err := this.mysqlInstance().Exec(querySql)
			util.CheckErr(err)
			arow, _ = dbRet.RowsAffected()
			tmpId, _ := dbRet.LastInsertId()
			id = util.Int642str(tmpId)
		}

		if arow > 0 {
			uinfo["username"]	=	username
			uinfo["email"]		=	email
			uinfo["active"]		=	active
			uinfo["id"]			=	id
		}
	}

	c.HTML(http.StatusOK, "admin/user/edit", map[string]interface{}{
		"uinfo"		:		uinfo,
		"desc"		:		desc,
	})
}