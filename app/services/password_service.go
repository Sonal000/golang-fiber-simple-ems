package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password string) (string, error)
	ValidatePassword(hashedPassword, plainPassword string) bool
}

type passwordService struct{}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (s *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func (s *passwordService) ValidatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
