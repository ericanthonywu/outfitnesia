package model

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	DefaultShowData struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	}
	DeleteDefault struct {
		Id string `json:"id"`
	}
	JenisShowData struct {
		DefaultShowData
		Kategori string `json:"kategori"`
	}
	JenisKategoriUpdateWithoutFile struct {
		Label string `json:"label"`
	}
)
