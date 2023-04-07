package di

import (
	"health_service/service/http/handler"
	"health_service/service/http/router"
	"health_service/service/repository"
	"health_service/service/usecase"

	"go.uber.org/fx"
)

// Module provided to fx
var Module = fx.Provide(
	// provide usecase
	ProvideUserUsecase,
	// handler
	provideUserHandler,

	// router
	provodeUserRouter,

	// provide repositories
	provideUserRepo,
)

// Handler
func provideUserHandler(userUsecase usecase.IuserUsecase) handler.UserHandler {
	return handler.NewUserhandler(userUsecase)
}

// Router
func provodeUserRouter(
	userHandlerr handler.UserHandler,
) router.UserRouter {
	return router.NewUserRouter(userHandlerr)
}

// Usecases
func ProvideUserUsecase(
	userRepo repository.IuserRepo,
) usecase.IuserUsecase {
	return usecase.NewUserUsecase(userRepo)
}

// Repositories
func provideUserRepo() repository.IuserRepo {
	return repository.NewUserRepo()
}
