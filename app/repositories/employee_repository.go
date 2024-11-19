package repository

import (
	"employeeManagement/app/dto"
	"employeeManagement/app/entities"
	db "employeeManagement/pkg"
	"fmt"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(employee dto.EmployeeRequestDto) (entities.Employee, error)
	GetEmployees() ([]entities.Employee, error)
	GetEmployee(id string) (dto.EmployeeUserResponseDto, error)
	GetEmployeeByUserID(id string) (dto.EmployeeUserResponseDto, error)
	UpdateEmployee(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error)
	UpdateEmployeeByUserID(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error)
	DeleteEmployee(id string) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{
		db: db.GetDBConnection(),
	}
}

func (r *employeeRepository) CreateEmployee(employee dto.EmployeeRequestDto) (entities.Employee, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	userEntity := dto.ToUserEmployeeEntity(employee)
	if err := tx.Create(&userEntity).Error; err != nil {
		tx.Rollback()
		return entities.Employee{}, err
	}
	employeeEntity := dto.ToEmployeeEntity(employee, userEntity)
	if err := tx.Create(&employeeEntity).Error; err != nil {
		tx.Rollback()
		return entities.Employee{}, err
	}
	if err := tx.Commit().Error; err != nil {
		return entities.Employee{}, err
	}
	return employeeEntity, nil
}

func (r *employeeRepository) GetEmployee(id string) (dto.EmployeeUserResponseDto, error) {
	employee := entities.Employee{}
	if err := r.db.Where("id = ?", id).First(&employee).Error; err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	user := entities.User{}
	if err := r.db.Where("id = ?", employee.UserID).First(&user).Error; err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return dto.ToEmployeeUserResponseDto(employee, user), nil
}

func (r *employeeRepository) GetEmployeeByUserID(id string) (dto.EmployeeUserResponseDto, error) {
	user := entities.User{}
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	employee := entities.Employee{}
	if err := r.db.Where("user_id = ?", id).First(&employee).Error; err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return dto.ToEmployeeUserResponseDto(employee, user), nil
}

func (r *employeeRepository) GetEmployees() ([]entities.Employee, error) {

	employees := []entities.Employee{}
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) UpdateEmployee(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	employeeEntity := entities.Employee{}
	if err := tx.First(&employeeEntity, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, fmt.Errorf("employee not found")
	}
	userEntity := entities.User{}
	if err := tx.First(&userEntity, "id = ?", employeeEntity.UserID).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, fmt.Errorf("user not found")
	}
	if err := r.db.Model(&employeeEntity).Updates(entities.Employee{
		Name:       employee.Name,
		Position:   employee.Position,
		Department: employee.Department,
		Salary:     employee.Salary,
	}).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, err
	}
	if err := r.db.Model(&userEntity).Updates(entities.User{
		Email: employee.Email,
	}).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, err
	}
	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return dto.ToEmployeeUserResponseDto(employeeEntity, userEntity), nil
}
func (r *employeeRepository) UpdateEmployeeByUserID(id string, employee dto.UpdateEmployeeRequestDto) (dto.EmployeeUserResponseDto, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	employeeEntity := entities.Employee{}
	if err := tx.First(&employeeEntity, "user_id = ?", id).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, fmt.Errorf("employee not found")
	}
	userEntity := entities.User{}
	if err := tx.First(&userEntity, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, fmt.Errorf("user not found")
	}
	if err := r.db.Model(&employeeEntity).Updates(entities.Employee{
		Name:       employee.Name,
		Position:   employee.Position,
		Department: employee.Department,
		Salary:     employee.Salary,
	}).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, err
	}
	if err := r.db.Model(&userEntity).Updates(entities.User{
		Email: employee.Email,
	}).Error; err != nil {
		tx.Rollback()
		return dto.EmployeeUserResponseDto{}, err
	}
	if err := tx.Commit().Error; err != nil {
		return dto.EmployeeUserResponseDto{}, err
	}
	return dto.ToEmployeeUserResponseDto(employeeEntity, userEntity), nil
}

func (r *employeeRepository) DeleteEmployee(id string) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	employeeEntity := entities.Employee{}
	if err := tx.First(&employeeEntity, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("employee not found")
	}
	userEntity := entities.User{}
	if err := tx.First(&userEntity, "id = ?", employeeEntity.UserID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("user not found")
	}
	if err := tx.Delete(&employeeEntity).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&userEntity).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
