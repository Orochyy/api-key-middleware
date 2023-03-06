package user

import (
	"api-key-middleware/internal/core/ports"
	"api-key-middleware/internal/requests"
	"github.com/gofiber/fiber/v2"
)

type ShowRequest struct {
	requests.Request
	ShowRequestData
}

type ShowRequestData struct {
	Name string `json:"username" validate:"-"`
}

func NewShowRequest(ctx *fiber.Ctx, validator ports.Validator) *ShowRequest {
	return &ShowRequest{Request: requests.Request{Ctx: ctx, Validator: validator}}
}

func (r *ShowRequest) Validate() error {
	r.Name = r.Ctx.Query("username")

	return r.Validator.Struct(r)
}
