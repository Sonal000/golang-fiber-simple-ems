package dto

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Token string          `json:"token"`
	User  UserResponseDto `json:"user"`
}
