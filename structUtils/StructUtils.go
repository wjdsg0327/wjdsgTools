package structUtils

import (
	"errors"
	"fmt"
	"reflect"
)

type StructUtils struct {
}

// CopyProperties 结构体转换
/**
* @params target：需要转换的目标
* @params source: 数据源
* 注意：必须传结构体指针&
 */
func (StructUtils) CopyProperties(target, source interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(target), reflect.ValueOf(target)
	srcType, srcValue := reflect.TypeOf(source), reflect.ValueOf(source)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}

// GetStructInfoList 获取结构体所有信息
/**
* @params structValue：结构体
 */
func (StructUtils) GetStructInfoList(structValue interface{}) []reflect.StructField {

	structType := reflect.TypeOf(structValue)

	b := structType.Kind() == reflect.Struct
	if !b {
		panic("传值不是一个结构体")
	}
	var list []reflect.StructField
	for i := 0; i < structType.NumField(); i++ {
		list = append(list, structType.Field(i))
	}

	return list
}

// GetStructFieldList 获取结构体所有字段
/**
* @params structValue：结构体
 */
func (StructUtils) GetStructFieldList(structValue interface{}) []string {

	structType := reflect.TypeOf(structValue)

	b := structType.Kind() == reflect.Struct
	if !b {
		panic("传值不是一个结构体")
	}
	var list []string
	//遍历结构体的字段
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		list = append(list, field.Name)
	}
	return list
}

// GetStructFieldInfo 根据结构体字段名字获取字段的详细
/**
 * @params structValue：结构体
 * @params field 字段名字
 */
func (StructUtils) GetStructFieldInfo(structValue interface{}, fieldName string) reflect.StructField {
	structType := reflect.TypeOf(structValue)

	b := structType.Kind() == reflect.Struct
	if !b {
		panic("传值不是一个结构体")
	}
	value, b2 := structType.FieldByName(fieldName)
	if !b2 {
		panic("没有此字段")
	}
	return value
}
