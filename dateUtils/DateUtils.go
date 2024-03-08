package DateUtils

import (
	"fmt"
	"strconv"
	"time"
)

type DateUtils struct {
}

// ParseYMDHMS 将数据转换为年月日时分秒
/**
* @params date：时间
 *@params format：格式 2006-01-02 15:04:05
*/
func (DateUtils) ParseYMDHMS(date time.Time, format string) string {

	return date.Format(format)

}

// ParseYMD 将数据转换为年月日
/**
 *@params date:时间
 *@params format：格式 2006-01-02 15:04:05
 */
func (DateUtils) ParseYMD(date time.Time, format string) string {

	return date.Format(format)

}

// Format 字符串时间转time
/**
 *@params dateTimeStr:时间
 *@params format：格式 2006-01-02 15:04:05
 */
func (DateUtils) Format(dateTimeStr string, format string) time.Time {

	dateTime, err := time.Parse(format, dateTimeStr)

	if err != nil {
		fmt.Println("转换失败:", err)
		panic("转换失败")
	}
	return dateTime

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
/**
 *@params date时间
 *@params DateField需要计算的类型
 *@params offset 未来时间就正整数，过去时间就负整数
 */
func (DateUtils) Offset(date time.Time, field int, offset int) time.Time {

	if field == 1 {
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

// Between 计算两个时间之间的年月日
func (DateUtils) Between(t1 time.Time, t2 time.Time, DateField YMD) int {
	//days := int(t2.Sub(t1).Hours() / 24)
	duration := t2.Sub(t1)
	var date int
	var err error
	switch DateField {
	case 1:
		n := fmt.Sprintf("%.0f", duration.Hours()/24/365)
		date, err = strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("计算错误：%s", err))
		}
		return date
	case 2:
		n := fmt.Sprintf("%.0f", duration.Hours()/24/31)
		date, err = strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("计算错误：%s", err))
		}
		return date
	case 3:
		n := fmt.Sprintf("%.0f", duration.Hours()/24)
		date, err = strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("计算错误：%s", err))
		}
		return date
	default:
		panic("类型传参错误：DateField")
	}
}

// GetBetweenDate 获取两个日期之间的所有日期，格式：yyyy-MM-dd
func (DateUtils) GetBetweenDate(start, end string) []string {
	var list []string

	// 时间格式为 yyyy-MM-dd
	startDate, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)

	distance := endDate.Sub(startDate).Hours() / 24
	if distance < 1 {
		list = append(list, start)
		return list
	}

	for d := startDate; d.Before(endDate.Add(24 * time.Hour)); d = d.Add(24 * time.Hour) {
		list = append(list, d.Format("2006-01-02"))
	}

	return list
}

// Yesterday 获取昨天的日期
func (DateUtils) Yesterday() time.Time {
	date := time.Now()
	t := time.Date(date.Year(), date.Month(), date.Day()-1, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return t
}

// Tomorrow 获取昨天的日期
func (DateUtils) Tomorrow() time.Time {
	date := time.Now()
	t := time.Date(date.Year(), date.Month(), date.Day()+1, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return t
}

// LastWeek 获取上周的日期
func (DateUtils) LastWeek() time.Time {
	date := time.Now()
	t := time.Date(date.Year(), date.Month(), date.Day()-7, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return t
}

// NextWeek 获取下周的日期
func (DateUtils) NextWeek() time.Time {
	date := time.Now()
	t := time.Date(date.Year(), date.Month(), date.Day()+7, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return t
}

// LastMonth 获取上个月的日期
func (DateUtils) LastMonth() time.Time {
	date := time.Now()
	t := time.Date(date.Year(), date.Month()-1, date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return t
}

// NextMonth 获取下个月的日期
func (DateUtils) NextMonth() time.Time {
	date := time.Now()
	t := time.Date(date.Year(), date.Month()+1, date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return t
}

// GetZodiac 根据月日计算星座
func (DateUtils) GetZodiac(mouth time.Month, day int) string {

	switch mouth {
	case 1:
		if day >= 20 {
			return "水瓶座"
		}
		return "摩羯座"
	case 2:
		if day >= 19 {
			return "双鱼座"
		}
		return "水瓶座"
	case 3:
		if day >= 21 {
			return "白羊座"
		}
		return "双鱼座"
	case 4:
		if day >= 20 {
			return "金牛座"
		}
		return "白羊座"
	case 5:
		if day >= 21 {
			return "双子座"
		}
		return "金牛座"
	case 6:
		if day >= 21 {
			return "巨蟹座"
		}
		return "双子座"
	case 7:
		if day >= 23 {
			return "狮子座"
		}
		return "巨蟹座"
	case 8:
		if day >= 23 {
			return "处女座"
		}
		return "狮子座"
	case 9:
		if day >= 23 {
			return "天秤座"
		}
		return "处女座"
	case 10:
		if day >= 23 {
			return "天蝎座"
		}
		return "天秤座"
	case 11:
		if day >= 22 {
			return "射手座"
		}
		return "天蝎座"
	case 12:
		if day >= 22 {
			return "摩羯座"
		}
		return "射手座"
	default:
		return "未知星座"
	}
}

// AgeOfNow 计算年龄
func (DateUtils) AgeOfNow(birthdate string) int {
	layout := "2006-01-02"
	birthday, err := time.Parse(layout, birthdate)

	if err != nil {
		fmt.Println("日期解析错误，需要格式是：yyyy-MM-dd:", err)
		return 0
	}

	now := time.Now()
	age := now.Year() - birthday.Year()

	// 如果生日还未过，年龄减1
	if now.YearDay() < birthday.YearDay() {
		age--
	}

	return age
}
