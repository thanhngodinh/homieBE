package service

import (
	"context"
	"database/sql"
	"hostel-service/internal/hostel/domain"
	"hostel-service/internal/hostel/port"
	"hostel-service/internal/util"
	"time"

	"github.com/google/uuid"
)

type HostelService interface {
	GetHostels(ctx context.Context, pageSize int, pageIdx int) (*domain.GetHostelsResponse, error)
	GetHostelById(ctx context.Context, code string) (*domain.Hostel, error)
	CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	DeleteHostel(ctx context.Context, code string) (int64, error)
}

func NewHostelService(db *sql.DB, repository port.HostelRepository) HostelService {
	return &hostelService{db: db, repository: repository}
}

type hostelService struct {
	db         *sql.DB
	repository port.HostelRepository
}

func (s *hostelService) GetHostels(ctx context.Context, pageSize int, pageIdx int) (*domain.GetHostelsResponse, error) {
	hostels, total, err := s.repository.GetHostels(ctx, pageSize, pageIdx)
	if err != nil {
		return nil, err
	}
	return &domain.GetHostelsResponse{
		Data: hostels,
		Pagin: util.Pagination{
			Total:    total,
			PageIdx:  pageIdx,
			PageSize: pageSize,
		},
	}, nil
}

func (s *hostelService) GetHostelById(ctx context.Context, code string) (*domain.Hostel, error) {
	return s.repository.GetHostelById(ctx, code)
}

func (s *hostelService) CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	hostel.CreatedAt = time.Now()
	hostel.Id = uuid.New().String()
	tx, err := s.db.Begin()
	if err != nil {
		return -1, err
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.CreateHostel(ctx, hostel)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

func (s *hostelService) UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	t := time.Now()
	hostel.UpdatedAt = &t
	tx, err := s.db.Begin()
	if err != nil {
		return -1, err
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.UpdateHostel(ctx, hostel)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

func (s *hostelService) DeleteHostel(ctx context.Context, code string) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, err
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.DeleteHostel(ctx, code)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
