package handlers

import (
	"api-key-middleware/internal/core/ports"
	"api-key-middleware/internal/requests/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.UserService
	validator   ports.Validator
}

var _ ports.UserHandlers = &UserHandlers{}

func NewUserHandlers(userService ports.UserService, validator ports.Validator) *UserHandlers {
	return &UserHandlers{
		userService: userService,
		validator:   validator,
	}
}

func (h *UserHandlers) Get(ctx *fiber.Ctx) error {
	request := user.NewShowRequest(ctx, h.validator)
	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if request.Name == "" {
		userData, err := h.userService.FindAllProfile(ctx.Context())
		if err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
			"data":    userData,
		})
	} else {
		userData, err := h.userService.FindByName(ctx.Context(), request.Name)
		if err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
			"data":    userData,
		})
	}

}
