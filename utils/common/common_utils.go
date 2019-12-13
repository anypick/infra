package common

import (
	"fmt"
	"reflect"
	"strings"
)

func StrIsNotBlank(str string) bool {
	if strings.TrimSpace(str) == "" {
		return false
	}
	return true
}

func StrIsBlank(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}

func InterfaceToStr(source interface{}) string {
	return fmt.Sprintf("%v", source)
}

func StructToMap(source interface{}) map[string]interface{} {
	typeOf := reflect.TypeOf(source)
	valueOf := reflect.ValueOf(source)
	count := valueOf.NumField()
	data := make(map[string]interface{}, count)
	for i := 0; i < count; i ++ {
		data[typeOf.Field(i).Name] = valueOf.Field(i).Interface()
	}
	return data
}
