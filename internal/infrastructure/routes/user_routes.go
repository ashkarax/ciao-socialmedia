package routes

import (
	"github.com/ashkarax/ciao-socialmedia/internal/infrastructure/handler"
	JWTmiddleware "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, jwtMiddleWare *JWTmiddleware.JWTmiddleware, user *handler.UserHandler, post *handler.PostHandler, relation *handler.RelationHandler, auth2o *handler.Auth2oHandler) {

	engin.POST("/signup", user.UserSignUp)
	engin.POST("/verify", user.UserOTPVerication)
	engin.POST("/login", user.UserLogin)
	engin.POST("/forgotpassword", user.ForgotPasswordRequest)
	engin.PATCH("/resetpassword", user.ResetPassword)
	engin.GET("/accessgenerator", jwtMiddleWare.AccessRegenerator)

	authmanagement := engin.Group("/auth")
	{
		authmanagement.GET("/credentials", auth2o.Auth2oCredentialsForMobileApp)
	}

	engin.Use(jwtMiddleWare.UserAuthorization)
	{

		profilemanagement := engin.Group("/profile")
		{
			profilemanagement.GET("/", user.GetUserProfile)
			profilemanagement.PATCH("/edit", user.EditUserProfile)
			profilemanagement.GET("/followers", relation.GetFollowersDetails)
			profilemanagement.GET("/following", relation.GetFollowingDetails)
		}

		postmanagement := engin.Group("/post")
		{
			postmanagement.POST("/", post.AddNewPost)
			postmanagement.GET("/", post.GetAllPostByUser)
			postmanagement.DELETE("/", post.DeletePost)
			postmanagement.PATCH("/", post.EditPost)

			postmanagement.GET("/userrelatedposts", post.GetAllRelatedPostsForHomeScreen)

			likemanagement := postmanagement.Group("/like")
			{
				likemanagement.POST("/:postid", post.LikePost)
				likemanagement.DELETE("/:postid", post.UnLikePost)
			}

		}
		exploremanagement := engin.Group("/explore")
		{
			exploremanagement.GET("/", post.GetMostLovedPostsFromGlobalUser)
			exploremanagement.GET("/profile/:userBid", user.GetAnotherUserProfile)

			searchmanagement := exploremanagement.Group("/search")
			{
				searchmanagement.GET("/user/:searchtext", user.SearchUser)

			}
		}
		followRelationshipManagement := engin.Group("/relation")
		{
			followRelationshipManagement.POST("/follow/:followingId", relation.Follow)
			followRelationshipManagement.DELETE("/unfollow/:followingId", relation.UnFollow)

		}

	}

}
