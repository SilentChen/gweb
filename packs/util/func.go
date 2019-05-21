package util

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func Int2str(num int) (string) {
	return strconv.Itoa(num)
}

func Str2int(str string) (int) {
	num, err := strconv.Atoi(str)
	if nil != err {
		num = 0
	}

	return  num
}

func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

/**
 * @param str, string number.
 * @param defaultNum, return this when str conv into in err.
 * @return int
 */
func Str2int2(str string, defaultNum int) (int) {
	num, err := strconv.Atoi(str)

	if nil != err {
		num = defaultNum
	}

	return num
}

func Int642str(num int64) (string) {
	return strconv.FormatInt(num,10)
}

func Str2int64 (str string) (int64) {
	num, err := strconv.ParseInt(str, 10, 64)

	if nil != err {
		num = 0
	}

	return num
}

func Str2int642 (str string, defaultNum int64) (int64) {
	num, err := strconv.ParseInt(str, 10, 64)

	if nil != err {
		num = defaultNum
	}

	return num
}

func DateFormat(t time.Time, format string) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return t.Format(format)
}

func Date2unix(date string) int64 {
	timezone, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(Gtime_layout, date, timezone)

	return tmp.Unix()
}

func Unix2date(utime int64) string {
	return time.Unix(utime, 0).Format(Gtime_layout)
}

func Str2byte(str string) []byte {
	return []byte(str)
}

func Byte2str(b []byte) string {
	return string(b)
}

func Byte2int(b []byte)int{
	var ret int = 0
	var len int = len(b)
	var i uint = 0
	for i=0; i<uint(len); i++{
		ret = ret | (int(b[i]) << (i*8))
	}
	return ret
}

func Int2byte(i int) []byte {
	var len uintptr = unsafe.Sizeof(i)
	ret := make([]byte, len)
	var tmp int = 0xff
	var index uint = 0
	for index=0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index*8) & i) >> (index*8))
	}

	return ret
}

func OnceTimerTask(second time.Duration, f func()) {
	timer := time.NewTimer(time.Second * second)
	go func() {
		<- timer.C
		f()
	}()
}

func Request(reqType string, url string, params string, headers map[string]string) (string,error) {
	ret := ""
	client := http.Client{}
	req, err := http.NewRequest(strings.ToUpper(reqType), url, strings.NewReader(params))
	if nil != err {
		return ret, err
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if nil != err {
		return ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return ret, err
	}

	ret = string(body)

	return ret,nil
}
