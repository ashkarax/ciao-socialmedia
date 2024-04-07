package repository

import (
	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	"gorm.io/gorm"
)

type RelationRepo struct {
	DB *gorm.DB
}

func NewRelationRepo(db *gorm.DB) interfaceRepository.IRelationRepo {
	return &RelationRepo{DB: db}
}

func (d *RelationRepo) InitiateFollowRelationship(data *requestmodels.FollowRequest) error {

	query := "INSERT INTO follow_relationships (follower_id,following_id) VALUES($1,$2)"
	err := d.DB.Exec(query, data.UserID, data.OppositeUserID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *RelationRepo) InitiateUnFollowRelationship(data *requestmodels.FollowRequest) error {

	query := "DELETE FROM follow_relationships WHERE follower_id=$1 AND following_id=$2"
	err := d.DB.Exec(query, data.UserID, data.OppositeUserID).Error
	if err != nil {
		return err
	}
	return nil

}
