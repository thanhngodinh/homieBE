package domain

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type HostelStatus string

const (
	HostelActive   HostelStatus = "A"
	HostelWaiting  HostelStatus = "W"
	HostelInActive HostelStatus = "I"
)

type Hostel struct {
	Id               string         `json:"id" gorm:"column:id;primary_key"`
	Name             string         `json:"name" gorm:"column:name"`
	Province         string         `json:"province,omitempty" gorm:"column:province"`
	District         string         `json:"district,omitempty" gorm:"column:district"`
	Ward             string         `json:"ward,omitempty" gorm:"column:ward"`
	Street           string         `json:"street,omitempty" gorm:"column:street"`
	Status           HostelStatus   `json:"status,omitempty" gorm:"column:status;type:text"`
	Cost             int            `json:"cost,omitempty" gorm:"column:cost"`
	Deposit          int            `json:"deposit,omitempty" gorm:"column:deposit"`
	ElectricityPrice int            `json:"electricityPrice,omitempty" gorm:"column:electricity_price"`
	WaterPrice       int            `json:"waterPrice,omitempty" gorm:"column:water_price"`
	ParkingPrice     int            `json:"parkingPrice,omitempty" gorm:"column:parking_price"`
	ServicePrice     int            `json:"servicePrice,omitempty" gorm:"column:service_price"`
	Capacity         int            `json:"capacity,omitempty" gorm:"column:capacity"`
	Area             int            `json:"area,omitempty" gorm:"column:area"`
	Phone            string         `json:"phone,omitempty" gorm:"column:phone"`
	ImageUrl         pq.StringArray `json:"imageUrl,omitempty" gorm:"column:image_url;type:text[]"`
	Utilities        pq.StringArray `json:"utilities,omitempty" gorm:"utilities;type:text[]"`
	IsLiked          bool           `json:"isLiked" gorm:"is_liked"`
	View             int            `json:"-" gorm:"column:view"`
	Description      string         `json:"description,omitempty" gorm:"column:description"`
	CreatedAt        time.Time      `json:"createdAt,omitempty" gorm:"colum:created_at"`
	CreatedBy        string         `json:"createdBy,omitempty" gorm:"colum:created_by"`
	UpdatedAt        *time.Time     `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	UpdatedBy        *string        `json:"updatedBy,omitempty" gorm:"colum:updated_by"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}

type GetHostelsResponse struct {
	Data  []Hostel `json:"data"`
	Total int64    `json:"total"`
}

type HostelUtilities struct {
	HostelId    string `gorm:"column:hostel_id"`
	UtilitiesId string `gorm:"column:utilities_id"`
}
