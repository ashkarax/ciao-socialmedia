package interfaceUseCase

import requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"

type IRelationUseCase interface {
	Follow(*requestmodels.FollowRequest) error
	UnFollow(*requestmodels.FollowRequest) error

}
