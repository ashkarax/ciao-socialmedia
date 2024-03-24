package usecase

import (
	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
)

type JWTUseCase struct {
	JWTRepo interfaceRepository.IJWTRepo
}

func NewJWTUseCase(JWTRepo interfaceRepository.IJWTRepo) interfaceUseCase.IJWTUseCase {
	return &JWTUseCase{JWTRepo: JWTRepo}
}

func (r *JWTUseCase) GetUserStatForGeneratingAccessToken(UserId *string) (*string, error) {
	userStat, err := r.JWTRepo.GetUserStatForGeneratingAccessToken(UserId)
	if err != nil {
		return userStat, err
	}
	return userStat, nil
}
