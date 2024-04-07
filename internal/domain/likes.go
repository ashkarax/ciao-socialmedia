package domain

import "time"

type PostLikes struct {
	LikeID uint `gorm:"primarykey"`

	UserID uint  `gorm:"not null"`
	Users  Users `gorm:"foreignKey:UserID"`

	PostID uint `gorm:"not null"`
	Posts  Post `gorm:"foreignKey:PostID"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UniqueConstraint struct {
		UserID uint `gorm:"uniqueIndex:idx_user_post"`
		PostID uint `gorm:"uniqueIndex:idx_user_post"`
	} `gorm:"embedded;uniqueIndex:idx_user_post"`
}
