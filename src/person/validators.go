package person

type PersonModel struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,max=100"`
	Gender    int    `json:"gender"`
}
