package services

import (
	"api-key-middleware/internal/core/domain"
	"api-key-middleware/internal/core/ports"
	"context"
)

type UserService struct {
	repository ports.UserRepository
}

var _ ports.UserService = &UserService{}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) FindAllProfile(ctx context.Context) ([]*domain.UserProfile, error) {
	return s.repository.FindAllProfile(ctx)
}

func (s *UserService) FindByName(ctx context.Context, name string) (*domain.UserProfile, error) {
	return s.repository.FindByName(ctx, name)
}
