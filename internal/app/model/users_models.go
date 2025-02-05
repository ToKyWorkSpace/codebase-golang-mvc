package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int                `bson:"age"`
}

type PostUsers struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Name     string `bson:"name,omitempty"`
	Age      int    `bson:"age"`
}
