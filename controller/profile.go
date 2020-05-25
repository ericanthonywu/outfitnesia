package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"outfitnesia/model"
)

func UpdateProfileToko(c echo.Context) error {
	request := new(model.Toko)

	if err := c.Bind(request); err != nil {
		return err
	}

	if err := model.TokoC.UpdateId(request.Id, request); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model.EmptyResponse{})
}

func GetProfileToko(c echo.Context) error{
	request := new(model.GetProfileToko)
	toko := new(model.Toko)
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := model.TokoC.FindId(request.Id).One(&toko); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, toko)
}
