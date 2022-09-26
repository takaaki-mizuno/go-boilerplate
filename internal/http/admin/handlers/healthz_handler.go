package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takaaki-mizuno/go-boilerplate/internal/http/admin/responses"
	"github.com/takaaki-mizuno/go-boilerplate/pkg/database"
)

// HealthzHandler ... endpoint for checking health
func (handler *Handler) HealthzHandler(c *gin.Context) {
	if database.PingDB(handler.db) {
		c.JSON(http.StatusOK, responses.NewSuccessStatus())
		return
	}

	c.JSON(http.StatusServiceUnavailable, responses.NewServiceUnavailableErrorStatus())
}
