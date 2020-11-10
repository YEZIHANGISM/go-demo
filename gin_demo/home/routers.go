package home

import (
	"gin_demo/jwt"

	"github.com/gin-gonic/gin"
)

// Routers urls
func Routers(e *gin.Engine) {
	e.GET("/hello/", HelloService)
	e.GET("/home/", jwt.JWTAuthMiddleWare(), HomeService)
}
