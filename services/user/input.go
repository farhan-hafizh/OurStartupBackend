package user

// this used for making sure the input type is correct
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
}

type LoginUserInput struct {
	Query    string `json:"query" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
