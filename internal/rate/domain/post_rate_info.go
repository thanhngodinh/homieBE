package domain

type PostRateInfo struct {
	Id       int64   `json:"id" gorm:"column:id;primary_key"`
	PostId   string  `json:"-" gorm:"column:post_id"`
	Star1    *int    `json:"star1" gorm:"column:star1"`
	Star2    *int    `json:"star2" gorm:"column:star2"`
	Star3    *int    `json:"star3" gorm:"column:star3"`
	Star4    *int    `json:"star4" gorm:"column:star4"`
	Star5    *int    `json:"star5" gorm:"column:star5"`
	Total    *int64  `json:"-" gorm:"column:total"`
	AvgRate  float64 `json:"avgRate" gorm:"-"`
	RateList []Rate  `json:"rateList" gorm:"-"`
}
