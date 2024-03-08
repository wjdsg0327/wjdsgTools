package InterfaceUtils

import "reflect"

type InterfaceUtils struct {
}

// GetType 获取接口实际类型
func (InterfaceUtils) GetType(any interface{}) string {

	r := reflect.TypeOf(any).Kind().String()

	return r
}

// IsType 检查接口的类型是否为某个特定类型
func (InterfaceUtils) IsType(i interface{}, targetType reflect.Kind) bool {
	return reflect.TypeOf(i).Kind() == targetType
}
