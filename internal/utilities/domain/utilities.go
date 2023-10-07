package domain

import (
	"time"

	"gorm.io/gorm"
)

type Utilities struct {
	Id        string         `json:"id" gorm:"column:id;primary_key"`
	Name      string         `json:"name" gorm:"column:name"`
	Icon      string         `json:"icon" gorm:"column:icon"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"colum:created_at"`
	CreatedBy string         `json:"createdBy,omitempty" gorm:"colum:created_by"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	UpdatedBy *string        `json:"updatedBy,omitempty" gorm:"colum:updated_by"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type GetUtilities struct {
	Id   string `json:"id" gorm:"column:id;primary_key"`
	Name string `json:"name" gorm:"column:name"`
	Icon string `json:"icon" gorm:"column:icon"`
}
