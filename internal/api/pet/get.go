package pet

import (
	"context"

	"github.com/webbsalad/pet/internal/convertor"
	desc "github.com/webbsalad/pet/internal/pb/github.com/webbsalad/pet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPet(ctx context.Context, req *desc.GetPetRequest) (*desc.Pet, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	pet, err := i.petService.GetPet(ctx)
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &desc.Pet{
		Ascii:       pet.Ascii,
		Description: pet.Description,
	}, nil

}
