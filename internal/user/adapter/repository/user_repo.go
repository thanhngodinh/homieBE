package repository

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	user_domain "hostel-service/internal/user/domain"
	"hostel-service/internal/util"

	"gorm.io/gorm"
)

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{DB: db}
}

type UserAdapter struct {
	DB *gorm.DB
}

func (r *UserAdapter) GetPostLikedByUser(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error) {
	hostels := []hostel_domain.Hostel{}
	res := r.DB.Table("hostels").
		Select("hostels.*").
		Joins("left join user_like_posts on user_like_posts.post_id = hostels.id").
		Where("user_like_posts.user_id = ?", userId).Scan(&hostels)
	return hostels, res.RowsAffected, res.Error
}

func (r *UserAdapter) UserLikePost(ctx context.Context, up user_domain.UserLikePosts) (int64, error) {
	res := r.DB.Table("user_like_posts").Create(up)
	if res.Error != nil && res.Error.Error() == util.ErrorDuplicateKey {
		res = r.DB.Table("user_like_posts").Delete(up)
	}
	return res.RowsAffected, res.Error
}
