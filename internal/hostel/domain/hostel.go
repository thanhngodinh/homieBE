package domain

import (
	"time"
)

type HostelStatus string

const (
	HostelActive   HostelStatus = "A"
	HostelInActive HostelStatus = "I"
	HostelDelete   HostelStatus = "D"
)

type Hostel struct {
	Id               string       `json:"id" gorm:"column:id;primary_key" example:"07e7a76c-1bbb-11ed-861d-0242ac120002" swaggerignore:"true"`
	Name             string       `json:"name" gorm:"column:name" example:"Robert Robertson"`
	Province         string       `json:"province,omitempty" gorm:"column:province" example:"Titao"`
	District         string       `json:"district,omitempty" gorm:"column:district" example:"Bedford"`
	Ward             string       `json:"ward,omitempty" gorm:"column:ward" example:"Bedford"`
	Street           string       `json:"street,omitempty" gorm:"column:street" example:"144 J B Hazra Road"`
	PostType         string       `json:"type,omitempty" gorm:"column:post_type" example:"H"`
	Status           HostelStatus `json:"status,omitempty" gorm:"column:status" example:"A"`
	Cost             int          `json:"cost,omitempty" gorm:"column:cost" example:"1000000"`
	ElectricityPrice int          `json:"electricity_price,omitempty" gorm:"column:electricity_price" example:"100000"`
	WaterPrice       int          `json:"water_price,omitempty" gorm:"column:water_price" example:"100000"`
	ParkingPrice     int          `json:"parking_price,omitempty" gorm:"column:parking_price" example:"100000"`
	WifiPrice        int          `json:"wifi_price,omitempty" gorm:"column:wifi_price" example:"100000"`
	Capacity         int          `json:"capacity,omitempty" gorm:"column:capacity" example:"1"`
	Area             int          `json:"area,omitempty" gorm:"column:area" example:"20"`
	Description      *string      `json:"description,omitempty" gorm:"column:description" example:"Nha tro sieu dep"`
	CreatedAt        time.Time    `json:"createdAt,omitempty" gorm:"colum:created_at" example:"2006-01-02 03:04:07" swaggerignore:"true"`
	CreatedBy        string       `json:"createdBy,omitempty" gorm:"colum:created_by" example:"07e7a76c-1bbb-11ed-861d-0242ac120002" swaggerignore:"true"`
	UpdatedAt        *time.Time   `json:"updatedAt,omitempty" gorm:"colum:updated_at" example:"2006-01-02 03:04:07" swaggerignore:"true"`
	UpdatedBy        *string      `json:"updatedBy,omitempty" gorm:"colum:updated_by" example:"07e7a76c-1bbb-11ed-861d-0242ac120002" swaggerignore:"true"`
	DeletedAt        *time.Time   `json:"-" gorm:"colum:deleted_at" example:"2006-01-02 03:04:07" swaggerignore:"true"`
}

type GetHostelsResponse struct {
	Data  []Hostel `json:"data"`
	Total int64    `json:"total"`
}
