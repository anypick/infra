package common

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
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
	for i := 0; i < count; i++ {
		data[typeOf.Field(i).Name] = valueOf.Field(i).Interface()
	}
	return data
}

// 获取绝对路径
func GetAbsolutePath(fileName string, skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	// 解析出文件路径
	dir := filepath.Dir(file)
	return filepath.Join(dir, fileName)
}
