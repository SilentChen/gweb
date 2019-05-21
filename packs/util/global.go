package util

import (
	"sync"
	"time"
)

var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06",   //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01",      // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1",       // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan",     // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2",  // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon",    // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3",  // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

const (
	Gapp_mode				=	"dev"
	Gapp_ver				=	"1.0.0"
	Gapp_port				=	":8080"
	Gapp_db					=	"mysql"
	Gmysql_host				=	"106.12.130.31"
	Gmysql_port				=	"3306"
	Gmysql_user				=	"blog"
	Gmysql_pwd				=	"qwer"
	Gmysql_dbname			=	"blog"
	Gmysql_max_idle_conns	=	"100"
	Gmysql_max_open_conns	=	"200"
	Gmysql_charset			=	"utf8"
	Gtime_layout			=	"2006-01-02 15:04:05"
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