package main

import (
	"paypath/config"
	"paypath/internal/user/handler"
	"paypath/internal/user/processor"
	"paypath/internal/user/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.SetEnv()
	userRepository, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}
	userProcessor := processor.NewProcessor(userRepository)
	userhandler := handler.NewUser(userProcessor)
	echoServer := echo.New()

	// Middleware
	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	echoServer.POST("/user", userhandler.Create)
	echoServer.GET("/confirm", userhandler.Confirm)
	echoServer.POST("/login", userhandler.Login)
	echoServer.Start(":8080")
}
