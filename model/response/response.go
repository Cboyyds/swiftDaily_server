package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Success = iota
	Error
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func OK(c *gin.Context) {
	Result(Success, nil, "success", c)
}
func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, nil, msg, c)
}
func OKWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "success", c)
}
func OKWithDetail(data interface{}, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, nil, msg, c)
}
