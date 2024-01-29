package middleware

import (
	"ginsdgo/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("asdq0912xw22"),
	}
}

type MyClaims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims MyClaims) (string ,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}
// get the user info from the token
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token_raw := c.Request.Header.Get("Authorization")
		// split token
		token_arr := strings.Split(token_raw, " ")
		token := ""
		if len(token_arr) == 2 {
			token = token_arr[1]
		}
		if token == "" {
			c.JSON(401, gin.H{
				"status": utils.TOKEN_ERROR,
				"msg":  utils.GetErrMsg(utils.TOKEN_ERROR),
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"status": utils.TOKEN_ERROR,
				"msg": utils.GetErrMsg(utils.TOKEN_ERROR),
			})
			c.Abort()
			return
		}
		c.Set("phone", claims.Phone)
		c.Next()
	}
}