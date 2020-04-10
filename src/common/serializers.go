package common

import (
	"github.com/gin-gonic/gin"
)

type PostSerializer struct {
	C  *gin.Context
	Id interface{}
}

type PostResponse struct {
	ID interface{} `json:"id"`
}

func (s *PostSerializer) Response() PostResponse {
	response := PostResponse{
		ID: s.Id,
	}
	return response
}
