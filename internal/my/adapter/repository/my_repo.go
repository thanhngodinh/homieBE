package repository

import (
	"context"
	hostel_domain "hostel-service/internal/hostel/domain"
	"hostel-service/internal/my/domain"
	"hostel-service/internal/package/util"

	"gorm.io/gorm"
)

func NewMyAdapter(db *gorm.DB) *MyAdapter {
	return &MyAdapter{DB: db}
}

type MyAdapter struct {
	DB *gorm.DB
}

func (r *MyAdapter) GetMyProfile(ctx context.Context, userId string) (*domain.User, error) {
	res := &domain.User{}
	r.DB.Table("users").Where("id = ?", userId).First(res)
	return res, nil
}

func (r *MyAdapter) GetMyPostLiked(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error) {
	hostels := []hostel_domain.Hostel{}
	res := r.DB.Table("hostels").
		Select(`hostels.*, true as "is_liked"`).
		Joins("left join user_like_posts on user_like_posts.post_id = hostels.id").
		Where("user_like_posts.user_id = ?", userId).Scan(&hostels)
	return hostels, res.RowsAffected, res.Error
}

func (r *MyAdapter) GetMyPosts(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error) {
	hostels := []hostel_domain.Hostel{}
	res := r.DB.Table("hostels").
		Where("created_by = ?", userId).Order("created_at desc").
		Find(&hostels)
	return hostels, res.RowsAffected, res.Error
}

func (r *MyAdapter) LikePost(ctx context.Context, up domain.LikePost) (int64, error) {
	res := r.DB.Table("user_like_posts").Create(up)
	if res.Error != nil && res.Error.Error() == util.ErrorDuplicateKey {
		res = r.DB.Table("user_like_posts").Delete(up)
	}
	return res.RowsAffected, res.Error
}

func (r *MyAdapter) UpdateMyProfile(ctx context.Context, user *domain.UpdateMyProfileReq) error {
	res := r.DB.Table("users").Updates(user)
	return res.Error
}

func (r *MyAdapter) UpdateMyAvatar(ctx context.Context, user *domain.UpdateMyAvatarReq) error {
	res := r.DB.Table("users").Updates(user)
	return res.Error
}
