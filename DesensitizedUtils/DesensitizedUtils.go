package DesensitizedUtils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type DesensitizedUtils struct {
}

func (DesensitizedUtils) ChineseName(name string) string {
	if name == "" {
		return ""
	}
	s := hide(name, 1, len(name)/3)
	return s
}

func (DesensitizedUtils) Email(email string) string {
	if email == "" {
		return ""
	}
	index := strings.Index(email, "@")
	emailValue := hide(email, 1, index)
	return emailValue
}

func (DesensitizedUtils) MobilePhone(phone string) string {
	if phone == "" {
		return ""
	}
	return hide(phone, 3, 7)
}

func (DesensitizedUtils) IdCard(idCard string, front int, end int) string {
	if idCard == "" {
		return ""
	}
	return hide(idCard, front, len(idCard)-end)
}

func (DesensitizedUtils) Address(address string, sensitiveSize int) string {
	if address == "" {
		return ""
	}
	length := len(strings.Split(address, ""))

	return hide(address, length-sensitiveSize, length)
}

func (DesensitizedUtils) PossWord(password string) string {
	if password == "" {
		return ""
	}
	return hide(password, 0, len(password))
}

//SecureStructList 结构体切片脱敏
/**
 *需要在结构体的tag标签备注stt[start:end],例如stt:"1:2",表示此字段需要脱敏从第0个字符到第2个字符
 *如果需要从第1个字符到最后一个都需要脱敏可以：stt:"1:a",反之，start=a表示0，end=a表示最后一个字符
 */
func (DesensitizedUtils) SecureStructList(structData []interface{}) []map[string]interface{} {

	frist := reflect.ValueOf(structData[0])
	b := frist.Kind() == reflect.Struct
	if !b {
		panic("不是一个结构体")
	}
	var list []map[string]interface{}
	for _, data := range structData {
		var m = make(map[string]interface{})
		structValue := reflect.ValueOf(data)
		for i := 0; i < structValue.NumField(); i++ {
			name := structValue.Type().Field(i).Name
			value := structValue.Field(i).Interface()
			stt := structValue.Type().Field(i).Tag.Get("stt")
			split := strings.Split(stt, ":")
			if len(split) != 2 {
				m[name] = value
			} else {
				fmt.Println("========", stt)
				start := 0
				end := 0
				if split[0] == "a" {
					start = 0
				} else {
					start, _ = strconv.Atoi(split[0])
				}
				if split[1] == "a" {
					end = len(strings.Split(value.(string), ""))
				} else {
					end, _ = strconv.Atoi(split[1])
				}
				m[name] = hide(value.(string), start, end)
			}

		}
		list = append(list, m)
	}

	return list

}

func hide(value string, start int, end int) string {

	split := strings.Split(value, "")
	for i := start; i < end; i++ {
		split[i] = "*"
	}
	return strings.Join(split, "")
}
