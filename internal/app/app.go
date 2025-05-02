package app

import (
	"github.com/webbsalad/pet/internal/config"
	"github.com/webbsalad/pet/internal/file_storage"
	"go.uber.org/fx"

	pet_api "github.com/webbsalad/pet/internal/api/pet"
	pb "github.com/webbsalad/pet/internal/pb/github.com/webbsalad/pet"
	pet_repository "github.com/webbsalad/pet/internal/repository/pet/storage"
	pet_service "github.com/webbsalad/pet/internal/service/pet/v1"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			file_storage.InitFileStorage,
		),

		grpcOption(),
		gatewayOption(),

		servicesOption(),
	)
}

func servicesOption() fx.Option {
	return fx.Options(
		fx.Provide(

			pet_api.NewImplementation,

			pet_service.NewService,

			pet_repository.NewRepository,
		),
		fx.Invoke(
			pb.RegisterPetServiceServer,
		),
	)
}
