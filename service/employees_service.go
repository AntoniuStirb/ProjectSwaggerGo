package service

import (
	"github.com/AntoniuStirb/ProjectSwaggerGo/dto"
	"github.com/AntoniuStirb/ProjectSwaggerGo/repository"
)

type EmployeeServiceInterface interface {
	GetAllEmployees() ([]dto.EmployeeDTO, error)
}

func GetAllEmployees() ([]dto.EmployeeDTO, error) {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	employeeDTOs := make([]dto.EmployeeDTO, len(employees))
	for i, employee := range employees {
		employeeDTOs[i] = dto.EmployeeDTO{
			UserId:        employee.UserId,
			FirstName:     employee.FirstName,
			LastName:      employee.LastName,
			PhoneNumber:   employee.PhoneNumber,
			WorkEmail:     employee.WorkEmail,
			PersonalEmail: employee.PersonalEmail,
			Status:        employee.Status,
		}
	}
	return employeeDTOs, nil
}
