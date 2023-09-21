package repository

import (
	"context"
	"fmt"
	"hostel-service/internal/hostel/domain"

	"gorm.io/gorm"
)

func NewHostelAdapter(db *gorm.DB) *HostelAdapter {
	return &HostelAdapter{DB: db}
}

type HostelAdapter struct {
	DB *gorm.DB
}

func (r *HostelAdapter) GetHostels(ctx context.Context, hostel *domain.HostelFilter) ([]domain.Hostel, int64, error) {
	var (
		tx     *gorm.DB
		hotels []domain.Hostel
	)

	tx = r.DB.Table("hostels")
	if hostel.Name != nil && len(*hostel.Name) > 0 {
		tx = tx.Where("name ilike ?", fmt.Sprintf("%%%v%%", *hostel.Name))
	}
	if hostel.Province != nil && len(*hostel.Province) > 0 {
		tx = tx.Where("province ilike ?", fmt.Sprintf("%%%v%%", *hostel.Province))
	}
	if hostel.District != nil && len(*hostel.District) > 0 {
		tx = tx.Where("district ilike ?", fmt.Sprintf("%%%v%%", *hostel.District))
	}
	if hostel.Ward != nil && len(*hostel.Ward) > 0 {
		tx = tx.Where("ward = ilike ?", fmt.Sprintf("%%%v%%", *hostel.Ward))
	}
	if hostel.Street != nil && len(*hostel.Street) > 0 {
		tx = tx.Where("street ilike ?", fmt.Sprintf("%%%v%%", *hostel.Street))
	}
	if hostel.PostType != nil && len(*hostel.PostType) > 0 {
		tx = tx.Where("post_type = ?", hostel.PostType)
	}
	if hostel.Status != nil && len(*hostel.Status) > 0 {
		tx = tx.Where("status = ?", hostel.Status)
	}
	if hostel.CostFrom != nil {
		tx = tx.Where("cost >= ?", hostel.CostFrom)
	}
	if hostel.CostTo != nil {
		tx = tx.Where("cost <= ?", hostel.CostTo)
	}
	if hostel.CreatedAt != nil {
		tx = tx.Where("created_at = ?", hostel.CreatedAt)
	}
	if hostel.CreatedBy != nil && len(*hostel.CreatedBy) > 0 {
		tx = tx.Where("created_by = ?", hostel.CreatedBy)
	}
	res1 := tx.Find(&hotels)
	total := res1.RowsAffected
	res2 := tx.Order(hostel.Sort).Limit(hostel.PageSize).Offset(hostel.PageIdx * hostel.PageSize).Find(&hotels)
	return hotels, total, res2.Error
}

func (r *HostelAdapter) GetHostelById(ctx context.Context, id string) (*domain.Hostel, error) {
	var hostel domain.Hostel
	r.DB.Table("hostels").Where("id = ?", id).First(&hostel)
	return &hostel, nil
}

func (r *HostelAdapter) CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	res := r.DB.Table("hostels").Create(hostel)
	return res.RowsAffected, res.Error
}

func (r *HostelAdapter) UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	res := r.DB.Table("hostels").Model(&hostel).Updates(hostel)
	return res.RowsAffected, res.Error
}

func (r *HostelAdapter) DeleteHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	res := r.DB.Table("hostels").Delete(hostel)
	return res.RowsAffected, res.Error
}
