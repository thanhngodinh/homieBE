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

func (r *HostelAdapter) GetHostels(ctx context.Context, filter *domain.HostelFilter, userId string) ([]domain.Hostel, int64, error) {
	var (
		tx     *gorm.DB
		hotels []domain.Hostel
	)

	tx = r.DB.Table("hostels").
		Select(fmt.Sprintf(`
	(select true from user_like_posts where user_like_posts.post_id = hostels.id and user_like_posts.user_id = '%v') as "is_liked",
	array_remove(array_agg(hostels_utilities.utilities_id), NULL) as utilities,
	hostels.*`, userId)).
		Joins("left join hostels_utilities on hostels_utilities.hostel_id = hostels.id").Group("hostels.id")
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
	if len(filter.Utilities) > 0 {
		tx = tx.Where("? <@ (select array_agg(hostels_utilities.utilities_id) from hostels_utilities where hostels_utilities.hostel_id = hostels.id)", filter.Utilities)
	}
	res1 := tx.Scan(&hotels)
	total := res1.RowsAffected
	res2 := tx.Order(filter.Sort).Limit(filter.PageSize).Offset(filter.PageIdx * filter.PageSize).Scan(&hotels)
	return hotels, total, res2.Error
}

func (r *HostelAdapter) GetHostelById(ctx context.Context, id string) (*domain.Hostel, error) {
	var hostel domain.Hostel

	r.DB.Table("hostels").
		Select("hostels.*, array_remove(array_agg(hostels_utilities.utilities_id), NULL) as utilities").
		Joins("left join hostels_utilities on hostels_utilities.hostel_id = hostels.id").Group("hostels.id").
		Where("id = ?", id).First(&hostel)
	r.DB.Table("hostels").Where("id = ?", id).Updates(map[string]interface{}{"view": hostel.View + 1})
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
