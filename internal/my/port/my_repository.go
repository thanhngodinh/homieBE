package port

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	"hostel-service/internal/my/domain"
)

type MyRepository interface {
	GetMyPostLiked(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error)
	GetMyPosts(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error)
	LikePost(ctx context.Context, up domain.LikePost) (int64, error)
}
