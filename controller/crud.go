package controller

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"outfitnesia/model"
)

func CreateKategori(c echo.Context) error {
	label := c.FormValue("label")
	gambar, err := c.FormFile("gambar")
	if err != nil {
		return err
	}

	filename, err := model.InsertImage(gambar, "./uploads/kategori/")
	if err != nil {
		return err
	}
	if err := model.KategoriC.Insert(bson.M{
		"label":  label,
		"gambar": filename,
		"jenis":  nil,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.EmptyResponse{})
}

func CreateJenis(c echo.Context) error {
	label := c.FormValue("label")
	kategori := c.FormValue("kategori")
	gambar, err := c.FormFile("gambar")
	if err != nil {
		return err
	}

	filename, err := model.InsertImage(gambar, "./uploads/jenis/")
	if err := model.KategoriC.UpdateId(kategori, bson.M{
		"$push": bson.M{
			"jenis": bson.M{
				"label":  label,
				"gambar": filename,
			},
		},
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, model.EmptyResponse{})
}
