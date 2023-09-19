package domain

type UserLikePosts struct {
	Id     string `json:"id" gorm:"column:id;primary_key"`
	UserId string `json:"userId" gorm:"column:user_id"`
	PostId string `json:"postId" gorm:"column:post_id"`
}
