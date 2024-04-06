package domain

type FollowRelationship struct {
	FollowerID uint  `gorm:"not null"`
	Follower   Users `gorm:"foreignKey:FollowerID"`

	FollowingID uint  `gorm:"not null"`
	Following   Users `gorm:"foreignKey:FollowingID"`
}
