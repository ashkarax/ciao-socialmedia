package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	"gorm.io/gorm"
)

type JWTRepo struct {
	DB *gorm.DB
}

func NewJWTRepo(db *gorm.DB) interfaceRepository.IJWTRepo {
	return &JWTRepo{DB: db}
}

func (d *JWTRepo) GetUserStatForGeneratingAccessToken(userId *string) (*string, error) {
	var userCurrentStatus string
	query := "SELECT status from users WHERE id=?"
	result := d.DB.Raw(query, userId).Scan(&userCurrentStatus)

	if result.RowsAffected == 0 {
		errMessage := fmt.Sprintf("No results found,No user with this id=%s found in db", *userId)
		return &userCurrentStatus, errors.New(errMessage)
	}
	if result.Error != nil {
		return &userCurrentStatus, result.Error
	}

	return &userCurrentStatus, nil
}
