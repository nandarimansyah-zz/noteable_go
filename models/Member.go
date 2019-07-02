package models

type Member struct {
	LoginID       string `json:"login_id"`
	AvatarURL     string `json:"avatar_url"`
	NoOfFollower  int    `json:"no_of_follower"`
	NoOfFollowing int    `json:"no_of_following"`
}
