package routes

import (
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/handler"
	JWTmiddleware "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, jwtMiddleWare *JWTmiddleware.JWTmiddleware, user *handler.UserHandler, post *handler.PostHandler) {

	engin.POST("/signup", user.UserSignUp)
	engin.POST("/verify", user.UserOTPVerication)
	engin.POST("/login", user.UserLogin)
	engin.POST("/forgotpassword", user.ForgotPasswordRequest)
	engin.PATCH("/resetpassword", user.ResetPassword)

	//engin.GET("/profile", user.GetUserProfile)

	engin.Use(jwtMiddleWare.UserAuthorization)
	{

		postmanagement := engin.Group("/post")
		{
			postmanagement.POST("/", post.AddNewPost)
			postmanagement.GET("/", post.GetAllPostByUser)
			postmanagement.DELETE("/", post.DeletePost)
			//postmanagement.PATCH("/", post.EditPost)

		}

	}

}
