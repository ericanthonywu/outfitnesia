package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"os"
	"time"
)

var (
	UserC     *mgo.Collection
	KategoriC *mgo.Collection
)

func InitDB() {
	session, err := mgo.Dial(os.Getenv("MONGOURL"))
	if err != nil {
		panic(err)
	}

	DB := session.DB(os.Getenv("DB"))
	//count, err := UserC.Find(bson.M{}).Count()
	//if err != nil {
	//	panic(err)
	//}
	//if count == 0 {
	//	hashed, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if err = UserC.Insert(bson.M{
	//		"username": "superadmin",
	//		"password": hashed,
	//	}); err != nil {
	//		panic(err)
	//	}
	//}

	UserC = DB.C("user")
	KategoriC = DB.C("kategori")
}

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username  string        `json:"username"`
	Password  []byte        `json:"password"`
	CreatedAt time.Time     `json:"created_at"`
}

type Kategori struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Label     string        `json:"label"`
	Gambar    string        `json:"gambar"`
	Jenis     []Jenis       `json:"jenis"`
	CreatedAt time.Time     `json:"created_at"`
}

type Jenis struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Label     string        `json:"label"`
	Gambar    string        `json:"gambar"`
	CreatedAt time.Time     `json:"created_at"`
}
