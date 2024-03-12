package wjdsgtools

import (
	"github.com/wjdsg0327/wjdsgtools/CalculateUtils"
	"github.com/wjdsg0327/wjdsgtools/DateUtils"
	"github.com/wjdsg0327/wjdsgtools/FileUtils"
	"github.com/wjdsg0327/wjdsgtools/IdCardUtils"
	"github.com/wjdsg0327/wjdsgtools/IdUtils"
	"github.com/wjdsg0327/wjdsgtools/InterfaceUtils"
	"github.com/wjdsg0327/wjdsgtools/RandomUtils"
	"github.com/wjdsg0327/wjdsgtools/SecureUtils"
	"github.com/wjdsg0327/wjdsgtools/StringUtils"
	"github.com/wjdsg0327/wjdsgtools/StructUtils"
	"wjdsgtools/DesensitizedUtils"
)

type Utils struct {
}

// DateUtils 日期处理
func (Utils) DateUtils() DateUtils.DateUtils {
	utils := DateUtils.DateUtils{}
	return utils
}

// IdUtils id生成处理
func (Utils) IdUtils() IdUtils.IdUtils {
	utils := IdUtils.IdUtils{}
	return utils
}

// RandomUtils 随机数处理
func (Utils) RandomUtils() RandomUtils.RandomUtils {
	utils := RandomUtils.RandomUtils{}
	return utils
}

// SecureUtil 加处理
func (Utils) SecureUtil() SecureUtils.SecureUtils {
	utils := SecureUtils.SecureUtils{}
	return utils
}

// StringUtils 字符串处理
func (Utils) StringUtils() StringUtils.StringUtils {
	utils := StringUtils.StringUtils{}
	return utils
}

// StructUtils 结构体处理
func (Utils) StructUtils() StructUtils.StructUtils {
	utils := StructUtils.StructUtils{}
	return utils
}

// CalculateUtils 计算操作
func (Utils) CalculateUtils() CalculateUtils.CalculateUtils {
	utils := CalculateUtils.CalculateUtils{}
	return utils
}

// FileUtils 文件操作
func (Utils) FileUtils() FileUtils.FileUtils {
	utils := FileUtils.FileUtils{}
	return utils
}

// InterfaceUtils 接口类型处理
func (Utils) InterfaceUtils() InterfaceUtils.InterfaceUtils {
	utils := InterfaceUtils.InterfaceUtils{}
	return utils
}

// IdCardUtils 身份证处理
func (Utils) IdCardUtils() IdCardUtils.IdCardUtils {
	utils := IdCardUtils.IdCardUtils{}
	return utils
}

// DesensitizedUtils 信息脱敏
func (Utils) DesensitizedUtils() DesensitizedUtils.DesensitizedUtils {
	utils := DesensitizedUtils.DesensitizedUtils{}
	return utils
}
