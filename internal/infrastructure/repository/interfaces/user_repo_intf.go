package interfaceRepository

import (
	"time"

	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
)

type IUserRepo interface {
	IsUserExist(string) bool
	IsUserExistWithSameUserName(string) bool
	CreateUser(*requestmodels.UserSignUpReq) error
	ChangeUserStatusActive(string) error
	GetUserId(string) (string, error)
	GetHashPassAndStatus(string) (string, string, string, error)
	DeleteRecentOtpRequestsBefore5min() error
	TemporarySavingUserOtp(int, string, time.Time) error
	GetOtpInfo(string) (string, time.Time, error)

	UpdateUserPassword(*string, *string) error

	GetUserDataLite(*string) (*responsemodels.UserProfile, error)

	SearchUserByNameOrUserName(*requestmodels.SearchRequest) (*[]responsemodels.SearchResp, error)

	UpdateUserDetails(*requestmodels.EditUserProfile) error
}
