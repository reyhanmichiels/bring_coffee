package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/bring_coffee/util"
)

func ForgetPassword(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "failed to get your jwt token",
			"err":     nil,
		})
	}

	email, tokenExpireTime, err := util.ParsesAndValidateTokenForgetPassword(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "failed to parse and validate jwt",
			"err":     err,
		})
	}

	if float64(time.Now().Unix()) > tokenExpireTime {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "your session has been expired",
			"error":   nil,
		})
		return
	}

	c.Set("email", email)
	c.Next()
}
