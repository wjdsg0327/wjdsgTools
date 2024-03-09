package IdCardUtils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type IdCardUtils struct {
}

// IsValidCard18 判断身份证是否合法
func (IdCardUtils) IsValidCard18(card string) bool {

	if len(card) != 18 {
		fmt.Println("======wjdsg======:此函数只能判断中国大陆第二代身份证18位")
		return false
	}
	code := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	cardArr := make([]int, len(card))

	for i, v := range card {
		num, _ := strconv.Atoi(string(v))
		cardArr[i] = num
	}
	sum := 0
	for i, i2 := range code {
		sum = sum + i2*cardArr[i]
	}
	i := sum % 11
	if i > 10 {
		return false
	}
	checkCode := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	s := checkCode[i]
	s2 := card[len(card)-1:]
	if s == strings.ToUpper(s2) {
		return true
	} else {
		return false
	}
}

// IdCardGetAge 根据身份证计算年龄
func (IdCardUtils) IdCardGetAge(card string) int {

	card18 := IdCardUtils{}.IsValidCard18(card)
	if !card18 {
		panic("身份证不合法")
	}

	birthdate := card[6:14]

	birthday, err := time.Parse("20060102", birthdate)

	if err != nil {

		panic(fmt.Sprintf("日期解析错误,身份证不合法:%s", err))
	}

	now := time.Now()
	age := now.Year() - birthday.Year()

	// 如果生日还未过，年龄减1
	if now.YearDay() < birthday.YearDay() {
		age--
	}

	return age
}

var area = map[int]string{
	11: "北京市",
	12: "天津市",
	13: "河北省",
	14: "山西省",
	15: "内蒙古自治区",
	21: "辽宁省",
	22: "吉林省",
	23: "黑龙江省",
	31: "上海市",
	32: "江苏省",
	33: "浙江省",
	34: "安徽省",
	35: "福建省",
	36: "江西省",
	37: "山东省",
	41: "河南省",
	42: "湖北省",
	43: "湖南省",
	44: "广东省",
	45: "广西壮族自治区",
	46: "海南省",
	50: "重庆市",
	51: "四川省",
	52: "贵州省",
	53: "云南省",
	54: "西藏自治区",
	61: "陕西省",
	62: "甘肃省",
	63: "青海省",
	64: "宁夏回族自治区",
	65: "新疆维吾尔自治区",
}

// IdCardGetAREA 根据身份证获取所属省份
func (IdCardUtils) IdCardGetAREA(card string) string {

	card18 := IdCardUtils{}.IsValidCard18(card)
	if !card18 {
		panic("身份证不合法")
	}
	areaCode := card[0:2]
	code, err := strconv.Atoi(areaCode)
	if err != nil {
		panic("身份证错误")
	}
	s := area[code]

	return s
}

// IdCardSEX 根据身份证获取性别
func (IdCardUtils) IdCardSEX(card string) string {

	card18 := IdCardUtils{}.IsValidCard18(card)
	if !card18 {
		panic("身份证不合法")
	}

	s := card[16:17]

	code, _ := strconv.Atoi(s)
	i := code % 2
	if i == 0 {
		return "女"
	}

	return "男"

}
