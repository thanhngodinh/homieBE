package service

import (
	"context"
	"hostel-service/internal/post/domain"
	"hostel-service/internal/post/port"
	user_domain "hostel-service/internal/user/domain"
	user_port "hostel-service/internal/user/port"
	"time"

	"github.com/google/uuid"
)

type PostService interface {
	GetPosts(ctx context.Context, hostel *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	SearchPosts(ctx context.Context, hostel *domain.PostFilter, userId string) ([]domain.Post, int64, error)
	GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error)
	GetPostById(ctx context.Context, code string, userId string) (*domain.Post, error)
	CreatePost(ctx context.Context, hostel *domain.Post) (int64, error)
	UpdatePost(ctx context.Context, hostel *domain.Post) (int64, error)
	DeletePost(ctx context.Context, code string) (int64, error)
}

func NewPostService(
	repository port.PostRepository,
	userRepo user_port.UserRepository,
) PostService {
	return &hostelService{
		repository: repository,
		userRepo:   userRepo,
	}
}

type hostelService struct {
	repository port.PostRepository
	userRepo   user_port.UserRepository
}

func (s *hostelService) GetPosts(ctx context.Context, hostel *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	return s.repository.GetPosts(ctx, hostel, userId)
}

func (s *hostelService) SearchPosts(ctx context.Context, hostel *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	return s.repository.GetPosts(ctx, hostel, userId)
}

func (s *hostelService) GetSuggestPosts(ctx context.Context, userId string) ([]domain.Post, int64, error) {
	if userId == "" {
		hostel := &domain.PostFilter{
			PageSize: 10,
			PageIdx:  0,
			Sort:     "view desc",
		}
		return s.repository.GetPosts(ctx, hostel, userId)
	}
	userSuggest, err := s.userRepo.GetUserSuggest(ctx, userId)
	if err != nil {
		return nil, 0, err
	}
	costFrom := userSuggest.Cost - 500000
	costTo := userSuggest.Cost + 500000
	capacityFrom := userSuggest.Capacity - 1
	capacityTo := userSuggest.Capacity + 1
	hostel := &domain.PostFilter{
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
	posts, total, err := s.repository.GetPosts(ctx, hostel, userId)
	if total < 10 {
		hostel := &domain.PostFilter{
			Province: &userSuggest.Province,
			PageSize: int(10 - total),
			PageIdx:  0,
			Sort:     "view asc",
		}
		addPosts, addTotal, _ := s.repository.GetPosts(ctx, hostel, userId)
		posts = append(posts, addPosts...)
		total += addTotal
	}
	return posts, total, err
}

func (s *hostelService) GetPostById(ctx context.Context, code string, userId string) (*domain.Post, error) {
	res, err := s.repository.GetPostById(ctx, code)
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

func (s *hostelService) CreatePost(ctx context.Context, hostel *domain.Post) (int64, error) {
	hostel.Id = "post-" + uuid.NewString()
	hostel.CreatedAt = time.Now()
	hostel.EndedAt = time.Now().AddDate(0, 1, 0)
	hostel.Status = domain.PostActive
	return s.repository.CreatePost(ctx, hostel)
}

func (s *hostelService) UpdatePost(ctx context.Context, hostel *domain.Post) (int64, error) {
	t := time.Now()
	hostel.UpdatedAt = &t
	return s.repository.UpdatePost(ctx, hostel)
}

func (s *hostelService) DeletePost(ctx context.Context, code string) (int64, error) {
	hostel := &domain.Post{Id: code}
	return s.repository.DeletePost(ctx, hostel)
}
