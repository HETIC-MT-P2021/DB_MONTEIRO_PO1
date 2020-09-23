package models

import (
	"fmt"
)

type Customer struct {
	ID                     int    `json:"customerNumber"`
	Name                   string `json:"customerName"`
	ContactLastName        string `json:"contactLastName"`
	ContactFirstName       string `json:"contactFirstName"`
	Phone                  string `json:"phone"`
	AddressLine1           string `json:"addressLine1"`
	AddressLine2           string `json:"addressLine2"`
	City                   string `json:"city"`
	State                  string `json:"state"`
	PostalCode             string `json:"postalCode"`
	Country                string `json:"country"`
}

func GetCustomer(customer_id int) (Customer, error) {
	var customer Customer

	// Execute the query
	search := "SELECT customers.customerNumber, customers.customerName, customers.contactLastName," +
			  "customers.contactFirstName, customers.phone, customers.addressLine1," +
			  "customers.addressLine2, customers.city, customers.state," +
			  "customers.postalCode, customers.country "
	from := "FROM customers "
	where := "WHERE customers.customerNumber = ?"

	err := DB.QueryRow(search + from + where, customer_id).Scan(&customer.ID,
																&customer.Name,
																&customer.ContactLastName,
																&customer.ContactFirstName,
																&customer.Phone,
																&customer.AddressLine1,
																&customer.AddressLine2,
																&customer.City,
																&customer.State,
																&customer.PostalCode,
																&customer.Country)
	if err != nil {
		fmt.Println(err.Error())
		return customer, err
	}

	customerOrders, err := GetLengthOrders(customer_id)
	if err != nil {
		fmt.Println(customerOrders)
		return customer, err
	}

	return customer, nil
}
