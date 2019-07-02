package interfaces

import "net/http"

type ICommentController interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	GetAllComments(w http.ResponseWriter, r *http.Request)
	DeleteComments(w http.ResponseWriter, r *http.Request)
}
