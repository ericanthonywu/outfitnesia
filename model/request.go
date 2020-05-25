package model

type (
	LoginAdmin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	RegisterUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	LoginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	LoginToko struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	RegisterToko struct {
		Merek string `json:"merek"`
		Password string `json:"password"`
		Instagram string `json:"instagram"`
		Line string `json:"line"`
		Whatsapp string `json:"whatsapp"`
		Alamat string `json:"alamat"`
		FotoKTP string `json:"foto_ktp"`
	}
	JenisShowData struct {
		DefaultShowData
		Kategori string `json:"kategori"`
	}
	JenisKategoriUpdateWithoutFile struct {
		Label string `json:"label"`
		Id    string `json:"id"`
	}
	GetProfileToko struct {
		Id string `json:"id"`
	}
	ToggleShowProduk struct {
		ShowStatus uint8 `json:"show_status"`
	}
)
