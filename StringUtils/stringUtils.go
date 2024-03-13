package StringUtils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type StringUtils struct {
}

func (StringUtils) ChinaToPinYin(s string) string {

	a := pinyin.LazyConvert(s, nil)

	var test string

	for _, name := range a {

		one := name[0:1]
		upper := strings.ToUpper(one)
		end := name[1:len(name)]
		name = upper + end
		test = test + name
	}

	return test
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
func (StringUtils) JsonToSting(data interface{}) string {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Json转String错误：%s", err))
	}
	return string(jsonBytes)
}

// EncodeHexStr 字符串转hex
func (StringUtils) EncodeHexStr(s string) string {

	inputBytes := []byte(s)

	// 使用hex包进行编码
	hexString := hex.EncodeToString(inputBytes)

	return hexString
}

// DecodeHexStr hex转字符串
func (StringUtils) DecodeHexStr(s string) string {

	// 使用hex包进行编码
	decodeString, err := hex.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("解析失败：%s", err))
	}
	return string(decodeString)
}
