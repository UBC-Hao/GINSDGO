package routes

import (
	v1 "ginsdgo/api/v1"
	"ginsdgo/middleware"

	"github.com/gin-gonic/gin"
)


func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.Static("/", "./static")
	
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
			return
		})
	}

	all := r.Group("/api/v1")
	{
		all.POST("/login", v1.Login)
	}
}