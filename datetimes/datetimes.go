package datetimes

import (
	"errors"
	"fmt"
	"time"
)

// 获取日期，YYYY-MM-DD，根据不同参数返回不同呈现形式
func GetDay(opts ...int) string {
	if len(opts) == 0 {
		return time.Now().Format("2006-01-02")
	}
	t := opts[0]
	switch t {
	case 1:
		return time.Now().Format("20060102")
	case 2:
		return time.Now().Format("2006年01月02日")
	case 3:
		return time.Now().Format("2006/01/02")
	case 4:
		return time.Now().Format("01/02/2006")
	}

	return time.Now().Format("2006-01-02")
}

// 获取日期，YYYY-MM-DD HH:ii:ss，根据不同参数返回不同呈现形式
func GetNowDate(opts ...int) string {
	if len(opts) == 0 {
		return time.Now().Format("2006-01-02 15:04:05")
	}

	t := opts[0]
	switch t {
	case 1:
		return time.Now().Format("20060102150405")
	case 2:
		return time.Now().Format("2006年01月02日 15时04分05秒")
	case 3:
		return time.Now().Format("2006/01/02 15/04/05")
	case 4:
		return time.Now().Format("01/02/2006 15/04/05")
	}

	return time.Now().Format("2006-01-02 15:04:05")
}

// 日期时间转换成时间戳
// _date 时间内容
// format 转换模板
func DateToTimes(_date string, format string) (int64, error) {
	var date string
	date = _date

	// 转化所需内定模板
	var layout string
	if format == "Y-m-d H:i:s" || format == "" {
		layout = "2006-01-02 15:04:05"
	} else if format == "YmdHis" {
		layout = "20060102150405"
	} else if format == "Y年m月d日 H:i:s" {
		layout = "2006年01月02日 15:04:05"
	} else if format == "Y年m月d日 H时i分s秒" {
		layout = "2006年01月02日 15时04分05秒"
	} else if format == "Y-m-d" {
		layout = "2006-01-02"
	} else if format == "Ymd" {
		layout = "20060102"
	} else {
		layout = "2006-01-02 15:04:05"
	}

	// 日期转换为时间戳
	local, _ := time.LoadLocation("Local") // 获取时区
	tmp, err := time.ParseInLocation(layout, date, local)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("转换失败, 错误信息=%s", err))
	}
	// 转化为时间戳 类型是int64
	timestamp := tmp.Unix()

	return timestamp, nil
}
