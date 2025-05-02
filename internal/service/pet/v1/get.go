package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pet/internal/model"
)

func (s *Service) GetPet(ctx context.Context) (model.Pet, error) {
	pet, err := s.petRepository.GetPet(ctx)
	if err != nil {
		return model.Pet{}, fmt.Errorf("get: %w", err)
	}

	return pet, nil
}
