package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/mojocn/base64Captcha"
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
		Height:             60,
		Width:              194,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    true,
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
	Id string   			`json:"id"`
	VerifyValue string		`json:"verify_value"`
}

//func CaptchaVerify(c echo.Context) error {
//	var postParameters verifyBody
//	if err:= c.Bind(&postParameters);err !=nil{
//		return c.JSON(http.StatusBadRequest,"")
//	}
//	//verify the captcha
//	//比较图像验证码
//	c.Logger().Debug(">>>> request", postParameters)
//	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)
//	body := map[string]interface{}{"status":verifyResult}
//	c.Logger().Debug(">>>> result", verifyResult)
//
//	return c.JSON(http.StatusOK,body)
//}

func AttrUpdate(c echo.Context) error {
	var postParameters verifyBody
	if err:= c.Bind(&postParameters);err !=nil{
		return c.JSON(http.StatusBadRequest,"")
	}
	//verify the captcha
	//比较图像验证码
	c.Logger().Debug(">>>> request", postParameters)
	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)
	body := map[string]interface{}{"status":verifyResult}
	c.Logger().Debug(">>>> result", verifyResult)

	return c.JSON(http.StatusOK,body)
}