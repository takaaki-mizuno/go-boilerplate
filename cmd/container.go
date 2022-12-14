package cmd

import (
	"github.com/takaaki-mizuno/go-boilerplate/config"
	adminHandler "github.com/takaaki-mizuno/go-boilerplate/internal/http/admin/handlers"
	appHandler "github.com/takaaki-mizuno/go-boilerplate/internal/http/app/handlers"
	appMiddlewares "github.com/takaaki-mizuno/go-boilerplate/internal/http/app/middlewares"
	// {{ REPLACE repository_import }}
	// {{ REPLACE END repository_import }}
	"github.com/takaaki-mizuno/go-boilerplate/pkg/database"
	"github.com/takaaki-mizuno/go-boilerplate/pkg/logger"
	commonMiddlewares "github.com/takaaki-mizuno/go-boilerplate/pkg/middlewares"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(logger.NewLogger)
	_ = container.Provide(database.InitDatabase)
	_ = container.Provide(database.NewTransaction)

	// Handlers for User APIs
	_ = container.Provide(appHandler.NewHandler)

	// Handler for Admin APIs
	_ = container.Provide(adminHandler.NewHandler)

	// Middlewares for All APIs
	_ = container.Provide(commonMiddlewares.SecurityHeaders, dig.Name("securityHeaders"))
	_ = container.Provide(commonMiddlewares.RequestID, dig.Name("requestID"))
	_ = container.Provide(commonMiddlewares.Logger, dig.Name("logger"))

	// Middlewares for App APIs
	_ = container.Provide(appMiddlewares.RequestHeaders, dig.Name("requestHeaders"))

	// Repositories
	// {{ REPLACE repository }}
	// {{ REPLACE END repository }}

	return container
}
