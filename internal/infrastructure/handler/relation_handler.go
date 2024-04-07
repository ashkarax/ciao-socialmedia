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

	if err := c.BindJSON(&FollowRequestData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FollowRequestData.UserID = userIdString

	err := u.RelationUseCase.Follow(&FollowRequestData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to follow userB", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully", nil, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *RelationHandler) UnFollow(c *gin.Context) {
	var FollowRequestData requestmodels.FollowRequest
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	if err := c.BindJSON(&FollowRequestData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

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


