package main

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/controller"
	_costumMiddleware "go-web-server/middleware"
	"go-web-server/router"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	config.InitConfig()
}

func main() {
	e := echo.New()

	//middleware
	customMiddleware := _costumMiddleware.HTTPMiddleware()
	e.Use(customMiddleware.CORS)
	e.Use(middleware.Logger())

	//router
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	router.AnimalRouter(e, controller.NewAnimalController())

	//start server
	host := viper.GetString("application.server.host")
	port := viper.GetInt("application.server.port")
	err := e.Start(fmt.Sprintf("%s:%d", host, port))
	e.Logger.Fatal(err)
}
