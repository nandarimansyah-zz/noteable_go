package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nandarimansyah/noteable_go/interfaces"
	"github.com/nandarimansyah/noteable_go/models"

	"github.com/gorilla/mux"
	"github.com/nandarimansyah/noteable_go/helpers"
)

// CommentController implementation of ICommentController
type CommentController struct {
	Service interfaces.ICommentService
}

func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgName := vars["name"]

	var cmtBody models.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmtBody); err != nil {
		helpers.APIResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := c.Service.CreateComment(orgName, cmtBody.Comment)

	if err != nil {
		helpers.APIResponse(w, http.StatusInternalServerError, "Something is not right")
		return
	}

	helpers.APIResponse(w, http.StatusCreated, result)
	return
}

func (c *CommentController) GetAllComments(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orgName := vars["name"]

	result, _ := c.Service.GetComments(orgName)
	helpers.APIResponse(w, http.StatusOK, result)
	return
}

func (c *CommentController) DeleteComments(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orgName := vars["name"]

	result := c.Service.DeleteComments(orgName)

	if !result {
		helpers.APIResponse(w, http.StatusInternalServerError, "Fail to delete")
		return
	}

	helpers.APIResponse(w, http.StatusOK, "Success to delete")
	return
}
