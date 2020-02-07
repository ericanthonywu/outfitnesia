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

func Login(c echo.Context) (err error) {
	request := new(model.Login)
	user := new(model.User)

	if err := c.Bind(request); err != nil {
		return err
	}

	if request.Password == "" || request.Username == "" {
		return echo.ErrBadRequest
	}

	//otomatis bos awkowko ok
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

func Migrate(c echo.Context) (err error) {
	collection := model.UserC

	hashed, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err = collection.Insert(bson.M{
		"username": "superadmin",
		"password": hashed,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "migrate successfully!"})
}
