package service

import (
	"context"
	"hostel-service/internal/utilities/domain"
	"hostel-service/internal/utilities/port"
	"time"
)

type UtilitiesService interface {
	GetAllUtilities(ctx context.Context) ([]domain.GetUtilities, error)
	CreateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error)
	UpdateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error)
	DeleteUtilities(ctx context.Context, code string) (int64, error)
}

func NewUtilitiesService(
	repository port.UtilitiesRepository,
) UtilitiesService {
	return &utilitiesService{
		repository: repository,
	}
}

type utilitiesService struct {
	repository port.UtilitiesRepository
}

func (s *utilitiesService) GetAllUtilities(ctx context.Context) ([]domain.GetUtilities, error) {
	return s.repository.GetAllUtilities(ctx)
}

func (s *utilitiesService) CreateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	utilities.CreatedAt = time.Now()
	return s.repository.CreateUtilities(ctx, utilities)
}

func (s *utilitiesService) UpdateUtilities(ctx context.Context, utilities *domain.Utilities) (int64, error) {
	t := time.Now()
	utilities.UpdatedAt = &t
	return s.repository.UpdateUtilities(ctx, utilities)
}

func (s *utilitiesService) DeleteUtilities(ctx context.Context, code string) (int64, error) {
	utilities := &domain.Utilities{Id: code}
	return s.repository.DeleteUtilities(ctx, utilities)
}
