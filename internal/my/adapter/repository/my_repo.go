package repository

import (
	"context"
	"fmt"
	"hostel-service/internal/my/domain"
	post_domain "hostel-service/internal/post/domain"
	"hostel-service/pkg/util"

	"gorm.io/gorm"
)

func NewMyRepo(db *gorm.DB) *MyRepo {
	return &MyRepo{DB: db}
}

type MyRepo struct {
	DB *gorm.DB
}

func (r *MyRepo) GetMyProfile(ctx context.Context, userId string) (*domain.UserProfile, error) {
	res := &domain.UserProfile{}
	r.DB.Table("users").
		Select(fmt.Sprintf(`users.*,
		(select count(*) from user_like_posts where user_like_posts.user_id = '%s') as "like",
		(select count(*) from posts where posts.created_by = '%s') as "post"`,
			userId, userId)).
		Where("users.id = ?", userId).First(res)
	return res, nil
}

func (r *MyRepo) GetMyPostLiked(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	posts := []post_domain.Post{}
	res := r.DB.Table("posts").
		Select(`posts.*, (select avg from post_rate_info where post_rate_info.post_id = posts.id) as "avg_rate"`).
		Joins("left join user_like_posts on user_like_posts.post_id = posts.id").
		Where("user_like_posts.user_id = ?", userId).Scan(&posts)
	return posts, res.RowsAffected, res.Error
}

func (r *MyRepo) GetMyPosts(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	posts := []post_domain.Post{}
	res := r.DB.Table("posts").
		Where("created_by = ?", userId).Order("created_at desc").
		Find(&posts)
	return posts, res.RowsAffected, res.Error
}

func (r *MyRepo) LikePost(ctx context.Context, up domain.LikePost) (int64, error) {
	res := r.DB.Table("user_like_posts").Create(up)
	if res.Error != nil && res.Error.Error() == util.ErrorDuplicateKey {
		res = r.DB.Table("user_like_posts").Delete(up)
	}
	return res.RowsAffected, res.Error
}

func (r *MyRepo) UpdateMyProfile(ctx context.Context, user *domain.UpdateMyProfileReq) error {
	res := r.DB.Table("users").Updates(user)
	return res.Error
}

func (r *MyRepo) UpdateMyAvatar(ctx context.Context, user *domain.UpdateMyAvatarReq) error {
	res := r.DB.Table("users").Updates(user)
	return res.Error
}
