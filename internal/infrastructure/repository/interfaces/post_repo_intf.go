package interfaceRepository

import (
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type IPostRepo interface {
	AddNewPost(postData *requestmodels.AddPostData) error
	GetAllActivePostByUser(*string) (*[]responsemodels.PostData, error)
	GetPostMediaById(*string) (*[]string, error)
	DeletePostById(*string, *string) error
	DeletePostMedias(*string) error
}
