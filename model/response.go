package model

import "github.com/globalsign/mgo/bson"

type LoginResponse struct {
	Token    string        `json:"token"`
	Id       bson.ObjectId `json:"id"`
	Username string        `json:"username"`
}

type ErrorResponse struct {
	Message interface{} `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

type FileResponse struct {
	Data interface{} `json:"data"`
	FilePrefix string `json:"file_prefix"`
}

type EmptyResponse struct {

}
