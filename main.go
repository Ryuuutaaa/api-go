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

	server.GET("/users", controllers.ReadAll)
	server.POST("/users/create", controllers.Create)
	server.GET("/users/:id", controllers.Read)
	server.PUT("/users/:id", controllers.Update)
	server.DELETE("/users/:id", controllers.Delete)

	server.Start(":1323")
}
