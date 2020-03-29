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
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
		middleware.Recover(),   //  recover server on production if it's stop
		middleware.Logger(),    //  logging
		middleware.RequestID(), //  add request ID in every route
		middleware.Secure(),    //  secure
	)

	e.Static("/","uploads")

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Logger().Error(report)

		if report.Message == "not found" {
			_ = c.JSON(http.StatusNotFound, report)
		} else {
			_ = c.JSON(report.Code, report)
		}
	}

	e.GET("/migrateAdmin", controller.Migrate)

	admin := e.Group("/admin")
	admin.POST("/login", controller.LoginAdmin)

	admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWTSECRETTOKEN")),
		TokenLookup: "header:token",
	}))

	admin.POST("/kategori",controller.ShowKategori)
	admin.POST("/add-kategori", controller.CreateKategori)
	admin.PUT("/update-kategori", controller.UpdateKategori)
	admin.DELETE("/delete-kategori", controller.DeleteKategori)

	admin.POST("/jenis", controller.ShowJenis)
	admin.POST("/add-jenis", controller.CreateJenis)
	admin.PUT("/update-jenis", controller.UpdateJenis)
	admin.DELETE("/delete-jenis", controller.DeleteJenis)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
