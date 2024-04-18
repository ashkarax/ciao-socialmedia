package handler

import (
	"net/http"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type Auth2oHandler struct {
	AuthCredentials *config.Auth2o
}

func NewAuth2oHandler(authCredentials *config.Auth2o) *Auth2oHandler {
	return &Auth2oHandler{AuthCredentials: authCredentials}
}

func (u *Auth2oHandler) Auth2oCredentialsForMobileApp(c *gin.Context) {

	type AuthCredentials struct {
		ClientId                string `json:"clientid"`
		ProjectId               string `json:"projectid"`
		AuthUri                 string `json:"authuri"`
		TokenUri                string `json:"tokenuri"`
		AuthProviderX509CentUrl string `json:"authproviderx509centurl"`
	}

	var AuthCred AuthCredentials

	AuthCred.ClientId = u.AuthCredentials.ClientId
	AuthCred.AuthUri = u.AuthCredentials.AuthUri
	AuthCred.ProjectId = u.AuthCredentials.ProjectId
	AuthCred.TokenUri = u.AuthCredentials.TokenUri
	AuthCred.AuthProviderX509CentUrl = u.AuthCredentials.AuthProviderX509CentUrl

	response := responsemodels.Responses(http.StatusOK, "success", AuthCred, nil)
	c.JSON(http.StatusOK, response)
}
