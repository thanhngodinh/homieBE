package domain

var SECRET_KEY = []byte("secret")

type User struct {
	Id              string  `json:"id" gorm:"column:id;primary_key"`
	Username        string  `json:"username" gorm:"column:username"`
	Password        string  `json:"password" gorm:"column:password"`
	Phone           *string `json:"phone,omitempty" gorm:"column:phone"`
	Email           *string `json:"email,omitempty" gorm:"column:email"`
	IsEmailVerified string  `json:"is_email_verified,omitempty" gorm:"column:is_email_verified"`
	Birthdate       *string `json:"birth_date,omitempty" gorm:"column:birth_date"`
	Gender          *string `json:"gender,omitempty" gorm:"column:gender"`
	FirstName       *string `json:"first_name,omitempty" gorm:"column:first_name"`
	LastName        *string `json:"last_nam,omitempty" gorm:"column:last_name"`
	MiddleName      *string `json:"middle_name,omitempty" gorm:"column:middle_name"`
	Address         *string `json:"address,omitempty" gorm:"column:address"`
	CreatedAt       *string `json:"createdAt,omitempty" gorm:"colum:created_at"`
	CreatedBy       *string `json:"createdBy,omitempty" gorm:"colum:created_by"`
	UpdatedAt       *string `json:"updatedAt,omitempty" gorm:"colum:updated_at"`
	UpdatedBy       *string `json:"updatedBy,omitempty" gorm:"colum:updated_by"`
}

type AccessToken struct {
	TokenString string `json:"access_token"`
}
