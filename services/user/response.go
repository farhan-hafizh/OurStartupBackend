package user

type UserBasicResponse struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Username   string `json:"username"`
}

type UserWithProfileResponse struct {
	UserBasicResponse        // embed
	AvatarFile        string `json:"avatar_url"`
}

type RegisterResponse struct {
	UserBasicResponse        //embed
	Email             string `json:"email"`
}

type LoginResponse struct {
	UserBasicResponse        //embed
	Token             string `json:"token"`
}

func FormatUserBasicResponse(user User) UserBasicResponse {
	return UserBasicResponse{
		Name:       user.Name,
		Occupation: user.Occupation,
		Username:   user.Username,
	}
}
func FormatUserWithProfileResponse(user User) UserWithProfileResponse {
	return UserWithProfileResponse{
		UserBasicResponse: FormatUserBasicResponse(user),
		AvatarFile:        user.AvatarFileName,
	}
}

func FormatRegisterResponse(user User) RegisterResponse {
	return RegisterResponse{
		UserBasicResponse: FormatUserBasicResponse(user),
		Email:             user.Email,
	}
}

func FormatLoginResponse(user User, token string) LoginResponse {
	return LoginResponse{
		UserBasicResponse: FormatUserBasicResponse(user),
		Token:             token,
	}
}
