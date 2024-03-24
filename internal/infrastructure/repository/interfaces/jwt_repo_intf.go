package interfaceRepository

type IJWTRepo interface {
	GetUserStatForGeneratingAccessToken(*string) (*string, error)
}
