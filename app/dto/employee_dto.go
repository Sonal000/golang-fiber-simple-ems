package dto

import "employeeManagement/app/entities"

type EmployeeRequestDto struct {
	Name       string  `json:"name" validate:"required"`
	Position   string  `json:"position" validate:"required"`
	Department string  `json:"department" validate:"required"`
	Salary     float64 `json:"salary" validate:"required,min=1"`
	Email      string  `json:"email" validate:"required,email"`
	Password   string  `json:"password" validate:"required"`
}
type UpdateEmployeeRequestDto struct {
	Name       string  `json:"name"`
	Position   string  `json:"position"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	Email      string  `json:"email"`
}

type EmployeeResponseDto struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Position   string  `json:"position"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	UserID     string  `json:"user_id"`
}
type EmployeeUserResponseDto struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Position   string  `json:"position"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	UserID     string  `json:"user_id"`
	Email      string  `json:"email"`
}

func ToEmployeeResponseDto(employee entities.Employee) EmployeeResponseDto {
	return EmployeeResponseDto{
		Id:         employee.Id.String(),
		Name:       employee.Name,
		Position:   employee.Position,
		Department: employee.Department,
		Salary:     employee.Salary,
		UserID:     employee.UserID.String(),
	}
}
func ToEmployeeUserResponseDto(employee entities.Employee, user entities.User) EmployeeUserResponseDto {
	return EmployeeUserResponseDto{
		Id:         employee.Id.String(),
		Name:       employee.Name,
		Position:   employee.Position,
		Department: employee.Department,
		Salary:     employee.Salary,
		UserID:     employee.UserID.String(),
		Email:      user.Email,
	}
}

func ToEmployeeEntity(employee EmployeeRequestDto, user entities.User) entities.Employee {
	return entities.Employee{
		Name:       employee.Name,
		Position:   employee.Position,
		Department: employee.Department,
		Salary:     employee.Salary,
		UserID:     user.Id,
		User:       user,
	}
}

func ToUserEmployeeEntity(employee EmployeeRequestDto) entities.User {
	return entities.User{
		Email:    employee.Email,
		Password: employee.Password}
}
