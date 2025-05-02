package pet

import (
	"context"

	"github.com/webbsalad/pet/internal/convertor"
	"github.com/webbsalad/pet/internal/model"
	desc "github.com/webbsalad/pet/internal/pb/github.com/webbsalad/pet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) PutPet(ctx context.Context, req *desc.PutPetRequest) (*desc.PutPetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	if err := i.petService.PutPet(ctx, model.Pet{
		Ascii:       req.GetAscii(),
		Description: req.GetDescription(),
	}); err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &desc.PutPetResponse{}, nil

}
