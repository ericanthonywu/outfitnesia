package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"os"
	"time"
)

var (
	UserC     *mgo.Collection
	AdminC    *mgo.Collection
	TokoC     *mgo.Collection
	KategoriC *mgo.Collection
	BannerC   *mgo.Collection
	ProdukC   *mgo.Collection
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
	AdminC = DB.C("admin")
	TokoC = DB.C("toko")
	KategoriC = DB.C("kategori")
	BannerC = DB.C("banner")
	ProdukC = DB.C("produk")
}

type (
	Admin struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Username  string        `json:"username"`
		Password  []byte        `json:"password"`
		CreatedAt time.Time     `json:"created_at"`
	}

	User struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Username  string        `json:"username"`
		Password  []byte        `json:"password"`
		CreatedAt time.Time     `json:"created_at"`
	}

	Banner struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Order     uint8         `json:"order"`
		Nama      string        `json:"nama"`
		CreatedAt time.Time     `json:"created_at"`
	}

	Toko struct {
		Id         bson.ObjectId   `json:"id" bson:"_id,omitempty"`
		Username   string          `json:"username"`
		Password   []byte          `json:"password"`
		NamaToko   string          `json:"nama_toko"`
		Merek      string          `json:"merek"`
		Deskripsi  string          `json:"deskripsi"`
		Follower   []bson.ObjectId `json:"follower"`
		Email      string          `json:"email"`
		Instagram  string          `json:"instagram"`
		Whatsapp   string          `json:"whatsapp"`
		Website    string          `json:"website"`
		Alamat     string          `json:"alamat"`
		FotoProfil string          `json:"foto_profil"`
		BukaLapak  string          `json:"buka_lapak"`
		Shopee     string          `json:"shopee"`
		Tokopedia  string          `json:"tokopedia"`
		FotoKTP    string          `json:"foto_ktp"`
		CreatedAt  time.Time       `json:"created_at"`
	}
	Follower struct {
		Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
		User bson.ObjectId `json:"user"`
	}

	Kategori struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Label     string        `json:"label"`
		Gambar    string        `json:"gambar"`
		Jenis     []Jenis       `json:"jenis"`
		CreatedAt time.Time     `json:"created_at"`
	}

	Jenis struct {
		Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Label     string        `json:"label"`
		Gambar    string        `json:"gambar"`
		CreatedAt time.Time     `json:"created_at"`
	}

	Produk struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		NamaProduk string        `json:"nama_produk"`
		Kategori   bson.ObjectId `json:"kategori"`
		Jenis      bson.ObjectId `json:"jenis"`
		Bahan      string        `json:"bahan"`
		Warna      string        `json:"warna"`
		Deskripsi  string        `json:"deskripsi"`
		FotoProduk []string      `json:"foto_produk"`
		ShowStatus uint8         `json:"show_status"`
		CreatedAt  time.Time     `json:"created_at"`
	}
)

func NewProduk() *Produk {
	return &Produk{CreatedAt: time.Now(), ShowStatus: 1}
}

func NewBanner() *Banner {
	return &Banner{CreatedAt: time.Now()}
}

func NewToko() *Toko {
	return &Toko{CreatedAt: time.Now()}
}

func NewJenis() *Jenis {
	return &Jenis{CreatedAt: time.Now()}
}

func NewKategori() *Kategori {
	return &Kategori{CreatedAt: time.Now()}
}

func NewUser() *User {
	return &User{CreatedAt: time.Now()}
}
