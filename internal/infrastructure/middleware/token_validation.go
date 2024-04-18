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
