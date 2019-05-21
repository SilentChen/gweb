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

	id := c.Query("id")

	if "" != id {
		row, err := this.mysqlInstance().GetRow(fmt.Sprintf("select * from `post` where id = %s", id))
		if nil == err {
			ainfo = *row
		}
	}

	if this.isPost(c) {
		var querySql, id string
		params, errMsg := this.getAndCheckParams(c, []string{"title", "color", "is_top", "tags", "url_name", "post_time", "status", "content", "url_type"})
		if len(errMsg) > 0 {
			this.errorShow(errMsg)
		}
	}

	log.Println(ainfo)

	c.HTML(http.StatusOK, "admin/article/edit",map[string]interface{}{
		"post"		:		ainfo,
	})
}
