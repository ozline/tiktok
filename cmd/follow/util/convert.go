package util

import (
	"reflect"

	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func ConvertStruct(source *user.User) *follow.User {
	sourceType := reflect.TypeOf(source).Elem() //源结构体类型
	target := follow.User{}
	targetType := reflect.TypeOf(target) //目标结构体类型
	// 遍历源结构体的字段
	for i := 0; i < sourceType.NumField(); i++ {
		// 获取源字段
		sourceField := sourceType.Field(i)
		// 在目标结构体中查找相同名称的字段
		targetField, ok := targetType.FieldByName(sourceField.Name)
		if ok {
			// 获取源结构体的字段值
			sourceValue := reflect.ValueOf(source).FieldByName(sourceField.Name)
			// 根据目标字段将源字段的值转换为目标类型并设置到目标结构体中
			targetField := reflect.ValueOf(&target).Elem().FieldByName(targetField.Name)
			targetField.Set(sourceValue)
		}
	}

	return &target
}
