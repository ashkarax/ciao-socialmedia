package handler

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase interfaceUseCase.IUserUseCase
}

func NewUserHandler(userUseCase interfaceUseCase.IUserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (u *UserHandler) UserSignUp(c *gin.Context) {
	var userSignupData requestmodels.UserSignUpReq
	if err := c.BindJSON(&userSignupData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "signup failed(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resSignup, err := u.UserUseCase.UserSignUp(&userSignupData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "signup failed", resSignup, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "signup success", resSignup, nil)
	c.JSON(http.StatusOK, response)

}

func (u *UserHandler) UserOTPVerication(c *gin.Context) {

	var otpData requestmodels.OtpVerification
	temptoken := c.Request.Header.Get("x-temp-token")

	if err := c.BindJSON(&otpData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "OTP verification failed(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, otpVerifErr := u.UserUseCase.VerifyOtp(&otpData, &temptoken)
	if otpVerifErr != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "OTP verification failed", result, otpVerifErr.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusAccepted, "OTP verification success", result, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	var loginData requestmodels.UserLoginReq

	if err := c.BindJSON(&loginData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "login failed(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resLogin, err := u.UserUseCase.UserLogin(&loginData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "login failed", resLogin, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "login success", resLogin, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) ForgotPasswordRequest(c *gin.Context) {
	var forgotReqData requestmodels.ForgotPasswordReq

	if err := c.BindJSON(&forgotReqData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resForgotPass, err := u.UserUseCase.ForgotPasswordRequest(&forgotReqData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed", resForgotPass, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "otp generated succesfully,continue with /resetpassword", resForgotPass, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) ResetPassword(c *gin.Context) {
	var requestData requestmodels.ForgotPasswordData

	temptoken := c.Request.Header.Get("x-temp-token")

	if err := c.BindJSON(&requestData); err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed request(possible-reason:no json input)", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resForgotPass, err := u.UserUseCase.ForgotPasswordActual(&requestData, &temptoken)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to reset password", resForgotPass, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "password updated success", resForgotPass, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetUserProfile(c *gin.Context) {
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	userData, err := u.UserUseCase.UserProfile(&userIdString)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to get user profile", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", userData, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *UserHandler) SearchUser(c *gin.Context) {

	var searchInput requestmodels.SearchRequest

	searchText := c.Param("searchtext")
	searchInput.SearchText = searchText

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	searchInput.MyId = userIdString

	usersSlice, err := u.UserUseCase.SearchUser(&searchInput)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed to find users", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "users found succesfully", usersSlice, nil)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetAnotherUserProfile(c *gin.Context) {
	var requestData requestmodels.FollowRequest

	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	userbId := c.Param("userBid")

	requestData.OppositeUserID = userbId
	requestData.UserID = userIdString

	userData, err := u.UserUseCase.UserProfileOfUserB(&requestData)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "failed to get user profile", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "succesfully retreival", userData, nil)
	c.JSON(http.StatusOK, finalReslt)
}


