package repository

import (
	"context"
	"fmt"
	hostel_domain "hostel-service/internal/hostel/domain"
	"hostel-service/internal/user/domain"

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

func (r *UserRepo) GetRoommateById(ctx context.Context, userId string) (*domain.Roommate, error) {
	roommate := &domain.Roommate{}
	res := r.DB.Table("users").Where("id = ?", userId).First(roommate)
	return roommate, res.Error
}

func (r *UserRepo) GetUserPosts(ctx context.Context, userId string) ([]hostel_domain.Hostel, int64, error) {
	hostels := []hostel_domain.Hostel{}
	res := r.DB.Table("hostels").
		Where("created_by = ?", userId).
		Find(&hostels)
	return hostels, res.RowsAffected, res.Error
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

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	res := r.DB.Where("username = ?", username).First(user)
	return user, res.Error
}

func (r *UserRepo) GetById(ctx context.Context, userId string) (*domain.User, error) {
	user := &domain.User{}
	res := r.DB.Where("id = ?", userId).First(user)
	return user, res.Error
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	res := r.DB.Create(user)
	return res.Error
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userId string, newPassword string) error {
	res := r.DB.Table("users").Where("id = ?", userId).Updates(map[string]interface{}{"password": newPassword})
	return res.Error
}
