package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type IUserUseCase interface {
	UserSignUp(*requestmodels.UserSignUpReq) (responsemodels.SignupData, error)
	VerifyOtp(*requestmodels.OtpVerification, *string) (responsemodels.OtpVerifResult, error)
	UserLogin(*requestmodels.UserLoginReq) (responsemodels.UserLoginRes, error)

	ForgotPasswordRequest(*requestmodels.ForgotPasswordReq) (responsemodels.ForgotPasswordRes, error)
	ForgotPasswordActual(*requestmodels.ForgotPasswordData, *string) (responsemodels.ForgotPasswordData, error)

	UserProfile(*string) (*responsemodels.UserProfile, error)
}
