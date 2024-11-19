package handler

import (
	"employeeManagement/app/dto"
	"employeeManagement/app/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	loginReq := dto.LoginRequestDto{}
	if err := c.BodyParser(&loginReq); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Message: "Invalid request",
		})
		return err
	}
	res, err := handler.authService.Login(loginReq)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
		})
		return err
	}
	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Login successful",
		Data:    res,
	})
	return nil
}

func (handler *AuthHandler) RegisterAdmin(c *fiber.Ctx) error {
	userReq := dto.UserRequestDto{}
	if err := c.BodyParser(&userReq); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Message: "Invalid request",
		})
		return err
	}
	res, err := handler.authService.RegisterAdmin(userReq)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
		})
		return err
	}
	c.Status(fiber.StatusCreated).JSON(dto.Response{
		Message: "Admin created successfully",
		Data:    res,
	})
	return nil
}
