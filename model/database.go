package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"os"
)

var (
	UserC    *mgo.Collection
	KategoriC *mgo.Collection
)

func InitDB() {
	session, err := mgo.Dial(os.Getenv("MONGOURL"))
	if err != nil {
		panic(err)
	}
	DB := session.DB(os.Getenv("DB"))

	UserC = DB.C("user")
	KategoriC = DB.C("kategori")
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password []byte        `json:"password"`
}

type Kategori struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Label  string        `json:"label"`
	Gambar string        `json:"gambar"`
	Jenis  []Jenis       `json:"jenis"`
}

type Jenis struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Label  string        `json:"label"`
	Gambar string        `json:"gambar"`
}
