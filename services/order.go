package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetOrder : Get order infos with products details
func GetOrder(orderID int) (structs.OrderView, error) {
	var orderRender structs.OrderView

	orderID, orderProducts, err := models.GetProductsCode(10123)
	if err != nil {
		return orderRender, err
	}

	orderRender = structs.OrderView{
		orderID,
		orderProducts,
	}

	return orderRender, err
}
