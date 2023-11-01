package service

import (
	"context"
	"hostel-service/internal/rate/domain"
	"hostel-service/internal/rate/port"
	"time"
)

type RateService interface {
	GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error)
	CreateRate(ctx context.Context, rate *domain.Rate) (int64, error)
	UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error)
}

func NewRateService(
	repository port.RateRepository,
) RateService {
	return &rateService{
		repository: repository,
	}
}

type rateService struct {
	repository port.RateRepository
}

func (s *rateService) GetPostRate(ctx context.Context, postId string) (*domain.PostRateInfo, error) {
	return s.repository.GetPostRate(ctx, postId)
}

func (s *rateService) CreateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	rate.CreatedAt = time.Now()
	res, err := s.repository.CreateRate(ctx, rate)
	if err != nil {
		return 0, err
	}

	rateInfo, err := s.repository.GetSimplePostRate(ctx, rate.PostId)
	if err != nil {
		return 0, err
	}
	*rateInfo.Total++
	switch rate.Star {
	case 1:
		*rateInfo.Star1++
	case 2:
		*rateInfo.Star2++
	case 3:
		*rateInfo.Star3++
	case 4:
		*rateInfo.Star4++
	case 5:
		*rateInfo.Star5++
	}
	err = s.repository.UpdatePostRateInfo(ctx, rateInfo)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *rateService) UpdateRate(ctx context.Context, rate *domain.Rate) (int64, error) {
	oldRate, err := s.repository.GetRate(ctx, rate.PostId, rate.UserId)
	if err != nil {
		return 0, err
	}

	t := time.Now()
	rate.CreatedAt = oldRate.CreatedAt
	rate.UpdatedAt = &t
	res, err := s.repository.UpdateRate(ctx, rate)
	if err != nil {
		return 0, err
	}

	rateInfo, err := s.repository.GetSimplePostRate(ctx, rate.PostId)
	if err != nil {
		return 0, err
	}

	switch oldRate.Star {
	case 1:
		*rateInfo.Star1--
	case 2:
		*rateInfo.Star2--
	case 3:
		*rateInfo.Star3--
	case 4:
		*rateInfo.Star4--
	case 5:
		*rateInfo.Star5--
	}

	switch rate.Star {
	case 1:
		*rateInfo.Star1++
	case 2:
		*rateInfo.Star2++
	case 3:
		*rateInfo.Star3++
	case 4:
		*rateInfo.Star4++
	case 5:
		*rateInfo.Star5++
	}

	err = s.repository.UpdatePostRateInfo(ctx, rateInfo)
	if err != nil {
		return 0, err
	}

	return res, nil
}
