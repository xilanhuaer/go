package assert

import (
	"reflect"
	"strings"
)

type AssertService struct {
}

// Equals compares two objects and returns true if they are equal
// param expected: the expected object
// param actual: the actual object
// return: true if the objects are equal, false otherwise
func (assertService *AssertService) DeepEquals(expected, actual interface{}) bool {
	return reflect.DeepEqual(expected, actual)
}

// IsNull returns true if the object is null
// param expected: the expected object
// return: true if the object is null, false otherwise
func (assertService *AssertService) IsNUll(expected interface{}) bool {
	return expected == nil
}

// IsNotNull returns true if the object is not null
// param expected: the expected object
// return: true if the object is not null, false otherwise
func (assertService *AssertService) IsNotNull(expected interface{}) bool {
	return expected != nil
}
func (assertService *AssertService) IsEmpty(path string, expected interface{}) bool {
	// 解析路径
	paths := strings.Split(path, ".")
	// 判断对象的类型是否为map
	if reflect.TypeOf(expected).Kind() == reflect.Map {
		// 将对象转换为map对象
		expectedMap := expected.(map[string]interface{})
		// 判断map对象是否包含路径中的所有key
		for _, key := range paths {
			if _, ok := expectedMap[key]; !ok {
				return false
			}
		}
		return true
	}

	return true
}
