package JWTmiddleware

import (
	"fmt"
	"net/http"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	jwttoken "github.com/ashkarax/ciao-socialmedia/pkg/jwt_token"
	"github.com/gin-gonic/gin"
)

type JWTmiddleware struct {
	keys       *config.Token
	JWTUseCase interfaceUseCase.IJWTUseCase
}

func NewJWTMiddleware(JwtUseCase interfaceUseCase.IJWTUseCase, keys *config.Token) *JWTmiddleware {
	return &JWTmiddleware{JWTUseCase: JwtUseCase, keys: keys}
}

func (r *JWTmiddleware) UserAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("x-access-token")

	if accessToken == "" {
		response := responsemodels.Responses(http.StatusUnauthorized, "no access  token found", nil, "In your request,The Required tokens to get into this page are not available.")
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	userId, err := jwttoken.VerifyAccessToken(accessToken, r.keys.UserSecurityKey)
	if err != nil {
		if userId == "" {
			response := responsemodels.Responses(http.StatusUnauthorized, "Token Tampared ,Id not accessible", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		response := responsemodels.Responses(http.StatusUnauthorized, "Token Tampared ,token verification failed", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	c.Set("userId", userId)
	fmt.Println("access token is upto date")
	c.Next()

}

func (r *JWTmiddleware) AccessRegenerator(c *gin.Context) {
	accessToken := c.Request.Header.Get("x-access-token")
	refreshToken := c.Request.Header.Get("x-refresh-token")

	if accessToken == "" || refreshToken == "" {
		response := responsemodels.Responses(http.StatusUnauthorized, "no access or refresh token found", nil, "In your request,The Required tokens to get into this page are not available.")
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	userId, err := jwttoken.VerifyAccessToken(accessToken, r.keys.UserSecurityKey)
	if err != nil {
		if userId == "" {
			response := responsemodels.Responses(http.StatusUnauthorized, "Access Token Tampared ,Id not accessible", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
	}

	errn := jwttoken.VerifyRefreshToken(refreshToken, r.keys.UserSecurityKey)
	if errn != nil {
		response := responsemodels.Responses(http.StatusUnauthorized, "Refresh Token Tampared", nil, errn.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	status, err1 := r.JWTUseCase.GetUserStatForGeneratingAccessToken(&userId)
	if err1 != nil || *status == "blocked" {
		response := responsemodels.Responses(http.StatusUnauthorized, "Id not accessible", status, err1.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	newAcessToken, err2 := jwttoken.GenerateAcessToken(r.keys.UserSecurityKey, userId)
	if err2 != nil {
		response := responsemodels.Responses(http.StatusUnauthorized, "Failed to generate New Access Token", nil, err2.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	response := responsemodels.Responses(http.StatusOK, "Generated New Access Token", newAcessToken, nil)
	c.JSON(http.StatusOK, response)
}
