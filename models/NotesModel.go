package models

import "gopkg.in/mgo.v2/bson"

type Note struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `bson:"title" json:"title"`
	Text  string        `bson:"text" json:"text"`
}
