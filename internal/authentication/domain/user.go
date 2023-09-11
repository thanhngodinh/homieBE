package domain

import "time"

var SECRET_KEY = []byte("secret")

type User struct {
	Id        string     `json:"id" gorm:"column:id;primary_key" swaggerignore:"true"`
	Username  string     `json:"username" gorm:"column:username"`
	Password  string     `json:"password" gorm:"column:password" swagger:"ignoreResponse"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"colum:created_at" swaggerignore:"true"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"colum:updated_at" swaggerignore:"true"`
}

type AccessToken struct {
	TokenString string `json:"access_token"`
}
