package controller

import (
	"go-web-server/domain"
	"go-web-server/service"
	_http "go-web-server/utils/http"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnimalController struct {
	animalService *service.AnimalService
}

func NewAnimalController() *AnimalController {
	return &AnimalController{service.NewAnimalService()}
}

func (h *AnimalController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, h.animalService.Get())
}

func (h *AnimalController) GetOne(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	an := h.animalService.GetOne(int(id))
	res := _http.Response{
		Message: "Success",
		Data:    an,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *AnimalController) Create(c echo.Context) (err error) {
	a := new(domain.Animal)
	if err = c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusPreconditionFailed)
	}

	newA, err := h.animalService.Create(*a)
	var res _http.Response

	if err != nil {
		res = _http.Response{
			Message: "Failed to create animal",
		}
	} else {
		res = _http.Response{
			Message: "Success creating animal",
			Data:    newA,
		}
	}

	return c.JSON(http.StatusOK, res)
}

func (h *AnimalController) Patch(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var err error
	a := new(domain.Animal)
	if err = c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusPreconditionFailed)
	}

	an, err := h.animalService.Update(int(id), *a)
	if err != nil {
		panic(err.Error())
	}

	res := _http.Response{
		Message: "Success updating animal",
		Data:    an,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *AnimalController) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := h.animalService.Delete(int(id))
	if err != nil {
		panic(err.Error())
	}

	res := _http.Response{
		Message: "Success deleting animal",
	}
	return c.JSON(http.StatusOK, res)
}
