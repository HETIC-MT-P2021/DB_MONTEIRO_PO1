package structs

import "github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"

// Customer : Architecture rendered to view
type CustomerView struct {
	models.Customer
	Orders []models.TotalAndPriceProduct `json:"orders"`
}
