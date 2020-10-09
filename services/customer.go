package services

import (
	"fmt"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetCustomer : Get customer infos with number of ordered products and total price of ordered products
func GetCustomer(customerID int) (structs.CustomerView, error) {
	var customerRender structs.CustomerView
	var totalAndPriceProductSlice []models.TotalAndPriceProduct

	customer, ordersID, err := models.GetCustomer(customerID)
	if err != nil {
		return structs.CustomerView{}, err
	}

	for _, orderID := range ordersID {
		fmt.Println(orderID)
		totalAndPriceProduct, err := models.GetTotalAndPriceProduct(orderID)
		if err != nil {
			return structs.CustomerView{}, err
		}

		totalAndPriceProductSlice = append(totalAndPriceProductSlice, totalAndPriceProduct)
	}

	customerRender = structs.CustomerView{
		customer,
		totalAndPriceProductSlice,
	}

	return customerRender, nil
}
