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
	model.InitDB()
	e := echo.New()
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
		middleware.Recover(),   //recover server on production if it's stop
		middleware.Logger(),    //logging
		middleware.RequestID(), //add request ID in every route
		middleware.Secure(),
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
			_ = c.JSON(report.Code, model.ErrorResponse{report})
		}
	}

	e.GET("/migrateAdmin", controller.Migrate)

	admin := e.Group("/admin")
	admin.POST("/login", controller.LoginAdmin)

	admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWTSECRETTOKEN")),
		TokenLookup: "header:token",
	}))

	admin.POST("/kategori", controller.CreateKategori)
	admin.POST("/jenis", controller.CreateJenis)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
