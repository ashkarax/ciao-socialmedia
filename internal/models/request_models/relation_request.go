package requestmodels

type FollowRequest struct {
	OppositeUserID string `json:"following_id"`
	UserID         string
}
