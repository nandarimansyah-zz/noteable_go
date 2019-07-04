package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nandarimansyah/noteable_go/interfaces/mocks"
	"github.com/nandarimansyah/noteable_go/models"
	"github.com/nandarimansyah/noteable_go/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockICommentRepository(mockCtrl)

	service := new(services.CommentService)
	service.Repository = repo

	expectedOrgName := "organName"
	expectedComment := "Please Remember This is Comment"

	expectedResult := models.Comment{
		OrgName: expectedOrgName,
		Comment: expectedComment,
	}
	repo.EXPECT().CreateComment(gomock.Any()).Return(expectedResult, nil)

	actualResult, _ := service.CreateComment(expectedOrgName, expectedComment)

	assert.Equal(t, expectedComment, actualResult.Comment)
	assert.Equal(t, expectedOrgName, actualResult.OrgName)
}

func TestGetComments(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockICommentRepository(mockCtrl)

	service := new(services.CommentService)
	service.Repository = repo

	expectedOrgName := "organName"

	expectedResult := []models.Comment{
		models.Comment{Comment: "Comment 1", OrgName: expectedOrgName},
		models.Comment{Comment: "Comment 2", OrgName: expectedOrgName},
	}

	repo.EXPECT().GetComments(gomock.Eq(expectedOrgName)).Return(expectedResult, nil)

	actualResult, _ := service.GetComments(expectedOrgName)

	assert.Equal(t, len(expectedResult), len(actualResult))
	assert.Equal(t, expectedResult[0].OrgName, expectedOrgName)
}

func TestDeleteComments(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mocks.NewMockICommentRepository(mockCtrl)

	service := new(services.CommentService)
	service.Repository = repo

	expectedOrgName := "organName"

	expectedResult := true
	repo.EXPECT().DeleteComments(gomock.Eq(expectedOrgName)).Return(expectedResult)

	actualResult := service.DeleteComments(expectedOrgName)

	assert.Equal(t, expectedResult, actualResult)
}
