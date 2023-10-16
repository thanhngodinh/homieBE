package service

import (
	"context"
	"hostel-service/internal/hostel/domain"
	"hostel-service/internal/hostel/port"
	user_domain "hostel-service/internal/user/domain"
	user_port "hostel-service/internal/user/port"
	"time"

	"github.com/google/uuid"
)

type HostelService interface {
	GetHostels(ctx context.Context, hostel *domain.HostelFilter, userId string) ([]domain.Hostel, int64, error)
	SearchHostels(ctx context.Context, hostel *domain.HostelFilter, userId string) ([]domain.Hostel, int64, error)
	GetSuggestHostels(ctx context.Context, userId string) ([]domain.Hostel, int64, error)
	GetHostelById(ctx context.Context, code string, userId string) (*domain.Hostel, error)
	CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error)
	DeleteHostel(ctx context.Context, code string) (int64, error)
}

func NewHostelService(
	repository port.HostelRepository,
	userRepo user_port.UserRepository,
) HostelService {
	return &hostelService{
		repository: repository,
		userRepo:   userRepo,
	}
}

type hostelService struct {
	repository port.HostelRepository
	userRepo   user_port.UserRepository
}

func (s *hostelService) GetHostels(ctx context.Context, hostel *domain.HostelFilter, userId string) ([]domain.Hostel, int64, error) {
	return s.repository.GetHostels(ctx, hostel, userId)
}

func (s *hostelService) SearchHostels(ctx context.Context, hostel *domain.HostelFilter, userId string) ([]domain.Hostel, int64, error) {
	return s.repository.GetHostels(ctx, hostel, userId)
}

func (s *hostelService) GetSuggestHostels(ctx context.Context, userId string) ([]domain.Hostel, int64, error) {
	if userId == "" {
		hostel := &domain.HostelFilter{
			PageSize: 10,
			PageIdx:  0,
			Sort:     "view desc",
		}
		return s.repository.GetHostels(ctx, hostel, userId)
	}
	userSuggest, err := s.userRepo.GetUserSuggest(ctx, userId)
	if err != nil {
		return nil, 0, err
	}
	costFrom := userSuggest.Cost - 500000
	costTo := userSuggest.Cost + 500000
	capacityFrom := userSuggest.Capacity - 1
	capacityTo := userSuggest.Capacity + 1
	hostel := &domain.HostelFilter{
		Province:     &userSuggest.Province,
		District:     &userSuggest.District,
		CostFrom:     &costFrom,
		CostTo:       &costTo,
		CapacityFrom: &capacityFrom,
		CapacityTo:   &capacityTo,
		PageSize:     10,
		PageIdx:      0,
		Sort:         "view desc",
	}
	hostels, total, err := s.repository.GetHostels(ctx, hostel, userId)
	if total < 10 {
		hostel := &domain.HostelFilter{
			Province: &userSuggest.Province,
			PageSize: int(10 - total),
			PageIdx:  0,
			Sort:     "view asc",
		}
		addHostels, addTotal, _ := s.repository.GetHostels(ctx, hostel, userId)
		hostels = append(hostels, addHostels...)
		total += addTotal
	}
	return hostels, total, err
}

func (s *hostelService) GetHostelById(ctx context.Context, code string, userId string) (*domain.Hostel, error) {
	res, err := s.repository.GetHostelById(ctx, code)
	if err != nil {
		return nil, err
	}
	if userId != "" {
		err = s.userRepo.UpdateUserSuggest(ctx, &user_domain.UpdateUserSuggest{
			Id:       userId,
			Province: res.Province,
			District: res.District,
			Cost:     res.Cost,
			Capacity: res.Capacity,
		})
	}
	return res, err
}

func (s *hostelService) CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	hostel.Id = uuid.New().String()
	hostel.CreatedAt = time.Now()
	hostel.EndedAt = time.Now().AddDate(0, 1, 0)
	hostel.Status = domain.HostelActive
	return s.repository.CreateHostel(ctx, hostel)
}

func (s *hostelService) UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	t := time.Now()
	hostel.UpdatedAt = &t
	return s.repository.UpdateHostel(ctx, hostel)
}

func (s *hostelService) DeleteHostel(ctx context.Context, code string) (int64, error) {
	hostel := &domain.Hostel{Id: code}
	return s.repository.DeleteHostel(ctx, hostel)
}
