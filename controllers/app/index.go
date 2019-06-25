package app

import (
	"fmt"
	"web/packs/gin"
	"web/packs/util"
)

type Index struct {
	Base
}

func (this *Base) Index(c *gin.Context) {
	totalNum :=	this.mysqlInstance().DefGetOne("select count(*) from `post`", "0")

	page := util.Str2int(c.DefaultQuery("page", "0"))

	_, list, _ := this.mysqlInstance().GetAll(fmt.Sprintf("select * from `post` limit %d,%d", this.pageOffset(page), this.pageSize()))

	pagebar := util.NewPager(page, util.Str2int(totalNum), this.pageSize(), "/", true).ToString()

	_, tags,_ := this.mysqlInstance().GetAll("select distinct tags from `post`")

	options := map[string]string{
		"sitename"		:		util.Gwebsetting.Get("webTitle"),
		"subtitle"		:		util.Gwebsetting.Get("webSubTitle"),
		"siteurl"		:		util.Gwebsetting.Get("webUrl"),
		"stat"			:		util.Gwebsetting.Get("webEmail"),
	}

	this.display(c, map[string]interface{}{
		"pagebar"		:		pagebar,
		"list"			:		list,
		"options"		:		options,
		"tagList"		:		tags,
	})
}

func (this *Base) Article(c *gin.Context) {
	var title, content string

	aid := c.DefaultQuery("aid", "-1")
	if "-1" != aid {
		row,  _ := this.mysqlInstance().GetRow("select title,content from `post` where id = " + aid)
		title  = (*row)["title"]
		content = (*row)["content"]

	}

	if "" == title || "" == content {
		this.errorShow(c, []string{"No Article!"})
		return
	}

	this.display(c, map[string]interface{}{
		"title"		:		title,
		"content"	:		content,
	})
}

func (this *Base) Category(c *gin.Context) {
	var list map[string]string
	tag := c.DefaultQuery("tag", "-1")
	if -1 != tag {
		totalNum, _ := this.mysqlInstance().GetOne(fmt.Sprintf("select count(*) from `post` where tags = '%s'", tag))
		page := util.Str2int(c.DefaultQuery())
		_, list, _ := this.mysqlInstance(fmt.Sprintf("select title,post_time from `post` where tags = '%s'", tag))
	}

	this.display(c, map[string]interface{}{
		"list"		:		list,
	})
}