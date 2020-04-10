package person

import "time"

type PersonModel struct {
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Age        int       `json:"age" binding:"required,max=100"`
	Gender     string    `json:"gender"`
	Date       time.Time `json:"createdAt"`
	Experience string    `json:"experience"`
}
