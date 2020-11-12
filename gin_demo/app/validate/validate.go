package validate

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
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
	// Date       string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

var trans ut.Translator

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

		// 注册自定义的字段方法校验器
		// if err := v.RegisterValidation("checkDate", customFieldValidation); err != nil {
		// 	return err
		// }

		// 错误信息翻译校验器
		zhT := zh.New()
		enT := en.New()

		// 第一个参数表示备用语言环境，后面的参数表示应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		var ok bool
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("uni.GetTranslator(zh) failed")
		}
		zhTranslations.RegisterDefaultTranslations(v, trans)

		// if err := v.RegisterTranslation(); err != nil {
		// 	return err
		// }
		return nil
	}
	return nil
}

// SignUpService handler
func SignUpService(c *gin.Context) {
	var signs SignUpParam
	if err := c.ShouldBind(&signs); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		// 非validator.ValidationErrors类型的错误直接返回
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  err.Error(),
			})
			return
		}

		// validator.ValidationErrors类型的错误翻译后返回
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  errs.Translate(trans),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
	})
}
