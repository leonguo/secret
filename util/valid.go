package util

import "regexp"

// 验证手机号码格式
func IsValidNumber(phone string) bool {
	if check, _ := regexp.MatchString("^([+]?\\d{1,4}[-\\s]?|)\\d{3}[-\\s]?\\d{3}[-\\s]?\\d{4}", phone); !check {
		return false
	}
	return true
}

// 检查map Key值是否存在
func CheckMapKeyExist(m map[string]interface{}, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
