package ports

import (
	"api-key-middleware/internal/core/domain"
	"context"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers interface {
	Get(ctx *fiber.Ctx) error
}

type UserService interface {
	FindAllProfile(ctx context.Context) ([]*domain.UserProfile, error)
	FindByName(ctx context.Context, name string) (*domain.UserProfile, error)
}

type UserRepository interface {
	FindAllProfile(ctx context.Context) ([]*domain.UserProfile, error)
	FindByName(ctx context.Context, name string) (*domain.UserProfile, error)
}
