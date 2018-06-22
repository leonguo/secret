package util

// 判断是不是bmp文件
func IsBmp(body []byte) bool {
	var BMP byte = 66
	if body[0] == BMP {
		return true
	}
	return false
}