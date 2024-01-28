package v1

import (
	"github.com/gin-gonic/gin"
	"ginsdgo/utils"
	"net/http"
)

func ErrorAndAbort(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": utils.GetErrMsg(code),
	})
	c.Abort()
}