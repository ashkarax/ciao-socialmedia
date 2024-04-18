package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type IRelationUseCase interface {
	Follow(*requestmodels.FollowRequest) error
	UnFollow(*requestmodels.FollowRequest) error

	GetFollowersDetailsOfUser(*string) (*[]responsemodels.SearchResp, error)
	GetFollowingDetailsOfUser(*string) (*[]responsemodels.SearchResp, error)
}
