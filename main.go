package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"outfitnesia/controller"
	"outfitnesia/model"
)

func main() {
	e := echo.New()
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}

	model.InitDB()

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
		middleware.Recover(),   //recover server on production if it's stop
		middleware.Logger(),    //logging
		middleware.RequestID(), //add request ID in every route
	)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Logger().Error(report)
		if report.Message == "not found" {
			_ = c.JSON(http.StatusNotFound, model.ErrorResponse{Message: report.Message})
		} else {
			_ = c.JSON(report.Code, model.ErrorResponse{Message: report.Message})
		}
	}

	e.GET("/migrate", controller.Migrate)

	user := e.Group("/user")
	user.POST("/login", controller.Login)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
