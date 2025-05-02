package convertor

import (
	"errors"

	"github.com/webbsalad/pet/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertError(err error) error {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())

	}

	return status.Errorf(codes.InvalidArgument, "internal server error: %v", err)
}
