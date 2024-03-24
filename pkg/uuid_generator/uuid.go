package uuidgenerator

import "github.com/google/uuid"

func ReturnUuid() *string{
	randomName := uuid.New().String()	
	return &randomName
}