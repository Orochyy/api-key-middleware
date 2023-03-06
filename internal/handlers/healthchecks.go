package handlers

import (
	"api-key-middleware/internal/core/ports"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HealthChecksHandlers struct {
	cache        ports.Cache
	dbConnection *sql.DB
}

var _ ports.HealthChecksHandlers = (*HealthChecksHandlers)(nil)

func NewHealthChecksHandlers(cache ports.Cache, dbConnection *sql.DB) *HealthChecksHandlers {
	return &HealthChecksHandlers{
		cache:        cache,
		dbConnection: dbConnection,
	}
}

// HealthChecks godoc
// @Tags healthChecks
// @Summary healthChecks
// @Description healthChecks
// @Accept  json
// @Produce  json
// @Success 200 {object} HTTPSuccess "ok"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /health_checks [get]
func (h *HealthChecksHandlers) HealthChecks(ctx *fiber.Ctx) error {
	var cacheStatus, dbStatus string
	cacheStatus = http.StatusText(http.StatusOK)
	dbStatus = http.StatusText(http.StatusOK)

	err := h.cache.Ping()
	if err != nil {
		cacheStatus = err.Error()
	}

	err = h.dbConnection.Ping()
	if err != nil {
		dbStatus = err.Error()
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"cache":    cacheStatus,
		"database": dbStatus,
		"app":      http.StatusText(http.StatusOK),
	})
}
