package jwt

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "empty auth",
			})
			c.Abort()
			return
		}

		// parts := strings.SplitN(authHeader, " ", 2)
		// if !(len(parts) == 2 && parts[0] == "Bearer") {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code": 2004,
		// 		"msg":  "wrong format with auth",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// mc, err := ParseToken(parts[1])
		mc, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("username", mc.Username)
		c.Next()
	}

}
