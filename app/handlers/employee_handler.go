package handler

import (
	"employeeManagement/app/dto"
	"employeeManagement/app/services"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	employeeService services.EmployeeService
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: services.NewEmployeeService(),
	}
}

func (handler *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	employee := dto.EmployeeRequestDto{}
	if err := c.BodyParser(&employee); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}
	newEmp, err := handler.employeeService.CreateEmployee(employee)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}
	c.Status(fiber.StatusCreated).JSON(dto.Response{
		Message: "Employee created successfully",
		Data:    newEmp,
	})
	return nil
}

func (handler *EmployeeHandler) GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := handler.employeeService.GetEmployee(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Employee fetched successfully",
		Data:    employee,
	})
	return nil
}
func (handler *EmployeeHandler) GetEmployeeProfile(c *fiber.Ctx) error {
	userClaims, ok := c.Locals("userClaims").(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(dto.Response{
			Message: "You need to be logged in to access this resource",
		})
	}
	id, ok := userClaims["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(dto.Response{
			Message: "user not found",
		})
	}
	employee, err := handler.employeeService.GetEmployeeProfile(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Employee fetched successfully",
		Data:    employee,
	})
	return nil
}

func (handler *EmployeeHandler) GetEmployees(c *fiber.Ctx) error {
	employees, err := handler.employeeService.GetEmployees()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Employees fetched successfully",
		Data:    employees,
	})
	return nil
}

func (handler *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	employee := dto.UpdateEmployeeRequestDto{}
	if err := c.BodyParser(&employee); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Message: "Invalid request",
			Data:    nil,
		})
		return err
	}
	id := c.Params("id")
	updatedEmployee, err := handler.employeeService.UpdateEmployee(id, employee)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}
	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Employee updated successfully",
		Data:    updatedEmployee,
	})
	return nil
}

func (handler *EmployeeHandler) UpdateEmployeeProfile(c *fiber.Ctx) error {
	employee := dto.UpdateEmployeeRequestDto{}
	if err := c.BodyParser(&employee); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Message: "Invalid request",
			Data:    nil,
		})
		return err
	}
	userClaims, ok := c.Locals("userClaims").(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(dto.Response{
			Message: "You need to be logged in to access this resource",
		})
	}
	id, ok := userClaims["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(dto.Response{
			Message: "user not found",
		})
	}
	updatedEmployee, err := handler.employeeService.UpdateEmployeeProfile(id, employee)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}
	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Employee updated successfully",
		Data:    updatedEmployee,
	})
	return nil
}

func (handler *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	err := handler.employeeService.DeleteEmployee(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
		})
		return err
	}
	c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "Employee deleted successfully",
	})
	return nil
}
