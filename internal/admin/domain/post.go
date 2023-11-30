package domain

import (
	"time"

	rate_domain "hostel-service/internal/rate/domain"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type PostStatus string

const (
	PostActive   PostStatus = "A"
	PostWaiting  PostStatus = "W"
	PostInActive PostStatus = "I"
)

type Post struct {
	Id               string                    `json:"id" gorm:"column:id;primary_key"`
	Name             string                    `json:"name" gorm:"column:name"`
	Province         string                    `json:"province,omitempty" gorm:"column:province"`
	District         string                    `json:"district,omitempty" gorm:"column:district"`
	Ward             string                    `json:"ward,omitempty" gorm:"column:ward"`
	Street           string                    `json:"street,omitempty" gorm:"column:street"`
	Status           PostStatus                `json:"status,omitempty" gorm:"column:status;type:text"`
	Cost             int                       `json:"cost,omitempty" gorm:"column:cost"`
	Deposit          int                       `json:"deposit,omitempty" gorm:"column:deposit"`
	ElectricityPrice int                       `json:"electricityPrice,omitempty" gorm:"column:electricity_price"`
	WaterPrice       int                       `json:"waterPrice,omitempty" gorm:"column:water_price"`
	ParkingPrice     int                       `json:"parkingPrice,omitempty" gorm:"column:parking_price"`
	ServicePrice     int                       `json:"servicePrice,omitempty" gorm:"column:service_price"`
	Capacity         int                       `json:"capacity,omitempty" gorm:"column:capacity"`
	Area             int                       `json:"area,omitempty" gorm:"column:area"`
	Phone            string                    `json:"phone,omitempty" gorm:"column:phone"`
	ImageUrl         pq.StringArray            `json:"imageUrl,omitempty" gorm:"column:image_url;type:text[]"`
	Utilities        pq.StringArray            `json:"utilities,omitempty" gorm:"utilities;type:text[];->"`
	Like             int                       `json:"like" gorm:"like;->"`
	View             int                       `json:"-" gorm:"column:view"`
	Description      string                    `json:"description,omitempty" gorm:"column:description"`
	AuthorId         string                    `json:"authorId" gorm:"column:author_id;->"`
	AuthorName       string                    `json:"authorName" gorm:"column:author_name;->"`
	AuthorAvatar     string                    `json:"authorAvatar" gorm:"column:author_avatar;->"`
	RateInfo         *rate_domain.PostRateInfo `json:"rateInfo" gorm:"column:rate_info;->"`
	CreatedAt        time.Time                 `json:"createdAt,omitempty" gorm:"colum:created_at"`
	EndedAt          time.Time                 `json:"endedAt,omitempty" gorm:"colum:ended_at"`
	CreatedBy        string                    `json:"createdBy,omitempty" gorm:"colum:created_by"`
	UpdatedAt        *time.Time                `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	UpdatedBy        *string                   `json:"updatedBy,omitempty" gorm:"colum:updated_by"`
	DeletedAt        gorm.DeletedAt            `json:"-"`
}

type PostUtilities struct {
	PostId      string `gorm:"column:post_id"`
	UtilitiesId string `gorm:"column:utility_id"`
}

type RateInfo struct {
	PostId string `gorm:"column:post_id"`
}

type Compare struct {
	Post1 Post `json:"post1"`
	Post2 Post `json:"post2"`
}
