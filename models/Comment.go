package models

import "gopkg.in/mgo.v2/bson"

type Comment struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	OrgName  string        `bson:"org_name" json:"org_name"`
	Comment  string        `bson:"comment" json:"comment"`
	IsDelete bool          `bson:"is_delete" json:"-"`
}
