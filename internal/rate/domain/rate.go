package domain

import (
	"time"

	"gorm.io/gorm"
)

type Rate struct {
	UserId    string         `json:"userId,omitempty" gorm:"column:user_id;primary_key"`
	PostId    string         `json:"postId,omitempty" gorm:"column:post_id;primary_key"`
	Star      int            `json:"star,omitempty" gorm:"column:star"`
	Comment   string         `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"colum:created_at"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`

	AuthorName   string `json:"authorName,omitempty" gorm:"column:display_name;->"`
	AuthorAvatar string `json:"authorAvatar,omitempty" gorm:"column:avatar_url;->"`
}
