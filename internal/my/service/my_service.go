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
