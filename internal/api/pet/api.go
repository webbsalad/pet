package pet

import (
	desc "github.com/webbsalad/pet/internal/pb/github.com/webbsalad/pet"
	"github.com/webbsalad/pet/internal/service/pet"
)

type Implementation struct {
	desc.UnimplementedPetServiceServer

	petService pet.Service
}

func NewImplementation(petService pet.Service) desc.PetServiceServer {
	return &Implementation{
		petService: petService,
	}
}
