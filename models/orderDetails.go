package models

// TotalAndPriceProduct : Calculated with sql query
// NbProductOrdered : product * quantity (product and quantity are stored in db)
// TotalPrice : price of a product * quantity (price and quantity are stored in db)
type TotalAndPriceProduct struct {
	OrderID          int
	NbProductOrdered int
	TotalPrice       float32
}
