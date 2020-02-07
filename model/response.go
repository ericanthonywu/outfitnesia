package model

import "gopkg.in/mgo.v2/bson"

type LoginResponse struct {
	Token    string        `json:"token"`
	Id       bson.ObjectId `json:"id"`
	Username string        `json:"username"`
}

type ErrorResponse struct {
	Message interface{} `json:"message"`
}
