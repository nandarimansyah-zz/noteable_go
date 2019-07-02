package services

import (
	"github.com/nandarimansyah/noteable_go/interfaces"
	"github.com/nandarimansyah/noteable_go/models"
	"github.com/sirupsen/logrus"
)

//NotesService implementation of INotesService
type NotesService struct {
	NotesRepository interfaces.INotesRepository
}

func (s *NotesService) CreateNote(Note models.Note) (models.Note, error) {
	return s.NotesRepository.CreateNote(Note)
}

func (s *NotesService) GetNote(ID string) (models.Note, error) {
	data, err := s.NotesRepository.GetNote(ID)
	if err != nil {
		logrus.Error(err)
	}
	return data, nil
}

func (s *NotesService) UpdateNote(Note models.Note) (models.Note, error) {
	return s.NotesRepository.UpdateNote(Note)
}

func (s *NotesService) DeleteNote(ID string) (models.Note, error) {
	data, err := s.NotesRepository.GetNote(ID)
	if err != nil {
		logrus.Error(err)
	} else {
		err = s.NotesRepository.DeleteNote(ID)
		if err != nil {
			logrus.Error(err)
		}
	}

	return data, err
}
