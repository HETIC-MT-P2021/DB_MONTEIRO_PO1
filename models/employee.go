package models

// Employee : Architecture in db
type Employee struct {
	EmployeeNumber int    `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Extension      string `json:"extension"`
	Email          string `json:"email"`
	OfficeCode     string `json:"officeCode"`
	JobTitle       string `json:"jobTitle"`
}

// GetEmployees : Get all amployees
func GetEmployees() ([]Employee, error) {
	var employees []Employee
	var employee Employee

	query := `
		SELECT
			employees.employeeNumber,
			employees.lastName,
			employees.firstName,
			employees.extension,
			employees.email,
			employees.officeCode,
			employees.jobTitle
		FROM employees
	`

	employeesResult, err := DB.Query(query)

	if err != nil {
		return employees, err
	}

	for employeesResult.Next() {
		err := employeesResult.Scan(
			&employee.EmployeeNumber,
			&employee.LastName,
			&employee.FirstName,
			&employee.Extension,
			&employee.Email,
			&employee.OfficeCode,
			&employee.JobTitle)

		if err != nil {
			return employees, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
