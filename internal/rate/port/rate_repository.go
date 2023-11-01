package port

import (
	"context"
	"hostel-service/internal/rate/domain"
)

type RateRepository interface {
	GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error)
	GetSimplePostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error)
	UpdatePostRateInfo(ctx context.Context, rateInfo *domain.PostRateInfo) error
	GetRate(ctx context.Context, postId string, userId string) (*domain.Rate, error)
	CreateRate(ctx context.Context, rate *domain.Rate) (int64, error)
	UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error)
}
