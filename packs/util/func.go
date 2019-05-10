package util

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func Date2unix(date string) int64 {
	timezone, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(Gtime_layout, date, timezone)

	return tmp.Unix()
}

func Unix2date(utime int64) string {
	return time.Unix(utime, 0).Format(Gtime_layout)
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
