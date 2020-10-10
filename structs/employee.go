package structs

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
)

// EmployeesView : Architecture rendered to view
type EmployeesView struct {
	Employees []models.Employee `json:"employees"`
}
