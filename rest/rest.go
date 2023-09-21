package rest

import (
	"github.com/gin-gonic/gin"
	user_handler "github.com/reyhanmichiels/bring_coffee/app/user/handler"
)

type Rest struct {
	gin *gin.Engine
}

func NewRest(gin *gin.Engine) Rest {
	return Rest{
		gin: gin,
	}
}

func (rest *Rest) RouteHealthCheck() {
	rest.gin.GET("/api/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "successed",
		})
	})
}

func (rest *Rest) RouteUser(userHandler *user_handler.UserHandler) {
	v1 := rest.gin.Group("/api/v1")

	v1.POST("/regist", userHandler.Registration)
	v1.POST("/users/verify", userHandler.VerifyAccount)
	v1.POST("/users/otp", userHandler.SendOTP)
}

func (rest *Rest) Serve() {
	rest.gin.Run()
}
