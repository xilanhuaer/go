package utils

import "encoding/json"

// 将字符串转换为map对象
// param str: 字符串
// return: map对象
func StringToMap(str string) (map[string]interface{}, error) {
	var result map[string]interface{}
	// 将字符串转换为map对象
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
