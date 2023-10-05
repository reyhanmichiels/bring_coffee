package main

import (
	"github.com/gin-gonic/gin"
	user_handler "github.com/reyhanmichiels/bring_coffee/app/user/handler"
	user_repository "github.com/reyhanmichiels/bring_coffee/app/user/repository"
	user_usecase "github.com/reyhanmichiels/bring_coffee/app/user/usecase"
	"github.com/reyhanmichiels/bring_coffee/infrastructure"
	"github.com/reyhanmichiels/bring_coffee/infrastructure/postgresql"
	"github.com/reyhanmichiels/bring_coffee/rest"
	"github.com/reyhanmichiels/bring_coffee/util"
)

func main() {
	//load env
	infrastructure.LoadEnv()

	//connect DB
	postgresql.ConnectDatabase()

	//migrate DB
	postgresql.Migrate()

	//init mail
	mail := util.NewMail()

	//init repo
	userRepo := user_repository.NewUserRepository(postgresql.DB)

	//init usecase
	userUsecase := user_usecase.NewUserUsecase(userRepo, mail)

	//init handler
	userHandler := user_handler.NewUserHandler(userUsecase)

	//init rest
	rest := rest.NewRest(gin.Default())

	//init healtch check
	rest.RouteHealthCheck()

	//init route
	rest.RouteUser(userHandler)

	//run server
	rest.Serve()
}
