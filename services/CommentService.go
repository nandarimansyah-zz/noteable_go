package services

import (
	"github.com/nandarimansyah/noteable_go/interfaces"
	"github.com/nandarimansyah/noteable_go/models"
	"gopkg.in/mgo.v2/bson"
)

type CommentService struct {
	Repository interfaces.ICommentRepository
}

func (s *CommentService) CreateComment(orgName, comment string) (models.Comment, error) {
	cm := models.Comment{
		ID:      bson.NewObjectId(),
		OrgName: orgName,
		Comment: comment,
	}
	return s.Repository.CreateComment(cm)
}

func (s *CommentService) GetComments(orgName string) ([]models.Comment, error) {
	return s.Repository.GetComments(orgName)
}

func (s *CommentService) DeleteComments(orgName string) bool {
	return s.Repository.DeleteComments(orgName)
}
