package domain

type FollowRelationship struct {
	FollowerID uint  `gorm:"not null"`
	Follower   Users `gorm:"foreignKey:FollowerID"`

	FollowingID uint  `gorm:"not null"`
	Following   Users `gorm:"foreignKey:FollowingID"`

	UniqueConstraint struct {
		FollowerID  uint `gorm:"uniqueIndex:idx_follower_following"`
		FollowingID uint `gorm:"uniqueIndex:idx_follower_following"`
	} `gorm:"embedded;uniqueIndex:idx_follower_following"`
}
