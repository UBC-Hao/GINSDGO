package v1

import (
	"ginsdgo/middleware"
	"ginsdgo/model"
	"ginsdgo/sdgo"
	"ginsdgo/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func Login(c *gin.Context) {
	// login using phone and username
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	ok := formData.Validate()
	if ok == false {
		ErrorAndAbort(c, utils.LOGIN_ERROR)
		return
	}
	// check login
	ok = sdgo.CheckLogin(formData.Phone, formData.Username)
	if ok == false {
		ErrorAndAbort(c, utils.LOGIN_ERROR)
		return
	}
	// generate token
	setToken(c, formData)
	return
}


// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Phone: user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(10 * time.Hour).Unix(),
			Issuer:    "GinSDGO",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		ErrorAndAbort(c, utils.TOKEN_ERROR)
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Phone,
		"message": utils.GetErrMsg(200),
		"token":   token,
	})
	return
}