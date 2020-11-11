package test

import "github.com/gin-gonic/gin"

// Routers test url
func Routers(e *gin.Engine) {
	e.GET("/test/shutdown/delay/", TestDelayService)
}
