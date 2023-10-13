package domain

import "time"

type User struct {
	Id              string     `json:"id" gorm:"column:id;primary_key"`
	Username        string     `json:"username" gorm:"column:username"`
	Password        string     `json:"password" gorm:"column:password"`
	Phone           string     `json:"phone,omitempty" gorm:"column:phone"`
	Email           string     `json:"email,omitempty" gorm:"column:email"`
	IsEmailVerified string     `json:"isVerifiedEmail,omitempty" gorm:"column:is_verified_email"`
	IsFindRoommate  bool       `json:"isFindRoommate,omitempty" gorm:"column:is_find_roommate"`
	DateOfBirth     *time.Time `json:"dateOfBirth,omitempty" gorm:"column:date_of_birth"`
	Avatar          string     `json:"avatar,omitempty" gorm:"column:avatar_url"`
	Gender          string     `json:"gender,omitempty" gorm:"column:gender"`
	FirstName       string     `json:"firstName,omitempty" gorm:"column:first_name"`
	LastName        string     `json:"lastName,omitempty" gorm:"column:last_name"`
	Province        string     `json:"province,omitempty" gorm:"column:province_profile"`
	District        string     `json:"district,omitempty" gorm:"column:district_profile"`
	CostFrom        int        `json:"costFrom,omitempty" gorm:"column:cost_from"`
	CostTo          int        `json:"costTo,omitempty" gorm:"column:cost_to"`
	CreatedAt       *time.Time `json:"createdAt,omitempty" gorm:"colum:created_at"`
	CreatedBy       string     `json:"createdBy,omitempty" gorm:"colum:created_by"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	UpdatedBy       string     `json:"updatedBy,omitempty" gorm:"colum:updated_by"`
}
