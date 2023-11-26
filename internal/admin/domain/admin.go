package domain

import "time"

type Admin struct {
	Id             string     `json:"id" gorm:"column:id;primary_key"`
	Username       string     `json:"username" gorm:"column:username"`
	Password       string     `json:"password" gorm:"column:password"`
	Phone          string     `json:"phone,omitempty" gorm:"column:phone"`
	Email          string     `json:"email,omitempty" gorm:"column:email"`
	Avatar         string     `json:"avatar,omitempty" gorm:"column:avatar_url"`
	Name           string     `json:"name,omitempty" gorm:"column:display_name"`
	CreatedAt      *time.Time `json:"createdAt,omitempty" gorm:"colum:created_at"`
	CreatedBy      string     `json:"createdBy,omitempty" gorm:"colum:created_by"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	UpdatedBy      string     `json:"updatedBy,omitempty" gorm:"colum:updated_by"`
}
