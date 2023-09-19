package port

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
)

type UserRepository interface {
	GetPostLikedByUser(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error)
	UserLikePost(ctx context.Context, userId string, postId string) (int64, error)
}
