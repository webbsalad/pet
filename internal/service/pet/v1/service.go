package v1

import (
	"github.com/webbsalad/pet/internal/repository/pet"
	pet_service "github.com/webbsalad/pet/internal/service/pet"
)

type Service struct {
	petRepository pet.Repository
}

func NewService(
	petRepository pet.Repository,
) pet_service.Service {
	return &Service{
		petRepository: petRepository,
	}
}
