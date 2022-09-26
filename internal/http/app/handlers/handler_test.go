package handlers

import (
	"github.com/takaaki-mizuno/go-boilerplate/config"
	"github.com/takaaki-mizuno/go-boilerplate/pkg/database"
	"github.com/takaaki-mizuno/go-boilerplate/pkg/logger"
)

func createHandler() HandlerInterface {
	configInstance, _ := config.LoadTestConfig()
	db, _, _ := database.GetMockDB()

	return NewHandler(
		db,
		configInstance,
		logger.NewTestLogger(),
	)
}
