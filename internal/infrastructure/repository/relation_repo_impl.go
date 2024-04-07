package repository

import (
	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	"gorm.io/gorm"
)

type RelationRepo struct {
	DB *gorm.DB
}

func NewRelationRepo(db *gorm.DB) interfaceRepository.IRelationRepo {
	return &RelationRepo{DB: db}
}

func (d *RelationRepo) InitiateFollowRelationship(data *requestmodels.FollowRequest) error {

	query := "INSERT INTO follow_relationships (follower_id, following_id) VALUES ($1, $2) ON CONFLICT (follower_id, following_id) DO NOTHING;"
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

func (d *RelationRepo) GetFollowersDetailsOfUserById(userId *string) (*[]responsemodels.SearchResp, error) {
	var followersData []responsemodels.SearchResp

	query := "SELECT * FROM follow_relationships JOIN users ON follow_relationships.following_id=users.id WHERE following_id=$1"
	err := d.DB.Raw(query, userId).Scan(&followersData).Error
	if err != nil {
		return nil, err
	}

	return &followersData, err
}

func (d *RelationRepo) GetFollowingDetailsOfUserById(userId *string) (*[]responsemodels.SearchResp, error) {
	var followersData []responsemodels.SearchResp

	query := "SELECT * FROM follow_relationships JOIN users ON follow_relationships.following_id=users.id WHERE follower_id=$1"
	err := d.DB.Raw(query, userId).Scan(&followersData).Error
	if err != nil {
		return nil, err
	}

	return &followersData, err
}

func (d *RelationRepo) GetFollowerAndFollowingCountofUser(userId *string) (*uint, *uint, error) {
	var counts struct {
		FollowersCount uint `gorm:"column:followers_count"`
		FollowingCount uint `gorm:"column:following_count"`
	}
	query := "SELECT (SELECT COUNT(*) FROM follow_relationships WHERE following_id = $1) AS followers_count,(SELECT COUNT(*) FROM follow_relationships WHERE follower_id = $1) AS following_count "
	err := d.DB.Raw(query, userId).Scan(&counts).Error
	if err != nil {
		return nil, nil, err
	}
	return &counts.FollowersCount, &counts.FollowingCount, nil

}

func (d *RelationRepo) UserAFollowingUserBorNot(requestData *requestmodels.FollowRequest) (bool, error) {
	var count uint

	query := "SELECT COUNT(*) FROM follow_relationships WHERE follower_id = ? AND following_id = ?"
	err := d.DB.Raw(query, requestData.UserID, requestData.OppositeUserID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}
