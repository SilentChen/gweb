package admin

import (
	"fmt"
	"log"
	"net/http"
	"web/packs/gin"
	"web/packs/util"
)

type Article struct {
	Base
}

func (this *Base) List(c *gin.Context) {
	searchtype := c.DefaultPostForm("searchtype", "title")
	status := c.DefaultQuery("status", "0")

	count1 := this.mysqlInstance().DefGetOne("select count(*) from `post` where status = 1", "0")
	count2 := this.mysqlInstance().DefGetOne("select count(*) from `post` where status = 2", "0")

	page := util.Str2int(c.DefaultQuery("page", "0"))

	count, articles, _ := this.mysqlInstance().GetAll(fmt.Sprintf("select ? from `post` where status = %s limit %d,%d", status, this.pageOffset(page), this.pageSize()), "*")

	pagebar := util.NewPager(page, count, this.pageSize(), "/admin/article/list", true).ToString()

	c.HTML(http.StatusOK, "admin/article/list", map[string]interface{}{
		"list"			:		*articles,
		"pagebar"		:		pagebar,
		"status"		:		status,
		"searchtype"	:		searchtype,
		"count1"		:		count1,
		"count2"		:		count2,
	})
}

func (this *Base) Edit(c *gin.Context) {
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
	gmsg := make(map[string]string)
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
		log.Println(params)
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
			gmsg["msg"] 	= 	"success !"
			gmsg["color"]	=	"success"
		}else{
			gmsg["msg"] 	= 	"fail, please try again"
			gmsg["color"]	=	"error"
		}
	}

	c.HTML(http.StatusOK, "admin/article/edit",map[string]interface{}{
		"post"		:		ainfo,
		"desc"		:		desc,
		"gmsg"		:		gmsg,
	})
}
