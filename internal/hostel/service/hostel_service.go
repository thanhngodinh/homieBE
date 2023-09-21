package service

import (
	"context"
	"hostel-service/internal/hostel/domain"
	"hostel-service/internal/hostel/port"
	"time"

	"github.com/google/uuid"
)

type HostelService interface {
	GetHostels(ctx context.Context, hostel *domain.HostelFilter) (*domain.GetHostelsResponse, error)
	GetHostelById(ctx context.Context, code string) (*domain.Hostel, error)
	CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	DeleteHostel(ctx context.Context, code string) (int64, error)
}

func NewHostelService(repository port.HostelRepository) HostelService {
	return &hostelService{repository: repository}
}

type hostelService struct {
	repository port.HostelRepository
}

func (s *hostelService) GetHostels(ctx context.Context, hostel *domain.HostelFilter) (*domain.GetHostelsResponse, error) {
	hostels, total, err := s.repository.GetHostels(ctx, hostel)
	if err != nil {
		return nil, err
	}
	return &domain.GetHostelsResponse{
		Data:  hostels,
		Total: total,
	}, nil
}

func (s *hostelService) GetHostelById(ctx context.Context, code string) (*domain.Hostel, error) {
	return s.repository.GetHostelById(ctx, code)
}

func (s *hostelService) CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	hostel.CreatedAt = time.Now()
	hostel.Id = uuid.New().String()
	// tx, err := s.db.Begin()
	// if err != nil {
	// 	return -1, err
	// }
	// ctx = context.WithValue(ctx, "tx", tx)
	// res, err := s.repository.CreateHostel(ctx, hostel)
	// if err != nil {
	// 	tx.Rollback()
	// 	return -1, err
	// }
	// err = tx.Commit()
	return s.repository.CreateHostel(ctx, hostel)
}

func (s *hostelService) UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	t := time.Now()
	hostel.UpdatedAt = &t
	return s.repository.UpdateHostel(ctx, hostel)
}

func (s *hostelService) DeleteHostel(ctx context.Context, code string) (int64, error) {
	hostel := &domain.Hostel{Id : code}
	return s.repository.DeleteHostel(ctx, hostel)
}
