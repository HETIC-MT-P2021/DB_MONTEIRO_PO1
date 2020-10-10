package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetOffices : Get offices infos with associated employees
func GetOffices() (structs.OfficeView, error) {
	var officesRender structs.OfficeView
	var officesWithEmployees []models.OfficeWithEmployees

	offices, err := models.GetOffices()
	if err != nil {
		return officesRender, err
	}

	for _, office := range offices {
		employees, err := models.GetEmployeesOffice(office.OfficeCode)
		if err != nil {
			return officesRender, err
		}

		officeWithEmployees := models.OfficeWithEmployees{
			office,
			employees,
		}

		officesWithEmployees = append(officesWithEmployees, officeWithEmployees)
	}

	officesRender = structs.OfficeView{
		officesWithEmployees,
	}

	return officesRender, nil
}
