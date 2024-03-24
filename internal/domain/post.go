package domain

import "time"

type postStatus string

const (
	Normal    postStatus = "normal"
	Archieved postStatus = "archieved"
)

type Post struct {
	PostId uint `gorm:"primarykey"`

	UserID uint  `gorm:"not null"`
	Users  Users `gorm:"foreignKey:UserID"`

	Caption string

	CreatedAt time.Time

	PostStatus postStatus `gorm:"default:normal"`
}

type PostMedia struct {
	MediaId uint `gorm:"primarykey"`

	PostId uint `gorm:"not null"`
	Order  Post `gorm:"foreignKey:PostId"`

	MediaUrl string
}
