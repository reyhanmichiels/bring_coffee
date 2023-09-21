package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/bring_coffee/domain"
	"github.com/reyhanmichiels/bring_coffee/infrastructure/postgresql"
	"github.com/reyhanmichiels/bring_coffee/util"
)

func JWTAuth(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	token := strings.Split(bearerToken, "")[1]
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "failed to get your jwt token",
			"err":     nil,
		})
	}

	userId, tokenExpireTime, err := util.ParsesAndValidateJWT(token)
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

	var user domain.User
	err = postgresql.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "failed to get user",
			"error":   err,
		})
		return
	}

	c.Set("user", user)
	c.Next()
}
