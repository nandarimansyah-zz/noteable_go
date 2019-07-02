package interfaces

import "net/http"

type INotesController interface {
	CreateNote(w http.ResponseWriter, r *http.Request)
	GetNote(w http.ResponseWriter, r *http.Request)
	UpdateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}
