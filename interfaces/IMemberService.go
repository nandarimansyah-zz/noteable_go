package interfaces

import (
	"github.com/nandarimansyah/noteable_go/models"
)

type IMemberService interface {
	GetAllMember(orgName string) []models.Member
}
