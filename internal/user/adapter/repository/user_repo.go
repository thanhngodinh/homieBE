package repository

import (
	"context"
	"fmt"
	post_domain "hostel-service/internal/post/domain"
	"hostel-service/internal/user/domain"
	"time"

	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) SearchRoommates(ctx context.Context, filter *domain.RoommateFilter) ([]domain.Roommate, int64, error) {
	var (
		tx        *gorm.DB
		roommates []domain.Roommate
	)

	tx = r.DB.Table("users").Select("users.*").Where("is_find_roommate = ?", true)
	if len(filter.Gender) > 0 {
		tx = tx.Where("gender = ?", filter.Gender)
	}
	if len(filter.Name) > 0 {
		tx = tx.Where("display_name ilike ? ", fmt.Sprintf("%%%v%%", filter.Name))
	}
	if len(filter.Province) > 0 {
		tx = tx.Where("province_profile ilike ?", fmt.Sprintf("%%%v%%", filter.Province))
	}
	if len(filter.District) > 0 {
		tx = tx.Where("district_profile <@ ?", filter.District)
	}
	if filter.CostFrom > 0 {
		tx = tx.Where("cost_from <= ?", filter.CostFrom)
	}
	if filter.CostTo > 0 {
		tx = tx.Where("cost_to >= ?", filter.CostTo)
	}

	res1 := tx.Scan(&roommates)
	total := res1.RowsAffected
	res2 := tx.Order(filter.Sort).Limit(filter.PageSize).Offset(filter.PageIdx * filter.PageSize).Scan(&roommates)
	return roommates, total, res2.Error
}

func (r *UserRepo) GetUserProfile(ctx context.Context, userId string) (*domain.UserProfile, error) {
	res := &domain.UserProfile{}
	r.DB.Table("users").Where("id = ?", userId).First(res)
	return res, nil
}

func (r *UserRepo) GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error) {
	roommate := &domain.Roommate{}
	res := r.DB.Table("users").Where("id = ?", userId).First(roommate)
	return roommate, res.Error
}

func (r *UserRepo) GetUserPosts(ctx context.Context, userId string) ([]post_domain.Post, int64, error) {
	posts := []post_domain.Post{}
	res := r.DB.Table("posts").
		Where("created_by = ?", userId).Order("created_at desc").
		Find(&posts)
	return posts, res.RowsAffected, res.Error
}

func (r *UserRepo) UpdateUserSuggest(ctx context.Context, us *domain.UpdateUserSuggest) error {
	res := r.DB.Table("users").Updates(us)
	return res.Error
}

func (r *UserRepo) GetUserSuggest(ctx context.Context, userId string) (*domain.UserSuggest, error) {
	user := &domain.UserSuggest{}
	res := r.DB.Table("users").Where("id = ?", userId).First(user)
	return user, res.Error
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*domain.UserProfile, error) {
	user := &domain.UserProfile{}
	res := r.DB.Table("users").Where("username = ?", username).First(user)
	return user, res.Error
}

func (r *UserRepo) GetById(ctx context.Context, userId string) (*domain.User, error) {
	user := &domain.User{}
	res := r.DB.Where("id = ?", userId).First(user)
	return user, res.Error
}

func (r *UserRepo) Create(ctx context.Context, user *domain.CreateUser) error {
	res := r.DB.Table("users").Create(user)
	return res.Error
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userId string, newPassword string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"password": newPassword, "is_verified_email": true})
	return res.Error
}

func (r *UserRepo) UpdateUserStatus(ctx context.Context, userId string, status string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"status": status})
	return res.Error
}

func (r *UserRepo) VerifyPhone(ctx context.Context, userId string, phone string, otp string, expirationTime time.Time) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"phone": phone, "expiration_time": expirationTime, "otp": otp})
	return res.Error
}

func (r *UserRepo) VerifyPhoneOTP(ctx context.Context, userId string, otp string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"is_verified_phone": true})
	return res.Error
}
