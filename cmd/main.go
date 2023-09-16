package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/bring_coffee/infrastructure"
	"github.com/reyhanmichiels/bring_coffee/infrastructure/postgresql"
	"github.com/reyhanmichiels/bring_coffee/rest"
)

func main() {
	//load env
	infrastructure.LoadEnv()

	//connect DB
	postgresql.ConnectDatabase()
	
	//init rest
	rest := rest.NewRest(gin.Default())
	
	//init healtch check	
	rest.RouteHealthCheck()

	//run server
	rest.Serve()
}
