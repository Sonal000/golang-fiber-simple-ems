package dto

import "employeeManagement/app/entities"

type UserResponseDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type UserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type AdminRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserRequestDto struct {
	Email string `json:"email"`
}

func ToUserResponseDto(user entities.User) UserResponseDto {
	return UserResponseDto{
		Id:    user.Id.String(),
		Email: user.Email,
		Role:  user.Role,
	}
}

func ToUserEntity(user UserRequestDto) entities.User {
	return entities.User{
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
