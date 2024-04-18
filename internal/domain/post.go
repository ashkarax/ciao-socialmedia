package domain

import "time"

type postStatus string

const (
	Normal    postStatus = "normal"
	Archieved postStatus = "archieved"
)

type Post struct {
	PostID uint `gorm:"primarykey"`

	UserID uint  `gorm:"not null"`
	Users  Users `gorm:"foreignKey:UserID"`

	Caption string

	CreatedAt time.Time

	PostStatus postStatus `gorm:"default:normal"`
}

type PostMedia struct {
	MediaID uint `gorm:"primarykey"`

	PostID uint `gorm:"not null"`
	Post   Post `gorm:"foreignKey:PostID"`

	MediaUrl string
}
