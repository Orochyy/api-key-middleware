package ports

import "github.com/gofiber/fiber/v2"

type HealthChecksHandlers interface {
	HealthChecks(ctx *fiber.Ctx) error
}
