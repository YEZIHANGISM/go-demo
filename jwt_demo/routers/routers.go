package routers

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	r := gin.Default()
	r.POST("/auth", authHandler)
	return r
}
