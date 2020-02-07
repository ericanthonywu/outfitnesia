package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"os"
)

var (
	UserC *mgo.Collection
)

func InitDB() {
	session, err := mgo.Dial(os.Getenv("MONGOURL"))
	if err != nil {
		panic(err)
	}
	DB := session.DB(os.Getenv("DB"))

	UserC = DB.C("user")
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password []byte        `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}
