package stringUtils

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StringUtils struct {
}

// Split 根据指定分割符，分割字符串
func (StringUtils) Split(content string, code string) []string {

	return strings.Split(content, code)

}

// StringToJson 字符串转json：map或者结构体
/**
 * @params jsonValue ：json字符串数据
 * @params dst ：接收数据的载体，最好是map或者结构体
 */
func (StringUtils) StringToJson(jsonValue string, dst interface{}) interface{} {

	err := json.Unmarshal([]byte(jsonValue), dst)
	if err != nil {
		panic(fmt.Sprintf("解析失败:%s", err))
	}

	return dst
}

// JsonToSting json转字符串
/**
 * @params data ：json数据
 */
func JsonToSting(data interface{}) string {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Json转String错误：%s", err))
	}
	return string(jsonBytes)
}
