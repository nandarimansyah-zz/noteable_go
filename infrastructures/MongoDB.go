package infrastructures

import (
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type MongoDB struct {
	DBCached *mgo.Database
}

//NewMongoDB new mongo db initialize
func NewMongoDB(url string, dbName string) *MongoDB {

	session, err := mgo.Dial(url)
	if err != nil {
		log.Error(err)
	}

	cDb := session.DB(dbName)

	mongoDB := &MongoDB{
		DBCached: cDb,
	}

	return mongoDB
}
