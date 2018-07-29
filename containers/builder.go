package containers

import (
	"authentication/config"
	controllers "authentication/controllers/v1"
	"authentication/handlers"
	"authentication/middleware"
	"authentication/repositories"
	"authentication/server"
	"authentication/services"

	"go.uber.org/dig"
)

// BuildContainer returns a container with all app dependencies built in
func BuildContainer() *dig.Container {
	container := dig.New()

	// config
	container.Provide(config.NewConfig)

	// persistance layer
	container.Provide(repositories.NewDBCollections)
	container.Provide(repositories.NewUserRepository)
	container.Provide(repositories.NewRoleRepository)

	// services
	container.Provide(services.NewUserService)
	container.Provide(services.NewRoleService)
	container.Provide(services.NewTokenService)

	// controllers
	container.Provide(controllers.NewRoleController)
	container.Provide(controllers.NewUserController)

	// generic http layer
	container.Provide(middleware.NewMiddleware)
	container.Provide(handlers.NewHttpHandlers)

	// server
	container.Provide(server.NewServer)

	return container
}
