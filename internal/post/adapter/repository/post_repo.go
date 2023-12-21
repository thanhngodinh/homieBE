package repository

import (
	"context"
	"fmt"
	"hostel-service/internal/post/domain"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{DB: db}
}

type PostRepo struct {
	DB *gorm.DB
}

func (r *PostRepo) GetPosts(ctx context.Context, filter *domain.PostFilter, userId string) ([]domain.Post, int64, error) {
	var (
		tx     *gorm.DB
		hotels []domain.Post
	)

	tx = r.DB.Table("posts").
		Select(`
	(select avg from post_rate_info where post_rate_info.post_id = posts.id) as "avg_rate",
	array_remove(array_agg(post_utilities.utility_id), NULL) as utilities,
	posts.*`).
		Joins("left join post_utilities on post_utilities.post_id = posts.id").Group("posts.id")
	if filter.Name != "" {
		tx = tx.Where("name ilike ?", fmt.Sprintf("%%%v%%", filter.Name))
	}
	if filter.Province != "" {
		tx = tx.Where("province ilike ?", fmt.Sprintf("%%%v%%", filter.Province))
	}
	if filter.District != "" {
		tx = tx.Where("district ilike ?", fmt.Sprintf("%%%v%%", filter.District))
	}
	if filter.Ward != "" {
		tx = tx.Where("ward = ilike ?", fmt.Sprintf("%%%v%%", filter.Ward))
	}
	if filter.Street != "" {
		tx = tx.Where("street ilike ?", fmt.Sprintf("%%%v%%", filter.Street))
	}
	if filter.CostFrom != 0 {
		tx = tx.Where("cost >= ?", filter.CostFrom)
	}
	if filter.CostTo != 0 {
		tx = tx.Where("cost <= ?", filter.CostTo)
	}
	if filter.DepositFrom != 0 {
		tx = tx.Where("deposit >= ?", filter.DepositFrom)
	}
	if filter.DepositTo != 0 {
		tx = tx.Where("deposit <= ?", filter.DepositTo)
	}
	if filter.Capacity != 0 {
		tx = tx.Where("capacity = ?", filter.Capacity)
	}
	if filter.CapacityFrom != 0 {
		tx = tx.Where("capacity >= ?", filter.CapacityFrom)
	}
	if filter.CapacityTo != 0 {
		tx = tx.Where("capacity <= ?", filter.CapacityTo)
	}
	if filter.CreatedAt != nil {
		tx = tx.Where("created_at = ?", filter.CreatedAt)
	}
	if filter.CreatedBy != "" {
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
	res2 := tx.Order(filter.Sort).Limit(filter.PageSize).Offset((filter.PageIdx - 1) * filter.PageSize).Scan(&hotels)
	return hotels, total, res2.Error
}

func (r *PostRepo) GetPostById(ctx context.Context, id string, userId string) (*domain.Post, error) {
	var post domain.Post

	r.DB.Table("posts").
		Select(fmt.Sprintf(`posts.*,
		(select true from user_like_posts where user_like_posts.post_id = posts.id and user_like_posts.user_id = '%v') as "is_liked",
		users.id as author_id, users.display_name as author_name, users.avatar_url as author_avatar, users.phone as phone,
		array_remove(array_agg(post_utilities.utility_id), NULL) as utilities`, userId)).
		Joins("left join users on users.id = posts.created_by").
		Joins("left join post_utilities on post_utilities.post_id = posts.id").
		Group("posts.id").Group("users.id").
		Where("posts.id = ?", id).First(&post)
	r.DB.Table("posts").Where("id = ?", id).Updates(map[string]interface{}{"view": post.View + 1})
	return &post, nil
}

func (r *PostRepo) GetPostByIds(ctx context.Context, ids []string) ([]domain.Post, error) {
	var post []domain.Post

	r.DB.Table("posts").
		Select(`posts.*,
		array_remove(array_agg(post_utilities.utility_id), NULL) as utilities`).
		Joins("left join post_utilities on post_utilities.post_id = posts.id").
		Group("posts.id").
		Where("posts.id = any(?)", pq.Array(ids)).Find(&post)
	return post, nil
}

func (r *PostRepo) CreatePost(ctx context.Context, post *domain.Post) (int64, error) {
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

func (r *PostRepo) UpdatePost(ctx context.Context, post *domain.UpdatePostReq) (int64, error) {
	res := r.DB.Table("posts").Model(&post).Updates(post)
	r.DB.Table("post_utilities").Where("post_id = ?", post.Id).Delete(domain.PostUtilities{})
	if post.Utilities != nil {
		hu := []domain.PostUtilities{}
		for _, u := range post.Utilities {
			hu = append(hu, domain.PostUtilities{PostId: post.Id, UtilitiesId: u})
		}
		r.DB.Table("post_utilities").Create(hu)
	}
	return res.RowsAffected, res.Error
}

func (r *PostRepo) DeletePost(ctx context.Context, post *domain.Post) (int64, error) {
	res := r.DB.Table("posts").Delete(post)
	return res.RowsAffected, res.Error
}
