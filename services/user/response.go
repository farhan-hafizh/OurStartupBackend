package user

type RegisterResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

func FormatRegisterResponse(user User) RegisterResponse {
	return RegisterResponse{
		ID:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
	}
}
