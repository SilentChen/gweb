package app

import (
	"fmt"
	"log"
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
	log.Println(list)
	this.display(c, map[string]interface{}{
		"pagebar"		:		pagebar,
		"list"			:		list,
		"options"		:		options,
		"tagList"		:		tags,
	})
}