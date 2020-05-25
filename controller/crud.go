package controller

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
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
	}).Limit(request.Limit).
		Skip(request.Offset).
		All(&kategori); err != nil {
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
				Id:        bson.NewObjectId(),
				Label:     label,
				Gambar:    filename,
				CreatedAt: time.Now(),
			},
		},
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, model.EmptyResponse{})
}

func UpdateKategoriGambar(c echo.Context) error {
	var filename string
	label := c.FormValue("label")
	gambar, err := c.FormFile("gambar")

	if err != nil {
		return err
	}
	id := bson.ObjectIdHex(c.FormValue("kategori"))

	updatedData := new(model.Kategori)
	updatedData.Label = label

	filename, err = model.InsertImage(gambar, "./uploads/kategori/")
	if err != nil {
		return err
	}
	updatedData.Gambar = filename

	kategori := new(model.Kategori)

	if err := model.KategoriC.FindId(id).Select(bson.M{
		"gambar": 1,
	}).One(&kategori); err != nil {
		return err
	}

	if err := model.KategoriC.UpdateId(id, updatedData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.EmptyResponse{})
}

func UpdateKategori(c echo.Context) error {
	request := new(model.JenisKategoriUpdateWithoutFile)
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := model.KategoriC.UpdateId(bson.ObjectIdHex(request.Id), bson.M{"label": request.Label}); err != nil {
		return err
	}

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

func ShowBanner(c echo.Context) error {
	banner := new(model.Banner)
	if err := model.BannerC.Find(bson.M{}).Sort("order").All(banner); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.FileResponse{
		Data:       banner,
		FilePrefix: "banner",
	})
}

func AddBanner(c echo.Context) error {
	gambar, err := c.FormFile("gambar")
	banner := model.NewBanner()
	if err != nil {
		return err
	}

	filestring, err := model.InsertImage(gambar, "./uploads/banner")
	if err != nil {
		return err
	}

	banner.Nama = filestring

	if err := model.BannerC.Insert(banner); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model.EmptyResponse{})
}

func DeleteBanner(c echo.Context) error {
	request := new(model.DeleteDefault)
	banner := new(model.Banner)

	if err := c.Bind(request); err != nil {
		return err
	}
	BannerC := model.BannerC

	if err := BannerC.FindId(request.Id).Select(bson.M{"nama": 1}).One(&banner); err != nil {
		return err
	}

	if err := os.Remove("./uploads/banner/" + banner.Nama); err != nil {
		return err
	}

	if err := BannerC.RemoveId(request.Id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.EmptyResponse{})
}

func ShowProduk(c echo.Context) error {
	request := new(model.DefaultShowData)
	produk := new(model.Produk)
	if err := c.Bind(request); err != nil {
		return err
	}

	if err := model.ProdukC.Find(bson.M{
		"show_status": 1,
	}).All(&produk); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, produk)
}

func AddProduk(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	produk := model.NewProduk()
	files := form.File["image"]

	var arrayOfFilename []string
	for _, file := range files {
		filename, err := model.InsertImage(file, "./uploads/produk/")
		if err != nil {
			return err
		}

		arrayOfFilename = append(arrayOfFilename, filename)
	}

	produk.FotoProduk = arrayOfFilename
	produk.Kategori = bson.ObjectIdHex(c.FormValue("kategori"))
	produk.Jenis = bson.ObjectIdHex(c.FormValue("jenis"))
	produk.Bahan = c.FormValue("bahan")
	produk.Deskripsi = c.FormValue("deskripsi")
	produk.NamaProduk = c.FormValue("namaproduk")

	if err := model.ProdukC.Insert(produk); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.EmptyResponse{})
}

func DeleteProduk(c echo.Context) error {
	request := new(model.DeleteDefault)
	produk := new(model.Produk)

	if err := c.Bind(request); err != nil {
		return err
	}

	produkCollection := model.ProdukC

	if err := produkCollection.FindId(request.Id).Select(bson.M{
		"foto_produk": 1,
	}).One(&produk); err != nil {
		return err
	}

	for _, file := range produk.FotoProduk {
		if err := os.Remove("./uploads/produk/" + file); err != nil {
			return err
		}
	}

	if err := produkCollection.RemoveId(request.Id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model.EmptyResponse{})
}

func ToggleShowProduk(c echo.Context) error {

}
