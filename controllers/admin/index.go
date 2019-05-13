package admin

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"web/packs/gin"
	"web/packs/util"
)

type Index struct {
	Base
}

func (_ *Index) Index(c *gin.Context) {
	c.HTML(util.Ghttp_status_success, "admin/index", map[string]interface{}{
		"admin_name"	:		"admin",
		"version"		:		"1.0.1",
	})
}

func (_ *Index) Main(c *gin.Context) {
	c.HTML(util.Ghttp_status_success, "admin/main", map[string]interface{}{
		"app_ver"	:	"1.0.1",
		"hostname"	:	getHostName(),
		"go_ver"	:	runtime.Version(),
		"os"		:	runtime.GOOS,
		"cpu_num"	:	runtime.NumCPU(),
		"arch"		:	runtime.GOARCH,
		"postnum"	:	0,
		"tagnum"	:	0,
		"usernum"	:	0,
	})
}

func (this *Index) WebSet(c *gin.Context) {
	if this.isPost(c) {
		util.Gwebsetting.BantchSet(map[string]string{
			"webTitle"		:	c.DefaultPostForm("webtitle", ""),
			"webUrl"		:	c.DefaultPostForm("weburl", ""),
			"webSubTitle"	:	c.DefaultPostForm("websubtitle", ""),
			"webEmail"		:	c.DefaultPostForm("webemail", ""),
			"webPageSize"	:	c.DefaultPostForm("webpagenum", ""),
			"webKeyWord"	:	c.DefaultPostForm("webkeyword", ""),
			"webDesc"		:	c.DefaultPostForm("webdesc", ""),
			"webTheme"		:	c.DefaultPostForm("webtheme", ""),
		})

		querySql := "replace into `option` (`name`, `value`) values"

		refType := reflect.TypeOf(util.Gwebsetting)
		refVal  := reflect.ValueOf(util.Gwebsetting)
		times 	:= refType.NumField() - 1
		for i := 0; i < times; i++ {
			if refVal.Field(i).Kind() == reflect.Int {
				querySql += fmt.Sprintf("(\"%s\",\"%s\"),", refType.Field(i).Name, util.Int642str(refVal.Field(i).Int()))
			}else{
				querySql += fmt.Sprintf("(\"%s\",\"%s\"),", refType.Field(i).Name, refVal.Field(i).String())
			}
		}

		util.OnceTimerTask(3, func(){
			this.mysqlInstance().Exec(querySql[0:len(querySql) - 1])

		})
	}

	c.HTML(util.Ghttp_status_success, "admin/system", map[string]interface{}{
		"webtitle"		:		util.Gwebsetting.Get("webTitle"),
		"websubtitle"	:		util.Gwebsetting.Get("webSubTitle"),
		"weburl"		:		util.Gwebsetting.Get("webUrl"),
		"webemail"		:		util.Gwebsetting.Get("webEmail"),
		"webpagenum"	:		util.Gwebsetting.Get("webPageSize"),
		"webkeyword"	:		util.Gwebsetting.Get("webKeyWord"),
		"webdesc"		:		util.Gwebsetting.Get("webDesc"),
		"webtheme"		:		util.Gwebsetting.Get("webTheme"),
	})
}

/**
 * Private Begin.
 */
func getHostName() string {
	hname, err := os.Hostname()
	if nil != err {
		hname = "localhost"
	}

	return hname
}