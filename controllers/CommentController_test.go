package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/nandarimansyah/noteable_go/controllers"
	"github.com/nandarimansyah/noteable_go/interfaces/mocks"
	"github.com/nandarimansyah/noteable_go/models"
	"github.com/stretchr/testify/assert"
)

func TestCommentController_CreateComment(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	service := mocks.NewMockICommentService(mockCtrl)
	controller := new(controllers.CommentController)
	controller.Service = service

	orgName := "orgname"
	comment := "this is comment"

	expectedResult := models.Comment{
		OrgName: orgName,
		Comment: comment,
	}
	service.EXPECT().CreateComment(orgName, comment).Return(expectedResult, nil)

	body := `{"comment": "` + comment + `"}`

	req := httptest.NewRequest(http.MethodPost, "/orgs/orgname/comments", bytes.NewBufferString(body))
	rw := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/orgs/{name}/comments", controller.CreateComment).Methods(http.MethodPost)
	router.ServeHTTP(rw, req)

	var actualResult models.Comment
	err := json.NewDecoder(rw.Body).Decode(&actualResult)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rw.Code)
}

func TestCommentController_CreateCommentFailBodyNotValid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	service := mocks.NewMockICommentService(mockCtrl)
	controller := new(controllers.CommentController)
	controller.Service = service

	body := `{`

	req := httptest.NewRequest(http.MethodPost, "/orgs/orgname/comments", bytes.NewBufferString(body))
	rw := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/orgs/{name}/comments", controller.CreateComment).Methods(http.MethodPost)
	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusBadRequest, rw.Code)
}

func TestCommentController_GetComments(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	service := mocks.NewMockICommentService(mockCtrl)
	controller := new(controllers.CommentController)
	controller.Service = service

	orgName := "orgname"

	expectedResult := []models.Comment{
		models.Comment{OrgName: orgName, Comment: "Comment 1"},
		models.Comment{OrgName: orgName, Comment: "Comment 2"},
	}
	service.EXPECT().GetComments(orgName).Return(expectedResult, nil)

	req := httptest.NewRequest(http.MethodGet, "/orgs/orgname/comments", nil)
	rw := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/orgs/{name}/comments", controller.GetAllComments).Methods(http.MethodGet)
	router.ServeHTTP(rw, req)

	var actualResult []models.Comment
	err := json.NewDecoder(rw.Body).Decode(&actualResult)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Equal(t, len(expectedResult), len(actualResult))
}

func TestCommentController_DeleteComments(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	service := mocks.NewMockICommentService(mockCtrl)
	controller := new(controllers.CommentController)
	controller.Service = service

	orgName := "orgname"

	expectedResult := true
	service.EXPECT().DeleteComments(orgName).Return(expectedResult)

	req := httptest.NewRequest(http.MethodDelete, "/orgs/orgname/comments", nil)
	rw := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/orgs/{name}/comments", controller.DeleteComments).Methods(http.MethodDelete)
	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
}

func TestCommentController_DeleteCommentsFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	service := mocks.NewMockICommentService(mockCtrl)
	controller := new(controllers.CommentController)
	controller.Service = service

	orgName := "orgname"

	expectedResult := false
	service.EXPECT().DeleteComments(orgName).Return(expectedResult)

	req := httptest.NewRequest(http.MethodDelete, "/orgs/orgname/comments", nil)
	rw := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/orgs/{name}/comments", controller.DeleteComments).Methods(http.MethodDelete)
	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)
}
