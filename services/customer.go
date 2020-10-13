package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetCustomer : Get customer infos with number of ordered products and total price of ordered products
func GetCustomer(customerID int) (structs.CustomerView, error) {
	var customerRender structs.CustomerView
	var totalAndPriceProducts []models.TotalAndPriceProduct

	customer, ordersID, err := models.GetCustomer(customerID)
	if err != nil {
		return structs.CustomerView{}, err
	}

	for _, orderID := range ordersID {
		totalAndPriceProduct, err := models.GetTotalAndPriceProduct(orderID)
		if err != nil {
			return structs.CustomerView{}, err
		}

		totalAndPriceProducts = append(totalAndPriceProducts, totalAndPriceProduct)
	}

	customerRender = structs.CustomerView{
		customer,
		totalAndPriceProducts,
	}

	return customerRender, nil
}
