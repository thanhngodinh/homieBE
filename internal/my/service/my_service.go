package service

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	hostel_port "hostel-service/internal/hostel/port"
	"hostel-service/internal/my/domain"
	"hostel-service/internal/my/port"
)

type MyService interface {
	GetMyPostLiked(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error)
	GetMyPosts(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error)
	LikePost(ctx context.Context, userId string, postId string) (int64, error)
	GetMyProfile(ctx context.Context, userId string) (*domain.User, error)
	UpdateMyProfile(ctx context.Context, user *domain.UpdateMyProfileReq) error
	UpdateMyAvatar(ctx context.Context, user *domain.UpdateMyAvatarReq) error
}

func NewMyService(
	myRepo port.MyRepository,
	hostelRepo hostel_port.HostelRepository,
) MyService {
	return &myService{
		myRepo:     myRepo,
		hostelRepo: hostelRepo,
	}
}

type myService struct {
	myRepo     port.MyRepository
	hostelRepo hostel_port.HostelRepository
}

func (s *myService) GetMyPostLiked(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error) {
	hostels, total, err := s.myRepo.GetMyPostLiked(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &hostel_domain.GetHostelsResponse{
		Data:  hostels,
		Total: total,
	}, nil
}

func (s *myService) GetMyPosts(ctx context.Context, userId string) (*hostel_domain.GetHostelsResponse, error) {
	hostels, total, err := s.myRepo.GetMyPosts(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &hostel_domain.GetHostelsResponse{
		Data:  hostels,
		Total: total,
	}, nil
}

func (s *myService) LikePost(ctx context.Context, userId string, postId string) (int64, error) {
	up := domain.LikePost{
		UserId: userId,
		PostId: postId,
	}
	return s.myRepo.LikePost(ctx, up)
}

func (s *myService) GetMyProfile(ctx context.Context, userId string) (*domain.User, error) {
	return s.myRepo.GetMyProfile(ctx, userId)
}

func (s *myService) UpdateMyProfile(ctx context.Context, user *domain.UpdateMyProfileReq) error {
	return s.myRepo.UpdateMyProfile(ctx, user)
}

func (s *myService) UpdateMyAvatar(ctx context.Context, user *domain.UpdateMyAvatarReq) error {
	return s.myRepo.UpdateMyAvatar(ctx, user)
}
