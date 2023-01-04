package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	respOk   = 0  // "OK"
	respFail = -1 // "FAIL"
)

type RespResult struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, RespResult{
		Data: data,
		Code: respOk,
		Msg:  "success",
	})
}

// Fail . return 400 and error msg
func Fail(c *gin.Context, e error) {
	c.JSON(http.StatusBadRequest, RespResult{
		Code: respFail,
		Msg:  e.Error(),
	})
	c.Abort()
}
