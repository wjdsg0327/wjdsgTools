package wjdsgtools

import (
	DateUtils "github.com/wjdsg0327/wjdsgtools/dateUtils"
	IdUtils "github.com/wjdsg0327/wjdsgtools/idUtils"
	Md5Utils "github.com/wjdsg0327/wjdsgtools/md5Utils"
	RandomUtils "github.com/wjdsg0327/wjdsgtools/randomUtils"
	StringUtils "github.com/wjdsg0327/wjdsgtools/stringUtils"
	StructUtils "github.com/wjdsg0327/wjdsgtools/structUtils"
)

type Utils struct {
}

func DateField() DateUtils.DateField {
	field := DateUtils.DateField{}
	return field
}

func (Utils) DateUtils() DateUtils.DateUtils {
	utils := DateUtils.DateUtils{}
	return utils
}

func (Utils) IdUtils() IdUtils.IdUtils {
	utils := IdUtils.IdUtils{}
	return utils
}

func (Utils) RandomUtils() RandomUtils.RandomUtils {
	utils := RandomUtils.RandomUtils{}
	return utils
}

func (Utils) Md5Utils() Md5Utils.Md5Utils {

	utils := Md5Utils.Md5Utils{}
	return utils
}

func (Utils) StringUtils() StringUtils.StringUtils {
	utils := StringUtils.StringUtils{}
	return utils
}

func (Utils) StructUtils() StructUtils.StructUtils {
	utils := StructUtils.StructUtils{}
	return utils
}
