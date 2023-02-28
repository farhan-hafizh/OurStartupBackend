package user

type RegisterResponse struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Username   string `json:"username"`
}

type LoginResponse struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Username   string `json:"username"`
	Token      string `json:"token"`
}

func FormatRegisterResponse(user User) RegisterResponse {
	return RegisterResponse{
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Username:   user.Username,
	}
}

func FormatLoginResponse(user User, token string) LoginResponse {
	return LoginResponse{
		Name:       user.Name,
		Occupation: user.Occupation,
		Username:   user.Username,
		Token:      token,
	}
}
