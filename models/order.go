package models

// Order : Architecture in database
type Order struct {
	OrderNumber  int    `json:"orderNumber"`
	OrderDate    string `json:"orderDate"`
	RequiredDate string `json:"requiredDate"`
	ShippedDate  string `json:"shippedDate"`
	Status       string `json:"status"`
	Comments     string `json:"comments"`
}

// GetTotalAndPriceProduct : Get total and price of all products ordered
func GetTotalAndPriceProduct(orderID int) (TotalAndPriceProduct, error) {
	var totalAndPriceProduct TotalAndPriceProduct

	query := `
		SELECT
		    orderdetails.orderNumber,
			SUM(orderdetails.quantityOrdered) AS nbProductOrdered,
			SUM(orderdetails.quantityOrdered * orderdetails.priceEach) AS totalPrice
		FROM orderdetails
		JOIN orders ON orderdetails.orderNumber = orders.orderNumber
		WHERE orders.orderNumber = ?
		GROUP BY orderdetails.orderNumber
	`

	err := DB.QueryRow(query, 10123).Scan(
		&totalAndPriceProduct.OrderID,
		&totalAndPriceProduct.NbProductOrdered,
		&totalAndPriceProduct.TotalPrice)

	if err != nil {
		return totalAndPriceProduct, err
	}

	return totalAndPriceProduct, nil
}
