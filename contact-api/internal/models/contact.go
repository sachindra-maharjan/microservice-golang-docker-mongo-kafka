package models

import "gopkg.in/mgo.v2/bson"

type Contact struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName string `bson:"lastname" json:"lastname"`
	Email string `bson:"email" json:"email"`
	Phone string `bson:"phone" json:"phone"`
}