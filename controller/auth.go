package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"outfitnesia/model"
)

func LoginAdmin(c echo.Context) (err error) {
	request := new(model.LoginAdmin)
	admin := new(model.Admin)

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	if err = model.UserC.Find(bson.M{
		"username": request.Username,
	}).Select(bson.M{
		"id":       1,
		"username": 1,
		"password": 1,
	}).One(&admin); err != nil {
		return err
	}

	if admin == new(model.Admin) {
		return echo.ErrNotFound
	}

	if err = bcrypt.CompareHashAndPassword(admin.Password, []byte(request.Password)); err != nil {
		return echo.ErrForbidden
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = admin.Username
	claims["id"] = admin.Id

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRETTOKEN")))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.LoginResponse{
		Token:    t,
		Username: admin.Username,
		Id:       admin.Id,
	})
}

func LoginUser(c echo.Context) (err error) {
	request := new(model.LoginUser)
	user := new(model.User)

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	if err = model.UserC.Find(bson.M{
		"username": request.Username,
	}).Select(bson.M{
		"id":       1,
		"username": 1,
		"password": 1,
	}).One(&user); err != nil {
		return err
	}

	if user == new(model.User) {
		return echo.ErrNotFound
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		return echo.ErrForbidden
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.Id

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRETTOKEN")))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.LoginResponse{
		Token:    t,
		Username: user.Username,
		Id:       user.Id,
	})
}

func RegisterUser(c echo.Context) error {
	request := new(model.RegisterUser)
	user := model.NewUser()

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	user.Username = request.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := model.UserC.Insert(user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.CustomResponse{})
}

func LoginToko(c echo.Context) (err error) {
	request := new(model.LoginToko)
	user := new(model.User)

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	if err = model.UserC.Find(bson.M{
		"username": request.Username,
	}).Select(bson.M{
		"id":       1,
		"username": 1,
		"password": 1,
	}).One(&user); err != nil {
		return err
	}

	if user == new(model.User) {
		return echo.ErrNotFound
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		return echo.ErrForbidden
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.Id

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRETTOKEN")))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.LoginResponse{
		Token:    t,
		Username: user.Username,
		Id:       user.Id,
	})
}

func RegisterToko(c echo.Context) (err error) {
	toko := model.NewToko()

	password := c.FormValue("password")
	merek := c.FormValue("merek")
	alamat := c.FormValue("alamat")
	whatsapp := c.FormValue("whatsapp")
	instagram := c.FormValue("ig")
	line := c.FormValue("line")

	fotoktp, err := c.FormFile("fotoktp")
	if err != nil {
		return err
	}

	if password == "" ||
		merek == "" ||
		alamat == "" ||
		instagram == "" ||
		line == "" ||
		whatsapp == "" {
		return echo.ErrBadRequest
	}

	fileNameKTP, err := model.InsertImage(fotoktp, "./uploads/fotoKTP")
	if err != nil {
		return err
	}

	toko.Whatsapp = whatsapp
	toko.Instagram = instagram
	toko.Merek = merek
	toko.Alamat = alamat
	toko.FotoKTP = fileNameKTP
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	toko.Password = hashedPassword

	if err := model.UserC.Insert(toko); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.EmptyResponse{})
}

func Migrate(c echo.Context) (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err = model.UserC.Insert(model.Admin{
		Username: "superadmin",
		Password: hashed,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "migrate successfully!"})
}
