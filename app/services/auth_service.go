package services

import (
	"employeeManagement/app/dto"
	repository "employeeManagement/app/repositories"
	"fmt"
)

type AuthService interface {
	Login(user dto.LoginRequestDto) (dto.LoginResponseDto, error)
	RegisterAdmin(user dto.UserRequestDto) (dto.UserResponseDto, error)
}

type authService struct {
	userRepository  repository.UserRepository
	jwtService      JWTService
	passwordService PasswordService
}

func NewAuthService() AuthService {
	return &authService{
		userRepository:  repository.NewUserRepository(),
		jwtService:      NewJWTService(),
		passwordService: NewPasswordService(),
	}
}

func (service *authService) Login(user dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	userEntity, err := service.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("user not found")
	}
	if !service.passwordService.ValidatePassword(userEntity.Password, user.Password) {
		return dto.LoginResponseDto{}, fmt.Errorf("password did not match")
	}

	token, err := service.jwtService.GenerateToken(userEntity.Id.String(), userEntity.Role)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("error while generating token")
	}
	response := dto.LoginResponseDto{
		Token: token,
		User:  dto.ToUserResponseDto(userEntity),
	}

	return response, nil
}

func (service *authService) RegisterAdmin(user dto.UserRequestDto) (dto.UserResponseDto, error) {
	if user.Role != "admin" {
		return dto.UserResponseDto{}, fmt.Errorf("this route is only for admin")
	}
	hashedPassword, err := service.passwordService.HashPassword(user.Password)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("could not hash password: %w", err)
	}
	user.Password = hashedPassword
	admin, err := service.userRepository.CreateUser(dto.ToUserEntity(user))
	if err != nil {
		return dto.UserResponseDto{}, err
	}
	return dto.ToUserResponseDto(admin), nil
}
