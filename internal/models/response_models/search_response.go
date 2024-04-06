package responsemodels

type SearchResp struct {
	UserId         uint   `json:"userid"  gorm:"column:id"`
	Name           string `json:"name"`
	UserName       string `json:"username"`
	UserProfileURL string `json:"userprofileimageurl,omitempty"`
	FollowedBy     string `json:"followedby,omitempty"`
}
