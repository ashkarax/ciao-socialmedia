package responsemodels

type SearchResp struct {
	UserId            uint   `json:"userid"  gorm:"column:id"`
	Name              string `json:"name"`
	UserName          string `json:"username"`
	UserProfileImgURL string `json:"userprofileimageurl,omitempty"`

	//for userB only
	FollowedBy      string `json:"followedby,omitempty"`
	FollowingStatus bool   `json:"following_status,omitempty"`
}
