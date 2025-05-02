package pet

import (
	"context"

	"github.com/webbsalad/pet/internal/model"
)

type Repository interface {
	PutPet(ctx context.Context, pet model.Pet) error
	GetPet(ctx context.Context) (model.Pet, error)
}
