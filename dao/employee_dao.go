package dao

import "time"

type EmployeeDAO struct {
	Id float64

	UserId float64

	FirstName string

	LastName string

	CNP string

	JobTitle string

	City string

	PhoneNumber string

	WorkEmail string

	PersonalEmail string

	EmployeeRateType string

	EmployeeRateValue float64

	ContractType string

	Currency string

	VacationDays float64

	StartingDate time.Time

	EndingDate time.Time

	HasMedicalPackage bool

	Status bool
}
