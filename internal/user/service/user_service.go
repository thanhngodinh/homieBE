package service

import (
	"context"
	hostel_port "hostel-service/internal/hostel/port"
	"hostel-service/internal/user/domain"
	"hostel-service/internal/user/port"
)

type UserService interface {
	UpdateUserSuggest(ctx context.Context, userUpdate *domain.UpdateUserSuggest) error
	GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (int64, error)
}

func NewUserService(
	userRepo port.UserRepository,
	hostelRepo hostel_port.HostelRepository,
) UserService {
	return &userService{
		userRepo:   userRepo,
		hostelRepo: hostelRepo,
	}
}

type userService struct {
	userRepo   port.UserRepository
	hostelRepo hostel_port.HostelRepository
}

func (s *userService) UpdateUserSuggest(ctx context.Context, userUpdate *domain.UpdateUserSuggest) error {
	return s.userRepo.UpdateUserSuggest(ctx, userUpdate)
}

func (s *userService) GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error) {
	return s.userRepo.GetUserSuggest(ctx, userId)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.userRepo.GetByUsername(ctx, username)
}

func (s *userService) Create(ctx context.Context, user *domain.User) (int64, error) {
	return s.userRepo.Create(ctx, user)
}
