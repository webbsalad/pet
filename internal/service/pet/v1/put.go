package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pet/internal/model"
)

func (s *Service) PutPet(ctx context.Context, pet model.Pet) error {
	if err := s.petRepository.PutPet(ctx, pet); err != nil {
		return fmt.Errorf("put pet: %w", err)
	}

	return nil
}
