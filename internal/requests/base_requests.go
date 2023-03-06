package requests

import (
	"api-key-middleware/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Ctx       *fiber.Ctx
	Validator ports.Validator
}
