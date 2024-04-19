package requestmodels

type UserSignUpReq struct {
	Name            string `json:"name" validate:"required,gte=3"`
	UserName        string `json:"username" validate:"required,gte=3"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,gte=3"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type OtpVerification struct {
	Otp string `json:"otp" validate:"required,len=4,number"`
}

type UserLoginReq struct {
	Email    string `json:"email"    validate:"required,gte=8"`
	Password string `json:"password" validate:"required,min=4"`
}

type ForgotPasswordReq struct {
	Email string `json:"email"    validate:"required,gte=8"`
}

type ForgotPasswordData struct {
	Otp             string `json:"otp" validate:"required,len=4,number"`
	Password        string `json:"password" validate:"required,gte=3"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type EditUserProfile struct {
	Name     string `json:"name" validate:"required,gte=3"`
	UserName string `json:"username" validate:"required,gte=3"`
	Bio      string `json:"bio" validate:"lte=50"`
	Links    string `json:"links" validate:"lte=25"`

	UserId string
}
