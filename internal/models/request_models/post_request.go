package requestmodels

import "mime/multipart"

type AddPostData struct {
	Caption string                  `form:"caption" validate:"lte=60"`
	Media   []*multipart.FileHeader `form:"media" validate:"required"`

	UserId    string `validate:"required"`
	MediaURLs []string
}

type PostId struct {
	PostId string `json:"postid" validate:"required,number"`
}

type LikeRequest struct {
	PostID string `json:"postid"`
	UserID string
}

type EditPost struct {
	Caption string `form:"caption" validate:"lte=60"`
	UserId  string `validate:"required"`
	PostId  string `json:"postid" validate:"required,number"`
}
