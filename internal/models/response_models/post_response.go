package responsemodels

import "time"

type AddPostResp struct {
	Caption string `json:"description,omitempty"`
	UserId  string `json:"userId,omitempty"`

	Media string `json:"media,omitempty"`
}

type PostData struct {
	UserId            uint   `json:"userid"  gorm:"column:id"`
	UserName          string `json:"username"`
	UserProfileImgURL string `json:"userprofileimageurl"`

	PostId     uint      `json:"postid"`
	LikeStatus bool      `json:"like_status" gorm:"column:is_liked"`
	Caption    string    `json:"caption"`
	CreatedAt  time.Time `json:"-"`

	LikesCount    string `json:"likes_count"`
	CommentsCount string `json:"comments_count"`

	PostAge  string   `json:"post-age"`
	MediaUrl []string `json:"media-urls" gorm:"type:text"`
}
