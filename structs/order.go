package structs

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
)

// OrderView : Architecture rendered to view
type OrderView struct {
	OrderID  int              `json:"orderID"`
	Products []models.Product `json:"products"`
}
