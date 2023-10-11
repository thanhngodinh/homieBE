package port

import (
	"context"
	"hostel-service/internal/hostel/domain"
)

type HostelRepository interface {
	GetHostels(ctx context.Context, hostel *domain.HostelFilter, userId string) ([]domain.Hostel, int64, error)
	GetHostelById(ctx context.Context, id string) (*domain.Hostel, error)
	CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	DeleteHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
}
