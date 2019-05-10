package util

import (
	"sync"
)

const (
	Gapp_mode				=	"dev"
	Gapp_ver				=	"1.0.0"
	Gapp_port				=	":8080"
	Gmysql_host				=	"106.12.130.31"
	Gmysql_port				=	"3306"
	Gmysql_user				=	"blog"
	Gmysql_pwd				=	"qwer"
	Gmysql_dbname			=	"blog"
	Gmysql_max_idle_conns	=	"100"
	Gmysql_max_open_conns	=	"200"
	Gmysql_charset			=	"utf8"
	Gtime_layout			=	"2006-01-02 15:04:05"
	Ghttp_status_success	=	200
)

type WebSetting struct {
	webTitle 	string
	webUrl  	string
	webSubTitle	string
	webEmail	string
	webPageSize	int
	webKeyWord	string
	webDesc		string
	webTheme	string
	webTimeZone	string
	lock		sync.Mutex
}

// some global setting
var Gwebsetting WebSetting

func init() {
	Gwebsetting.webTitle	= 	"SilentBlog"
	Gwebsetting.webUrl 		= 	"http://127.0.0.1:8080"
	Gwebsetting.webSubTitle = 	"PersionalBlog"
	Gwebsetting.webEmail	=	"silent@go.com"
	Gwebsetting.webPageSize	=	10
	Gwebsetting.webKeyWord	=	"Silent, SilentBlog"
	Gwebsetting.webDesc		=	"SilentChenPersionalBlog"
	Gwebsetting.webTheme	=	"default"
	Gwebsetting.webTimeZone	=	"Asia/Shanghai"
}

func (ws *WebSetting) Set(key string, val string) {
	ws.lock.Lock()
	defer ws.lock.Unlock()

	updateWs(key, val)
}

func (ws *WebSetting) BantchSet(params map[string]string) {
	ws.lock.Lock()
	defer  ws.lock.Unlock()

	for k, v := range params {
		updateWs(k, v)
	}
}

func (_ *WebSetting)Get(key string) string {
	var ret string

	switch key {
	case "webTitle":
		ret = Gwebsetting.webTitle
		break
	case "webUrl":
		ret = Gwebsetting.webUrl
		break
	case "webSubTitle":
		ret = Gwebsetting.webSubTitle
		break
	case "webEmail":
		ret = Gwebsetting.webEmail
		break
	case "webPageSize":
		ret = Int2str(Gwebsetting.webPageSize)
		break
	case "webKeyWord":
		ret = Gwebsetting.webKeyWord
		break
	case "webDesc":
		ret = Gwebsetting.webDesc
		break
	case "webTheme":
		ret = Gwebsetting.webTheme
		break
	case "webTimeZone":
		ret = Gwebsetting.webTimeZone
		break
	default:
		ret = ""
		break
	}

	return ret
}

// internal function
func updateWs(key string, val string) {
	switch key {
	case "webTitle":
		Gwebsetting.webTitle = val
		break
	case "webUrl":
		Gwebsetting.webUrl = val
		break
	case "webSubTitle":
		Gwebsetting.webSubTitle = val
		break
	case "webEmail":
		Gwebsetting.webEmail = val
		break
	case "webPageSize":
		Gwebsetting.webPageSize = Str2int(val)
		break
	case "webKeyWord":
		Gwebsetting.webKeyWord = val
		break
	case "webDesc":
		Gwebsetting.webDesc = val
		break
	case "webTheme":
		Gwebsetting.webTheme = val
		break
	case "webTimeZone":
		Gwebsetting.webTimeZone = val
		break
	default:
		break
	}
}