package DateUtils

import (
	"time"
)

type DateUtils struct {
}

// ParseYMDHMS 将数据转换为年月日时分秒
func (DateUtils) ParseYMDHMS(date time.Time) string {

	return date.Format("2006-01-02 15:04:05")

}

// ParseYMD 将数据转换为年月日
func (DateUtils) ParseYMD(date time.Time) string {

	return date.Format("2006-01-02")

}

// BeginOfDay 获取指定日期的开始时间
func (DateUtils) BeginOfDay(date time.Time) time.Time {

	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// EndOfDay 获取指定日期的结束时间
func (DateUtils) EndOfDay(date time.Time) time.Time {

	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
}

// Offset 计算偏移后的日期时间
// @params date时间
// @params DateField需要计算的类型
// @params offset 未来时间就正整数，过去时间就负整数
func (DateUtils) Offset(date time.Time, field int, offset int) time.Time {

	if field == 1 {
		// 计算去年的时间
		return date.AddDate(offset, 0, 0)
	} else if field == 2 {
		return date.AddDate(0, offset, 0)
	} else if field == 3 {
		return date.AddDate(0, 0, offset)
	} else {
		return date
	}

}

// BeginOfMonth 获取指定月开始时间
func (DateUtils) BeginOfMonth(date time.Time) time.Time {

	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)

}

// EndOfMonth 获取指定月结束时间
func (DateUtils) EndOfMonth(date time.Time) time.Time {

	nextMonth := time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	return nextMonth.Add(-time.Nanosecond)

}
