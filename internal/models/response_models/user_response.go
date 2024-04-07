package responsemodels

type SignupData struct {
	Name            string `json:"name,omitempty"`
	UserName        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	OTP             string `json:"otp,omitempty"`
	Token           string `json:"token,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	IsUserExist     string `json:"isUserExist,omitempty"`
}
type OtpVerifResult struct {
	Email        string `json:"email,omitempty"`
	Otp          string `json:"otp,omitempty"`
	Result       string `json:"result,omitempty"`
	Token        string `json:"token,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}

type UserLoginRes struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}

type ForgotPasswordRes struct {
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

type ForgotPasswordData struct {
	Token           string `json:"token,omitempty"`
	Otp             string `json:"otp,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
}

type UserProfile struct {
	UserId uint `json:"userid"  gorm:"column:id"`

	Name              string `json:"name"`
	UserName          string `json:"username"`
	UserProfileImgURL string `json:"userprofileimageurl,omitempty"`

	PostsCount     uint `json:"posts_count"`
	FollowersCount uint `json:"followers_count"`
	FollowingCount uint `json:"following_count"`

	//for userB only
	FollowedBy      string `json:"followedby,omitempty"`
	FollowingStatus bool   `json:"following_status,omitempty"`
}
