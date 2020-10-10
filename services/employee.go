package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetEmployees : Get employees infos with associated offices
func GetEmployees() (structs.EmployeesView, error) {
	var employeesRender structs.EmployeesView

	employees, err := models.GetEmployees()
	if err != nil {
		return employeesRender, err
	}

	employeesRender = structs.EmployeesView{
		employees,
	}

	return employeesRender, nil
}
