package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type IPostUseCase interface {
	AddNewPost(*requestmodels.AddPostData) (*responsemodels.AddPostResp, error)
	GetAllPostByUser(*string) (*[]responsemodels.PostData, error)

	DeletePost(*requestmodels.PostId, *string) error

	LikePost(*requestmodels.LikeRequest) error
	UnLikePost(*requestmodels.LikeRequest) error

	GetAllRelatedPostsForHomeScreen(*string) (*[]responsemodels.PostData, error)
	GetMostLovedPostsFromGlobalUser(*string) (*[]responsemodels.PostData, error)
}
