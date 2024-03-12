package DesensitizedUtils

import (
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

func (DesensitizedUtils) Struct(stt interface{}) {

}

func hide(value string, start int, end int) string {

	split := strings.Split(value, "")
	for i := start; i < end; i++ {
		split[i] = "*"
	}
	return strings.Join(split, "")
}
