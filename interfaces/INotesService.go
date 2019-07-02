package interfaces

import "github.com/nandarimansyah/noteable_go/models"

//INotesService note service interface
type INotesService interface {
	CreateNote(Note models.Note) (models.Note, error)
	GetNote(ID string) (models.Note, error)
	UpdateNote(Note models.Note) (models.Note, error)
	DeleteNote(ID string) (models.Note, error)
}
