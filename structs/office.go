package structs

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
)

// OfficeView : Architecture rendered to view
type OfficeView struct {
	Offices []models.OfficeWithEmployees `json:"offices"`
}
