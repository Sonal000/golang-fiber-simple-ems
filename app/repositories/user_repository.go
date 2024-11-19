package repository

import (
	"employeeManagement/app/dto"
	"employeeManagement/app/entities"
	db "employeeManagement/pkg"
	"errors"
	"fmt"

	// "log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}
type UserRepository interface {
	CreateUser(user entities.User) (entities.User, error)
	GetUsers() ([]entities.User, error)
	GetUser(id string) (entities.User, error)
	UpdateUser(id string, user dto.UpdateUserRequestDto) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetDBConnection(),
	}
}

func (r *userRepository) CreateUser(user entities.User) (entities.User, error) {

	if err := r.db.Create(&user).Error; err != nil {
		// log.Printf("Error while creating user: %+v, err: %+v", user, err)
		return entities.User{}, err
	}

	return user, nil
}
func (r *userRepository) GetUser(id string) (entities.User, error) {
	userEntity := entities.User{}
	if err := r.db.Where("id = ?", id).First(&userEntity).Error; err != nil {
		return userEntity, err
	}
	return userEntity, nil
}

func (u *userRepository) GetUsers() ([]entities.User, error) {
	users := []entities.User{}
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) UpdateUser(id string, user dto.UpdateUserRequestDto) (entities.User, error) {
	existingUser := entities.User{}
	if err := u.db.First(&existingUser, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, fmt.Errorf("user not found")
		}
		return entities.User{}, err
	}
	if err := u.db.Model(&existingUser).Updates(entities.User{
		Email: user.Email,
	}).Error; err != nil {
		return entities.User{}, err
	}
	return existingUser, nil
}

func (r *userRepository) GetUserByEmail(email string) (entities.User, error) {
	user := entities.User{}
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
