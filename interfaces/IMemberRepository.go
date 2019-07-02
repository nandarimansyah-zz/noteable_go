package interfaces

import "github.com/nandarimansyah/noteable_go/models"

type IMemberRepository interface {
	GetAllMember(orgName string) []models.Member
}
