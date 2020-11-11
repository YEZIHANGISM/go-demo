package validate

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	Age        string `json:"age" binding:"required,gte=1,lte=130"`
	Date       string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

func customFieldValidation(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// InitValidator 自定义校验器
func InitValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("checkDate", customFieldValidation); err != nil {
			return err
		}
		return nil
	}
	return nil
}

// SignUpService handler
func SignUpService(c *gin.Context) {
	if err := InitValidator(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  err.Error(),
		})
		return
	}
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
