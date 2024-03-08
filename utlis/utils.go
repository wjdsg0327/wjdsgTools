package wjdsgtools

import (
	"github.com/wjdsg0327/wjdsgtools/CalculateUtils"
	"github.com/wjdsg0327/wjdsgtools/DateUtils"
	"github.com/wjdsg0327/wjdsgtools/FileUtils"
	"github.com/wjdsg0327/wjdsgtools/IdUtils"
	"github.com/wjdsg0327/wjdsgtools/Md5Utils"
	"github.com/wjdsg0327/wjdsgtools/RandomUtils"
	"github.com/wjdsg0327/wjdsgtools/StringUtils"
	"github.com/wjdsg0327/wjdsgtools/StructUtils"
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

// Md5Utils 加处理
func (Utils) Md5Utils() Md5Utils.Md5Utils {

	utils := Md5Utils.Md5Utils{}
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
