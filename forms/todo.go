package forms

type PostTodo struct {
	Title    string `json:"title" binding:"required"`
	Status   int    `json:"status" binding:"gte=0,lte=5"`
	UpdateAt int64  `json:"update_at"`
}
