package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DefaultShowData struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type DeleteDefault struct {
	Id string `json:"id"`
}

type JenisShowData struct {
	DefaultShowData
	Kategori string `json:"kategori"`
}
