package responsemodels

import "time"

type AddPostResp struct {
	Caption string `json:"description,omitempty"`
	UserId  string `json:"userId,omitempty"`

	Media string `json:"media,omitempty"`
}

type PostData struct {
	PostId     uint      `json:"postid"`
	LikeStatus bool      `json:"like_status" gorm:"column:is_liked"`
	Caption    string    `json:"caption"`
	CreatedAt  time.Time `json:"-"`

	PostAge  string   `json:"post-age"`
	MediaUrl []string `json:"media-urls" gorm:"type:text"`
}
