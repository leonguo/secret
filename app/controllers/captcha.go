package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/mojocn/base64Captcha"
	"strings"
	"io/ioutil"
	"regexp"
)

type configJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigCharacter base64Captcha.ConfigCharacter
}

func GenerateCaptcha(c echo.Context) error {
	//创建base64图像验证码
	//config struct for Character
	//字符,公式,验证码配置
	var config = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  194,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     true,
		CaptchaLen:         6,
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha("", config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	body := map[string]interface{}{"code": 1, "data": base64Png, "captchaId": captchaId, "msg": "success"}
	return c.JSON(http.StatusOK, body)
}

type verifyBody struct {
	Id          string `json:"id"`
	VerifyValue string `json:"verify_value"`
	Phone       string `json:"phone"`
	Region      string `json:"region"`
}

func AttrUpdate(c echo.Context) error {
	var postParameters verifyBody
	if err := c.Bind(&postParameters); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	c.Logger().Debug(">>>> request xx ", postParameters)
	province := postParameters.Region
	cookieString := "ace_settings=%7B%22navbar-fixed%22%3A1%2C%22sidebar-fixed%22%3A1%2C%22breadcrumbs-fixed%22%3A1%7D; BIGipServerpool_SD_DaiLiShang=755060746.36895.0000; JSESSIONID=4A5F238C9ACAD3036A12B665CFDCB68C"
	client := &http.Client{}
	formData := "mercOprMbl="+postParameters.Phone
	req, err := http.NewRequest("POST", "http://119.18.194.36/miniIf/MinilistSs", strings.NewReader(formData))
	if err != nil {
		c.Logger().Error("new request error ", err)
		return c.JSON(http.StatusInternalServerError, "")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookieString)
	req.Header.Set("User-Agent", "Mozilla/5.0(Macintosh;IntelMacOSX10_7_0)AppleWebKit/535.11(KHTML,likeGecko)Chrome/17.0.963.56Safari/535.11")
	req.Header.Set("Referer", "http://119.18.194.36/miniIf/minilistViewSs.do")
	req.Header.Set("Origin", "http://119.18.194.36")
	req.Header.Set("Host", "119.18.194.36")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		c.Logger().Error("new request error  do >>> ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger().Error("new request error  read body  >>>", err)
	}
	c.Logger().Debug("body code is  ", resp.StatusCode)
	reg := regexp.MustCompile("\\d{15}") //连续的数字
	mercId := reg.FindString(string(body))
	c.Logger().Debug("body is ", mercId)
	if len(mercId) == 15 {
		b := "province="+province+"&mercId="+mercId
		req1, err := http.NewRequest("POST", "http://119.18.194.36/miniIf/miniSavePlaceData", strings.NewReader(b))
		if err != nil {
			c.Logger().Error("new request error ", err)
			return c.JSON(http.StatusInternalServerError, "")
		}
		req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req1.Header.Set("Cookie", cookieString)
		req1.Header.Set("User-Agent", "Mozilla/5.0(Macintosh;IntelMacOSX10_7_0)AppleWebKit/535.11(KHTML,likeGecko)Chrome/17.0.963.56Safari/535.11")
		req1.Header.Set("Referer", "http://119.18.194.36/miniIf/miniModifyPlace?mno="+mercId)
		req1.Header.Set("Origin", "http://119.18.194.36")
		req1.Header.Set("Host", "119.18.194.36")
		req1.Header.Set("Upgrade-Insecure-Requests", "1")

		resp1, err := client.Do(req1)
		if err != nil {
			c.Logger().Error("new request error  do >>> ", err)
		}
		defer resp1.Body.Close()
		body1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			c.Logger().Error("new request error  read body  >>>", err)
		}
		c.Logger().Debug("body1 code is  ", resp.StatusCode)
		c.Logger().Debug("body1 is ", string(body1))
	}
	//verify the captcha
	//verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)
	//body := map[string]interface{}{"status":verifyResult}
	//if verifyResult {
	//	c.Logger().Debug(">>>> request", postParameters)
	//
	//} else {
	//	return c.JSON(http.StatusBadRequest, "")
	//}
	return c.JSON(http.StatusOK, "")
}
