package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandarimansyah/noteable_go/interfaces"
	"github.com/urfave/negroni"
)

// Route wrapper
type Route struct {
	NotesController   interfaces.INotesController
	CommentController interfaces.ICommentController
	MemberController  interfaces.IMemberController
}

// NewRoute instances
func NewRoute() *Route {
	return &Route{}
}

// GetRouter Build the all router
func (r *Route) GetRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	v1Root := router.PathPrefix("/v1").Subrouter()
	v1Root.HandleFunc("/notes/{id}", r.NotesController.GetNote).Methods("GET")
	v1Root.HandleFunc("/notes", r.NotesController.CreateNote).Methods("POST")
	v1Root.HandleFunc("/notes", r.NotesController.UpdateNote).Methods("PUT")
	v1Root.HandleFunc("/notes", r.NotesController.DeleteNote).Methods("DELETE")

	orgsPath := router.PathPrefix("/orgs").Subrouter()
	orgsPath.HandleFunc("/{name}/comments", r.CommentController.GetAllComments).Methods("GET")
	orgsPath.HandleFunc("/{name}/comments", r.CommentController.CreateComment).Methods("POST")
	orgsPath.HandleFunc("/{name}/comments", r.CommentController.DeleteComments).Methods("DELETE")

	orgsPath.HandleFunc("/{name}/members", r.MemberController.GetAllMember).Methods("GET")

	n := negroni.New()
	n.Use(NewLoggerMiddleware())
	n.UseHandler(router)

	return n
}
