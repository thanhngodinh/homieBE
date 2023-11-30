package port

import (
	"context"
	"hostel-service/internal/admin/domain"
)

type AdminRepository interface {
	GetByUsername(ctx context.Context, username string) (*domain.Admin, error)
	GetById(ctx context.Context, id string) (*domain.Admin, error)

	GetPosts(ctx context.Context, filter *domain.PostFilter) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, id string) (*domain.Post, error)
	UpdatePostStatus(ctx context.Context, id string, status string) (int64, error)

	SearchUsers(ctx context.Context, filter *domain.UserFilter) ([]domain.User, int64, error)
	GetUserById(ctx context.Context, id string) (*domain.User, error)
	UpdateUserStatus(ctx context.Context, userId string, status string) error
	UpdatePassword(ctx context.Context, userId string, newPassword string) error
}
