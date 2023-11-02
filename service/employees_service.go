package service

import (
	"github.com/AntoniuStirb/ProjectSwaggerGo/database"
	"github.com/AntoniuStirb/ProjectSwaggerGo/models"
	"log"
)

func GetAllEmployees() ([]models.Employee, error) {
	rows, err := database.DB.Query("SELECT * FROM employees")
	if err != nil {
		log.Printf("Database error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
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
