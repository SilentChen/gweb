package admin

import (
	"fmt"
	"web/packs/gin"
	"web/packs/util"
)

type User struct {
	Base
}

func (this *User) List(c *gin.Context) {
	page := util.Str2int(c.DefaultQuery("page", "0"))
	totalNum := this.mysqlInstance().DefGetOne("select count(*) from user", "0")

	_, users, _ := this.mysqlInstance().GetAll(fmt.Sprintf("select * from user limit %d,%d", this.pageOffset(page), this.pageSize()))

	pagebar := util.NewPager(page, util.Str2int(totalNum), this.pageSize(), "/admin/user/list", true).ToString()

	this.display(c, map[string]interface{}{
		"list"			:		users,
		"pagebar"		:		pagebar,
	})
}

func (this *User) Edit(c *gin.Context) {
	var gmsg string

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
		var querySql, id string

		params, errMsg := this.getAndCheckParams(c, []string{"username", "password", "password2", "email", "active"})
		if len(errMsg) > 0 {
			this.errorShow(c, errMsg)
			return
		}

		id = c.DefaultPostForm("id", "")

		if params["password"] != params["password2"] {
			this.errorShow(c, []string{"bad two params"})
			return
		}

		var arow int64
		if "" != id {
			querySql = fmt.Sprintf("update user set user_name = '%s', password = password('%s'), email = '%s', active = '%s' where id = %s", params["username"],  params["password"], params["email"], params["active"], id)
			dbRet, err := this.mysqlInstance().Exec(querySql)
			util.CheckErr(err)
			arow, _ = dbRet.RowsAffected()

		}else{
			querySql = fmt.Sprintf("insert into user (user_name, password, email, active) values('%s',password('%s'),'%s','%s')", params["username"], params["password"], params["email"], params["active"])
			dbRet, err := this.mysqlInstance().Exec(querySql)
			util.CheckErr(err)
			arow, _ = dbRet.RowsAffected()
			tmpId, _ := dbRet.LastInsertId()
			id = util.Int642str(tmpId)
		}

		if arow > 0 {
			uinfo["username"]	=	params["username"]
			uinfo["email"]		=	params["email"]
			uinfo["active"]		=	params["active"]
			uinfo["id"]			=	id
			gmsg				=	MSG_SUCCESS
		}else{
			gmsg				=	MSG_ERROR
		}
	}

	this.display(c, map[string]interface{}{
		"uinfo"		:		uinfo,
		"desc"		:		desc,
		"gmsg"		:		gmsg,
	})
}