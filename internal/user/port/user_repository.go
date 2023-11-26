package port

import (
	"context"
	"hostel-service/internal/user/domain"
)

type UserRepository interface {
	SearchRoommates(ctx context.Context, filter *domain.RoommateFilter) ([]domain.Roommate, int64, error)
	GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error)
	GetUserProfile(ctx context.Context, userId string) (*domain.UserProfile, error)
	GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error)
	GetByUsername(ctx context.Context, username string) (*domain.UserProfile, error)
	GetById(ctx context.Context, userId string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	UpdateUserSuggest(ctx context.Context, us *domain.UpdateUserSuggest) error
	UpdatePassword(ctx context.Context, userId string, newPassword string) error
	UpdateUserStatus(ctx context.Context, userId string, status string) error
}
