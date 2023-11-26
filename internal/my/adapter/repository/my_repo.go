package repository

import (
	"context"
	"hostel-service/internal/my/domain"
	"hostel-service/internal/package/util"
	post_domain "hostel-service/internal/post/domain"

	"gorm.io/gorm"
)

func NewMyAdapter(db *gorm.DB) *MyAdapter {
	return &MyAdapter{DB: db}
}

type MyAdapter struct {
	DB *gorm.DB
}

func (r *MyAdapter) GetMyProfile(ctx context.Context, userId string) (*domain.UserProfile, error) {
	res := &domain.UserProfile{}
	r.DB.Table("users").Where("id = ?", userId).First(res)
	return res, nil
}

func (r *MyAdapter) GetMyPostLiked(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	posts := []post_domain.Post{}
	res := r.DB.Table("posts").
		Select(`posts.*, true as "is_liked"`).
		Joins("left join user_like_posts on user_like_posts.post_id = posts.id").
		Where("user_like_posts.user_id = ?", userId).Scan(&posts)
	return posts, res.RowsAffected, res.Error
}

func (r *MyAdapter) GetMyPosts(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	posts := []post_domain.Post{}
	res := r.DB.Table("posts").
		Where("created_by = ?", userId).Order("created_at desc").
		Find(&posts)
	return posts, res.RowsAffected, res.Error
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
