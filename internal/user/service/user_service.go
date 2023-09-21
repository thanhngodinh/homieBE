package service

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	hostel_port "hostel-service/internal/hostel/port"
	user_domain "hostel-service/internal/user/domain"
	"hostel-service/internal/user/port"
)

type UserService interface {
	GetPostLikedByUser(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error)
	UserLikePost(ctx context.Context, userId string, postId string) (int64, error)
}

func NewUserService(
	userRepo port.UserRepository,
	hostelRepo hostel_port.HostelRepository,
) UserService {
	return &userService{
		userRepo:   userRepo,
		hostelRepo: hostelRepo,
	}
}

type userService struct {
	userRepo   port.UserRepository
	hostelRepo hostel_port.HostelRepository
}

func (s *userService) GetPostLikedByUser(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error) {
	hostels, total, err := s.userRepo.GetPostLikedByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &hostel_domain.GetHostelsResponse{
		Data:  hostels,
		Total: total,
	}, nil
}

func (s *userService) UserLikePost(ctx context.Context, userId string, postId string) (int64, error) {
	up := user_domain.UserLikePosts{
		UserId: userId,
		PostId: postId,
	}
	return s.userRepo.UserLikePost(ctx, up)
}
