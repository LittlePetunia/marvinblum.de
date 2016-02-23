package blog

import (
	//"db"
	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
}

type Comment struct {
}
