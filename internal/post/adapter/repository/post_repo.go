package repository

import (
	"context"
	"fmt"
	"hostel-service/internal/post/domain"
	"time"

	"gorm.io/gorm"
)

func NewPostAdapter(db *gorm.DB) *PostAdapter {
	return &PostAdapter{DB: db}
}

type PostAdapter struct {
	DB *gorm.DB
}

func (r *PostAdapter) GetPosts(ctx context.Context, filter *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	var (
		tx     *gorm.DB
		hotels []domain.Post
	)

	tx = r.DB.Table("posts").
		Select(fmt.Sprintf(`
	(select true from user_like_posts where user_like_posts.post_id = posts.id and user_like_posts.user_id = '%v') as "is_liked",
	array_remove(array_agg(post_utilities.utility_id), NULL) as utilities,
	posts.*`, userId)).
		Joins("left join post_utilities on post_utilities.post_id = posts.id").Group("posts.id")
	if filter.Name != nil && len(*filter.Name) > 0 {
		tx = tx.Where("name ilike ?", fmt.Sprintf("%%%v%%", *filter.Name))
	}
	if filter.Province != nil && len(*filter.Province) > 0 {
		tx = tx.Where("province ilike ?", fmt.Sprintf("%%%v%%", *filter.Province))
	}
	if filter.District != nil && len(*filter.District) > 0 {
		tx = tx.Where("district ilike ?", fmt.Sprintf("%%%v%%", *filter.District))
	}
	if filter.Ward != nil && len(*filter.Ward) > 0 {
		tx = tx.Where("ward = ilike ?", fmt.Sprintf("%%%v%%", *filter.Ward))
	}
	if filter.Street != nil && len(*filter.Street) > 0 {
		tx = tx.Where("street ilike ?", fmt.Sprintf("%%%v%%", *filter.Street))
	}
	if filter.Status != nil && len(*filter.Status) > 0 {
		tx = tx.Where("status = ?", filter.Status)
	}
	if filter.CostFrom != nil {
		tx = tx.Where("cost >= ?", filter.CostFrom)
	}
	if filter.CostTo != nil {
		tx = tx.Where("cost <= ?", filter.CostTo)
	}
	if filter.DepositFrom != nil {
		tx = tx.Where("deposit >= ?", filter.DepositFrom)
	}
	if filter.DepositTo != nil {
		tx = tx.Where("deposit <= ?", filter.DepositTo)
	}
	if filter.Capacity != nil {
		tx = tx.Where("capacity = ?", filter.Capacity)
	}
	if filter.CapacityFrom != nil {
		tx = tx.Where("capacity >= ?", filter.CapacityFrom)
	}
	if filter.CapacityTo != nil {
		tx = tx.Where("capacity <= ?", filter.CapacityTo)
	}
	if filter.CreatedAt != nil {
		tx = tx.Where("created_at = ?", filter.CreatedAt)
	}
	if filter.CreatedBy != nil && len(*filter.CreatedBy) > 0 {
		tx = tx.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.IsIncludeEnded == false {
		tx = tx.Where("ended_at > ?", time.Now())
	}
	if len(filter.Utilities) > 0 {
		tx = tx.Where("? <@ (select array_agg(post_utilities.utility_id) from post_utilities where post_utilities.post_id = posts.id)", filter.Utilities)
	}
	res1 := tx.Scan(&hotels)
	total := res1.RowsAffected
	res2 := tx.Order(filter.Sort).Limit(filter.PageSize).Offset(filter.PageIdx * filter.PageSize).Scan(&hotels)
	return hotels, total, res2.Error
}

func (r *PostAdapter) GetPostById(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post

	r.DB.Table("posts").
		Select(`posts.*,
		users.display_name as author, users.avatar_url as author_avatar,
		array_remove(array_agg(post_utilities.utility_id), NULL) as utilities`).
		Joins("join users on users.id = posts.created_by").
		Joins("join post_utilities on post_utilities.post_id = posts.id").
		Group("posts.id").Group("users.id").
		Where("posts.id = ?", id).First(&post)
	r.DB.Table("posts").Where("id = ?", id).Updates(map[string]interface{}{"view": post.View + 1})
	return &post, nil
}

func (r *PostAdapter) CreatePost(ctx context.Context, post *domain.Post) (int64, error) {
	res := r.DB.Table("posts").Create(post)
	if post.Utilities != nil {
		hu := []domain.PostUtilities{}
		for _, u := range post.Utilities {
			hu = append(hu, domain.PostUtilities{PostId: post.Id, UtilitiesId: u})
		}
		r.DB.Table("post_utilities").Create(hu)
	}
	rate := &domain.RateInfo{
		PostId: post.Id,
	}
	res = r.DB.Table("post_rate_info").Create(rate)
	return res.RowsAffected, res.Error
}

func (r *PostAdapter) UpdatePost(ctx context.Context, post *domain.Post) (int64, error) {
	res := r.DB.Table("posts").Model(&post).Updates(post)
	r.DB.Table("post_utilities").Where("post_id = ?", post.Id).Delete(domain.PostUtilities{})
	if post.Utilities != nil {
		hu := []domain.PostUtilities{}
		for _, u := range post.Utilities {
			hu = append(hu, domain.PostUtilities{PostId: post.Id, UtilitiesId: u})
		}
	}
	return res.RowsAffected, res.Error
}

func (r *PostAdapter) DeletePost(ctx context.Context, post *domain.Post) (int64, error) {
	res := r.DB.Table("posts").Delete(post)
	return res.RowsAffected, res.Error
}
