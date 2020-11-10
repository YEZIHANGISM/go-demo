package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloService Controller
func HelloService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world!",
	})
}

func HomeService(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code":    2000,
		"message": "success",
		"data":    gin.H{"username": username},
	})
}
