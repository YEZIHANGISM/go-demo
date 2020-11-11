package test

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TestDelayService(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code":    2000,
		"message": "success",
	})
}
