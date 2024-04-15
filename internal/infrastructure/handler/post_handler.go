package handler

import (
	"fmt"
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

func (u *PostHandler) LikePost(c *gin.Context) {
	var LikeRequestData requestmodels.LikeRequest
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	postID := c.Param("postid")

	if postID == "" {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:empty input)", nil, "no post id found as param")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	LikeRequestData.PostID = postID
	LikeRequestData.UserID = userIdString

	err := u.PostUseCase.LikePost(&LikeRequestData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to like post", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "like added succesfully", nil, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *PostHandler) UnLikePost(c *gin.Context) {
	var LikeRequestData requestmodels.LikeRequest
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	postID := c.Param("postid")

	if postID == "" {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:empty input)", nil, "no post id found as param")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	LikeRequestData.PostID = postID
	LikeRequestData.UserID = userIdString

	err := u.PostUseCase.UnLikePost(&LikeRequestData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to unlike post", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "unliked post succesfully", nil, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *PostHandler) GetAllRelatedPostsForHomeScreen(c *gin.Context) {
	UserId, _ := c.Get("userId")
	UserIdString, _ := UserId.(string)

	resPostData, err := u.PostUseCase.GetAllRelatedPostsForHomeScreen(&UserIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch Posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Post fetched succesfully", resPostData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *PostHandler) GetMostLovedPostsFromGlobalUser(c *gin.Context) {
	UserId, _ := c.Get("userId")
	UserIdString, _ := UserId.(string)

	resPostData, err := u.PostUseCase.GetMostLovedPostsFromGlobalUser(&UserIdString)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch Posts", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Post fetched succesfully", resPostData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *PostHandler) EditPost(c *gin.Context) {
	var editInput requestmodels.EditPost

	UserId, _ := c.Get("userId")
	UserIdString, _ := UserId.(string)

	editInput.UserId = UserIdString

	if err := c.ShouldBind(&editInput); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't edit post", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println("from handler:", editInput)
	editResp, err := u.PostUseCase.EditPost(&editInput)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't edit Post", editResp, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "Post edited succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}
