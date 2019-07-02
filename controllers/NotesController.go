package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandarimansyah/noteable_go/helpers"
	"github.com/nandarimansyah/noteable_go/interfaces"
	"github.com/nandarimansyah/noteable_go/models"
	"gopkg.in/mgo.v2/bson"
)

// NotesController implementation of INotesController
type NotesController struct {
	NotesService interfaces.INotesService
}

func (c *NotesController) CreateNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		helpers.APIResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	note.ID = bson.NewObjectId()

	data, err := c.NotesService.CreateNote(note)
	if err != nil {
		helpers.APIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.APIResponse(w, http.StatusCreated, data)
}
func (c *NotesController) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := c.NotesService.GetNote(id)
	if err != nil {
		helpers.APIResponse(w, http.StatusBadRequest, nil)
		return
	}

	helpers.APIResponse(w, http.StatusOK, data)
	return
}
func (c *NotesController) UpdateNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		helpers.APIResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	data, err := c.NotesService.UpdateNote(note)
	if err != nil {
		helpers.APIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.APIResponse(w, http.StatusOK, data)
}
func (c *NotesController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		helpers.APIResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	data, err := c.NotesService.DeleteNote(note.ID.Hex())
	if err != nil {
		helpers.APIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.APIResponse(w, http.StatusOK, data)
}
