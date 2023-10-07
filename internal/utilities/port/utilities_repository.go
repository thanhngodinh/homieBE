package port

import (
	"context"
	"hostel-service/internal/utilities/domain"
)

type UtilitiesRepository interface {
	GetAllUtilities(ctx context.Context) ([]domain.GetUtilities, error)
	CreateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error)
	UpdateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error)
	DeleteUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error)
}
