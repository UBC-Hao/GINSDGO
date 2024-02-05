package routes

import (
	v1 "ginsdgo/api/v1"
	"ginsdgo/middleware"

	"github.com/gin-gonic/gin"
)


func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.Cors())
	_ = r.SetTrustedProxies(nil)

	r.Static("/static", "./static")
	
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("info", v1.Info)
		auth.GET("matches",v1.Matches)
	}

	all := r.Group("/api/v1")
	{
		all.POST("login", v1.Login)
	}
	_ = r.Run(":8080")
}