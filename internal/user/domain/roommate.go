package domain

import (
	"time"

	"github.com/lib/pq"
)

type Roommate struct {
	Id          string         `json:"id" gorm:"column:id;primary_key"`
	Phone       string         `json:"phone,omitempty" gorm:"column:phone"`
	Email       string         `json:"email,omitempty" gorm:"column:email"`
	DateOfBirth *time.Time     `json:"dateOfBirth,omitempty" gorm:"column:date_of_birth"`
	Avatar      string         `json:"avatar,omitempty" gorm:"column:avatar_url"`
	Gender      string         `json:"gender,omitempty" gorm:"column:gender"`
	FirstName   string         `json:"firstName,omitempty" gorm:"column:first_name"`
	LastName    string         `json:"lastName,omitempty" gorm:"column:last_name"`
	Province    string         `json:"province,omitempty" gorm:"column:province_profile"`
	District    pq.StringArray `json:"district,omitempty" gorm:"column:district_profile;type:text[]"`
	CostFrom    int            `json:"costFrom,omitempty" gorm:"column:cost_from"`
	CostTo      int            `json:"costTo,omitempty" gorm:"column:cost_to"`
}

type RoommateFilter struct {
	Id       string         `json:"-" gorm:"column:id;primary_key"`
	Gender   string         `json:"gender,omitempty" gorm:"column:gender"`
	Name     string         `json:"name,omitempty" gorm:"column:first_name"`
	Province string         `json:"province,omitempty" gorm:"column:province_profile"`
	District pq.StringArray `json:"district,omitempty" gorm:"column:district_profile"`
	CostFrom int            `json:"costFrom,omitempty" gorm:"column:cost_from"`
	CostTo   int            `json:"costTo,omitempty" gorm:"column:cost_to"`
	PageSize int            `json:"pageSize,omitempty"`
	PageIdx  int            `json:"pageIdx,omitempty"`
	Sort     string         `json:"sort,omitempty"`
}
