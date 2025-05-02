package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrPetNotFound = fmt.Errorf("pet not found: %w", ErrNotFound)
)
