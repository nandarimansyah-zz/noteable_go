package interfaces

import "github.com/nandarimansyah/noteable_go/models"

//INotesRepository note repository interface
type INotesRepository interface {
	CreateNote(Note models.Note) (models.Note, error)
	GetNote(ID string) (models.Note, error)
	UpdateNote(Note models.Note) (models.Note, error)
	DeleteNote(ID string) error
}
