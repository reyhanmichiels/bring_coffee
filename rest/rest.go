package rest

import (
	"github.com/gin-gonic/gin"
	user_handler "github.com/reyhanmichiels/bring_coffee/app/user/handler"
	"github.com/reyhanmichiels/bring_coffee/middleware"
)

type Rest struct {
	gin *gin.Engine
}

func NewRest(gin *gin.Engine) Rest {
	return Rest{
		gin: gin,
	}
}

func (rest *Rest) HandleCORS() {
	rest.gin.Use(middleware.CORS)
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

	rest.gin.POST("/api/auth/regist", userHandler.Registration)
	v1.POST("/users/verify", userHandler.VerifyAccount)
	v1.POST("/users/otp", userHandler.SendOTP)
	rest.gin.POST("/api/auth/basic/login", userHandler.BasicLogin)
	v1.POST("/users/verify-fp", userHandler.VerifyForgetPassword)
	v1.POST("/users/forget-password", middleware.ForgetPassword, userHandler.ForgetPassword)
}

func (rest *Rest) Serve() {
	rest.gin.Run()
}
