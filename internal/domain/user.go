package domain

import "gorm.io/gorm"

type status string

const (
	Blocked  status = "blocked"
	Deleted  status = "deleted"
	Pending  status = "pending"
	Active   status = "active"
	verified status = "verified"
	Rejected status = "rejected"
)

type Users struct {
	gorm.Model
	Name          string
	UserName      string
	Email         string
	Password      string
	Bio           string
	ProfileImgUrl string
	Links         string
	Status        status `gorm:"default:pending"`
}
