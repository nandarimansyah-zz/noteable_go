package interfaces

import "github.com/nandarimansyah/noteable_go/models"

type ICommentRepository interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetComments(orgName string) ([]models.Comment, error)
	DeleteComments(orgName string) bool
}
