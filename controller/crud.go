package controller

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	"net/http"
	"outfitnesia/model"
	"time"
)

func ShowKategori(c echo.Context) error {
	request := new(model.DefaultShowData)
	var kategori []model.Kategori

	if err := c.Bind(request); err != nil {
		return err
	}

	if err := model.KategoriC.Find(bson.M{}).Select(bson.M{
		"label":  1,
		"gambar": 1,
	}).Limit(request.Limit).Skip(request.Offset).All(&kategori); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.FileResponse{
		Data:       kategori,
		FilePrefix: "uploads/kategori/",
	})
}

func ShowJenis(c echo.Context) error {
	request := new(model.JenisShowData)
	var jenis model.Kategori

	if err := c.Bind(request); err != nil {
		return err
	}

	if err := model.KategoriC.FindId(bson.ObjectIdHex(request.Kategori)).Select(bson.M{
		"jenis": 1,
	}).Skip(request.Offset).Limit(request.Limit).One(&jenis); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.FileResponse{
		Data:       jenis.Jenis,
		FilePrefix: "uploads/jenis/",
	})
}

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

	if err := model.KategoriC.Insert(model.Kategori{
		Label:     label,
		Gambar:    filename,
		Jenis:     []model.Jenis{},
		CreatedAt: time.Now(),
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
	if err := model.KategoriC.UpdateId(bson.ObjectIdHex(kategori), bson.M{
		"$push": bson.M{
			"jenis": model.Jenis{
				Id:     bson.NewObjectId(),
				Label:  label,
				Gambar: filename,
				CreatedAt: time.Now(),
			},
		},
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, model.EmptyResponse{})
}

func UpdateKategori(c echo.Context) error {
	var filename string
	label := c.FormValue("label")
	gambar, err := c.FormFile("gambar")
	fmt.Println(gambar)
	if err != nil {
		return err
	}
	id := c.FormValue("kategori")

	if gambar != nil {
		filename, err = model.InsertImage(gambar, "./uploads/kategori/")
		if err != nil {
			return err
		}
	}
	updatedData := new(model.Kategori)
	updatedData.Label = label

	if filename != "" {
		updatedData.Gambar = filename
	}

	if err := model.KategoriC.UpdateId(bson.ObjectIdHex(id), updatedData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.EmptyResponse{})
}

func UpdateJenis(c echo.Context) error {
	var filename string

	label := c.FormValue("label")

	gambar, err := c.FormFile("gambar")
	if err != nil {
		return err
	}

	idKategori := bson.ObjectIdHex(c.FormValue("idKategori"))
	idJenis := bson.ObjectIdHex(c.FormValue("idJenis"))


	updatedData := new(model.Jenis)
	updatedData.Label = label

	if gambar != nil {
		filename, err = model.InsertImage(gambar, "./uploads/jenis/")
		updatedData.Gambar = filename
		if err != nil {
			return err
		}
	}

	if err := model.KategoriC.Update(bson.M{
		"_id":       idKategori,
		"jenis._id": idJenis,
	}, bson.M{
		"$set": bson.M{
			"jenis": updatedData,
		},
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.EmptyResponse{})
}

func DeleteKategori(c echo.Context) error {
	request := new(model.DeleteDefault)

	if err := model.KategoriC.RemoveId(bson.ObjectIdHex(request.Id)); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model.EmptyResponse{})
}

func DeleteJenis(c echo.Context) error {
	request := new(model.DeleteDefault)

	if err := model.KategoriC.Update(bson.M{
		"jenis._id": request.Id,
	}, bson.M{
		"$pull": bson.M{
			"jenis": bson.M{"_id": request.Id},
		},
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, model.EmptyResponse{})
}
