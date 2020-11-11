package validate

import "github.com/gin-gonic/gin"

// Routers urls
func Routers(e *gin.Engine) {
	e.POST("/signup/", SignUpService)
}
