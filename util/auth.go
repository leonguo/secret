package util

import (
	"github.com/labstack/echo"
	"strings"
	"sort"
	"bytes"
	"net/url"
	"crypto/sha1"
)

type headerSort struct {
	Keys   []string
	Values []string
}

const Secret = "xxf223"

/**
	URL签名规则
	参数 url 请求参数 access_key timestamp platform
  */

func SignAuth(c echo.Context) (check bool, err error) {
	params := make(map[string]string)
	uri := c.Request().RequestURI
	check = false
	u, err := url.Parse(uri)
	if err != nil {
		return
	}
	paramStr := ""
	params["Path"] = u.Path
	params["Method"] = c.Request().Method
	hs, checkSignStr := HeaderParams(c)
	if checkSignStr == "" {
		return
	}
	if len(hs.Keys) > 0 {
		for i, param := range hs.Keys {
			params[param] = hs.Values[i]
		}
	}
	params["SecretKey"] = Secret
	for k , v := range params {
		paramStr += k + "=" + v
	}
	h := sha1.New()
	h.Write([]byte(paramStr))
	signStr := string(h.Sum(nil))
	if signStr == checkSignStr {
		check = true
	}
	c.Logger().Printf("paramStr >>> %v", paramStr)
	c.Logger().Printf("signStr >>> %x", signStr)
	return
}

func HeaderParams(c echo.Context) (hs *headerSort, checkSignStr string) {
	hs = &headerSort{}
	header := c.Request().Header
	for k := range header {
		if k == "E-Sign" {
			checkSignStr = header.Get(k)
			continue
		}
		isSignKey := strings.HasPrefix(k, "E-")
		if isSignKey {
			hs.Keys = append(hs.Keys, k)
			hs.Values = append(hs.Values, header.Get(k))
		}
	}
	hs.Sort()
	return
}

func (hs *headerSort) Sort() {
	sort.Sort(hs)
}

// Additional function for function SignHeader.
func (hs *headerSort) Len() int {
	return len(hs.Values)
}

// Additional function for function SignHeader.
func (hs *headerSort) Less(i, j int) bool {
	return bytes.Compare([]byte(hs.Keys[i]), []byte(hs.Keys[j])) < 0
}

// Additional function for function SignHeader.
func (hs *headerSort) Swap(i, j int) {
	hs.Values[i], hs.Values[j] = hs.Values[j], hs.Values[i]
	hs.Keys[i], hs.Keys[j] = hs.Keys[j], hs.Keys[i]
}
