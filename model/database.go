package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

var DB *mgo.Database

func InitDB() {
	session, err := mgo.Dial(os.Getenv("MONGOURL"))
	if err != nil {
		panic(err)
	}
	DB = session.DB(os.Getenv("DB"))
}

func GetDB() *mgo.Database {
	if DB == new(mgo.Database) {
		InitDB()
	}
	return DB
}

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}
