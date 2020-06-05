package router

import (
	"go-web-server/controller"

	"github.com/labstack/echo/v4"
)

func AnimalRouter(e *echo.Echo, animalController *controller.AnimalController) {

	g := e.Group("/api/animals")

	g.GET("", animalController.Get)
	g.GET("/:id", animalController.GetOne)
	g.POST("", animalController.Create)
	g.PATCH("/:id", animalController.Patch)
	g.DELETE("/:id", animalController.Delete)
}
