package port

import (
	"context"
	"hostel-service/internal/authentication/domain"
)

type AuthenticationRepository interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (int64, error)
}
