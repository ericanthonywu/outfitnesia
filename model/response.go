package model

import "github.com/globalsign/mgo/bson"

type (
	LoginResponse struct {
		Token    string        `json:"token"`
		Id       bson.ObjectId `json:"id"`
		Username string        `json:"username"`
	}
	ErrorResponse struct {
		Message interface{} `json:"message"`
	}
	Response struct {
		Message string `json:"message"`
	}
	FileResponse struct {
		Data       interface{} `json:"data"`
		FilePrefix string      `json:"file_prefix"`
	}
	EmptyResponse struct {

	}
)
