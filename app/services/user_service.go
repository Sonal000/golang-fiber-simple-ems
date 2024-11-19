package services

import (
	"employeeManagement/app/dto"
	repository "employeeManagement/app/repositories"
)

type UserService interface {
	CreateUser(user dto.UserRequestDto) (dto.UserResponseDto, error)
	GetUsers() ([]dto.UserResponseDto, error)
	GetUser(id string) (dto.UserResponseDto, error)
	UpdateUser(id string, user dto.UpdateUserRequestDto) (dto.UserResponseDto, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}

func (service *userService) CreateUser(user dto.UserRequestDto) (dto.UserResponseDto, error) {
	userEntity := dto.ToUserEntity(user)
	newUser, err := service.userRepository.CreateUser(userEntity)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return dto.ToUserResponseDto(newUser), nil
}

func (service *userService) RegisterAdmin(user dto.UserRequestDto) (dto.UserResponseDto, error) {
	userEntity := dto.ToUserEntity(user)
	newUser, err := service.userRepository.CreateUser(userEntity)
	if err != nil {
		return dto.UserResponseDto{}, err
	}
	return dto.ToUserResponseDto(newUser), nil
}

func (service *userService) GetUsers() ([]dto.UserResponseDto, error) {
	userEntities, err := service.userRepository.GetUsers()
	if err != nil {
		return []dto.UserResponseDto{}, err
	}
	var users []dto.UserResponseDto
	for _, entity := range userEntities {
		users = append(users, dto.ToUserResponseDto(entity))
	}

	return users, nil
}

func (service *userService) GetUser(id string) (dto.UserResponseDto, error) {
	user, err := service.userRepository.GetUser(id)
	if err != nil {
		return dto.UserResponseDto{}, err
	}
	return dto.ToUserResponseDto(user), nil
}

func (service *userService) UpdateUser(id string, user dto.UpdateUserRequestDto) (dto.UserResponseDto, error) {
	updatedUser, err := service.userRepository.UpdateUser(id, user)
	if err != nil {
		return dto.UserResponseDto{}, err
	}
	return dto.ToUserResponseDto(updatedUser), nil
}
