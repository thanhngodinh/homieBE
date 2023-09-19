package service

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	"hostel-service/internal/user/port"

	"gorm.io/gorm"
)

type UserService interface {
	GetPostLikedByUser(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error)
	UserLikePost(ctx context.Context, userId string, postId string) (int64, error)
}

func NewUserService(db *gorm.DB, repository port.UserRepository) UserService {
	return &userService{db: db, repository: repository}
}

type userService struct {
	db         *gorm.DB
	repository port.UserRepository
}

func (s *userService) GetPostLikedByUser(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error) {
	hostels, total, err := s.repository.GetPostLikedByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &hostel_domain.GetHostelsResponse{
		Data:  hostels,
		Total: total,
	}, nil
}

func (s *userService) UserLikePost(ctx context.Context, userId string, postId string) (int64, error) {
	return s.repository.UserLikePost(ctx, userId, postId)
}
