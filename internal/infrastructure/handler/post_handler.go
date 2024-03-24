package handler

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostUseCase interfaceUseCase.IPostUseCase
}

func NewPostHandler(postUseCase interfaceUseCase.IPostUseCase) *PostHandler {
	return &PostHandler{PostUseCase: postUseCase}
}

func (u *PostHandler) AddNewPost(c *gin.Context) {
	var postData requestmodels.AddPostData

	UserId, _ := c.Get("userId")
	UserIdString, _ := UserId.(string)
	postData.UserId = UserIdString

	if err := c.ShouldBind(&postData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add post", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	resPostData, err := u.PostUseCase.AddNewPost(&postData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add Post", resPostData, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Post added succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *PostHandler) GetAllPostByUser(c *gin.Context) {
	UserId, _ := c.Get("userId")
	UserIdString, _ := UserId.(string)

	resPostData, err := u.PostUseCase.GetAllPostByUser(&UserIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch Posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Post fetched succesfully", resPostData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *PostHandler) DeletePost(c *gin.Context) {
	var postId requestmodels.PostId

	UserId, _ := c.Get("userId")
	UserIdString, _ := UserId.(string)

	if err := c.ShouldBind(&postId); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete post", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if postId.PostId == "" {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete post", nil, "no input recieved")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := u.PostUseCase.DeletePost(&postId, &UserIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete Post", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Post deleted succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}
