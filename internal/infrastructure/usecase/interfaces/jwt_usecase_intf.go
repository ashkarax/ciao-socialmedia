package interfaceUseCase

type IJWTUseCase interface {
	GetUserStatForGeneratingAccessToken(*string) (*string, error)
}
