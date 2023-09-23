package port

import (
	"context"
	"hostel-service/internal/user/domain"
)

type UserRepository interface {
	UpdateUserSuggest(ctx context.Context, us *domain.UpdateUserSuggest) error
	GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error)
}
