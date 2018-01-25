package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionStuff = "stuff"
	CollectionArticle = "article"

)

// @def Thing
type Thing struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string `binding:"required"`
	Value string `binding:"required"`
}

// @def Things
type Thines struct {
	Data []Thing         `json:"data" `
}

// @def Pets
type Pets struct {
	Id  bson.ObjectId `json:"id"`
	Tag []Tag         `json:"tag" `
}

// @def Tag
type Tag struct {
	Id   bson.ObjectId `json:"id"`
	Name string        `json:"name"`
}

// @def Error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
