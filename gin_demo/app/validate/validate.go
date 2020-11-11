package validate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUpParam validator
type SignUpParam struct {
	Name       string `json:"name" binding:"required"`
	NickName   string `json:"nick_name" binding:"required"`
	Gender     bool   `json:"gender" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Address    string `json:"address"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	REPassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// SignUpService handler
func SignUpService(c *gin.Context) {
	var signs SignUpParam
	if err := c.ShouldBind(&signs); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
	})
}
