package v1

import (
	"ginsdgo/sdgo"

	"github.com/gin-gonic/gin"
)


func Info(c *gin.Context) {
	// get phone
	value, _ := c.Get("phone")
	user,_ := value.(string)
	_,u := sdgo.GetUserInfo(user)
	// send the json of u to the client
	c.JSON(200, gin.H{
		"status":  200,
		"data":    u,
		"message": "",
	})
}