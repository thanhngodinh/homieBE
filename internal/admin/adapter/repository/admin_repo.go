package repository

import (
	"context"
	"fmt"
	"hostel-service/internal/admin/domain"
	"hostel-service/internal/admin/port"
	"time"

	"gorm.io/gorm"
)

func NewAdminRepo(db *gorm.DB) port.AdminRepository {
	return &adminRepo{DB: db}
}

type adminRepo struct {
	DB *gorm.DB
}

// Account
func (r *adminRepo) GetByUsername(ctx context.Context, username string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	res := r.DB.Where("username = ?", username).First(admin)
	return admin, res.Error
}

func (r *adminRepo) GetById(ctx context.Context, adminId string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	res := r.DB.Table("admins").Where("id = ?", adminId).First(admin)
	return admin, res.Error
}

// Post
func (r *adminRepo) GetPosts(ctx context.Context, filter *domain.PostFilter) ([]domain.Post, int64, error) {
	var (
		tx     *gorm.DB
		hotels []domain.Post
	)

	tx = r.DB.Table("posts").Select(`
	(select count(*) from user_like_posts where user_like_posts.post_id = posts.id) as like,
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
	if filter.Status != "" {
		tx = tx.Where("status = ?", filter.Status)
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
	res2 := tx.Order(filter.Sort).Limit(filter.PageSize).Offset(filter.PageIdx * filter.PageSize).Scan(&hotels)
	return hotels, total, res2.Error
}

func (r *adminRepo) GetPostById(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post

	r.DB.Table("posts").
		Select(`posts.*,
		users.id as author_id, users.display_name as author_name, users.avatar_url as author_avatar,
		array_remove(array_agg(post_utilities.utility_id), NULL) as utilities`).
		Joins("left join users on users.id = posts.created_by").
		Joins("left join post_utilities on post_utilities.post_id = posts.id").
		Group("posts.id").Group("users.id").
		Where("posts.id = ?", id).First(&post)
	return &post, nil
}

func (r *adminRepo) UpdatePostStatus(ctx context.Context, id string, status string) (int64, error) {
	res := r.DB.Table("posts").Where("id = ?", id).Updates(map[string]interface{}{"status": status})
	return res.RowsAffected, res.Error
}

// User
func (r *adminRepo) SearchUsers(ctx context.Context, filter *domain.UserFilter) ([]domain.User, int64, error) {
	var (
		tx        *gorm.DB
		users []domain.User
	)

	tx = r.DB.Table("users").Select("users.*")
	if len(filter.Gender) > 0 {
		tx = tx.Where("gender = ?", filter.Gender)
	}
	if len(filter.Name) > 0 {
		tx = tx.Where("display_name ilike ? ", fmt.Sprintf("%%%v%%", filter.Name))
	}
	if len(filter.Status) > 0 {
		tx = tx.Where("status = ?", filter.Status)
	}

	res1 := tx.Scan(&users)
	total := res1.RowsAffected
	res2 := tx.Order(filter.Sort).Limit(filter.PageSize).Offset(filter.PageIdx * filter.PageSize).Scan(&users)
	return users, total, res2.Error
}

func (r *adminRepo) GetUserById(ctx context.Context, userId string) (*domain.User, error) {
	user := &domain.User{}
	res := r.DB.Where("id = ?", userId).First(user)
	return user, res.Error
}

func (r *adminRepo) UpdateUserStatus(ctx context.Context, userId string, status string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"status": status})
	return res.Error
}

func (r *adminRepo) UpdatePassword(ctx context.Context, userId string, newPassword string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"password": newPassword, "is_verified_email": true})
	return res.Error
}
