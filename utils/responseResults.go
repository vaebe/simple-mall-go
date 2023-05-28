package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmptyInfo struct {
}

// ResponseResultInfo 响应信息
type ResponseResultInfo struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// ResponseResultsError 返回响应错误信息
func ResponseResultsError(c *gin.Context, err string) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"data": nil,
		"msg":  err,
	})
}

func ResponseResultsSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "请求成功！",
	})
}
