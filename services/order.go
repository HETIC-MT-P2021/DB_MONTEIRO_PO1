package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetOrder : Get order infos with products details
func GetOrder(orderID int) (structs.OrderView, error) {
	var orderRender structs.OrderView
	var products []models.Product

	orderID, orderProductsCode, err := models.GetProductsCode(10123)
	if err != nil {
		return orderRender, err
	}

	for _, productCode := range orderProductsCode {
		product, err := models.GetProduct(productCode)
		if err != nil {
			return orderRender, err
		}

		products = append(products, product)
	}

	orderRender = structs.OrderView{
		orderID,
		products,
	}

	return orderRender, err
}
