package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetOffices : Get offices infos with associated employees
func GetOffices() (structs.OfficeView, error) {
	var officesRender structs.OfficeView

	offices, err := models.GetOffices()
	if err != nil {
		return officesRender, err
	}

	officesRender = structs.OfficeView{
		offices,
	}

	return officesRender, nil
}
