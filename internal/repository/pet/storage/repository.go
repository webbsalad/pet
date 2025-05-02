package storage

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/webbsalad/pet/internal/file_storage"
	"github.com/webbsalad/pet/internal/model"
	"github.com/webbsalad/pet/internal/repository/pet"
)

type Repository struct {
	fs *file_storage.FileStorage
}

func NewRepository(fs *file_storage.FileStorage) (pet.Repository, error) {
	return &Repository{fs: fs}, nil
}

func (r *Repository) PutPet(ctx context.Context, pet model.Pet) error {
	if err := r.fs.Set(pet); err != nil {
		return fmt.Errorf("save pet to storage: %w", err)
	}
	return nil
}

func (r *Repository) GetPet(ctx context.Context) (model.Pet, error) {
	var p model.Pet

	err := r.fs.Get(&p)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return model.Pet{}, model.ErrPetNotFound
		}
		return model.Pet{}, fmt.Errorf("cannot read pet from storage: %w", err)
	}

	if p.Ascii == "" && p.Description == "" {
		return model.Pet{}, model.ErrPetNotFound
	}

	return p, nil
}
