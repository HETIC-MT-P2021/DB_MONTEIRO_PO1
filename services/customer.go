package services

import (
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/structs"
)

// GetCustomer : Get customer infos with number of ordered products and total price of ordered products
func GetCustomer(customerID int) (structs.Customer, error) {
	var customerRender structs.Customer

	customer, err := models.GetCustomer(customerID)
	if err != nil {
		return structs.Customer{}, err
	}

	totalPriceAndProduct, err := models.GetTotalAndPriceProduct(customer.OrderNumber)
	if err != nil {
		return structs.Customer{}, err
	}

	customerRender = structs.Customer{
		customer,
		totalPriceAndProduct.NbProductOrdered,
		totalPriceAndProduct.TotalPrice,
	}

	return customerRender, nil
}
