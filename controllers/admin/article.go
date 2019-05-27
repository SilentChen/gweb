package admin

import (
	"fmt"
	"net/http"
	"web/packs/gin"
	"web/packs/util"
)

type Article struct {
	Base
}

func (this *Base) List(c *gin.Context) {
	var querySql string
	searchtype := c.DefaultPostForm("searchtype", "")
	keyword    := c.DefaultPostForm("keyword", "")
	status := c.DefaultQuery("status", "0")

	count1 := this.mysqlInstance().DefGetOne("select count(*) from `post` where status = 1", "0")
	count2 := this.mysqlInstance().DefGetOne("select count(*) from `post` where status = 2", "0")

	page := util.Str2int(c.DefaultQuery("page", "0"))

	if "" != searchtype && "" != keyword {
		querySql = fmt.Sprintf("select ? from `post` where status = %s and %s like '%s%%' limit %d,%d", status, searchtype, keyword, this.pageOffset(page), this.pageSize())
	}else{
		querySql = fmt.Sprintf("select ? from `post` where status = %s limit %d,%d", status, this.pageOffset(page), this.pageSize())
	}

	count, articles, _ := this.mysqlInstance().GetAll(querySql, "*")

	pagebar := util.NewPager(page, count, this.pageSize(), "/admin/article/list", true).ToString()

	this.display(c, map[string]interface{}{
		"list"			:		*articles,
		"pagebar"		:		pagebar,
		"status"		:		status,
		"searchtype"	:		searchtype,
		"count1"		:		count1,
		"count2"		:		count2,
		"keyword"		:		keyword,
	})
}

func (this *Base) Edit(c *gin.Context) {
	var gmsg string
	ainfo := map[string]string{
		"title"		:		"",
		"color"		:		"",
		"is_top"	:		"",
		"tags"		:		"",
		"url_name"	:		"",
		"post_time"	:		"",
		"status"	:		"",
		"content"	:		"",
		"url_type"	:		"",
	}

	desc := "add"
	id := c.Query("id")

	if "" != id {
		row, err := this.mysqlInstance().GetRow(fmt.Sprintf("select * from `post` where id = %s", id))
		if nil == err {
			ainfo = *row
		}

		if "" != ainfo["id"] {
			desc = "edit"
		}
	}

	if this.isPost(c) {
		var querySql string
		params, errMsg := this.getAndCheckParams(c, []string{"title", "tags", "post_time", "status", "content"})
		if len(errMsg) > 0 {
			this.errorShow(c, errMsg)
			return
		}

		params["is_top"] 	= c.DefaultPostForm("is_top", "0")
		params["url_name"] 	= c.DefaultPostForm("url_name", "")
		params["url_type"] 	= c.DefaultPostForm("url_type", "0")
		params["color"] 	= c.DefaultPostForm("color", "")
		var arow int64
		if "" != id {
			querySql = fmt.Sprintf("update `post` set title = '%s', tags = '%s', post_time = '%s', status = %s, content = '%s', is_top = %s, url_name = '%s', url_type = %s, color = '%s' where id = %s", params["title"],params["tags"],params["post_time"],params["status"],params["content"],params["is_top"],params["url_name"],params["url_type"],params["color"], id)
			dbRet, err := this.mysqlInstance().Exec(querySql)
			util.CheckErr(err)
			arow, _ = dbRet.RowsAffected()

		}else{
			querySql = fmt.Sprintf("insert into `post` (title, tags, post_time, status, content, is_top, url_name, url_type, color) values('%s','%s','%s','%s','%s','%s','%s','%s','%s')", params["title"],params["tags"],params["post_time"],params["status"],params["content"],params["is_top"],params["url_name"],params["url_type"],params["color"])
			dbRet, err := this.mysqlInstance().Exec(querySql)
			util.CheckErr(err)
			arow, _ = dbRet.RowsAffected()
			tmpId, _ := dbRet.LastInsertId()
			id = util.Int642str(tmpId)
		}

		if arow > 0 {
			ainfo = params
			gmsg 	= 	MSG_SUCCESS
		}else{
			gmsg	=	MSG_ERROR
		}
	}

	this.display(c, map[string]interface{}{
		"post"		:		ainfo,
		"desc"		:		desc,
		"gmsg"		:		gmsg,
	})
}

func (this *Base) Del(c *gin.Context) {
	ret := make(map[string]interface{})

	id := c.PostForm("id")

	if "" == id {
		ret["state"] = APISTATUS_ERRPARAMS
	}else{
		dbRet, err := this.mysqlInstance().Exec(fmt.Sprintf("delete from `post` where id = %s", id))
		if nil != err {
			ret["state"] = APISTATUS_MYSQLQUERY
		}else{
			arow, _ := dbRet.RowsAffected()
			if arow > 0 {
				ret["state"] = APISTATUS_OK
			}else{
				ret["state"] = APISTATUS_MYSQLAROW
			}
		}
	}

	c.JSON(http.StatusOK, ret)
}

func (this *Base) Bupdate(c *gin.Context) {
	ret := make(map[string]string)

	params, errMsg := this.getAndCheckParams(c, []string{
		"opt",
		"ops",
	})

	if nil != errMsg {
		ret["state"] 	= 	APISTATUS_ERRPARAMS
		ret["msg"]		=	util.Join(errMsg, ",")
	}

	flag := 0
	var status int

	switch params["opt"] {
	case "topub":
		status = 0
	case "todrafts":
		status = 1
	case "totrash":
		status = 2
	case "delete":
		flag = 1
	default:
		flag = -1
	}

	if flag < 0 {
		ret["state"] = APISTATUS_ERRPARAMS
	}else if flag > 0 {
		dbRet, err := this.mysqlInstance().Exec(fmt.Sprintf("delete from `post` where id in (%s)", params["ops"]))
		if nil != err {
			ret["state"] = APISTATUS_MYSQLQUERY
		}else{
			arow, _ := dbRet.RowsAffected()
			if arow > 0 {
				ret["state"] = APISTATUS_OK
			}else{
				ret["state"] = APISTATUS_MYSQLAROW
			}
		}
	}else{
		dbRet, err := this.mysqlInstance().Exec(fmt.Sprintf("update `post` set status = %d where id in (%s)", status, params["ops"]))
		if nil != err {
			ret["state"] = APISTATUS_MYSQLQUERY
		}else{
			arow, _ := dbRet.RowsAffected()
			if arow > 0 {
				ret["state"] = APISTATUS_OK
			}else{
				ret["state"] = APISTATUS_MYSQLAROW
			}
		}
	}

	c.JSON(http.StatusOK, ret)
}
