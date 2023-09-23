package domain

type UserSuggest struct {
	Id       string `json:"id,omitempty" gorm:"column:id;primary_key"`
	Province string `json:"province,omitempty" gorm:"column:province_suggest"`
	District string `json:"distric,omitemptyt" gorm:"column:district_suggest"`
	Cost     int    `json:"cost,omitempty" gorm:"column:cost_suggest"`
	Capacity int    `json:"capacity,omitempty" gorm:"column:capacity_suggest"`
	Gender   string `json:"gender,omitempty" gorm:"column:gender"`
}

type UpdateUserSuggest struct {
	Id       string `json:"id,omitempty" gorm:"column:id;primary_key"`
	Province string `json:"province,omitempty" gorm:"column:province_suggest"`
	District string `json:"distric,omitemptyt" gorm:"column:district_suggest"`
	Cost     int    `json:"cost,omitempty" gorm:"column:cost_suggest"`
	Capacity int    `json:"capacity,omitempty" gorm:"column:capacity_suggest"`
}
