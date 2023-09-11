package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hostel-service/internal/hostel/domain"

	q "github.com/core-go/sql"
	"github.com/sirupsen/logrus"
)

func NewHostelAdapter(db *sql.DB) *HostelAdapter {
	return &HostelAdapter{DB: db}
}

type HostelAdapter struct {
	DB *sql.DB
}

func (r *HostelAdapter) GetHostels(ctx context.Context, pageSize int, pageIdx int) ([]domain.Hostel, int64, error) {
	var total int64
	stmt, err := r.DB.Prepare(`
		select 
			id, name, province, district, ward, street, status, cost,
			electricity_price, water_price, parking_price, wifi_price, capacity, area, decription,
			created_at, created_by, updated_at, updated_by
		from posts limit $1 offset $2`)
	if err != nil {
		return nil, total, err
	}
	rows, err := stmt.Query(pageSize, pageIdx*pageSize)
	if err != nil {
		return nil, total, err
	}
	defer rows.Close()
	var hostels []domain.Hostel
	for rows.Next() {
		var hostel domain.Hostel
		err := rows.Scan(
			&hostel.Id,
			&hostel.Name,
			&hostel.Province,
			&hostel.District,
			&hostel.Ward,
			&hostel.Street,
			&hostel.Status,
			&hostel.Cost,
			&hostel.ElectricityPrice,
			&hostel.WaterPrice,
			&hostel.ParkingPrice,
			&hostel.WifiPrice,
			&hostel.Capacity,
			&hostel.Area,
			&hostel.Description,
			&hostel.CreatedAt,
			&hostel.CreatedBy,
			&hostel.UpdatedAt,
			&hostel.UpdatedBy,
		)
		if err != nil {
			return hostels, total, err
		}
		hostels = append(hostels, hostel)
	}
	if err != nil {
		return hostels, total, err
	}
	stmt, err = r.DB.Prepare(`
		select 
			count(*) as total
		from posts`)
	if err != nil {
		return nil, total, err
	}
	if err := stmt.QueryRow().Scan(&total); err != nil {
		return nil, total, err
	}
	return hostels, total, nil
}

func (r *HostelAdapter) GetHostelById(ctx context.Context, id string) (*domain.Hostel, error) {
	var hostel domain.Hostel
	stmt, err := r.DB.Prepare(`
		select 
			id, name, province, district, ward, street, status, cost,
			electricity_price, water_price, parking_price, wifi_price, capacity, area, decription,
			created_at, created_by, updated_at, updated_by
		from posts where id = $1 and post_type = 'H' limit 1`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(
		&hostel.Id,
		&hostel.Name,
		&hostel.Province,
		&hostel.District,
		&hostel.Ward,
		&hostel.Street,
		&hostel.Status,
		&hostel.Cost,
		&hostel.ElectricityPrice,
		&hostel.WaterPrice,
		&hostel.ParkingPrice,
		&hostel.WifiPrice,
		&hostel.Capacity,
		&hostel.Area,
		&hostel.Description,
		&hostel.CreatedAt,
		&hostel.CreatedBy,
		&hostel.UpdatedAt,
		&hostel.UpdatedBy,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return &hostel, nil
}

func (r *HostelAdapter) CreateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	tx := q.GetTx(ctx)
	stmt, err := tx.Prepare(`
		insert 
			into posts (
			id, name, province, district, ward, street, post_type, status, cost,
			electricity_price, water_price, parking_price, wifi_price, capacity, area, decription,
			created_at, created_by) 
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)`)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(
		&hostel.Id,
		&hostel.Name,
		&hostel.Province,
		&hostel.District,
		&hostel.Ward,
		&hostel.Street,
		&hostel.PostType,
		&hostel.Status,
		&hostel.Cost,
		&hostel.ElectricityPrice,
		&hostel.WaterPrice,
		&hostel.ParkingPrice,
		&hostel.WifiPrice,
		&hostel.Capacity,
		&hostel.Area,
		&hostel.Description,
		&hostel.CreatedAt,
		&hostel.CreatedBy,
	)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *HostelAdapter) UpdateHostel(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	tx := q.GetTx(ctx)
	stmt, err := tx.Prepare(`
		update posts 
		set 
			name = $1, province = $2, district = $3, ward = $4, street = $5, status = $6, cost = $7,
			electricity_price = $8, water_price = $9, parking_price = $10, wifi_price = $11, capacity = $12, area = $13, decription = $14,
			updated_at = $15, updated_by = $16
		where id = $17`)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(
		&hostel.Name,
		&hostel.Province,
		&hostel.District,
		&hostel.Ward,
		&hostel.Street,
		&hostel.Status,
		&hostel.Cost,
		&hostel.ElectricityPrice,
		&hostel.WaterPrice,
		&hostel.ParkingPrice,
		&hostel.WifiPrice,
		&hostel.Capacity,
		&hostel.Area,
		&hostel.Description,
		&hostel.UpdatedAt,
		&hostel.UpdatedBy,
	)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *HostelAdapter) DeleteHostel(ctx context.Context, id string) (int64, error) {
	tx := q.GetTx(ctx)
	stmt, err := tx.Prepare(`update posts set status = 'D' where id = $1`)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logrus.Warn(err)
		return -1, err
	}
	return res.RowsAffected()
}

func inValidField(obj interface{}, field string, value string, excludeValue string) (bool, error) {
	var notValid bool
	var stmt *sql.Stmt
	var err error
	query := fmt.Sprintf(`select if(count(*), 'true', 'false') as no_valid from posts where %s = ? and not id = ?`, field)
	if db, ok := obj.(*sql.DB); ok {
		stmt, err = db.Prepare(query)
		if err != nil {
			return false, err
		}
	} else if tx, ok := obj.(*sql.Tx); ok {
		stmt, err = tx.Prepare(query)
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("unknow db handler type")
	}
	if err = stmt.QueryRow(value, excludeValue).Scan(&notValid); err != nil {
		return false, err
	}
	return notValid, nil
}
