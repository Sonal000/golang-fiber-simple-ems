package services

import (
	"employeeManagement/app/dto"
	repository "employeeManagement/app/repositories"
	"fmt"
)

type EmployeeService interface {
	CreateEmployee(e dto.EmployeeRequestDto) (dto.EmployeeResponseDto, error)
	GetEmployees() ([]dto.EmployeeResponseDto, error)
	GetEmployee(id string) (dto.EmployeeUserResponseDto, error)
	GetEmployeeProfile(id string) (dto.EmployeeUserResponseDto, error)
	UpdateEmployee(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error)
	UpdateEmployeeProfile(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error)
	DeleteEmployee(id string) error
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
	passwordService    PasswordService
	userRepository     repository.UserRepository
}

func NewEmployeeService() EmployeeService {
	return &employeeService{
		employeeRepository: repository.NewEmployeeRepository(),
		passwordService:    NewPasswordService(),
		userRepository:     repository.NewUserRepository(),
	}
}

func (service *employeeService) CreateEmployee(employee dto.EmployeeRequestDto) (dto.EmployeeResponseDto, error) {
	hashedPassword, err := service.passwordService.HashPassword(employee.Password)
	if err != nil {
		return dto.EmployeeResponseDto{}, fmt.Errorf("could not hash password: %w", err)
	}
	employee.Password = hashedPassword
	newEmp, err := service.employeeRepository.CreateEmployee(employee)
	if err != nil {
		return dto.EmployeeResponseDto{}, err
	}
	return dto.ToEmployeeResponseDto(newEmp), nil

}

func (e *employeeService) GetEmployee(id string) (dto.EmployeeUserResponseDto, error) {
	employee, err := e.employeeRepository.GetEmployee(id)
	if err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return employee, nil
}
func (e *employeeService) GetEmployeeProfile(id string) (dto.EmployeeUserResponseDto, error) {
	employee, err := e.employeeRepository.GetEmployeeByUserID(id)
	if err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return employee, nil
}

func (e *employeeService) GetEmployees() ([]dto.EmployeeResponseDto, error) {
	emps, err := e.employeeRepository.GetEmployees()
	if err != nil {
		return []dto.EmployeeResponseDto{}, err
	}

	employees := []dto.EmployeeResponseDto{}
	for _, ep := range emps {
		employees = append(employees, dto.ToEmployeeResponseDto(ep))
	}
	return employees, nil
}

func (e *employeeService) UpdateEmployee(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error) {
	updatedEmployee, err := e.employeeRepository.UpdateEmployee(id, employee)
	if err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return updatedEmployee, nil
}
func (e *employeeService) UpdateEmployeeProfile(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error) {
	updatedEmployee, err := e.employeeRepository.UpdateEmployeeByUserID(id, employee)
	if err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return updatedEmployee, nil
}

func (e *employeeService) DeleteEmployee(id string) error {
	err := e.employeeRepository.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}
