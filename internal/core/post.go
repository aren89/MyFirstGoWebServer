package core

type Post struct {
	Id string `json:"id" binding:"required"`
}
