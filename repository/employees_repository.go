package repository

import (
	"fmt"
	"github.com/AntoniuStirb/ProjectSwaggerGo/dao"
	"github.com/AntoniuStirb/ProjectSwaggerGo/database"
	"log"
)

type EmployeeRepositoryInterface interface {
	GetAllEmployees() ([]dao.EmployeeDAO, error)
}

func GetAllEmployees() ([]dao.EmployeeDAO, error) {
	rows, err := database.DB.Query("SELECT * FROM employees")
	if err != nil {
		log.Printf("Database error: %v", err)
		return nil, err
	}
	defer rows.Close()
	fmt.Println(rows)
	var employees []dao.EmployeeDAO
	for rows.Next() {
		var employee dao.EmployeeDAO
		err := rows.Scan(
			&employee.Id,
			&employee.UserId,
			&employee.FirstName,
			&employee.LastName,
			&employee.CNP,
			&employee.JobTitle,
			&employee.City,
			&employee.PhoneNumber,
			&employee.WorkEmail,
			&employee.PersonalEmail,
			&employee.EmployeeRateType,
			&employee.EmployeeRateValue,
			&employee.ContractType,
			&employee.Currency,
			&employee.VacationDays,
			&employee.StartingDate,
			&employee.EndingDate,
			&employee.HasMedicalPackage,
			&employee.Status,
		)
		if err != nil {
			log.Printf("Database error: %v", err)
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
