package interfaces

import "github.com/nandarimansyah/noteable_go/models"

type ICommentService interface {
	CreateComment(orgName, comment string) (models.Comment, error)
	GetComments(orgName string) ([]models.Comment, error)
	DeleteComments(orgName string) bool
}
