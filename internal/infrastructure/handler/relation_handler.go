package handler

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type RelationHandler struct {
	RelationUseCase interfaceUseCase.IRelationUseCase
}

func NewRelationHandler(relationUseCase interfaceUseCase.IRelationUseCase) *RelationHandler {
	return &RelationHandler{RelationUseCase: relationUseCase}
}

func (u *RelationHandler) Follow(c *gin.Context) {
	var FollowRequestData requestmodels.FollowRequest
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	followingId := c.Param("followingId")
	if followingId == "" {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:empty input)", nil, "no following id found as param")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FollowRequestData.OppositeUserID = followingId
	FollowRequestData.UserID = userIdString

	err := u.RelationUseCase.Follow(&FollowRequestData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to follow userB", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully followed", nil, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *RelationHandler) UnFollow(c *gin.Context) {
	var FollowRequestData requestmodels.FollowRequest
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	followingId := c.Param("followingId")
	if followingId == "" {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:empty input)", nil, "no following id found as param")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FollowRequestData.OppositeUserID = followingId
	FollowRequestData.UserID = userIdString

	err := u.RelationUseCase.UnFollow(&FollowRequestData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to unfollow userB", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully unfollowed", nil, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *RelationHandler) GetFollowersDetails(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	followersInfo, err := u.RelationUseCase.GetFollowersDetailsOfUser(&userIdString)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to get followers info", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully fetched followers info", followersInfo, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *RelationHandler) GetFollowingDetails(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	followersInfo, err := u.RelationUseCase.GetFollowingDetailsOfUser(&userIdString)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to get following info", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully fetched following info", followersInfo, nil)
	c.JSON(http.StatusOK, finalReslt)
}
