package usecase

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/ashkarax/ciao-socialmedia/internal/config"
	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	gosmtp "github.com/ashkarax/ciao-socialmedia/pkg/go_smtp"
	hashpassword "github.com/ashkarax/ciao-socialmedia/pkg/hash_password"
	jwttoken "github.com/ashkarax/ciao-socialmedia/pkg/jwt_token"
	randomnumbergenerator "github.com/ashkarax/ciao-socialmedia/pkg/random_number_generator"
	"github.com/go-playground/validator/v10"
)

type UserUseCase struct {
	UserRepo         interfaceRepository.IUserRepo
	tokenSecurityKey *config.Token
	PostRepo         interfaceRepository.IPostRepo
	RelationRepo     interfaceRepository.IRelationRepo
}

func NewUserUseCase(userRepo interfaceRepository.IUserRepo, tokenSecurityKey *config.Token, postRepo interfaceRepository.IPostRepo, relationRepo interfaceRepository.IRelationRepo) interfaceUseCase.IUserUseCase {
	return &UserUseCase{UserRepo: userRepo, tokenSecurityKey: tokenSecurityKey, PostRepo: postRepo, RelationRepo: relationRepo}
}

func (r *UserUseCase) UserSignUp(userData *requestmodels.UserSignUpReq) (responsemodels.SignupData, error) {

	var resSignUp responsemodels.SignupData

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					resSignUp.Name = "should be a valid Name. "
				case "UserName":
					resSignUp.UserName = "should be a valid username. "
				case "Email":
					resSignUp.Email = "should be a valid email address. "
				case "Password":
					resSignUp.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					resSignUp.ConfirmPassword = "should match the first password"
				}
			}
		}
		return resSignUp, errors.New("did't fullfill the signup requirement ")
	}

	if isUserExist := r.UserRepo.IsUserExist(userData.Email); isUserExist {
		resSignUp.IsUserExist = "User Exist ,change email"
		return resSignUp, errors.New("user exists, try again with another email id")
	}
	if isUserExistUserName := r.UserRepo.IsUserExistWithSameUserName(userData.UserName); isUserExistUserName {
		resSignUp.IsUserExist = "User Exist ,change username"
		return resSignUp, errors.New("user exists, try again with another username")
	}

	errRemv := r.UserRepo.DeleteRecentOtpRequestsBefore5min()
	if errRemv != nil {
		return resSignUp, errRemv
	}

	otp := randomnumbergenerator.RandomNumber()
	errOtp := gosmtp.SendVerificationEmailWithOtp(otp, userData.Email, userData.Name)
	if errOtp != nil {
		return resSignUp, errOtp
	}

	expiration := time.Now().Add(5 * time.Minute)

	errTempSave := r.UserRepo.TemporarySavingUserOtp(otp, userData.Email, expiration)
	if errTempSave != nil {
		fmt.Println("Cant save temporary data for otp verification in db")
		return resSignUp, errors.New("OTP verification down,please try after some time")
	}

	hashedPassword := hashpassword.HashPassword(userData.ConfirmPassword)
	userData.Password = hashedPassword

	errCreateUsr := r.UserRepo.CreateUser(userData)
	if errCreateUsr != nil {
		return resSignUp, errCreateUsr
	}

	tempToken, err := jwttoken.TempTokenForOtpVerification(r.tokenSecurityKey.TempVerificationKey, userData.Email)
	if err != nil {
		resSignUp.Token = "error creating temp token for otp verification"
		return resSignUp, errors.New("error creating token")
	}

	resSignUp.Token = tempToken

	return resSignUp, nil

}

func (r *UserUseCase) VerifyOtp(otpData *requestmodels.OtpVerification, TempVerificationToken *string) (responsemodels.OtpVerifResult, error) {
	var otpveriRes responsemodels.OtpVerifResult

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(otpData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					otpData.Otp = "otp should be a 4 digit number"
				}
			}
		}
		return otpveriRes, errors.New("did't fullfill the login requirement ")
	}
	email, unbindErr := jwttoken.UnbindEmailFromClaim(*TempVerificationToken, r.tokenSecurityKey.TempVerificationKey)
	if unbindErr != nil {
		otpveriRes.Token = "invalid token"
		return otpveriRes, unbindErr
	}

	userOTP, expiration, errGetInfo := r.UserRepo.GetOtpInfo(email)
	if errGetInfo != nil {
		return otpveriRes, errGetInfo
	}

	if otpData.Otp != userOTP {
		return otpveriRes, errors.New("invalid OTP")
	}
	if time.Now().After(expiration) {
		return otpveriRes, errors.New("OTP expired")
	}

	changeStatErr := r.UserRepo.ChangeUserStatusActive(email)
	if changeStatErr != nil {
		return otpveriRes, changeStatErr
	}

	userId, fetchErr := r.UserRepo.GetUserId(email)
	if fetchErr != nil {
		return otpveriRes, fetchErr
	}

	accessToken, aTokenErr := jwttoken.GenerateAcessToken(r.tokenSecurityKey.UserSecurityKey, userId)
	if aTokenErr != nil {
		otpveriRes.AccessToken = aTokenErr.Error()
		return otpveriRes, aTokenErr
	}
	refreshToken, rTokenErr := jwttoken.GenerateRefreshToken(r.tokenSecurityKey.UserSecurityKey)
	if rTokenErr != nil {
		otpveriRes.RefreshToken = rTokenErr.Error()
		return otpveriRes, rTokenErr
	}

	otpveriRes.Otp = "verified"
	otpveriRes.AccessToken = accessToken
	otpveriRes.RefreshToken = refreshToken

	return otpveriRes, nil
}

func (r *UserUseCase) UserLogin(loginData *requestmodels.UserLoginReq) (responsemodels.UserLoginRes, error) {
	var resLogin responsemodels.UserLoginRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(loginData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					resLogin.Email = "Enter a valid email"
				case "Password":
					resLogin.Password = "Password should have four or more digit"
				}
			}
			return resLogin, errors.New("did't fullfill the login requirement ")
		}
	}

	hashedPassword, userId, status, errr := r.UserRepo.GetHashPassAndStatus(loginData.Email)
	if errr != nil {
		return resLogin, errr
	}

	passwordErr := hashpassword.CompairPassword(hashedPassword, loginData.Password)
	if passwordErr != nil {
		return resLogin, passwordErr
	}

	if status == "blocked" {
		return resLogin, errors.New("user is blocked by the admin")
	}

	if status == "pending" {
		return resLogin, errors.New("user is on status pending,OTP not verified")
	}

	accessToken, accessTokenerr := jwttoken.GenerateAcessToken(r.tokenSecurityKey.UserSecurityKey, userId)
	if err != accessTokenerr {
		return resLogin, accessTokenerr
	}

	refreshToken, refreshTokenerr := jwttoken.GenerateRefreshToken(r.tokenSecurityKey.UserSecurityKey)
	if err != refreshTokenerr {
		return resLogin, refreshTokenerr
	}

	resLogin.AccessToken = accessToken
	resLogin.RefreshToken = refreshToken
	return resLogin, nil

}

func (r *UserUseCase) ForgotPasswordRequest(userData *requestmodels.ForgotPasswordReq) (responsemodels.ForgotPasswordRes, error) {
	var resData responsemodels.ForgotPasswordRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					resData.Email = "Enter a valid email"
				}
				return resData, errors.New("did't fullfill the login requirement")
			}
		}
	}
	_, _, status, errr := r.UserRepo.GetHashPassAndStatus(userData.Email)
	if errr != nil {
		return resData, errr
	}

	if status == "blocked" {
		return resData, errors.New("user is blocked by the admin")
	}

	if status == "pending" {
		return resData, errors.New("user is on status pending,OTP not verified")
	}

	errRemv := r.UserRepo.DeleteRecentOtpRequestsBefore5min()
	if errRemv != nil {
		return resData, errRemv
	}

	otp := randomnumbergenerator.RandomNumber()
	errOtp := gosmtp.SendRestPasswordEmailOtp(otp, userData.Email)
	if errOtp != nil {
		return resData, errOtp
	}

	expiration := time.Now().Add(5 * time.Minute)

	errTempSave := r.UserRepo.TemporarySavingUserOtp(otp, userData.Email, expiration)
	if errTempSave != nil {
		fmt.Println("Cant save temporary data for otp verification in db")
		return resData, errors.New("OTP verification down,please try after some time")
	}

	tempToken, err := jwttoken.TempTokenForOtpVerification(r.tokenSecurityKey.TempVerificationKey, userData.Email)
	if err != nil {
		resData.Token = "error creating temp token for otp verification"
		return resData, errors.New("error creating token")
	}

	resData.Token = tempToken
	return resData, nil
}

func (r *UserUseCase) ForgotPasswordActual(userData *requestmodels.ForgotPasswordData, TempVerificationToken *string) (responsemodels.ForgotPasswordData, error) {
	var resData responsemodels.ForgotPasswordData

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					resData.Otp = "otp should be a 4 digit number"
				case "Password":
					resData.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					resData.ConfirmPassword = "should match the first password"
				}
				return resData, errors.New("did't fullfill the login requirement")
			}
		}
	}

	email, unbindErr := jwttoken.UnbindEmailFromClaim(*TempVerificationToken, r.tokenSecurityKey.TempVerificationKey)
	if unbindErr != nil {
		resData.Token = "invalid token"
		return resData, unbindErr
	}

	userOTP, expiration, errGetInfo := r.UserRepo.GetOtpInfo(email)
	if errGetInfo != nil {
		return resData, errGetInfo
	}

	if userData.Otp != userOTP {
		return resData, errors.New("invalid OTP")
	}
	if time.Now().After(expiration) {
		return resData, errors.New("OTP expired")
	}

	hashedPassword := hashpassword.HashPassword(userData.ConfirmPassword)

	updateErr := r.UserRepo.UpdateUserPassword(&email, &hashedPassword)
	if updateErr != nil {
		return resData, updateErr
	}

	return resData, nil

}

func (r *UserUseCase) UserProfile(userId *string) (*responsemodels.UserProfile, error) {
	userData, errU := r.UserRepo.GetUserDataLite(userId)
	if errU != nil {
		return nil, errU
	}
	PostCount, errP := r.PostRepo.GetPostCountOfUser(userId)
	if errP != nil {
		return nil, errP
	}
	followerCount, followingCount, errC := r.RelationRepo.GetFollowerAndFollowingCountofUser(userId)
	if errC != nil {
		return nil, errC
	}
	userData.FollowersCount = *followerCount
	userData.FollowingCount = *followingCount
	userData.PostsCount = PostCount

	return userData, nil
}

func (r *UserUseCase) SearchUser(searchInput *requestmodels.SearchRequest) (*[]responsemodels.SearchResp, error) {

	if searchInput.SearchText == "" {
		return nil, fmt.Errorf("search input cannot be empty")
	}

	// Ensure search input length is within acceptable limits
	if len(searchInput.SearchText) < 3 || len(searchInput.SearchText) > 50 {
		return nil, fmt.Errorf("search input length must be between 3 and 50 characters")
	}

	// Ensure search input contains only alphanumeric characters, spaces, or underscores
	validSearch := regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`).MatchString
	if !validSearch(searchInput.SearchText) {
		return nil, fmt.Errorf("search input can only contain letters, numbers, spaces, or underscores")
	}
	usersSlice, errP := r.UserRepo.SearchUserByNameOrUserName(searchInput)
	if errP != nil {
		return nil, errP
	}

	return usersSlice, nil
}

func (r *UserUseCase) UserProfileOfUserB(requestData *requestmodels.FollowRequest) (*responsemodels.UserProfile, error) {

	userData, errU := r.UserRepo.GetUserDataLite(&requestData.OppositeUserID)
	if errU != nil {
		return nil, errU
	}
	PostCount, errP := r.PostRepo.GetPostCountOfUser(&requestData.OppositeUserID)
	if errP != nil {
		return nil, errP
	}
	followerCount, followingCount, errC := r.RelationRepo.GetFollowerAndFollowingCountofUser(&requestData.OppositeUserID)
	if errC != nil {
		return nil, errC
	}
	relationStat, err := r.RelationRepo.UserAFollowingUserBorNot(requestData)
	if err != nil {
		return nil, err
	}

	userData.FollowingStatus = relationStat
	userData.FollowersCount = *followerCount
	userData.FollowingCount = *followingCount
	userData.PostsCount = PostCount

	return userData, nil
}

func (r *UserUseCase) EditUserDetails(editInput *requestmodels.EditUserProfile) (*responsemodels.EditUserProfileResp, error) {
	var respEditUsr responsemodels.EditUserProfileResp

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(editInput)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					respEditUsr.Name = "should be a valid Name. "
				case "UserName":
					respEditUsr.UserName = "should be a valid username. "
				case "Bio":
					respEditUsr.Bio = "Bio can't exceed 60 characters "
				case "Links":
					respEditUsr.Links = "Links can't exceed 20 characters"
				}
			}
		}
		return &respEditUsr, errors.New("did't fullfill the signup requirement ")
	}

	userData, errU := r.UserRepo.GetUserDataLite(&editInput.UserId)
	if errU != nil {
		return nil, errU
	}

	if userData.UserName != editInput.UserName {

		if isUserExistUserName := r.UserRepo.IsUserExistWithSameUserName(editInput.UserName); isUserExistUserName {
			respEditUsr.IsUserExist = "User Exist ,change username"
			return &respEditUsr, errors.New("user exists, try again with another username")
		}

	}

	err = r.UserRepo.UpdateUserDetails(editInput)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
