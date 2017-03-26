package dbmodels

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id        bson.ObjectId `json:"id" bson: _id`
	FirstName string        `json:"firstName" bson:"firstName`
	LastName  string        `json:"lastName" bson:"lastName`
	Email     string        `json:"email" bson:"email`
	Groups    []string      `json:"groups" bson:"groups`
	Password  string        `json:"password" bson:"password`
}
