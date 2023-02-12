package helpers

import (
	"fmt"
	"reflect"
)

func IsNil(data string) bool {
	return data == ""
}

func IsEmpty(data interface{}) bool {
	return data == nil || data == ""
}

func IsFunc(function interface{}, funcName string) bool {
	fmt.Println("IsFunc =", reflect.ValueOf(function).MethodByName(funcName).IsValid())
	method := reflect.ValueOf(function).MethodByName(funcName)

	if method.IsValid() {
		return true
	}

	return false
}
