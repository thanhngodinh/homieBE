package domain

type LikePost struct {
	UserId string `json:"userId" gorm:"column:user_id;primary_key"`
	PostId string `json:"postId" gorm:"column:post_id;primary_key"`
}
