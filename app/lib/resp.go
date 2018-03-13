package lib

import (
	"github.com/labstack/echo"
)

/**
	返回适当的状态代码
		在每个响应中返回适当的HTTP状态代码。应根据本指南对成功的答复进行编码：
		200：请求成功了GET，POST，DELETE，或PATCH调用同步完成，或者PUT是同步更新的现有资源调用
		201：请求成功POST，或PUT同步创建新资源的调用。提供指向新创建资源的“位置”标题也是最佳做法。这在POST上下文中特别有用，因为新资源将具有与原始请求不同的URL。
		202：接受请求的POST，PUT，DELETE，或PATCH致电将被异步处理
		206：请求成功GET，但只返回部分响应：请参阅上面的范围

		请注意使用身份验证和授权错误代码：
		401 Unauthorized：请求失败，因为用户未通过身份验证
		403 Forbidden：请求失败，因为用户没有访问特定资源的权限

		出现错误时，请返回合适的代码以提供附加信息：
		422 Unprocessable Entity：您的请求已被理解，但包含无效的参数
		429 Too Many Requests：您的速度有限，稍后再试
		500 Internal Server Error：服务器出现问题，请检查状态网站和/或报告问题
		有关用户错误和服务器错误情况的状态代码指南
 */
type msg struct {
	ErrCode int         `json:"err_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Resp(c echo.Context, code int, message string, data interface{}) error {
	errorCode := 10000
	msg := msg{ErrCode: errorCode, Message: message, Data: data}
	return c.JSON(code, msg)
}
