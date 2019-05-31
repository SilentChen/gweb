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

	this.display(c, map[string]interface{}{
		"pagebar"		:		pagebar,
		"list"			:		list,
	})
}