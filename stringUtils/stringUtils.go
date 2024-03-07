package stringUtils

import "strings"

type StringUtils struct {
}

// Split 根据指定分割符，分割字符串
func (StringUtils) Split(content string, code string) []string {

	return strings.Split(content, code)

}
