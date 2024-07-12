//go:build wireinject
// +build wireinject

// Package di contains dependency injection setup using Wire for github.com/nuno-bastos/gin-gonic-wire-api application.
// Wire is used to automate the setup of dependencies, reducing boilerplate and ensuring
// type safety in the initialization of application components.
package wire_di

import (
	"github.com/google/wire"

	server "github.com/nuno-bastos/gin-gonic-wire-api/api"
	controller "github.com/nuno-bastos/gin-gonic-wire-api/api/controller"
	db "github.com/nuno-bastos/gin-gonic-wire-api/db"
	repository "github.com/nuno-bastos/gin-gonic-wire-api/repo"
	service "github.com/nuno-bastos/gin-gonic-wire-api/service"
	filters "github.com/nuno-bastos/gin-gonic-wire-api/service/filters"
)

// Inject initializes the github.com/nuno-bastos/gin-gonic-wire-api application by setting up all required dependencies
// and returns an instance of the HTTP server.
func Inject() (*server.ServerHTTP, error) {
	// Wire.Build sets up the dependency graph
	wire.Build(
		// Database connection setup.
		db.ConnectDatabase,

		// Repositories setup.
		repository.NewManagerChainRepository,
		repository.NewUserRepository,
		repository.NewSecurityMatrixRuleRepository,

		// Filters setup.
		filters.SecurityFiltersSet,

		// Services setup.
		service.NewSecurityMatrixCalculatorService,

		// Controllers setup.
		controller.NewCalculateGoSecurityMatrixController,

		// HTTP Server setup.
		server.StartServer,
	)

	return &server.ServerHTTP{}, nil // Initialize and return the HTTP server instance
}
