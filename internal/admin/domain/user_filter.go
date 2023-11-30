package domain

type UserFilter struct {
	Id       string `json:"-" gorm:"column:id;primary_key"`
	Gender   string `json:"gender,omitempty" gorm:"column:gender"`
	Name     string `json:"name,omitempty" gorm:"column:display_name"`
	Status   string `json:"status,omitempty" gorm:"column:status"`
	PageSize int    `json:"pageSize,omitempty"`
	PageIdx  int    `json:"pageIdx,omitempty"`
	Sort     string `json:"sort,omitempty"`
}
