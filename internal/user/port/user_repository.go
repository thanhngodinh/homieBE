package port

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	user_domain "hostel-service/internal/user/domain"
)

type UserRepository interface {
	GetPostLikedByUser(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error)
	UserLikePost(ctx context.Context, up user_domain.UserLikePosts) (int64, error)
}
