package repositories

import (
	"github.com/nandarimansyah/noteable_go/models"
)

type MemberRepository struct {
}

func (r *MemberRepository) GetAllMember(orgName string) []models.Member {
	return []models.Member{
		models.Member{LoginID: "abc123", AvatarURL: "http://web.id/img.jpg", NoOfFollower: 50, NoOfFollowing: 100},
		models.Member{LoginID: "def456", AvatarURL: "http://web.id/img.jpg", NoOfFollower: 20, NoOfFollowing: 50},
	}
}
