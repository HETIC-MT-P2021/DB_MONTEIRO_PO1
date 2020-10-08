package models

// Customer : Architecture in database
type Customer struct {
	CustomerNumber   int        `json:"customerNumber"`
	Name             string     `json:"customerName"`
	ContactLastName  string     `json:"contactLastName"`
	ContactFirstName string     `json:"contactFirstName"`
	Phone            string     `json:"phone"`
	AddressLine1     string     `json:"addressLine1"`
	AddressLine2     NullString `json:"addressLine2"`
	City             string     `json:"city"`
	State            NullString `json:"state"`
	PostalCode       string     `json:"postalCode"`
	Country          string     `json:"country"`
	Order
}

// GetCustomer : Get customer infos and associated orders id
func GetCustomer(customerID int) (Customer, error) {
	var customer Customer

	query := `
		SELECT 	
			customers.customerNumber,
			customers.customerName,
			customers.contactLastName,
			customers.contactFirstName,
			customers.phone,
			customers.addressLine1,
			customers.addressLine2,
			customers.city,
			customers.state,
			customers.postalCode,
			customers.country,
			orders.orderNumber
		FROM customers
		JOIN orders ON customers.customerNumber = orders.customerNumber
		WHERE customers.customerNumber = ?
	`

	err := DB.QueryRow(query, customerID).Scan(
		&customer.CustomerNumber,
		&customer.Name,
		&customer.ContactLastName,
		&customer.ContactFirstName,
		&customer.Phone,
		&customer.AddressLine1,
		&customer.AddressLine2,
		&customer.City,
		&customer.State,
		&customer.PostalCode,
		&customer.Country,
		&customer.Order.OrderNumber)

	if err != nil {
		return customer, err
	}

	return customer, nil
}
