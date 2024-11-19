package handler

import (
	"employeeManagement/app/dto"
	"employeeManagement/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler() UserHandler {
	return UserHandler{
		userService: services.NewUserService(),
	}
}

// CreateUser - Handles user creation
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user dto.UserRequestDto
	if err := c.BodyParser(&user); err != nil {
		log.Println("requesting user", user)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}
	newUser, err := h.userService.CreateUser(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
		})
		return err
	}

	c.Status(fiber.StatusCreated).JSON(dto.Response{
		Message: "User created successfully",
		Data:    newUser,
	})
	return nil
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.GetUser(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "User fetched successfully",
		Data:    user,
	})
	return nil
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	c.Status(fiber.StatusOK).JSON(dto.Response{
		Data:    users,
		Message: "Users fetched successfully",
	})
	return nil
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := dto.UpdateUserRequestDto{}
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Message: "Invalid request",
			Data:    nil,
		})
		return err
	}
	updatedUser, err := h.userService.UpdateUser(id, user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}
	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "User updated successfully",
		Data:    updatedUser,
	})
	return nil
}
