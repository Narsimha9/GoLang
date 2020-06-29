package models

//created by H.G Nuwan Indika

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name   string        `json:"name" bson:"name" binding:"required"`
		Gender string        `json:"gender" bson:"gender" binding:"required"`
		Age    int           `json:"age" bson:"age" binding:"required"`
	}
)
