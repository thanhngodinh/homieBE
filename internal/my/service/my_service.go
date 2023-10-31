package service

import (
	"context"
	"hostel-service/internal/my/domain"
	"hostel-service/internal/my/port"
	post_domain "hostel-service/internal/post/domain"
	post_port "hostel-service/internal/post/port"
)

type MyService interface {
	GetMyPostLiked(ctx context.Context, userId string) ([]post_domain.Post, int64, error)
	GetMyPosts(ctx context.Context, userId string) ([]post_domain.Post, int64, error)
	LikePost(ctx context.Context, userId string, postId string) (int64, error)
	GetMyProfile(ctx context.Context, userId string) (*domain.User, error)
	UpdateMyProfile(ctx context.Context, user *domain.UpdateMyProfileReq) error
	UpdateMyAvatar(ctx context.Context, user *domain.UpdateMyAvatarReq) error
}

func NewMyService(
	myRepo port.MyRepository,
	hostelRepo post_port.PostRepository,
) MyService {
	return &myService{
		myRepo:     myRepo,
		hostelRepo: hostelRepo,
	}
}

type myService struct {
	myRepo     port.MyRepository
	hostelRepo post_port.PostRepository
}

func (s *myService) GetMyPostLiked(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	return s.myRepo.GetMyPostLiked(ctx, userId)
}

func (s *myService) GetMyPosts(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	return s.myRepo.GetMyPosts(ctx, userId)
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
