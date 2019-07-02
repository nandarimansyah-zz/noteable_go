package services

import (
	"github.com/nandarimansyah/noteable_go/interfaces"
	"github.com/nandarimansyah/noteable_go/models"
)

type MemberService struct {
	Repository interfaces.IMemberRepository
}

func (s *MemberService) GetAllMember(orgName string) []models.Member {
	return s.Repository.GetAllMember(orgName)
}
