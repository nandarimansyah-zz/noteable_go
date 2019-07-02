package repositories

import (
	"github.com/nandarimansyah/noteable_go/models"
	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CommentRepository struct {
	MongoDB *mgo.Database
}

const (
	COMMENTS = "comments"
)

func (r *CommentRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	err := r.MongoDB.C(COMMENTS).Insert(&comment)
	return comment, err
}

func (r *CommentRepository) GetComments(orgName string) ([]models.Comment, error) {
	var response = []models.Comment{}
	err := r.MongoDB.C(COMMENTS).Find(bson.M{"org_name": orgName, "is_delete": false}).All(&response)
	return response, err
}

func (r *CommentRepository) DeleteComments(orgName string) bool {
	colQuerier := bson.M{"org_name": orgName, "is_delete": false}
	change := bson.M{"$set": bson.M{"is_delete": true}}
	_, err := r.MongoDB.C(COMMENTS).UpdateAll(colQuerier, change)
	if err != nil {
		logrus.Error(err)
		return false
	}

	return true
}
