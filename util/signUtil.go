package util

import (
	"sort"
	"bytes"
	"github.com/labstack/echo"
	"strings"
	"net/url"
	"io/ioutil"
	"github.com/davecgh/go-spew/spew"
	"../db/redis"
	"encoding/base64"
	"strconv"
	"time"
)

//const AppKey = "focaltv"
const AppSecret = "16c86816ab0cfa1493da230cdf356478"

type headerSort struct {
	Keys   []string
	Values []string
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

/**
	URL签名规则
	uri + args + headers + body + app_key
  */

func SignAuth(c echo.Context) (check bool, err error) {
	check = false
	params := make(map[string]string)
	uri := c.Request().RequestURI
	uriDetail, err := url.Parse(uri)
	spew.Dump(uriDetail)
	if err != nil {
		return
	}
	params["path"] = uriDetail.Path
	//params["Method"] = c.Request().Method
	params["args"] = c.QueryParams().Encode()
	checkSignStr, canonicalHeaders, signedHeaders := headerParams(c)
	if checkSignStr == "" {
		return
	}
	// get body
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		// Restore
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	signStr := params["path"] + "\n" +
		params["args"] + "\n" +
		canonicalHeaders + "\n" +
		signedHeaders + "\n" +
		Sha1hex(bodyBytes) + "\n" +
		AppSecret
	spew.Dump(signStr)
	hashSignStr := Sha1hex([]byte(signStr))
	spew.Dump(hashSignStr)
	if hashSignStr == checkSignStr {
		check = true
	}
	return
}

func headerParams(c echo.Context) (checkSignStr string, headerValues string, signedHeaders string) {
	hs := &headerSort{}
	header := c.Request().Header
	for k := range header {
		k = strings.ToLower(k)
		if k == "x-ft-sign" {
			checkSignStr = header.Get(k)
			continue
		}
		isSignKey := strings.HasPrefix(k, "x-ft-")
		if isSignKey {
			hs.Keys = append(hs.Keys, k)
			hs.Values = append(hs.Values, header.Get(k))
		}
	}
	hs.Sort()
	for i, k := range hs.Keys {
		if i != 0 {
			headerValues += "\n"
		}
		headerValues += k + "=" + strings.TrimSpace(hs.Values[i])
	}
	signedHeaders = strings.Join(hs.Keys, ";")
	return
}

// test
func AuthorizationHeader(c echo.Context) (bool, error) {
	return true, nil
}

func TokenSave(userId int64, username string) (token string, err error) {
	// TODO 有效期
	//expires := time.Now().Unix() + int64(time.Duration.Seconds())
	str := []byte(strconv.FormatInt(userId, 10) + ":" + username + ":" + strconv.FormatInt(time.Now().Unix(), 10))
	token = base64.StdEncoding.EncodeToString(str)
	key := "token:" + token
	err = redis.SetValue(key, userId)
	return
}

func TokenCheck(c echo.Context, token string) bool {
	userId, err := redis.GetValue("token:" + token)
	if err != nil {
		return false
	}
	if userId.(int64) > 0 {
		return true
	}
	return true
}
