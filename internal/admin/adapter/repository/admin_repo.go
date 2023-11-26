package repository

import (
	"context"
	"hostel-service/internal/admin/domain"

	"gorm.io/gorm"
)

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

type AdminRepo struct {
	DB *gorm.DB
}

func (r *AdminRepo) GetByUsername(ctx context.Context, adminname string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	res := r.DB.Where("adminname = ?", adminname).First(admin)
	return admin, res.Error
}

func (r *AdminRepo) GetById(ctx context.Context, adminId string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	res := r.DB.Table("admins").Where("id = ?", adminId).First(admin)
	return admin, res.Error
}

func (r *AdminRepo) UpdatePassword(ctx context.Context, adminId string, newPassword string) error {
	res := r.DB.Table("admins").Where("id = ?", adminId).Updates(map[string]interface{}{"password": newPassword})
	return res.Error
}
