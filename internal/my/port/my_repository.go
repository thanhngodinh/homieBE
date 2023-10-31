package port

import (
	"context"
	"hostel-service/internal/my/domain"
	post_domain "hostel-service/internal/post/domain"
)

type MyRepository interface {
	GetMyPostLiked(ctx context.Context, userId string) ([]post_domain.Post, int64, error)
	GetMyPosts(ctx context.Context, userId string) ([]post_domain.Post, int64, error)
	LikePost(ctx context.Context, up domain.LikePost) (int64, error)
	GetMyProfile(ctx context.Context, userId string) (*domain.User, error)
	UpdateMyProfile(ctx context.Context, user *domain.UpdateMyProfileReq) error
	UpdateMyAvatar(ctx context.Context, user *domain.UpdateMyAvatarReq) error
}
