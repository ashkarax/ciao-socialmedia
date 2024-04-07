package interfaceRepository

import (
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type IRelationRepo interface {
	InitiateFollowRelationship(*requestmodels.FollowRequest) error
	InitiateUnFollowRelationship(*requestmodels.FollowRequest) error

	GetFollowersDetailsOfUserById(userId *string) (*[]responsemodels.SearchResp, error)
	GetFollowingDetailsOfUserById(userId *string) (*[]responsemodels.SearchResp, error)

}
