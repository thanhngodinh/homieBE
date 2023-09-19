package repository

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"

	"gorm.io/gorm"
)

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{DB: db}
}

type UserAdapter struct {
	DB *gorm.DB
}

func (r *UserAdapter) GetPostLikedByUser(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error) {
	return nil, 0, nil
}

func (r *UserAdapter) UserLikePost(ctx context.Context, userId string, postId string) (int64, error) {
	return 1, nil
}
