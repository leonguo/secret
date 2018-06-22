package util

import (
	"crypto/sha1"
	"io"
	"fmt"
	"encoding/base64"
	"crypto/hmac"
)

func randomSlat() string {
	tab := "abcdefghijklmnopqrstuvwxyz0123456789"
	return RandomSample(tab, 8)
}

func HMacSha1(password, slat string) []byte {
	mac := hmac.New(sha1.New, []byte(slat))
	encPwd := mac.Sum([]byte(password))
	return encPwd
}

func HashPassword(password string) string {
	h := sha1.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MakePwd(password string) string {
	slat := randomSlat()
	encPwd := HMacSha1(password, slat)
	return base64.StdEncoding.EncodeToString([]byte(slat + string(encPwd)))
}

/**
    检测密码是否正确
	pwd为原始密码(即用户输入的密码)
	encPwd为已经编码后的密码(即保存到数据库里面的密码)
 */
func CheckPwd(pwd, encPwd string) bool {
	decStr, err := base64.StdEncoding.DecodeString(encPwd)
	if err != nil {
		return false
	}
	slat := string(decStr[0:8])
	encStr := string(decStr[8:])
	return string(HMacSha1(pwd, slat)) == encStr
}
