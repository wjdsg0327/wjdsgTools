package wjdsgtools

import (
	DateUtils "wjdsgtools/dateUtils"
	IdUtils "wjdsgtools/idUtils"
	Md5Utils "wjdsgtools/md5Utils"
	RandomUtils "wjdsgtools/randomUtils"
	StringUtils "wjdsgtools/stringUtils"
	StructUtils "wjdsgtools/structUtils"
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
