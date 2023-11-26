package port

import (
	"context"
	"hostel-service/internal/admin/domain"
)

type AdminRepository interface {
	UpdatePassword(ctx context.Context, userId string, newPassword string) error
	GetByUsername(ctx context.Context, username string) (*domain.Admin, error)
	GetById(ctx context.Context, id string) (*domain.Admin, error)
}
