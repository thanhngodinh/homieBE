package port

import (
	"context"
	"hostel-service/internal/user/domain"
	"time"
)

type UserRepository interface {
	SearchRoommates(ctx context.Context, filter *domain.RoommateFilter) ([]domain.Roommate, int64, error)
	GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error)
	GetUserProfile(ctx context.Context, userId string) (*domain.UserProfile, error)
	GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error)
	GetByUsername(ctx context.Context, username string) (*domain.UserProfile, error)
	GetById(ctx context.Context, userId string) (*domain.User, error)
	Create(ctx context.Context, user *domain.CreateUser) error
	UpdateUserSuggest(ctx context.Context, us *domain.UpdateUserSuggest) error
	UpdatePassword(ctx context.Context, userId string, newPassword string) error
	UpdateUserStatus(ctx context.Context, userId string, status string) error
	VerifyPhone(ctx context.Context, userId string, phone string, otp string, expirationTime time.Time) error
	VerifyPhoneOTP(ctx context.Context, userId string, otp string) error
}
