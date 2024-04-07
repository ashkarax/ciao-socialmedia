package usecase

import (
	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type RelationUseCase struct {
	RelationRepo interfaceRepository.IRelationRepo
}

func NewRelationUseCase(relationRepo interfaceRepository.IRelationRepo) interfaceUseCase.IRelationUseCase {
	return &RelationUseCase{RelationRepo: relationRepo}
}

func (r *RelationUseCase) Follow(data *requestmodels.FollowRequest) error {
	err := r.RelationRepo.InitiateFollowRelationship(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *RelationUseCase) UnFollow(data *requestmodels.FollowRequest) error {
	err := r.RelationRepo.InitiateUnFollowRelationship(data)
	if err != nil {
		return err
	}
	return nil

}

func (r *RelationUseCase) GetFollowersDetailsOfUser(userId *string) (*[]responsemodels.SearchResp, error) {
	data, err := r.RelationRepo.GetFollowersDetailsOfUserById(userId)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (r *RelationUseCase) GetFollowingDetailsOfUser(userId *string) (*[]responsemodels.SearchResp, error) {
	data, err := r.RelationRepo.GetFollowingDetailsOfUserById(userId)
	if err != nil {
		return nil, err
	}
	return data, nil

}
