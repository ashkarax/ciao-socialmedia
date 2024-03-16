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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, otpVerifErr := u.UserUseCase.VerifyOtp(&otpData, temptoken)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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