package structs

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
)

// Customer : Architecture rendered to view
type Customer struct {
	models.Customer
	NbProductOrdered int     `json:"nbProductOrdered"`
	TotalPrice       float32 `json:"totalPrice"`
}
