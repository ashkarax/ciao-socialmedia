package interfaceRepository

import requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"

type IRelationRepo interface {
	InitiateFollowRelationship(*requestmodels.FollowRequest) error
	InitiateUnFollowRelationship(*requestmodels.FollowRequest) error

}
