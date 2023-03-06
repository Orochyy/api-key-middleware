package middleware

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AuthMiddleware(db *sql.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("api-key")
		if apiKey == "" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM auth WHERE `api-key` = ?)", apiKey).Scan(&exists)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		if !exists {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
		return ctx.Next()
	}
}
