package repositories

import (
	"github.com/nandarimansyah/noteable_go/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NotesRepository struct {
	MongoDB *mgo.Database
}

const (
	NOTES = "notes"
)

func (r *NotesRepository) CreateNote(Note models.Note) (models.Note, error) {
	err := r.MongoDB.C(NOTES).Insert(&Note)
	return Note, err
}

func (r *NotesRepository) GetNote(ID string) (models.Note, error) {
	var note models.Note
	err := r.MongoDB.C(NOTES).FindId(bson.ObjectIdHex(ID)).One(&note)
	return note, err
}

func (r *NotesRepository) UpdateNote(Note models.Note) (models.Note, error) {
	err := r.MongoDB.C(NOTES).UpdateId(Note.ID, &Note)
	return Note, err
}

func (r *NotesRepository) DeleteNote(ID string) error {
	err := r.MongoDB.C(NOTES).RemoveId(bson.ObjectIdHex(ID))
	return err
}
