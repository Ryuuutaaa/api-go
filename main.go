package main

import (
	"demo/config"
	"demo/controllers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	config.SetupDB()

	server := echo.New()

	server.POST("/users", controllers.Create)
	server.GET("/users/:id", controllers.Read)
	server.PUT("/users/:id", controllers.Update)

	server.Start(":1323")
}
